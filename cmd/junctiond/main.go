package main

import (
	"fmt"
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/airchains-network/junction/app"
	"github.com/airchains-network/junction/cmd/junctiond/cmd"
)

var (
	Version   = "v0.1.0"
	Commit    = "aab9725"
	BuildDate = "2024-06-26T11:44:10Z"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		fmt.Printf("Version: %s\nCommit: %s\nBuild Date: %s\n", Version, Commit, BuildDate)
		return
	}
	rootCmd := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
