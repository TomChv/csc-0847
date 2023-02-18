package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/TomChv/csc-0847/project_1/backend/db"
	"github.com/TomChv/csc-0847/project_1/backend/server"
)

var (
	port string
	host string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run project 1 backend server",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := db.New()
		if err != nil {
			log.Fatalln(err)
		}

		s := server.New(c)

		// Format URL
		url := fmt.Sprintf("%s:%s", host, port)

		// Retrieve mode to update server's behavior.
		mode := os.Getenv("GIN_MODE")
		if mode == "release" {
			fmt.Printf("Launch server in production mode on %s.\n", url)
		}

		if err := s.Run(url); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	serverCmd.Flags().StringVar(&port, "port", "9000", "Set port to listen on.")
	serverCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Set host to listen on.")
}
