package generator

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/TanLeYang/go-ssg/models"
	"github.com/TanLeYang/go-ssg/parser"
)

type BaseTemplateWriter struct {
	models.SiteConfig
	BaseTemplate *template.Template
}

type IndexData struct {
	HTMLTitle string
	PageTitle string
	BaseURL   template.URL
	Content   template.HTML
	CustomCSS template.CSS
}

// InsertIntoBase inserts the given HTML content into the shared basic
// HTML template and writes it into the given path
func (w BaseTemplateWriter) InsertIntoBase(content template.HTML, pageTitle string, fullPath string) error {
	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("error creating file %s: %v", fullPath, err)
	}
	writer := bufio.NewWriter(f)

	syntaxHighlightingCSS, err := parser.GetSyntaxHighlightCSS()
	if err != nil {
		log.Printf("error when getting syntax highlighting CSS")
	}

	indexData := IndexData{
		HTMLTitle: getHTMLTitle(pageTitle, w.BlogTitle),
		PageTitle: pageTitle,
		BaseURL:   template.URL(w.BaseURL),
		Content:   content,
		CustomCSS: template.CSS(syntaxHighlightingCSS),
	}

	if err := w.BaseTemplate.Execute(writer, indexData); err != nil {
		return fmt.Errorf("error executing template %s: %v", fullPath, err)
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("error writing to file %s: %v", fullPath, err)
	}

	return nil
}

func getHTMLTitle(pageTitle, blogTitle string) string {
	if pageTitle == "" {
		return blogTitle
	}

	return fmt.Sprintf("%s - %s", pageTitle, blogTitle)
}
