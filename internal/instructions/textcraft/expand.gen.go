/*
MIT License

Copyright (c) 2025 Fluora

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated automatically. DO NOT EDIT.

package textcraft

import (
	"fmt"

	"github.com/Ja7ad/fluora/internal/types"
)

const rawExpandInstruction = `
### Description
  You are an advanced language assistant fluent in input language, specializing in expanding short or concise sentences 
into more detailed, descriptive, and informative versions while maintaining clarity and coherence.


### Tasks
  - Expand short or simple sentences into more detailed, informative versions.
  - Provide additional context, explanations, or supporting details as needed.
  - Ensure the expanded sentence remains clear, natural, and grammatically correct.


### Rules
  - Do not alter the original meaning of the sentence.
  - Maintain logical flow and readability.
  - Avoid unnecessary repetition or filler words.
  - Ensure that expansions remain relevant and do not introduce unrelated ideas.


### Input Format (JSON)
{
  "sentence": "The original short sentence that needs to be expanded.",
  "language": "Language of the provided sentence (e.g., English, Spanish, French)"
}


### Output Format (JSON)
{
  "original_sentence": "The original short sentence provided by the user.",
  "expand_sentence": "The more detailed and descriptive version of the sentence."
}


### Additional Notes (optional)
  - Ensure expansions add meaningful and relevant details without unnecessary complexity.
  - Keep readability high while enhancing the descriptive nature of the sentence.
  - Adapt the level of detail based on the sentence type (formal, casual, academic, etc.).

`

// ExpandInstruction Make sentences longer while keeping the meaning.
type ExpandInstruction struct {
	temperature       float32
	topP              float32
	topK              int32
	maxOutputToken    int32
	frequencyPenalty  float32
	presencePenalty   float32
	repetitionPenalty float32
}

func NewExpandInstruction() types.Instructor {
	return &ExpandInstruction{
		temperature:       0.5,
		topP:              0.85,
		topK:              50,
		maxOutputToken:    512,
		frequencyPenalty:  0.2,
		presencePenalty:   0,
		repetitionPenalty: 1,
	}
}

func (*ExpandInstruction) Title() string                { return "Expand" }
func (*ExpandInstruction) Info() string                 { return "Make sentences longer while keeping the meaning." }
func (*ExpandInstruction) Version() string              { return "v1.0.0" }
func (x *ExpandInstruction) Temperature() float32       { return x.temperature }
func (x *ExpandInstruction) TopP() float32              { return x.topP }
func (x *ExpandInstruction) TopK() int32                { return x.topK }
func (x *ExpandInstruction) MaxOutputToken() int32      { return x.maxOutputToken }
func (x *ExpandInstruction) FrequencyPenalty() float32  { return x.frequencyPenalty }
func (x *ExpandInstruction) PresencePenalty() float32   { return x.presencePenalty }
func (x *ExpandInstruction) RepetitionPenalty() float32 { return x.repetitionPenalty }
func (x *ExpandInstruction) Route() string              { return "/expand" }
func (*ExpandInstruction) Instruction() string          { return rawExpandInstruction }

func (x *ExpandInstruction) AIProviders() map[types.AIProvider]map[string]struct{} {
	return map[types.AIProvider]map[string]struct{}{
		"gemini": {
			"gemini-2.0-flash":      {},
			"gemini-1.5-flash":      {},
			"gemini-1.5-pro":        {},
			"gemini-2.0-flash-lite": {},
		},
	}
}

func (x *ExpandInstruction) ValidateProvider(provider, model string) error {
	models, ok := x.AIProviders()[types.AIProvider(provider)]
	if !ok {
		return fmt.Errorf("invalid AI provider: %s", provider)
	}

	_, ok = models[model]
	if !ok {
		return fmt.Errorf("invalid model: %s", model)
	}

	return nil
}
