package main

import (
	"log"

	"github.com/a23667788/m800-assignment/internal/service"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "m800",
		Short: "M800-assignment",
		Long:  "M800-assignment is a web server using gin framework, which can communicate between your Line offical account and Line clients via line messaging API.",
	}

	var port int
	var filePath string

	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start web service",
		Run: func(cmd *cobra.Command, args []string) {
			setting := InitializeSetting(filePath)
			server := service.NewM800Service(setting, port)
			if server != nil {
				server.StartWebServer()
			}
			defer server.Mongo.Close()
		},
	}

	startCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to listen on")
	startCmd.Flags().StringVarP(&filePath, "config", "f", "config.yaml", "Configuration file path")

	rootCmd.AddCommand(startCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
