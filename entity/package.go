package entity

import gitlab "github.com/xanzy/go-gitlab"

type Package struct {
	Name string
	Repo string
}

func NewPackage(gitlabProject *gitlab.Project) Package {
	return Package{
		Name: gitlabProject.Path,
		Repo: gitlabProject.HTTPURLToRepo,
	}
}
