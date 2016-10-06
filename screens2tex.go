package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type exercise struct {
	Title      string
	ScreenList []string
}

func main() {
	exerciseDirectory := os.Args[1]
	exerciseTitle := os.Args[2]

	screenList, err := ioutil.ReadDir(exerciseDirectory + "/screens")
	if err != nil {
		log.Fatal("Couldn't read files inside "+exerciseDirectory+"/screens"+"!\n", err)
	}

	exerciseData := exercise{exerciseTitle, nil}
	for _, s := range screenList {
		exerciseData.ScreenList = append(exerciseData.ScreenList, strings.TrimSuffix(s.Name(), filepath.Ext(s.Name())))
	}

	t, err := template.ParseFiles("template.tex")
	if err != nil {
		log.Fatal("Template couldn't be parsed!\n", err)
	}

	f, err := os.Create(exerciseDirectory + "/final.tex")
	defer f.Close()
	if err != nil {
		log.Fatal("File "+f.Name()+" couldn't be created!\n", err)
	}

	t.Execute(f, exerciseData)
}
