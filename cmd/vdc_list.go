package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/common/profile"
	"github.com/capitalonline/cds-gic-sdk-go/vdc"
	"github.com/landoop/tableprinter"
	"github.com/nevermosby/cds-gic-cli/model"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// vdclistCmd represents the vdc list command
var vdclistCmd = &cobra.Command{
	Use:   "list",
	Short: "list all the virtual datacenters of your account - 显示所有的虚拟数据中心",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vdc list called")
		cfg := model.ConfigFile{}
		err := viper.Unmarshal(&cfg)
		if err != nil {
			log.Fatalln("failed to load config:", err)
		}
		ListVDC(&cfg)
	},
}

func init() {
	vdcCmd.AddCommand(vdclistCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ListVDC to list the virtual datacenters
func ListVDC(cfg *model.ConfigFile) {
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
			// fmt.Println(descVdcResponse.ToJsonString())
			vdcList := make([]model.Vdc, 0)
			var vdcTmp model.Vdc
			for _, vdc := range descVdcResponse.Data {
				vdcTmp = model.Vdc{
					VdcId:    *vdc.VdcId,
					VdcName:  *vdc.VdcName,
					RegionId: *vdc.RegionId,
				}
				vdcList = append(vdcList, vdcTmp)
			}

			tableprinter.Print(os.Stdout, vdcList)
		}
	}
}
