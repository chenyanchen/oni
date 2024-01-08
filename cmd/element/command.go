package element

import (
	"github.com/spf13/cobra"

	"github.com/chenyanchen/oni/cmd/element/liquid"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "element",
		Short: "element",
		Long:  "element",
	}

	cmd.AddCommand(liquid.New())

	return cmd
}
