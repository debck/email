package main

import (
	"os"
	"fmt"
	"github.com/debck/email/cmd"
)

func main() {
	must(cmd.RootCmd.Execute())
}


func must(err error ) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
