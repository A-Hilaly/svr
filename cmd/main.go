package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/a-hilaly/svr/pkg/http2https"
	"github.com/a-hilaly/svr/pkg/httpredirect"
	"github.com/a-hilaly/svr/pkg/httpserver"
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
		- Simple redirection server
		- HTTP to HTTPS proxy server`,
	Run: func(cmd *cobra.Command, args []string) {
		switch Server {
		case "default":
			log.Println("Simple server")
			httpserver.Default(Port, Pattern)
		case "redirect":
			log.Println("Proxy server")
			httpserver.New(Port, Pattern, httpredirect.Handler(RedirectURL))
		case "http2https":
			log.Println("Http to https")
			httpserver.New(Port, Pattern, http2https.Handler())
		default:
			log.Fatalf("Unknown server name %v\n", Server)
		}
	},
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
