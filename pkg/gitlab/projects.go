package gitlab

import (
	gitlab "github.com/xanzy/go-gitlab"
)

// ListGroupProjects get a list of group projects
func (g *GitLab) ListGroupProjects(groupID int) ([]*gitlab.Project, error) {
	var projects []*gitlab.Project

	opt := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 100,
		},
		Archived:         gitlab.Bool(false),
		IncludeSubGroups: gitlab.Bool(false),
		Simple:           gitlab.Bool(true),
	}

	for {
		// Get the first page with projects.
		ps, resp, err := g.Client.Groups.ListGroupProjects(groupID, opt)
		if err != nil {
			return []*gitlab.Project{}, err
		}

		for _, v := range ps {
			includeProject := false
			langs, _, e := g.Client.Projects.GetProjectLanguages(v.ID, nil)
			if e != nil {
				continue
			}

			for lang, value := range *langs {
				if lang == "Go" && value > 10 {
					includeProject = true
				}
			}

			if includeProject {
				projects = append(projects, v)
			}
		}

		// Exit the loop when we've seen all pages.
		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		// Update page
		opt.Page = resp.NextPage
	}

	return projects, nil
}
