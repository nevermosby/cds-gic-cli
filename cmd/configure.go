package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"
	_ "github.com/nevermosby/cds-gic-cli/cli"
	"github.com/nevermosby/cds-gic-cli/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure your AccessKey and SecretKey locally - 配置本地AccessKey和SecretKey",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("configure called")
		cfg := model.ConfigFile{}
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("failed to load config:", err)
		}
		// fmt.Println("ak: ", cfg.AccessKey)
		// fmt.Println("sk: ", cfg.SecretKey)
		// fmt.Println("outputformat: ", cfg.OutputFormat)

		// cli.PrintfD("Access Key [%s]: ", MosaicString(cfg.AccessKey, 3))
		fmt.Printf("--- Access Key [%s] ---\n", MosaicString(cfg.AccessKey, 3))
		cfg.AccessKey = ReadInput(cfg.AccessKey)
		// cli.PrintfD("Secret Key [%s]: ", MosaicString(cfg.SecretKey, 3))
		fmt.Printf("--- Secret Key [%s] ---\n", MosaicString(cfg.SecretKey, 3))
		cfg.SecretKey = ReadInput(cfg.SecretKey)

		// save the config file
		err = SaveConfig(&cfg)
		if err != nil {
			log.Fatalln("unable to save conf file:", err)
		}

		PingCDS(&cfg)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func PingCDS(cfg *model.ConfigFile) {
	// load sdk to get datacenter
	// Init a credential with Access Key Id and Secret Access Key
	// You can apply them from the CDS web portal
	credential := common.NewCredential(
		cfg.AccessKey,
		cfg.SecretKey,
	)

	// init a client profile with method type
	cpf := profile.NewClientProfile()

	// Example: Get vdc
	cpf.HttpProfile.ReqMethod = "GET"
	vdcClient, _ := vdc.NewClient(credential, "", cpf)
	descVdcRequest := vdc.DescribeVdcRequest()
	descVdcResponse, err := vdcClient.DescribeVdc(descVdcRequest)
	if err != nil {
		fmt.Println("API request fail:", err)
	} else {
		if *descVdcResponse.Code != "Success" {
			fmt.Println("Reponse Code is not OK:", err)
		} else {
			fmt.Println("Your configuration is done.")
		}
	}
}
func MosaicString(s string, lastChars int) string {
	r := len(s) - lastChars
	if r > 0 {
		return strings.Repeat("*", r) + s[r:]
	} else {
		return strings.Repeat("*", len(s))
	}
}

func ReadInput(defaultValue string) string {
	var s string
	fmt.Scanf("%s\n", &s)
	if s == "" {
		return defaultValue
	}
	return s
}

func SaveConfig(config *model.ConfigFile) error {
	viper.Set("accesskey", config.AccessKey)
	viper.Set("secretkey", config.SecretKey)
	viper.Set("outputformat", config.OutputFormat)
	// viper.Set("",config)

	viper.WriteConfig()
	err := viper.WriteConfig()
	if err != nil {
		// fmt.Println("unable to write conf file:", err)
		return err
	}

	return nil
}
