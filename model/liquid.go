package model

type Liquid struct {
	ElementId                   string   `yaml:"elementId"`
	MaxMass                     int      `yaml:"maxMass"`
	LiquidCompression           float64  `yaml:"liquidCompression"`
	Speed                       int      `yaml:"speed"`
	MinHorizontalFlow           int      `yaml:"minHorizontalFlow"`
	MinVerticalFlow             int      `yaml:"minVerticalFlow"`
	SpecificHeatCapacity        float64  `yaml:"specificHeatCapacity"`
	ThermalConductivity         int      `yaml:"thermalConductivity"`
	SolidSurfaceAreaMultiplier  int      `yaml:"solidSurfaceAreaMultiplier"`
	LiquidSurfaceAreaMultiplier int      `yaml:"liquidSurfaceAreaMultiplier"`
	GasSurfaceAreaMultiplier    int      `yaml:"gasSurfaceAreaMultiplier"`
	LowTemp                     int      `yaml:"lowTemp"`
	HighTemp                    int      `yaml:"highTemp"`
	LowTempTransitionTarget     string   `yaml:"lowTempTransitionTarget"`
	HighTempTransitionTarget    string   `yaml:"highTempTransitionTarget"`
	DefaultTemperature          int      `yaml:"defaultTemperature"`
	DefaultMass                 int      `yaml:"defaultMass"`
	MolarMass                   float64  `yaml:"molarMass"`
	Toxicity                    int      `yaml:"toxicity"`
	LightAbsorptionFactor       int      `yaml:"lightAbsorptionFactor"`
	RadiationAbsorptionFactor   float64  `yaml:"radiationAbsorptionFactor"`
	RadiationPer1000Mass        int      `yaml:"radiationPer1000Mass"`
	Tags                        []string `yaml:"tags"`
	IsDisabled                  bool     `yaml:"isDisabled"`
	State                       string   `yaml:"state"`
	LocalizationID              string   `yaml:"localizationID"`
	DlcId                       string   `yaml:"dlcId"`
}
