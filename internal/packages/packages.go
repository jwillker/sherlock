package packages

import (
	"github.com/jwillker/sherlock/entity"
	"github.com/jwillker/sherlock/pkg/gitlab"
)

func List(g *gitlab.GitLab, groupID int) (*[]entity.Package, error) {
	pkgs := &[]entity.Package{}
	listProjects, err := g.ListGroupProjects(groupID)

	if err != nil {
		return pkgs, nil
	}

	for _, p := range listProjects {
		newPackage := entity.NewPackage(p)

		*pkgs = append(*pkgs, newPackage)
	}

	return pkgs, nil
}
