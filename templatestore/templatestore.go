package templatestore

import (
	"html/template"
	"path/filepath"

	"github.com/TanLeYang/go-ssg/directoryhelper"
)

func PostTemplate() (*template.Template, error) {
	templateName := "posts.html"
	return read(templateName)
}

func BaseTemplate() (*template.Template, error) {
	templateName := "template.html"
	return read(templateName)
}

func ShortTemplate() (*template.Template, error) {
	templateName := "short.html"
	return read(templateName)
}

func HomeTemplate() (*template.Template, error) {
	templateName := "home.html"
	return read(templateName)
}

func read(name string) (*template.Template, error) {
	fullPath := filepath.Join(getBaseTemplateFolder(), name)
	template, err := template.ParseFiles(fullPath)

	if err != nil {
		return nil, err
	}

	return template, err
}

func getBaseTemplateFolder() string {
	return filepath.Join(directoryhelper.RootDir(), "templates")
}
