package main

import "log"
import "github.com/spf13/cobra"

import "github.com/Vocalocity/distributed-call-capturer/client"
import "github.com/Vocalocity/distributed-call-capturer/server"

func main() {

	var serverCommand = &cobra.Command{
		Use:   "server",
		Short: "server mode",
		Long: `Execute distributed-call-capturer in server mode, accepting HTTP 
    requests and publishing commands to client agents for call capture
    `,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("initializing in server mode")
			s := server.Server{}
			s.Init()
		},
	}

	var clientCommand = &cobra.Command{
		Use:   "client",
		Short: "client mode",
		Long: `Execute distributed-call-capturer in client mode, listening for
    server rpc commands and capturing pcap for SIP/RTP
    `,
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("initializing in client mode")
			c := client.Client{}
			c.Init()
		},
	}

	var rootCmd = &cobra.Command{Use: "distributed-call-capturer"}
	rootCmd.AddCommand(serverCommand, clientCommand)
	rootCmd.Execute()
}
