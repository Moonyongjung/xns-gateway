package cmd

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/Moonyongjung/xns-gateway/cmd/client"
	"github.com/Moonyongjung/xns-gateway/cmd/param"
	"github.com/Moonyongjung/xns-gateway/gw"
	"github.com/Moonyongjung/xns-gateway/types"
	"github.com/Moonyongjung/xns-gateway/util"

	"github.com/spf13/cobra"
)

const serverStartTime = 3 * time.Second

func StartCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "start",
		Aliases: []string{"s"},
		Short:   "Start XNS gateway",
		Args:    param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s start
$ %s s 
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			xnsContext, err = checkRunInit(xnsContext)
			if err != nil {
				return err
			}

			xplac, err := client.NewXplaClient(xnsContext, true)
			if err != nil {
				return err
			}
			xnsContext.WithXplaClient(xplac)

			go func() {
				if err := gw.StartHttpServer(xnsContext); err != nil {
					xnsContext.GetChannels().ErrChan <- err
				}
			}()

			select {
			case err := <-xnsContext.GetChannels().ErrChan:
				return err
			case <-time.After(serverStartTime):
			}

			return gw.GatewayQuit()
		},
	}

	return cmd
}

func XnsContractsInfoCmd(xnsContext *types.XnsContext) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xns-contracts-info",
		Short: "Get XNS contracts info that the gateway has",
		Args:  param.WithUsage(cobra.NoArgs),
		Example: strings.TrimSpace(fmt.Sprintf(`
$ %s query xns-contracts-info
$ %s q xns-contracts-info
		`, types.DefaultAppName, types.DefaultAppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			contracts := xnsContext.GetApp().Contracts

			jsonBytes, err := json.Marshal(contracts)
			if err != nil {
				return err
			}

			util.LogInfo(string(jsonBytes))

			return nil
		},
	}

	return cmd
}
