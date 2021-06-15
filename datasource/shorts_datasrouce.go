package datasource

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"

	"github.com/TanLeYang/go-ssg/models"
)

type ShortsDatasource interface {
	GetAllShorts() ([]models.Short, error)
}

// LocalShortsDatasource is a ShortsDatasource that fetches shorts stored locally.
//
// Shorts are not explicitly stored but are built from the meta information attached to
// each individual Post.
type LocalShortsDatasource struct {
	PostsDirectory string
}

func (ds LocalShortsDatasource) GetAllShorts() ([]models.Short, error) {
	files, err := ioutil.ReadDir(ds.PostsDirectory)
	if err != nil {
		return nil, err
	}

	var shorts []models.Short
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		short, err := getShort(ds.PostsDirectory, file)
		if err != nil {
			log.Println(err)
			continue
		}

		shorts = append(shorts, *short)
	}

	return shorts, nil
}

func getShort(directory string, postFolder fs.FileInfo) (*models.Short, error) {
	meta, err := getPostMeta(directory, postFolder)
	if err != nil {
		return nil, err
	}

	link := fmt.Sprintf("./%s.html", postFolder.Name())

	short := models.Short{
		PostMeta: *meta,
		Link:     link,
	}

	return &short, nil
}
