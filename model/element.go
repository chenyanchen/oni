package model

type Element struct {
	ElementId                   string   `yaml:"elementId"`
	MaxMass                     float64  `yaml:"maxMass"`
	LiquidCompression           float64  `yaml:"liquidCompression"`
	Speed                       float64  `yaml:"speed"`
	MinHorizontalFlow           float64  `yaml:"minHorizontalFlow"`
	MinVerticalFlow             float64  `yaml:"minVerticalFlow"`
	SpecificHeatCapacity        float64  `yaml:"specificHeatCapacity"`
	ThermalConductivity         float64  `yaml:"thermalConductivity"`
	SolidSurfaceAreaMultiplier  float64  `yaml:"solidSurfaceAreaMultiplier"`
	LiquidSurfaceAreaMultiplier float64  `yaml:"liquidSurfaceAreaMultiplier"`
	GasSurfaceAreaMultiplier    float64  `yaml:"gasSurfaceAreaMultiplier"`
	Strength                    float64  `yaml:"strength"`
	LowTemp                     float64  `yaml:"lowTemp"`
	HighTemp                    float64  `yaml:"highTemp"`
	LowTempTransitionTarget     string   `yaml:"lowTempTransitionTarget"`
	HighTempTransitionTarget    string   `yaml:"highTempTransitionTarget"`
	DefaultTemperature          float64  `yaml:"defaultTemperature"`
	DefaultMass                 float64  `yaml:"defaultMass"`
	Hardness                    float64  `yaml:"hardness"`
	MolarMass                   float64  `yaml:"molarMass"`
	Toxicity                    float64  `yaml:"toxicity"`
	LightAbsorptionFactor       float64  `yaml:"lightAbsorptionFactor"`
	RadiationAbsorptionFactor   float64  `yaml:"radiationAbsorptionFactor"`
	RadiationPer1000Mass        float64  `yaml:"radiationPer1000Mass"`
	MaterialCategory            string   `yaml:"materialCategory"`
	Tags                        []string `yaml:"tags"`
	IsDisabled                  bool     `yaml:"isDisabled"`
	State                       string   `yaml:"state"`
	LocalizationID              string   `yaml:"localizationID"`
	DlcId                       string   `yaml:"dlcId"`
}
