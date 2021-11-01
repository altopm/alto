package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/altopm/alto/errors"
	"github.com/altopm/alto/utils"
	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:     "init",
	Short:   "Initialize a new project",
	Long:    `Initialize a new project to be published to the alto registry, using "sudo alto publish -now"`,
	Example: "sudo alto init altopm",
	Run: func(cmd *cobra.Command, args []string) {
		initPackage(args[0])
	},
}
var qs = []*survey.Question{
	{
		Name:      "deps",
		Prompt:    &survey.Input{Message: "Any needed dependencies? (e.g. nodejs, python3, or express.js)"},
		Transform: survey.Title,
	},
	{
		Name: "furtherConfig",
		Prompt: &survey.Select{
			Message: "Any further configuration? (e.g. a shellfile to set the user's OS in a config file)",
			Options: []string{"no", "yes"},
			Default: "yes",
		},
	},
}

func initPackage(packageTitle string) {
	qsAns := struct {
		Deps          string `survey:"deps"`
		FurtherConfig string `survey:"furtherConfig"`
	}{}
	err := survey.Ask(qs, &qsAns)
	if err != nil {
		errors.Handle(err.Error())
		os.Exit(1)
	}
	deps := utils.SplitString(qsAns.Deps, " ")
	initWheel := utils.Loader("%s Initializing...")
	initWheel.Start()
	if err != nil {
		errors.Handle(err.Error())
	}
	for i := 0; i < 4; i++ {
		err := os.MkdirAll(fmt.Sprintf("/var/alto/tmp/%s", packageTitle), 0755)
		if err != nil {
			initWheel.Stop()
			errors.Handle(err.Error())
		}
	}
	for i := 0; i < 2; i++ {
		err := os.RemoveAll(fmt.Sprintf("/var/alto/tmp/%s", packageTitle))
		if err != nil {
			initWheel.Stop()
			errors.Handle(err.Error())
		}
	}
	err = os.MkdirAll(".alton", 0755)
	if err != nil {
		initWheel.Stop()
		errors.Handle(err.Error())
	}
	initfile, err := os.Create("./.alton/alto.json")
	if err != nil {
		initWheel.Stop()
		errors.Handle(err.Error())
	}
	defer initfile.Close()
	// this might actually be worse then regexp
	fmt.Fprintln(initfile, fmt.Sprintf("{\n\t\"title\": \"%s\",\n\t\"deps\": [\n\t\t\"%s\",\n\t\t\"%s\"\n\t]\n}", packageTitle, deps[0], deps[1]))
	initWheel.Stop()
	utils.MessageSuccess(fmt.Sprintf("%s initialized successfully", packageTitle))
	utils.MessageNeutral("We strongly suggest adding .alton directory to your .gitignore")
}
func init() {
	rootCmd.AddCommand(initCommand)
}
