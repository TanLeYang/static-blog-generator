package models

// Post holds the data for a post
type Post struct {
	Title     string
	FileName  string
	Meta      PostMeta
	Content   []byte
	ImagesDir string
}

type PostMeta struct {
	Title string   `yaml:"title"`
	Short string   `yaml:"short"`
	Date  string   `yaml:"date"`
	Tags  []string `yaml:"tags"`
}
