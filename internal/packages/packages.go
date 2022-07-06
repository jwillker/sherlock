package packages

import (
	"github.com/jwillker/sherlock/entity"
	"github.com/jwillker/sherlock/pkg/gitlab"
)

func List(g *gitlab.GitLab, groupID int) (empty []entity.Package, err error) {
	pkgs := []entity.Package{}
	listProjects, err := g.ListGroupProjects(groupID)
	if err != nil {
		return empty, err
	}

	for _, p := range listProjects {
		newPackage := entity.NewPackage(p)

		pkgs = append(pkgs, newPackage)
	}

	return pkgs, nil
}
