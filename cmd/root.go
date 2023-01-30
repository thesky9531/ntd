package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/thesky9531/network-traffic-detection/config"
	"github.com/thesky9531/network-traffic-detection/entity"
	"os"
	"strings"
)

const Name = "NetworkTrafficDetection"

func NetworkTrafficDetection(inputs string) (err error) {
	fmt.Println(inputs)
	//var runFlag = pflag.NewFlagSet(Name,pflag.ExitOnError)
	//runFlag.AddGoFlag(flag.CommandLine.Lookup("once"))
	//runFlag.AddGoFlag(flag.CommandLine.Lookup("modules"))
	//settings := form.Settings{
	//	Name:                  Name,
	//	HasDashboards:         true,
	//	RunFlags:              runFlag,
	//}
	//command := GenRootCmdWithSetting()

	//配置文件 init
	cf := config.ConfigInit()

	//flag 覆盖配置文件

	//开始流量包采集
	entity.FlowCollection(cf)

	return
}

func init() {
	// backwards compatibility workaround, convert -flags to --flags:
	for i, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) > 2 {
			os.Args[1+i] = "-" + arg
		}
	}
}

// BeatsRootCmd handles all application command line interface, parses user
// flags and runs subcommands
type RootCmd struct {
	cobra.Command
	RunCmd        *cobra.Command
	SetupCmd      *cobra.Command
	VersionCmd    *cobra.Command
	CompletionCmd *cobra.Command
	ExportCmd     *cobra.Command
	TestCmd       *cobra.Command
	KeystoreCmd   *cobra.Command
}
