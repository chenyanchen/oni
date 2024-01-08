package liquid

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/chenyanchen/oni/model"
)

func New() *cobra.Command {
	h := &handler{}
	cmd := &cobra.Command{
		Use:   "liquid",
		Short: "liquid",
		Long:  "liquid",
		Run:   h.run,
	}
	cmd.Flags().StringVarP(&h.format, "format", "f", "json", "optional format: json, csv")
	cmd.Flags().StringVarP(&h.output, "output", "o", "", "output file")
	return cmd
}

type handler struct {
	// target format
	format string

	// output file
	output string
}

func (h *handler) run(cmd *cobra.Command, args []string) {
	// TODO: read liquid data from
	// 	macOS: ~/Library/Application Support/Steam/steamapps/common/OxygenNotIncluded/OxygenNotIncluded.app
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("get home dir: %v", err)
	}

	liquidFilePath := filepath.Join(
		homeDir,
		"Library/Application Support/Steam/steamapps/common/OxygenNotIncluded/OxygenNotIncluded.app",
		"Contents/Resources/Data/StreamingAssets/elements/liquid.yaml",
	)

	file, err := os.Open(liquidFilePath)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer file.Close()

	var result struct {
		Elements []*model.Liquid `yaml:"elements"`
	}

	if err := yaml.NewDecoder(file).Decode(&result); err != nil {
		log.Fatalf("decode yaml: %v", err)
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
	case "json":
		if err := json.NewEncoder(writer).Encode(result.Elements); err != nil {
			log.Fatalf("encode json: %v", err)
		}
	case "csv":
		csvWriter := csv.NewWriter(writer)

		// Write header.
		header := []string{
			"ElementId",
			"MaxMass",
			"LiquidCompression",
			"Speed",
			"MinHorizontalFlow",
			"MinVerticalFlow",
			"SpecificHeatCapacity",
			"ThermalConductivity",
			"SolidSurfaceAreaMultiplier",
			"LiquidSurfaceAreaMultiplier",
			"GasSurfaceAreaMultiplier",
			"LowTemp",
			"HighTemp",
			"LowTempTransitionTarget",
			"HighTempTransitionTarget",
			"DefaultTemperature",
			"DefaultMass",
			"MolarMass",
			"Toxicity",
			"LightAbsorptionFactor",
			"RadiationAbsorptionFactor",
			"RadiationPer1000Mass",
			"Tags",
			"IsDisabled",
			"State",
			"LocalizationID",
			"DlcId",
		}
		if err := csvWriter.Write(header); err != nil {
			log.Fatalf("write csv: %v", err)
		}

		// Write data.
		for _, element := range result.Elements {
			if err := csvWriter.Write([]string{
				element.ElementId,
				strconv.Itoa(element.MaxMass),
				strconv.FormatFloat(element.LiquidCompression, 'f', -1, 64),
				strconv.Itoa(element.Speed),
				strconv.Itoa(element.MinHorizontalFlow),
				strconv.Itoa(element.MinVerticalFlow),
				strconv.FormatFloat(element.SpecificHeatCapacity, 'f', -1, 64),
				strconv.Itoa(element.ThermalConductivity),
				strconv.Itoa(element.SolidSurfaceAreaMultiplier),
				strconv.Itoa(element.LiquidSurfaceAreaMultiplier),
				strconv.Itoa(element.GasSurfaceAreaMultiplier),
				strconv.Itoa(element.LowTemp),
				strconv.Itoa(element.HighTemp),
				element.LowTempTransitionTarget,
				element.HighTempTransitionTarget,
				strconv.Itoa(element.DefaultTemperature),
				strconv.Itoa(element.DefaultMass),
				strconv.FormatFloat(element.MolarMass, 'f', -1, 64),
				strconv.Itoa(element.Toxicity),
				strconv.Itoa(element.LightAbsorptionFactor),
				strconv.FormatFloat(element.RadiationAbsorptionFactor, 'f', -1, 64),
				strconv.Itoa(element.RadiationPer1000Mass),
				strings.Join(element.Tags, ","),
				strconv.FormatBool(element.IsDisabled),
				element.State,
				element.LocalizationID,
				element.DlcId,
			}); err != nil {
				log.Fatalf("write csv: %v", err)
			}

			csvWriter.Flush()
		}
	}
}
