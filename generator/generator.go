package generator

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/TanLeYang/go-ssg/datasource"
	"github.com/TanLeYang/go-ssg/directoryhelper"
	"github.com/TanLeYang/go-ssg/templatestore"
)

// Generator is the interface that wraps the Generate method.
//
// Generate generates a HTML file and writes it to a underlying destination folder
type Generator interface {
	Generate() error
}

// GenConfig contains the template for a Generator to build the HTML file
// the destination to write it to and the BaseTemplateWriter to do the
// writing
type GenConfig struct {
	Template    *template.Template
	Destination string
	Writer      *BaseTemplateWriter
}

// SiteGenerator generates the entire site by calling
// other generators to generate their respective content
type SiteGenerator struct {
	Config           GenConfig
	PostsDatasource  datasource.PostsDatasource
	ShortsDatasource datasource.ShortsDatasource
}

func (sg SiteGenerator) Generate() error {
	createOutputDir()
	sg.genHomepage()
	sg.genPosts()
	sg.genShorts()

	return nil
}

func (sg SiteGenerator) genHomepage() error {
	template, err := templatestore.HomeTemplate()
	if err != nil {
		log.Printf("error when getting homepage template %s", err.Error())
		return err
	}

	homepageGenerator := HomepageGenerator{
		GenConfig: GenConfig{
			Template:    template,
			Destination: directoryhelper.OutputDir(),
			Writer:      sg.Config.Writer,
		},
	}

	if err := homepageGenerator.Generate(); err != nil {
		log.Printf("error when generating homepage: %s", err.Error())
		return err
	}

	return nil
}

func (sg SiteGenerator) genShorts() error {
	shorts, err := sg.ShortsDatasource.GetAllShorts()
	if err != nil {
		log.Printf("error when getting all shorts: %s", err.Error())
		return err
	}

	shortTemplate, err := templatestore.ShortTemplate()
	if err != nil {
		log.Printf("error when getting short template: %s", err.Error())
		return err
	}

	shortGenerator := ShortGenerator{
		GenConfig: GenConfig{
			Template:    shortTemplate,
			Destination: directoryhelper.OutputDir(),
			Writer:      sg.Config.Writer,
		},
		Shorts: shorts,
	}

	if err := shortGenerator.Generate(); err != nil {
		log.Printf("error when generating shorts: %s", err.Error())
		return err
	}

	return nil
}

func (sg SiteGenerator) genPosts() error {
	posts, err := sg.PostsDatasource.GetAllPosts()
	if err != nil {
		log.Printf("error when getting all posts: %s", err.Error())
		return err
	}

	postTemplate, err := templatestore.PostTemplate()
	if err != nil {
		log.Printf("error when getting post template: %s", err.Error())
		return err
	}

	for _, post := range posts {
		postGenerator := PostGenerator{
			GenConfig: GenConfig{
				Template:    postTemplate,
				Destination: directoryhelper.OutputDir(),
				Writer:      sg.Config.Writer,
			},
			Post: post,
		}

		if err := postGenerator.Generate(); err != nil {
			log.Printf("error when generating posts: %s", err.Error())
			return err
		}
	}

	return nil
}

func createOutputDir() error {
	outputDir := directoryhelper.OutputDir()
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory at %s: %v", outputDir, err)
	}

	return nil
}
