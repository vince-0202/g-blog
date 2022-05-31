/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vince-0202/g-blog/client/bootstrap"
	"go.uber.org/zap"
)

var (
	clientOptions bootstrap.ClientOptions
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run g-blog client server",
	Long: `Use run common to run a client server.
	You should give a port which will the client server start,make sure this port is unused by other application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		s, err := bootstrap.NewServer(clientOptions)
		if err != nil {
			return err
		}

		s.InitSrver()

		zap.L().Info("start g-blog client", zap.Int("port", s.Port))
		if err := s.Server.Run(":" + strconv.Itoa(s.Port)); err != nil {
			zap.L().Error("startup service failed", zap.String("err", err.Error()))
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	clientOptions = bootstrap.ClientOptions{}
	runCmd.Flags().IntVar(&clientOptions.Port, "port", 1217, "Help message for toggle")
}
