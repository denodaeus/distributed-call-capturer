package main

import "fmt"

import "github.com/spf13/cobra"

func main() {

	var serverCommand = &cobra.Command{
		Use:   "server",
		Short: "server mode",
		Long: `Execute distributed-call-capturer in server mode, accepting HTTP 
    requests and publishing commands to client agents for call capture
    `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("initializing in server mode")
			s := Server{}
			s.init()
		},
	}

	var clientCommand = &cobra.Command{
		Use:   "client",
		Short: "client mode",
		Long: `Execute distributed-call-capturer in client mode, listening for
    server rpc commands and capturing pcap for SIP/RTP
    `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("initializing in client mode")
			c := Client{}
			c.init()
		},
	}

	var rootCmd = &cobra.Command{Use: "distributed-call-capturer"}
	rootCmd.AddCommand(serverCommand, clientCommand)
	rootCmd.Execute()
}
