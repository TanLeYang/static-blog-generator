package generator

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/TanLeYang/go-ssg/helper"
	"github.com/TanLeYang/go-ssg/models"
)

// PostVars holds the variables used by the post template
type PostVars struct {
	Title       string
	PublishDate string
	Content     template.HTML
}

type PostGenerator struct {
	GenConfig
	Post models.Post
}

func (g PostGenerator) Generate() error {
	post := g.Post
	templateToUse := g.Template
	buffer := bytes.Buffer{}

	// Generate template
	postVar := g.createPostVar()
	if err := templateToUse.Execute(&buffer, postVar); err != nil {
		return fmt.Errorf("error: failed to execute template for post with title: %s, %s", postVar.Title, err.Error())
	}

	// Copy images
	imagesDest := filepath.Join(g.Destination, fmt.Sprintf("%s_images", post.FileName))
	if err := copyImages(post.ImagesDir, imagesDest); err != nil {
		return err
	}

	html := template.HTML(buffer.String())
	fullpath := filepath.Join(g.Destination, fmt.Sprintf("%s.html", post.FileName))
	if err := g.Writer.InsertIntoBase(html, postVar.Title, fullpath); err != nil {
		return err
	}

	return nil
}

func (g PostGenerator) createPostVar() PostVars {
	return PostVars{
		Title:       g.Post.Title,
		PublishDate: g.Post.Meta.Date,
		Content:     template.HTML(g.Post.Content),
	}
}

func copyImages(source, destination string) error {
	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		return fmt.Errorf("error creating images directory at %s: %v", destination, err)
	}

	images, err := ioutil.ReadDir(source)
	if err != nil {
		return fmt.Errorf("error reading images directory at %s: %v", source, err)
	}

	for _, file := range images {
		src := filepath.Join(source, file.Name())
		dest := filepath.Join(destination, file.Name())
		if err := helper.CopyFile(src, dest); err != nil {
			return fmt.Errorf("error copying image from %s to %s: %v", src, dest, err)
		}
	}

	return nil
}
