package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/dsrvlabs/love/app"
	"github.com/dsrvlabs/love/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		"love",
		"cosmos",
		app.DefaultNodeHome,
		"love-1",
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
