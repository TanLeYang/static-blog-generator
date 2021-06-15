package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/TanLeYang/go-ssg/models"
)

// ShortGenerator generates the page showing a list of shorts, each short links
// to a corresponding full length post
type ShortGenerator struct {
	GenConfig
	Shorts []models.Short
}

func (g ShortGenerator) Generate() error {
	shorts := g.Shorts
	templateToUse := g.Template

	allShorts := bytes.Buffer{}
	for _, short := range shorts {
		buffer := bytes.Buffer{}

		if err := templateToUse.Execute(&buffer, short); err != nil {
			return fmt.Errorf("failed to execute template for short with title: %s, %s", short.Title, err.Error())
		}

		allShorts.Write(buffer.Bytes())
		allShorts.WriteString("<br>")
	}

	html := template.HTML(allShorts.String())
	fullPath := filepath.Join(g.Destination, "blog.html")
	if err := g.Writer.InsertIntoBase(html, "Blog", fullPath); err != nil {
		return err
	}

	return nil
}
