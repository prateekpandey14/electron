package cmd

import (
	"github.com/lileio/lile"
	"github.com/prateekpandey14/electron/electron/subscribers"
	"github.com/spf13/cobra"
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Subscribe to and process queue messages",
	Run: func(cmd *cobra.Command, args []string) {
		lile.Subscribe(&subscribers.ElectronServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(subscribeCmd)
}
