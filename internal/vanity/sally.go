package vanity

import (
	"io/ioutil"

	"github.com/jwillker/sherlock/entity"
	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	URL      string         `yaml:"url"`
	Packages map[string]Pkg `yaml:"packages"`
	Godoc    struct {
		Host string `yaml:"host"`
	} `yaml:"godoc"`
}

type Pkg struct {
	Repo string `yaml:"repo"`
}

// GenSallyConf generate Sally packages config format
func GenSallyConf(output, godoc, url string, pkgs []entity.Package) error {
	packages := map[string]Pkg{}

	for _, p := range pkgs {
		packages[p.Name] = Pkg{
			Repo: p.Repo,
		}
	}

	config := Config{
		URL:      url,
		Packages: packages,
		Godoc: struct {
			Host string "yaml:\"host\""
		}{
			Host: godoc,
		},
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(output, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
