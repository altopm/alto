package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/Delta456/box-cli-maker"
	"github.com/altopm/alto/errors"
	"github.com/altopm/alto/logs"
	"github.com/altopm/alto/ui"
	aurora "github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var installCommand = &cobra.Command{
	Use:     "install",
	Short:   "Install a package",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		var pkgName string = args[0]
		initWheel := ui.Loader("%s Initializing...")
		initWheel.Start()
		logs.CreateLogFile()
		initWheel.Stop()
		ui.SuccessText("Created logfile & temporary directory")
		wheel := ui.Loader("%s Searching for package " + pkgName)
		wheel.Start()
		resp, err := http.Get("https://registry.altopkg.com/package/" + pkgName)
		if err != nil {
			errors.Handle(err.Error())
		}
		defer resp.Body.Close()
		if strings.Contains(resp.Status, "404") {
			wheel.Stop()
			errors.Handle("Package not found. Double check your spelling and/or that you have the correct registry installed!")
		}
		wheel.Stop()
		ui.SuccessText("Package found!")
		wheel1 := ui.Loader("%s Finding package install location")
		wheel1.Start()
		s, err := os.Stat("/var/alto/installs")
		if os.IsNotExist(err) {
			os.MkdirAll("/var/alto/installs", os.ModePerm)
			fmt.Print(s)
			wheel1.Stop()
			ui.SuccessText("Created install location!")
		} else {
			wheel1.Stop()
			ui.SuccessText("Found install location!")
		}
		wheel1.Stop()
		wheel2 := ui.Loader("%s Downloading package")
		wheel2.Start()
		// Get the data
		var url string = fmt.Sprintf("https://registry.altopkg.com/repo/src/%s/binaries/%s.zip", pkgName, pkgName)
		pkg, err := http.Get(url)
		if err != nil {
			errors.Handle("Could not download package!")
		}
		defer pkg.Body.Close()
		if pkg.StatusCode == 404 {
			wheel2.Stop()
			errors.Handle("Please report this bug, as well as the following information on GitHub: https://github.com/altopm/alto/issues/Loader")
			fmt.Println(aurora.Yellow(fmt.Sprintf("\n\t%s %s %s %s %s", pkg.Proto, pkg.Status, pkg.Header, pkg.Request.Method, pkg.Request.URL)))
		}
		out, err := os.Create("/var/alto/installs/" + pkgName)
		if err != nil {
			wheel2.Stop()
			errors.Handle("Permission denied. Please run as sudo!")
		}
		_, err = io.Copy(out, pkg.Body)
		if err != nil {
			wheel2.Stop()
			fmt.Print(err)
		}
		wheel2.Stop()
		ui.SuccessText("Package downloaded!")
		wheel3 := ui.Loader("%s Unpacking")
		wheel3.Start()
		unpkg := exec.Command("tar", "-xvf", "/var/alto/installs/"+pkgName)
		err = unpkg.Run()
		if err != nil {
			wheel3.Stop()
			errors.Handle("Could not unpack package!")
		}
		wheel3.Stop()
		ui.SuccessText("Package unpacked!")
		wheel4 := ui.Loader("%s Adding to PATH")
		wheel4.Start()
		var path string = os.Getenv("PATH")
		var appPath string = fmt.Sprintf("/var/alto/installs/bin/%s/", pkgName)
		err = os.Setenv("PATH", fmt.Sprintf("%s:%s", appPath, path))
		if err != nil {
			wheel4.Stop()
			errors.Handle("Could not add to PATH!")
			fmt.Println(err)
		}
		wheel4.Stop()
		ui.SuccessText("Package added to PATH!")
		wheel5 := ui.Loader("%s Cleaning up")
		wheel5.Start()
		os.Remove("logs.log")
		wheel5.Stop()
		ui.SuccessText("All done!")
		Box := box.New(box.Config{Px: 2, Py: 4, Type: "Double", Color: "Green"})
		Box.Print("Installed successfully!", "Thanks for using alto!")
	},
}

func init() {
	rootCmd.AddCommand(installCommand)
}
