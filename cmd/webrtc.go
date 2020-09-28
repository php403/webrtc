package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type allowCmd struct {

}

var webRtcCmd = &cobra.Command{
	Use:                        "webrtc",
	Aliases:                    nil,
	SuggestFor:                 nil,
	Short:                      "1",
	Long:                       "2",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(args[0])
	},
}

func Execute() {
	if err := webRtcCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init()  {

}