package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/nevermosby/cds-gic-cli/cli"
	"github.com/nevermosby/cds-gic-cli/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configPath          = ".cds"
	configFile          = "config.json"
	DefaultOutputFormat = "json"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cds",
	Short: "CapitalOnline Cloud Command Line Tool",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Version: fmt.Sprintf("%s, build %s", cli.Version, cli.GitCommit),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cds/config.json)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configDir := filepath.Join(home, configPath)
		confPath := filepath.Join(home, configPath, configFile)
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			fmt.Println("start to create config file...")
			// create config dir
			err = os.MkdirAll(configDir, 0755)
			if err != nil {
				panic(err)
			}
			// init config file
			cfg := model.ConfigFile{
				AccessKey:    "",
				SecretKey:    "",
				OutputFormat: DefaultOutputFormat,
			}
			bytes, err := json.MarshalIndent(cfg, "", "\t")
			if err != nil {
				panic(err)
			}
			err = ioutil.WriteFile(confPath, bytes, 0600)
			if err != nil {
				panic(err)
			}
		}
		// Search config in home directory with name ".cds" (without extension).
		viper.SetConfigType("json")
		viper.SetConfigFile(confPath)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Fatalln("failed to find the config file")
		} else {
			fmt.Println("failed to read config file:", err)
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
