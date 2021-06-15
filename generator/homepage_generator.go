package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"
)

type HomepageGenerator struct {
	GenConfig
}

// TODO: Possibly add variables for homepage?
type HomePageVar struct{}

func (g HomepageGenerator) Generate() error {
	templateToUse := g.Template

	buffer := bytes.Buffer{}
	if err := templateToUse.Execute(&buffer, HomePageVar{}); err != nil {
		return fmt.Errorf("failed to execute template for homepage")
	}

	html := template.HTML(buffer.String())
	fullPath := filepath.Join(g.Destination, "index.html")
	if err := g.Writer.InsertIntoBase(html, "Welcome!", fullPath); err != nil {
		return err
	}

	return nil
}
