package cmd

import (
	"os"
	"strconv"

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
	investigateCmd.Flags().IntVarP(&GroupID, "group-id", "g", 0, "GitLab group id")
	investigateCmd.Flags().StringVarP(&BaseURL, "base-url", "u", "", "GitLab api url. Default https://gitlab.com/api/v4")
	investigateCmd.Flags().StringVarP(&Output, "output", "o", "packages.yaml", "Vanity packages config. Default: packages.yaml")
	investigateCmd.Flags().StringVarP(&Godoc, "godoc", "d", "", "Godoc URL")
	investigateCmd.Flags().StringVarP(&VanityURL, "vanity-url", "v", "", "Sally vanity URL")
	_ = investigateCmd.MarkFlagRequired("group-id")
	_ = investigateCmd.MarkFlagRequired("godoc")
	_ = investigateCmd.MarkFlagRequired("vanity-url")
}

func investigate() {
	token := os.Getenv("GITLAB_TOKEN")
	err, g := gitlab.New(token, BaseURL)

	if err != nil {
		log.Error(err)
	}

	log.Info("Searching projects in group: " + strconv.Itoa(GroupID))
	pkgs, err := packages.List(g, GroupID)

	if err != nil {
		log.Error(err)
	}

	log.Info("Generating config in " + Output)
	err = vanity.GenSallyConf(Output, Godoc, VanityURL, pkgs)

	if err != nil {
		log.Error(err)
	}

	log.Info("Done !")
}
