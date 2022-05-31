/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vince-0202/g-blog/client/bootstrap"
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

		s := bootstrap.NewServer(clientOptions)

		s.InitSrver()

		if err := s.Server.Run(":" + strconv.Itoa(s.Port)); err != nil {
			// fmt.Println("startup service failed, err:%v\n", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	clientOptions = bootstrap.ClientOptions{}
	runCmd.Flags().IntVar(&clientOptions.Port, "port", 1217, "Help message for toggle")
}
