/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vince-0202/g-blog/admin/bootstrap"
	"go.uber.org/zap"
)

var (
	adminOptions bootstrap.AdminOptions
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := bootstrap.NewServer(adminOptions)
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
	adminOptions = bootstrap.AdminOptions{}
	runCmd.Flags().IntVar(&adminOptions.Port, "port", 1218, "Help message for toggle")
}
