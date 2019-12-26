package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

const (
	suffix = "_sqlutil.go"
)

var (
	src          string
	dst          string
	mod          string
	curStruct    string
	curPkg       string
	srcPkg       string
	curColumns   []string
	curValues    []string
	curFields    []string
	curTableName string
	curBys       []by
	curImported  bool
	tableNameReg = regexp.MustCompile(`table:(\w+)`)
)

type by struct {
	Name string
	Type string
}

type tpl struct {
	Pkg         string
	Name        string
	Scan        string
	Column      string
	ColumnCount int
	TableName   string
	Value       string
	PlaceHolder string
	Bys         []by
}

func init() {
	flag.StringVar(&src, "src", ".", "-src=.")
	flag.StringVar(&mod, "mod", "", "-mod=.")
	flag.StringVar(&dst, "dst", "", "-dst=.")
	flag.Parse()
}

func main() {
	if src == "" {
		flag.PrintDefaults()
		return
	}
	if dst != "" && mod == "" {
		flag.PrintDefaults()
		return
	}
	genStruct()
}

func genStruct() {
	if err := filepath.Walk(src, func(curPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		name := info.Name()
		if strings.HasSuffix(name, suffix) || !strings.HasSuffix(name, ".go") {
			return nil
		}
		p, err := build.ImportDir(src, 0)
		if err != nil {
			return err
		}
		curPkg = p.Name
		srcPkg = p.Name
		var baseDir string
		if dst == "" {
			baseDir = filepath.Dir(curPath)
		} else {
			curPkg = path.Base(dst)
			baseDir = dst
		}
		dst := filepath.Join(baseDir, strings.Replace(name, ".go", suffix, -1))

		buf := bytes.NewBuffer(nil)

		if err := generate(curPath, buf); err != nil {
			return err
		}
		if buf.Len() != 0 {
			if err := ioutil.WriteFile(dst, buf.Bytes(), 0644); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

func generate(path string, buf *bytes.Buffer) error {
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	for _, decl := range file.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		if len(gd.Specs) == 0 {
			continue
		}
		if len(gd.Doc.List) != 0 {
			curTableName = getTableName(gd.Doc.List[0].Text)
		}
		if err := gen(gd.Specs, buf); err != nil {
			return err
		}
		curStruct = ""
		curPkg = ""
		curColumns = nil
		curFields = nil
		curTableName = ""
		curBys = nil
		curValues = nil
	}
	curImported = false
	return nil
}

func getTableName(comment string) string {
	subMatch := tableNameReg.FindStringSubmatch(comment)
	if len(subMatch) != 0 {
		return subMatch[1]
	}
	return ""
}

func gen(specs []ast.Spec, buf io.Writer) error {
	if !curImported {
		curImported = true
		if dst != "" {
			io.WriteString(buf, fmt.Sprintf(`package %s
			import(
				"database/sql"
				"context"
				"fmt"
				"strings"
				"%s"
			)`, curPkg, path.Join(mod, src)))
		} else {
			io.WriteString(buf, fmt.Sprintf(`package %s
			import(
				"database/sql"
				"context"
				"fmt"
				"strings"
			)`, curPkg))
		}

	}
	for _, spec := range specs {
		ts, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		st, ok := ts.Type.(*ast.StructType)
		if !ok {
			continue
		}
		curStruct = ts.Name.Name
		if curTableName == "" {
			curTableName = strings.ToLower(curStruct)
		}

		for _, field := range st.Fields.List {
			if field.Tag != nil {
				trimedValue := strings.Trim(field.Tag.Value, "`")
				tagValue := reflect.StructTag(trimedValue).Get("sqlutil")
				tagValues := strings.Split(tagValue, ",")
				if len(tagValues) == 0 {
					continue
				}
				ident, ok := field.Type.(*ast.Ident)
				if !ok {
					continue
				}
				curColumns = append(curColumns, tagValues[0])
				curValues = append(curValues, "obj."+field.Names[0].Name)
				curFields = append(curFields, `&result.`+field.Names[0].Name)
				if len(tagValues) > 1 && tagValues[1] == "by" {
					curBys = append(curBys, by{Name: tagValues[0], Type: ident.Name})
				}
			}
		}
		if err := execTpl(buf); err != nil {
			return err
		}
	}
	return nil
}

func execTpl(buf io.Writer) error {
	var pkg string
	if dst != "" {
		pkg = srcPkg + "."
	}
	tpl := tpl{Name: curStruct, Pkg: pkg, Column: strings.Join(curColumns, ", "), Bys: curBys, Scan: strings.Join(curFields, ", "), TableName: curTableName, PlaceHolder: strings.Repeat("?, ", len(curColumns)-1) + "?", Value: strings.Join(curValues, ", "), ColumnCount: len(curColumns)}
	t, err := template.New("sqlutil").Funcs(template.FuncMap{
		"raw":   raw,
		"title": title,
	}).Parse(tplStr)
	if err != nil {
		return err
	}
	if err := t.Execute(buf, tpl); err != nil {
		return err
	}
	return nil
}

func raw(prev string) template.HTML {
	return template.HTML(prev)
}

func title(prev string) string {
	return strings.Title(strings.Replace(prev, "id", "ID", -1))
}
