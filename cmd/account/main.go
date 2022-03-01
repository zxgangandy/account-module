package main

import (
	"account-module/internal/app/service"
	"account-module/internal/server"
	"account-module/pkg/conf"
	"account-module/pkg/datasource"
	"account-module/pkg/logger"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	logger.Init()
	conf.Init()

	rootCmd := &cobra.Command{
		Use:   conf.Config.Application.Name,
		Short: "Account service",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			datasource.Init()
			service.Init()
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			server.Run()
		},
	}

	//config.SetFlags(rootCmd.Flags(), cfg)
	//rootCmd.Flags().AddGoFlagSet(goflag.CommandLine)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("err: %+v", err)
	}
}
