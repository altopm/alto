package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Version struct {
	Version    string `json:"version"`
	httpStatus string `json:"httpStatus"`
}

var client = &http.Client{}
var versionCommand = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Alto and its server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Alto CLI v0.1 -- HEAD")
		v := &Version{}
		getJson("https://registry.altopkg.com/Status", v)
		fmt.Printf("Registry: %s\n", v.Version)
	},
}

func getJson(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
func init() {
	rootCmd.AddCommand(versionCommand)
}
