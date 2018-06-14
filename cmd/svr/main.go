package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a-hilaly/svr/httpserver"
	"github.com/spf13/cobra"
)

var (
	Server  string
	Port    int
	Pattern string
)

func init() {
	cmd.Flags().IntVarP(&Port, "port", "p", 8080, "Server listening port")
	cmd.Flags().StringVarP(&Server, "server", "s", "default", "Server type")
	cmd.Flags().StringVarP(&Pattern, "pattern", "r", "/", "pattern url")
}

var cmd = &cobra.Command{
	Use:   "svr",
	Short: "A collection of simple HTTP servers",
	Long: `A collection of simple HTTP servers
		- Simple HTTP server
		- Simple HTTP/HTTPS proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		switch Server {
		case "default":
			httpserver.Default(Port, Pattern)
		case "proxy":
			log.Panic("Not implemented")
		default:
			httpserver.Default(Port, Pattern)
		}
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
