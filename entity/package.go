package entity

import (
	"strings"

	gitlab "github.com/xanzy/go-gitlab"
)

type Package struct {
	Name string
	Repo string
}

func NewPackage(gitlabProject *gitlab.Project) Package {
	repo := strings.TrimPrefix(gitlabProject.HTTPURLToRepo, "https://")

	return Package{
		Name: gitlabProject.Path,
		Repo: repo,
	}
}
