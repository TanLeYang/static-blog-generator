package datasource

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/TanLeYang/go-ssg/models"
	"github.com/TanLeYang/go-ssg/parser"
	"gopkg.in/yaml.v2"
)

const metaFileName = "meta.yml"

type PostsDatasource interface {
	GetAllPosts() ([]models.Post, error)
}

// LocalPostsDatasource is a PostsDatasource that fetches posts stored locally
type LocalPostsDatasource struct {
	Directory string // The directory that contains all the posts
}

func (ds LocalPostsDatasource) GetAllPosts() ([]models.Post, error) {
	files, err := ioutil.ReadDir(ds.Directory)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		post, err := getPost(ds.Directory, file)
		if err != nil {
			log.Println(err)
			continue
		}

		posts = append(posts, *post)
	}

	return posts, nil
}

func getPost(directory string, postFolder fs.FileInfo) (*models.Post, error) {
	meta, err := getPostMeta(directory, postFolder)
	if err != nil {
		return nil, err
	}

	content, err := getPostContent(directory, postFolder)
	if err != nil {
		return nil, err
	}

	imagesDir := getPostImageDir(directory, postFolder)

	post := models.Post{
		Title:     meta.Title,
		FileName:  postFolder.Name(),
		Meta:      *meta,
		Content:   content,
		ImagesDir: imagesDir,
	}

	return &post, nil
}

func getPostMeta(directory string, postFolder fs.FileInfo) (*models.PostMeta, error) {
	metaFilePath := filepath.Join(directory, postFolder.Name(), metaFileName)

	b, err := ioutil.ReadFile(metaFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading meta file %v, %s", metaFilePath, err.Error())
	}

	var postMeta models.PostMeta
	yaml.Unmarshal(b, &postMeta)

	return &postMeta, nil
}

func getPostContent(directory string, postFolder fs.FileInfo) ([]byte, error) {
	postFilePath := filepath.Join(directory, postFolder.Name(), postFolder.Name()) + ".md"

	b, err := ioutil.ReadFile(postFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading post file %v, %s", postFilePath, err.Error())
	}

	htmlBytes, err := parser.ParseMarkdownToHTML(b)
	if err != nil {
		return nil, err
	}

	return htmlBytes, nil
}

func getPostImageDir(directory string, postFolder fs.FileInfo) string {
	imagesPath := filepath.Join(directory, postFolder.Name(), fmt.Sprintf("%s_images", postFolder.Name()))
	if _, err := os.Stat(imagesPath); os.IsNotExist(err) {
		return ""
	}

	return imagesPath
}
