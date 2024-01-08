package elementsdb

import (
	"context"
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	"github.com/chenyanchen/oni/model"
)

type csvKV struct {
	writer io.Writer
}

func NewCsvKV(writer io.Writer) *csvKV {
	return &csvKV{writer: writer}
}

func (s *csvKV) Get(ctx context.Context, k struct{}) ([]*model.Element, error) {
	// TODO implement me
	panic("implement me")
}

func (s *csvKV) Set(ctx context.Context, k struct{}, elements []*model.Element) error {
	writer := csv.NewWriter(s.writer)

	// Write header.
	if err := writer.Write(buildHeader()); err != nil {
		return err
	}

	// Write elements.
	for _, element := range elements {
		if err := writer.Write(formatElementForCSV(element)); err != nil {
			return err
		}
	}

	writer.Flush()

	return nil
}

func buildHeader() []string {
	return []string{
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
		"Strength",
		"LowTemp",
		"HighTemp",
		"LowTempTransitionTarget",
		"HighTempTransitionTarget",
		"DefaultTemperature",
		"DefaultMass",
		"Hardness",
		"MolarMass",
		"Toxicity",
		"LightAbsorptionFactor",
		"RadiationAbsorptionFactor",
		"RadiationPer1000Mass",
		"MaterialCategory",
		"Tags",
		"IsDisabled",
		"State",
		"LocalizationID",
		"DlcId",
	}
}

func formatElementForCSV(e *model.Element) []string {
	return []string{
		e.ElementId,
		formatFloat64(e.MaxMass),
		formatFloat64(e.LiquidCompression),
		formatFloat64(e.Speed),
		formatFloat64(e.MinHorizontalFlow),
		formatFloat64(e.MinVerticalFlow),
		formatFloat64(e.SpecificHeatCapacity),
		formatFloat64(e.ThermalConductivity),
		formatFloat64(e.SolidSurfaceAreaMultiplier),
		formatFloat64(e.LiquidSurfaceAreaMultiplier),
		formatFloat64(e.GasSurfaceAreaMultiplier),
		formatFloat64(e.Strength),
		formatFloat64(e.LowTemp),
		formatFloat64(e.HighTemp),
		e.LowTempTransitionTarget,
		e.HighTempTransitionTarget,
		formatFloat64(e.DefaultTemperature),
		formatFloat64(e.DefaultMass),
		formatFloat64(e.Hardness),
		formatFloat64(e.MolarMass),
		formatFloat64(e.Toxicity),
		formatFloat64(e.LightAbsorptionFactor),
		formatFloat64(e.RadiationAbsorptionFactor),
		formatFloat64(e.RadiationPer1000Mass),
		e.MaterialCategory,
		strings.Join(e.Tags, ","),
		strconv.FormatBool(e.IsDisabled),
		e.State,
		e.LocalizationID,
		e.DlcId,
	}
}

func formatFloat64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func (s *csvKV) Del(ctx context.Context, k struct{}) error {
	return nil
}
