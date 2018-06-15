package main

import (
	"fmt"
	"log"
	"os"

	"github.com/a-hilaly/svr/http2https"
	"github.com/a-hilaly/svr/httpproxy"
	"github.com/a-hilaly/svr/httpserver"
	"github.com/spf13/cobra"
)

var (
	Server      string
	Port        int
	Pattern     string
	RedirectURL string
)

func init() {
	cmd.Flags().IntVarP(&Port, "port", "p", 8080, "Server listening port")
	cmd.Flags().StringVarP(&Server, "server", "s", "default", "Server type")
	cmd.Flags().StringVarP(&Pattern, "pattern", "n", "/", "pattern url")
	cmd.Flags().StringVarP(&RedirectURL, "redirect", "r", "http://www.google.com", "pattern url")
}

var cmd = &cobra.Command{
	Use:   "svr",
	Short: "A collection of simple HTTP servers",
	Long: `A collection of simple HTTP servers
		- Simple HTTP server
		- Simple HTTP proxy server
		- HTTP to HTTPS proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		switch Server {
		case "default":
			log.Println("Simple server")
			httpserver.Default(Port, Pattern)
		case "proxy":
			log.Println("Proxy server")
			httpserver.New(Port, Pattern, httpproxy.Handler(RedirectURL))
		case "http2s":
			log.Println("Http to https")
			httpserver.New(Port, Pattern, http2https.Default())
		default:
			log.Println("Simple server")
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
