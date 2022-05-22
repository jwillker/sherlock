package vanity

import (
	"io/ioutil"

	"github.com/jwillker/sherlock/entity"
	yaml "gopkg.in/yaml.v3"
)

type Pkg struct {
	Repo string `yaml:"repo"`
}

// GenSallyConf generate Sally packages config format
func GenSallyConf(output string, pkgs *[]entity.Package) error {

	packages := map[string]Pkg{}

	for _, p := range *pkgs {
		packages[p.Name] = Pkg{
			Repo: p.Repo,
		}
	}

	data, err := yaml.Marshal(&packages)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(output, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
