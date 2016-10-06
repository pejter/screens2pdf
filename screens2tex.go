package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

type exercise struct {
	Title      string
	ScreenList []string
}

func main() {
	validExtensions := []string{".jpg", ".png", ".pdf"}
	sort.Strings(validExtensions)
	var (
		exerciseDirectory    string
		exerciseTitle        string
		screenshotsDirectory string
	)

	flag.StringVar(&exerciseDirectory, "dir", "", "Location of the exercise")
	flag.StringVar(&exerciseDirectory, "d", "", "Location of the exercise")
	flag.StringVar(&exerciseTitle, "title", "", "Title of the exercise")
	flag.StringVar(&exerciseTitle, "t", "", "Title of the exercise")
	flag.StringVar(&screenshotsDirectory, "scrotdir", "", "Location of the screenshots")
	flag.StringVar(&screenshotsDirectory, "s", "", "Location of the screenshots")

	flag.Parse()
	if exerciseTitle == "" {
		log.Fatal("You must specify a name for the exercise")
	}
	if exerciseDirectory == "" {
		log.Fatal("You must specify a directory for the exercise")
	}
	if screenshotsDirectory == "" {
		screenshotsDirectory = exerciseDirectory + "/screens"
	}

	screenList, err := ioutil.ReadDir(screenshotsDirectory)
	if err != nil {
		log.Fatal("Couldn't read files inside "+screenshotsDirectory+"!\n", err)
	}

	exerciseData := exercise{exerciseTitle, nil}
	for _, s := range screenList {
		ext := filepath.Ext(s.Name())
		i := sort.SearchStrings(validExtensions, ext)
		if i < len(validExtensions) && ext == validExtensions[i] {
			exerciseData.ScreenList = append(exerciseData.ScreenList, strings.TrimSuffix(s.Name(), ext))
		}
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
