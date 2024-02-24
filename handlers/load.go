package handlers

import (
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"

	"todolist/util"
)

// go's filepath.Match doesn't support double globs (**)
func getPathsOfTemplates() []string {
	templs := []string{}
	re := regexp.MustCompile(".*\\.htm[l]?$")
	err := filepath.WalkDir(filepath.Join(util.Must(os.Getwd()), "views"), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			if matches := re.FindStringSubmatch(path); len(matches) > 0 {
				templs = append(templs, matches[0])
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return templs
}

// Load all templates
var T = template.Must(template.ParseFiles(getPathsOfTemplates()...))
