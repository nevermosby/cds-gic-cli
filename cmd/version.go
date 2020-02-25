package cmd

import (
	"fmt"

	"github.com/nevermosby/cds-gic-cli/cli"
	"github.com/spf13/cobra"
)

//
// This variable is replaced in compile time
// `-ldflags "-X 'github.com/.../cli.Version=${VERSION}'"`
var (
	Version = "0.0.1"
)

// func GetVersion() string {
// 	return strings.Replace(Version, " ", "-", -1)
// }

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print current version - 打印当前版本号",
	Long: `A longer description for version For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(Version)
		fmt.Println(cli.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
