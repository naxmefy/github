package main

import (
	"github.com/naxmefy/github/pkg/cli"
	"log"
	"os"
)

func main() {
	err := cli.App().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
