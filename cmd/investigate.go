package cmd

import (
	"os"
	"strconv"

	"github.com/jwillker/sherlock/entity"
	"github.com/jwillker/sherlock/internal/packages"
	"github.com/jwillker/sherlock/internal/vanity"
	"github.com/jwillker/sherlock/pkg/gitlab"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// investigateCmd represents the investigate command
var investigateCmd = &cobra.Command{
	Use:   "investigate",
	Short: "Discovery Golang projects to index",
	Run: func(cmd *cobra.Command, args []string) {
		investigate()
	},
}

func init() {
	rootCmd.AddCommand(investigateCmd)
	investigateCmd.Flags().IntSliceVarP(&GroupIDs, "group-ids", "g", []int{}, "List of Gitlab group ids")
	investigateCmd.Flags().
		StringVarP(&BaseURL, "base-url", "u", "", "GitLab api url. Default https://gitlab.com/api/v4")
	investigateCmd.Flags().
		StringVarP(&Output, "output", "o", "packages.yaml", "Vanity packages config. Default: packages.yaml")
	investigateCmd.Flags().StringVarP(&Godoc, "godoc", "d", "", "Godoc URL")
	investigateCmd.Flags().StringVarP(&VanityURL, "vanity-url", "v", "", "Sally vanity URL")
	_ = investigateCmd.MarkFlagRequired("group-ids")
	_ = investigateCmd.MarkFlagRequired("godoc")
	_ = investigateCmd.MarkFlagRequired("vanity-url")
}

func investigate() {
	token := os.Getenv("GITLAB_TOKEN")
	err, g := gitlab.New(token, BaseURL)

	if err != nil {
		log.Error(err)
	}

	allPkgs := []entity.Package{}
	for _, group := range GroupIDs {
		log.Info("Searching projects in group: " + strconv.Itoa(group))
		pkgs, err := packages.List(g, group)
		if err != nil {
			log.Error(err)
			continue
		}

		allPkgs = append(allPkgs, pkgs...)

	}

	log.Info("Generating config in " + Output)
	err = vanity.GenSallyConf(Output, Godoc, VanityURL, allPkgs)

	if err != nil {
		log.Error(err)
	}
	log.Info("Done !")
}
