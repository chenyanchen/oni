package element

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/chenyanchen/oni/db/elementsdb"
	"github.com/chenyanchen/oni/model"
)

type handler struct {
	// target format
	format string

	// output file
	output string
}

func (h *handler) run(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()

	// TODO: read liquid data from
	// 	macOS: ~/Library/Application Support/Steam/steamapps/common/OxygenNotIncluded/OxygenNotIncluded.app
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("get home dir: %v", err)
	}

	elementsPath := filepath.Join(
		homeDir,
		"Library/Application Support/Steam/steamapps/common/OxygenNotIncluded/OxygenNotIncluded.app",
		"Contents/Resources/Data/StreamingAssets/elements",
	)

	elementFiles := []string{
		"gas.yaml",
		"liquid.yaml",
		"solid.yaml",
	}

	// Read elements from files.
	var elements []*model.Element
	for _, file := range elementFiles {
		elementsFromFile, err := readElementsFromYamlFile(filepath.Join(elementsPath, file))
		if err != nil {
			log.Fatalf("read elements from file: %v", err)
		}

		elements = append(elements, elementsFromFile...)
	}

	writer := os.Stdout
	if h.output != "" {
		outputFile, err := os.Create(h.output)
		if err != nil {
			log.Fatalf("create file: %v", err)
		}
		defer outputFile.Close()

		writer = outputFile
	}

	switch h.format {
	case "csv":
		if err = elementsdb.NewCsvKV(writer).Set(ctx, struct{}{}, elements); err != nil {
			log.Fatalf("write csv: %v", err)
		}
	case "json":
		if err = elementsdb.NewJsonKV(writer).Set(ctx, struct{}{}, elements); err != nil {
			log.Fatalf("write json: %v", err)
		}
	case "yaml":
		if err = elementsdb.NewYamlKV(writer).Set(ctx, struct{}{}, elements); err != nil {
			log.Fatalf("write yaml: %v", err)
		}
	default:
		log.Fatalf("unknown format: %s", h.format)
	}
}

func readElementsFromYamlFile(filepath string) ([]*model.Element, error) {
	// Open file.
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read elements from file.
	kv := elementsdb.NewYamlKV(file)
	elements, err := kv.Get(context.Background(), struct{}{})
	if err != nil {
		return nil, err
	}
	return elements, nil
}
