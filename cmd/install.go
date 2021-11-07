package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/altopm/alto/errors"
	"github.com/altopm/alto/logs"
	"github.com/altopm/alto/utils"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

var Verbose bool
var Hackery bool
var installCommand = &cobra.Command{
	Use:     "install",
	Short:   "Install a package",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		if Verbose {
			fmt.Println("Verbose output enabled")
			installVerbose(args)
		}
		if Hackery {
			fmt.Println("HACKER-Y output enabled")
			installHackerModeEnabled(args)
		}
		if len(args) == 0 {
			errors.Handle("Please specify a package to install!")
		} else if (Verbose == false) && (Hackery == false) && (len(args) == 1) {
			installNormal(args)
		}
	},
}

func installNormal(args []string) {
	var pkgName string = args[0]
	initWheel := utils.Loader("%s Initializing...")
	initWheel.Start()
	logs.CreateLogFile()
	initWheel.Stop()
	utils.MessageSuccess("Created logfile & temporary directory")
	wheel := utils.Loader("%s Searching for package " + pkgName)
	wheel.Start()
	logs.AppendLog(fmt.Sprintf("Pinging registry... (url: %s) GET", pkgName))
	resp, err := http.Get("https://registry.altopkg.com/package/" + pkgName)
	if err != nil {
		logs.AppendLog(fmt.Sprintf("Error: %s || StatusCode: %d", err, resp.StatusCode))
		errors.Handle(err.Error())
	}
	defer resp.Body.Close()
	if strings.Contains(resp.Status, "404") {
		logs.AppendLog(fmt.Sprintf("Error: Package not found"))
		wheel.Stop()
		errors.Handle("Package not found. Double check your spelling and/or that you have the correct registry installed!")
	}
	logs.AppendLog(fmt.Sprintf("StatusCode: %d", resp.StatusCode))
	wheel.Stop()
	utils.MessageSuccess("Package found!")
	wheel1 := utils.Loader("%s Finding package install location")
	wheel1.Start()
	logs.AppendLog(fmt.Sprintf("Finding package install location. Planned location: /var/alto/installs/%s", pkgName))
	s, err := os.Stat("/var/alto/installs")
	if os.IsNotExist(err) {
		os.MkdirAll("/var/alto/installs", os.ModePerm)
		fmt.Print(s)
		wheel1.Stop()
		logs.AppendLog("Created install directory")
		utils.MessageSuccess("Created install location!")
	} else {
		wheel1.Stop()
		logs.AppendLog("Install directory already exists")
		utils.MessageSuccess("Found install location!")
	}
	wheel1.Stop()
	wheel2 := utils.Loader("%s Downloading package")
	wheel2.Start()
	// Get the data
	var url string = fmt.Sprintf("https://flags.altopkg.com/%s/%s.tar.gz", pkgName, pkgName)
	logs.AppendLog(fmt.Sprintf("Downloading package from %s", url))
	pkg, err := http.Get(url)
	if err != nil {
		logs.AppendLog(fmt.Sprintf("Error: %s || StatusCode: %d", err, resp.StatusCode))
		errors.Handle("Could not download package!")
	}
	defer pkg.Body.Close()
	if pkg.StatusCode == 404 {
		logs.AppendLog(fmt.Sprintf("There was an unknown error. Details follow:\n%s %s %s %s %s\n\n%s", pkg.Proto, pkg.Status, pkg.Header, pkg.Request.Method, pkg.Request.URL, err.Error()))
		wheel2.Stop()
		errors.Handle("Please report this bug, as well as the following information on GitHub: https://github.com/altopm/alto/issues/Loader")
		utils.MessageWarning(fmt.Sprintf("\n\t%s %s %s %s %s", pkg.Proto, pkg.Status, pkg.Header, pkg.Request.Method, pkg.Request.URL))
	}
	logs.AppendLog(fmt.Sprintf("StatusCode: %d", pkg.StatusCode))
	out, err := os.Create("/var/alto/installs/" + pkgName)
	if err != nil {
		logs.AppendLog("User forgot to run as sudo")
		wheel2.Stop()
		errors.Handle("Permission denied. Please run as sudo!")
	}
	_, err = io.Copy(out, pkg.Body)
	if err != nil {
		wheel2.Stop()
		fmt.Print(err)
	}
	wheel2.Stop()
	utils.MessageSuccess("Package downloaded!")
	wheel3 := utils.Loader("%s Unpacking")
	wheel3.Start()
	logs.AppendLog(fmt.Sprintf("Unpacking package"))
	unpkg := exec.Command("tar", "-xvf", "/var/alto/installs/"+pkgName)
	err = unpkg.Run()
	if err != nil {
		logs.AppendLog("An unknown error occurred: " + err.Error())
		wheel3.Stop()
		errors.Handle("Could not unpack package!")
	}
	wheel3.Stop()
	utils.MessageSuccess("Package unpacked!")
	logs.AppendLog("Adding package to PATH envvar")
	wheel4 := utils.Loader("%s Adding to PATH")
	wheel4.Start()
	var path string = os.Getenv("PATH")
	var appPath string = fmt.Sprintf("/var/alto/installs/bin/%s/", pkgName)
	err = os.Setenv("PATH", fmt.Sprintf("%s:%s", appPath, path))
	if err != nil {
		logs.AppendLog("An unknown error occurred: " + err.Error())
		wheel4.Stop()
		errors.Handle("Could not add to PATH!")
		fmt.Println(err)
	}
	wheel4.Stop()
	utils.MessageSuccess("Package added to PATH!")
	wheel5 := utils.Loader("%s Cleaning up")
	wheel5.Start()
	logs.AppendLog("Removing this logfile")
	err = os.Remove("./logs.log")
	if err != nil {
		logs.AppendLog("An unknown error occurred: " + err.Error())
		wheel5.Stop()
		errors.Handle(err.Error())
	}
	err = os.RemoveAll("./bin")
	if err != nil {
		logs.AppendLog("An unknown error occurred: " + err.Error())
		wheel5.Stop()
		errors.Handle(err.Error())
	}
	wheel5.Stop()
	utils.MessageSuccess("All done!")
	Box := box.New(box.Config{Px: 12, Py: 2, Type: "Double", Color: "Green"})
	Box.Print("Installed successfully!", "Thanks for using alto!")
}
func installHackerModeEnabled(args []string) {
	var pkgName string = args[0]
	for i := 0; i < 100; i++ {
		fmt.Println(aurora.Green(gofakeit.HackerPhrase()))
	}
	for i := 0; i < 100; i++ {
		fmt.Println(aurora.Green(gofakeit.NewCrypto()))
	}
	logs.CreateLogFile()
	utils.MessageSuccess("Created logfile & temporary directory")
	resp, err := http.Get("https://registry.altopkg.com/package/" + pkgName)
	for i := 0; i < 100000; i++ {
		fmt.Print(gofakeit.BitcoinAddress())
	}
	if err != nil {
		errors.Handle(err.Error())
	}
	defer resp.Body.Close()
	if strings.Contains(resp.Status, "404") {
		errors.Handle("Package not found. Double check your spelling and/or that you have the correct registry installed!")
	}
	fmt.Print("\n")
	utils.MessageSuccess("Generated heck adress!")
	s, err := os.Stat("/var/alto/installs")
	if os.IsNotExist(err) {
		os.MkdirAll("/var/alto/installs", os.ModePerm)
		fmt.Print(s)
		utils.MessageSuccess("Created install location!")
	} else {
		utils.MessageSuccess("Posted heck address to hecking API!")
	}
	// Get the data
	var url string = fmt.Sprintf("https://registry.altopkg.com/repo/src/%s/binaries/%s.zip", pkgName, pkgName)
	pkg, err := http.Get(url)
	if err != nil {
		errors.Handle("Could not download package!")
	}
	defer pkg.Body.Close()
	if pkg.StatusCode == 404 {
		errors.Handle("Please report this bug, as well as the following information on GitHub: https://github.com/altopm/alto/issues/Loader")
		utils.MessageWarning(fmt.Sprintf("\n\t%s %s %s %s %s", pkg.Proto, pkg.Status, pkg.Header, pkg.Request.Method, pkg.Request.URL))
	}
	out, err := os.Create("/var/alto/installs/" + pkgName)
	if err != nil {
		errors.Handle("Permission denied. Please run as sudo!")
	}
	_, err = io.Copy(out, pkg.Body)
	if err != nil {
		fmt.Print(err)
	}
	utils.MessageSuccess("Hecking software installed!")
	unpkg := exec.Command("tar", "-xvf", "/var/alto/installs/"+pkgName)
	err = unpkg.Run()
	if err != nil {
		errors.Handle("Could not unpack package!")
	}
	utils.MessageSuccess("Hecker unpacked!")
	var path string = os.Getenv("PATH")
	var appPath string = fmt.Sprintf("/var/alto/installs/bin/%s/", pkgName)
	err = os.Setenv("PATH", fmt.Sprintf("%s:%s", appPath, path))
	if err != nil {
		errors.Handle("Could not add to PATH!")
		fmt.Println(err)
	}
	utils.MessageSuccess("Hecker is ready to use!")
	err = os.Remove("./logs.log")
	if err != nil {
		errors.Handle(err.Error())
	}
	err = os.RemoveAll("./bin")
	if err != nil {
		errors.Handle(err.Error())
	}
	utils.MessageSuccess("HECKING READY!")
}
func installVerbose(args []string) {

}
func init() {
	installCommand.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Enable verbose output")
	installCommand.Flags().BoolVarP(&Hackery, "hacker-y", "a", false, "Why be a normy pleb when you can be a HACKKKEERRRR.>?>@>?!")
	rootCmd.AddCommand(installCommand)
}
