package main

import (
	"github.com/tanqiuliu/greeting-grpc/cli"
	"github.com/tanqiuliu/greeting-grpc/server"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		target := os.Args[1]
		switch target {
		case "server":
			if err := server.Start(50001); err != nil {
				os.Exit(1)
			}
			return
		case "cli":
			if err := cli.Run(os.Args[2]); err != nil {
				os.Exit(1)
			}
		}
	}

}
