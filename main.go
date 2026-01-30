package main

import (
	"os"

	"github.com/agentio/macaroo/cmd/create"
	"github.com/agentio/macaroo/cmd/evaluate"
	"github.com/agentio/macaroo/cmd/extend"
	"github.com/agentio/macaroo/cmd/print"
	"github.com/agentio/macaroo/cmd/verify"
	"github.com/spf13/cobra"
)

func main() {
	if err := cmd().Execute(); err != nil {
		os.Exit(1)
	}
}

func cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "macaroo",
		Short: "a tool for working with macaroons",
	}
	cmd.AddCommand(create.Cmd())
	cmd.AddCommand(extend.Cmd())
	cmd.AddCommand(evaluate.Cmd())
	cmd.AddCommand(print.Cmd())
	cmd.AddCommand(verify.Cmd())
	return cmd
}
