package types

type Instruction struct {
	InstructionInfo *Info    `yaml:"instruction_info"`
	Config          *Config  `yaml:"config"`
	Description     string   `yaml:"description"`
	Tasks           []string `yaml:"tasks"`
	Rules           []string `yaml:"rules"`
	InputFormat     string   `yaml:"input_format"`
	OutputFormat    string   `yaml:"output_format"`
	AdditionalNotes []string `yaml:"additional_notes"`
}

type Info struct {
	Title   string `yaml:"title"`
	Info    string `yaml:"info"`
	Version string `yaml:"version"`
}

type Config struct {
	Models            map[AIProvider][]string `yaml:"models"`
	Route             string                  `yaml:"route"`
	Temperature       float64                 `yaml:"temperature"`
	TopP              float64                 `yaml:"top_p"`
	TopK              int                     `yaml:"top_k"`
	MaxOutputTokens   int                     `yaml:"max_output_tokens"`
	FrequencyPenalty  float64                 `yaml:"frequency_penalty"`
	PresencePenalty   float64                 `yaml:"presence_penalty"`
	RepetitionPenalty float64                 `yaml:"repetition_penalty"`
}

func (i *Instruction) ToMap(title, packageName string) map[string]any {
	return map[string]any{
		"Package":                 packageName,
		"Title":                   title,
		"Description":             i.Description,
		"Tasks":                   i.Tasks,
		"Rules":                   i.Rules,
		"InputFormat":             i.InputFormat,
		"OutputFormat":            i.OutputFormat,
		"AdditionalNotes":         i.AdditionalNotes,
		"InstructionInfo":         i.InstructionInfo.Info,
		"InstructionTitle":        i.InstructionInfo.Title,
		"InstructionVersion":      i.InstructionInfo.Version,
		"TemperatureConfig":       i.Config.Temperature,
		"TopPConfig":              i.Config.TopP,
		"TopKConfig":              i.Config.TopK,
		"MaxOutputTokenConfig":    i.Config.MaxOutputTokens,
		"FrequencyPenaltyConfig":  i.Config.FrequencyPenalty,
		"PresencePenaltyConfig":   i.Config.PresencePenalty,
		"RepetitionPenaltyConfig": i.Config.RepetitionPenalty,
		"InstructionRoute":        i.Config.Route,
		"Models":                  i.Config.Models,
	}
}

var Providers = map[string]map[string]struct{}{
	"gemini": {
		"gemini-2.0-flash":      {},
		"gemini-1.5-flash":      {},
		"gemini-1.5-pro":        {},
		"gemini-2.0-flash-lite": {},
	},
}

type AIProvider string

const (
	GEMINI AIProvider = "gemini"
)

func (a AIProvider) String() string {
	return string(a)
}

type Instructor interface {
	Title() string
	Info() string
	Version() string
	Temperature() float32
	TopP() float32
	TopK() int32
	MaxOutputToken() int32
	FrequencyPenalty() float32
	PresencePenalty() float32
	RepetitionPenalty() float32
	Route() string
	Instruction() string
	AIProviders() map[AIProvider]map[string]struct{}
	ValidateProvider(provider, model string) error
}

type GenerativeParams struct {
	Instruction string
	Text        string
	Model       string
	Gemini
}

type Gemini struct {
	Temperature     float32
	TopP            float32
	TopK            int32
	MaxOutputTokens int32
}

func (g Gemini) IsZero() bool {
	return g.Temperature == 0 && g.TopP == 0 && g.TopK == 0 && g.MaxOutputTokens == 0
}

func (g Gemini) Default() Gemini {
	g.Temperature = 0.9
	g.TopP = 0.1
	g.TopK = 100
	g.MaxOutputTokens = 8192
	return g
}
