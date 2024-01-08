package element

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	h := &handler{}

	cmd := &cobra.Command{
		Use:   "element",
		Short: "element",
		Long:  "element",
		Run:   h.run,
	}

	cmd.Flags().StringVarP(&h.format, "format", "f", "json", "optional format: json, yaml, csv")
	cmd.Flags().StringVarP(&h.output, "output", "o", "", "output file")

	return cmd
}
