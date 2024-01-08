package main

import (
	"github.com/spf13/cobra"

	"github.com/chenyanchen/oni/cmd/element"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oni",
		Short: "Oxygen Not Included tools",
	}

	cmd.AddCommand(element.New())

	return cmd
}
