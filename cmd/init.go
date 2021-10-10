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
var depsS = []*survey.Question{
	{
		Name:      "deps",
		Prompt:    &survey.Input{Message: "Any needed dependencies? (e.g. nodejs, python3, or express.js)"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
}

func initPackage(packageTitle string) {
	depsSAns := struct {
		Deps string `survey:"deps"`
	}{}
	err := survey.Ask(depsS, &depsSAns)
	deps := utils.SplitString(depsSAns.Deps, " ")
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