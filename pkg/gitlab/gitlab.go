package gitlab

import (
	"fmt"

	gitlab "github.com/xanzy/go-gitlab"
)

const baseURL = "https://gitlab.com/api/v4"

type GitLab struct {
	Client  *gitlab.Client
	Token   string
	BaseURL string
}

func New(token, bURL string) (error, *GitLab) {
	g := &GitLab{}
	if token == "" {
		return fmt.Errorf("Missing GitLab token"), g
	}

	if bURL == "" {
		bURL = baseURL
	}

	g.BaseURL = bURL
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(bURL))

	if err != nil {
		return err, g
	}
	g.Client = client

	return nil, g
}
