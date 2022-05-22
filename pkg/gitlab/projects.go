package gitlab

import gitlab "github.com/xanzy/go-gitlab"

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
		// Append projects
		projects = append(projects, ps...)

		// Exit the loop when we've seen all pages.
		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		// Update page
		opt.Page = resp.NextPage
	}

	return projects, nil
}
