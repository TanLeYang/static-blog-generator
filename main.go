package main

import (
	"log"
	"path/filepath"

	"github.com/TanLeYang/go-ssg/config"
	"github.com/TanLeYang/go-ssg/datasource"
	"github.com/TanLeYang/go-ssg/directoryhelper"
	"github.com/TanLeYang/go-ssg/generator"
	"github.com/TanLeYang/go-ssg/templatestore"
)

func main() {
	config := getBaseConfig()
	postsDS := getPostsDS()
	shortsDS := getShortsDS()
	siteGenerator := generator.SiteGenerator{
		Config:           config,
		PostsDatasource:  postsDS,
		ShortsDatasource: shortsDS,
	}

	err := siteGenerator.Generate()
	if err != nil {
		log.Fatalf("error when generating site: %s", err.Error())
	}
}

func getBaseConfig() generator.GenConfig {
	template, err := templatestore.BaseTemplate()
	if err != nil {
		log.Fatalf("failed to get base template, %s", err.Error())
	}

	destination := filepath.Join(directoryhelper.RootDir(), "output")

	siteConfig, err := config.SiteConfig()
	if err != nil {
		log.Fatalf("failed to get site config, %s", err.Error())
	}

	writer := generator.BaseTemplateWriter{
		SiteConfig:   *siteConfig,
		BaseTemplate: template,
	}

	return generator.GenConfig{
		Template:    template,
		Destination: destination,
		Writer:      &writer,
	}
}

func getPostsDS() datasource.PostsDatasource {
	directory := filepath.Join(directoryhelper.RootDir(), "data", "posts")
	return datasource.LocalPostsDatasource{
		Directory: directory,
	}
}

func getShortsDS() datasource.LocalShortsDatasource {
	postsDirectory := filepath.Join(directoryhelper.RootDir(), "data", "posts")
	return datasource.LocalShortsDatasource{
		PostsDirectory: postsDirectory,
	}
}
