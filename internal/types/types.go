package types

type Instruction struct {
	Package         string
	Title           string
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
	Temperature       float64 `yaml:"temperature"`
	TopP              float64 `yaml:"top_p"`
	TopK              int     `yaml:"top_k"`
	MaxOutputTokens   int     `yaml:"max_output_tokens"`
	FrequencyPenalty  float64 `yaml:"frequency_penalty"`
	PresencePenalty   float64 `yaml:"presence_penalty"`
	RepetitionPenalty float64 `yaml:"repetition_penalty"`
}

func (i *Instruction) ToMap() map[string]any {
	return map[string]any{
		"Package":                 i.Package,
		"Title":                   i.Title,
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
	}
}
