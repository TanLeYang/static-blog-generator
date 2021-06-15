package models

type SiteConfig struct {
	BlogTitle  string `yaml:"blog_title"`
	BlogAuthor string `yaml:"blog_author"`
	BaseURL    string `yaml:"base_url"`
}
