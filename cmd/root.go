package cmd

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/uaa-cli/config"
	"code.cloudfoundry.org/uaa-cli/help"
	"code.cloudfoundry.org/uaa-cli/uaa"
	"github.com/spf13/cobra"
)

var cfgFile uaa.Config

// User flags
var (
	password string
	username string
)

// Global flags
var (
	skipSSLValidation bool
	trace             bool
)

// Client flags
var (
	clientSecret         string
	authorizedGrantTypes string
	authorities          string
	accessTokenValidity  int32
	refreshTokenValidity int32
	displayName          string
	scope                string
	redirectUri          string
	clone                string
	zoneSubdomain        string
	port                 int
)

var RootCmd = cobra.Command{
	Use:   "uaa",
	Short: "A cli for interacting with UAAs",
	Long:  help.Root(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().BoolVarP(&trace, "trace", "t", false, "See additional info on HTTP requests")
}

func initConfig() {
	// Startup tasks
}

func GetSavedConfig() uaa.Config {
	cfgFile = config.ReadConfig()
	cfgFile.Trace = trace
	cfgFile.ZoneSubdomain = zoneSubdomain
	return cfgFile
}
