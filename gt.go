// gt stands for go touch is a tool to create
// a new file in go programming language format
// Copyright 2013 Li Bin <libin_dba@xiaomi.com>.
// All rights reserved.

package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"text/template"
	"time"
)

type Temp struct {
	PackageName string
	Year        int
}

const LETTER = `//Copyright {{.Year}} Li Bin <libin_dba@xiaomi.com>.
//All rights reserved.

package {{.PackageName}}

import (
	""
)
`

var file = flag.String("f", "", "Name used to create a new file")
var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	var name string

	flag.Parse()

	if len(*file) > 0 {
		name = getFileName(*file)
	} else {
		Usage()
	}

	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	path, _ := os.Getwd()
	pkgName := getPkgName(path)

	year, _, _ := time.Now().Date()

	c := &Temp{pkgName, year}

	t := template.Must(template.New("letter").Parse(LETTER))
	err = t.Execute(file, c)
	if err != nil {
		panic(err)
	}
}

func getPkgName(path string) string {
	patten := regexp.MustCompile(`/(\w+)$`)
	return patten.FindStringSubmatch(path)[1]
}

func getFileName(name string) string {
	patten := regexp.MustCompile(`(\w+)(\.go)?$`)
	return patten.FindStringSubmatch(name)[1] + ".go"
}
