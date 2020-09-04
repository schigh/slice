package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

type typeBucket struct {
	BaseType    string
	PackageType string
	ZeroValue   string
	BaseFile    string
	TestFile    string
}

func main() {
	dest, err := filepath.Abs("../../")
	if err != nil {
		panic(err)
	}
	buckets := []typeBucket{
		{"int", "Int", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"int8", "Int8", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"int16", "Int16", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"int32", "Int32", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"int64", "Int64", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"uint", "UInt", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"uint8", "UInt8", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"uint16", "UInt16", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"uint32", "UInt32", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"uint64", "UInt64", "0", "./tmpl/base.go.tpl", "./tmpl/integertype_test.go.tpl"},
		{"float64", "Float64", "0", "./tmpl/base.go.tpl", "./tmpl/floattype_test.go.tpl"},
		{"float32", "Float32", "0", "./tmpl/base.go.tpl", "./tmpl/floattype_test.go.tpl"},
		{"string", "String", `""`, "./tmpl/base.go.tpl", "./tmpl/stringtype_test.go.tpl"},
	}

	for _, bucket := range buckets {
		baseFileBytes, readBaseErr := ioutil.ReadFile(bucket.BaseFile)
		if readBaseErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Opening integertype base file failed with error: %v", readBaseErr)
			os.Exit(1)
		}

		testFileBytes, readTestErr := ioutil.ReadFile(bucket.TestFile)
		if readTestErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Opening integertype test file failed with error: %v", readTestErr)
			os.Exit(1)
		}

		baseTmpl := template.Must(template.New("base").Parse(string(baseFileBytes)))
		testTmpl := template.Must(template.New("test").Parse(string(testFileBytes)))

		var execErr, writeErr, sourceErr error
		var baseBytes, testBytes, baseSourceBytes, testSourceBytes []byte
		baseBuff := bytes.NewBuffer(baseBytes)
		testBuff := bytes.NewBuffer(testBytes)
		if execErr = baseTmpl.Execute(baseBuff, bucket); execErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "executing template failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, execErr)
			os.Exit(1)
		}

		if execErr = testTmpl.Execute(testBuff, bucket); execErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "executing template failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, execErr)
			os.Exit(1)
		}

		baseSourceBytes, sourceErr = format.Source(baseBuff.Bytes())
		if sourceErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "canonical source formatting failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, sourceErr)
			os.Exit(1)
		}
		baseBuff = bytes.NewBuffer(baseSourceBytes)

		testSourceBytes, sourceErr = format.Source(testBuff.Bytes())
		if sourceErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "canonical source formatting failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, sourceErr)
			os.Exit(1)
		}
		testBuff = bytes.NewBuffer(testSourceBytes)

		fileName := filepath.Join(dest, fmt.Sprintf("%s.go", bucket.BaseType))
		if writeErr = ioutil.WriteFile(fileName, baseBuff.Bytes(), 0644); writeErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "writing base file failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, writeErr)
			os.Exit(1)
		}

		testName := filepath.Join(dest, fmt.Sprintf("%s_test.go", bucket.BaseType))
		if writeErr = ioutil.WriteFile(testName, testBuff.Bytes(), 0644); writeErr != nil {
			_, _ = fmt.Fprintf(os.Stderr, "writing test file failed for %s/%s: %v", bucket.BaseType, bucket.PackageType, writeErr)
			os.Exit(1)
		}
	}
}
