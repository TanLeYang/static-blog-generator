package config

import (
	"io/ioutil"
	"path/filepath"

	"github.com/TanLeYang/go-ssg/directoryhelper"
	"github.com/TanLeYang/go-ssg/models"
	"gopkg.in/yaml.v3"
)

const configFileName = "config.yml"

func SiteConfig() (*models.SiteConfig, error) {
	fullPath := filepath.Join(directoryhelper.DataDir(), configFileName)

	b, err := ioutil.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	var siteConfig models.SiteConfig
	if err := yaml.Unmarshal(b, &siteConfig); err != nil {
		return nil, err
	}

	return &siteConfig, nil
}
