package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/altopm/alto/ui"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var installCommand = &cobra.Command{
	Use:   "install",
	Short: "Install a package",
	Run: func(cmd *cobra.Command, args []string) {
		var pkgName string = args[0]
		wheel := ui.New("%s Searching for package " + pkgName)
		wheel.Start()
		resp, err := http.Get("https://registry-production.up.railway.app/package/" + pkgName)
		if err != nil {
			fmt.Print(err)
		}
		defer resp.Body.Close()
		if strings.Contains(resp.Status, "404") {
			wheel.Stop()
			fmt.Println(Red("\tError!"), "Package not found. Double check your spelling and/or that you have the correct registry installed!")
			os.Exit(1)
		}
		wheel.Stop()
		fmt.Println(Green("\tSuccess!"), "Package found!")
		wheel1 := ui.New("%s Finding package install location")
		wheel1.Start()
		s, err := os.Stat("/var/alto/installs")
		if os.IsNotExist(err) {
			os.MkdirAll("/var/alto/installs", os.ModePerm)
			fmt.Print(s)
			wheel1.Stop()
			fmt.Println(Green("\tSuccess!"), "Created install location!")
		} else {
			wheel1.Stop()
			fmt.Println(Green("\tSuccess!"), "Found install location!")
		}
		wheel1.Stop()
		wheel2 := ui.New("%s Downloading package")
		wheel2.Start()
		// Get the data
		var url string = fmt.Sprintf("https://registry.altopkg.com/repo/src/%s/binary/%s/", pkgName, pkgName)
		pkg, err := http.Get(url)
		if err != nil {
			fmt.Println(Red("\tError!"), "Could not download package!")
		}
		defer pkg.Body.Close()
		if pkg.StatusCode == 404 {
			fmt.Println(Red("\tError!"), "Please report this bug, as well as the following information on GitHub: https://github.com/altopm/alto/issues/new")
			fmt.Println(Yellow(fmt.Sprintf("\n\t%s %s %s %s %s", pkg.Proto, pkg.Status, pkg.Header, pkg.Request.Method, pkg.Request.URL)))
		}
		out, err := os.Create("/var/alto/installs/" + pkgName)
		if err != nil {
			wheel2.Stop()
			fmt.Println(Red("\tError!"), "Permission denied. Please run as sudo!")
			os.Exit(1)
		}
		_, err = io.Copy(out, pkg.Body)
		if err != nil {
			fmt.Print(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCommand)
}
