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

const rawShortenInstruction = `
### Description
  You are an advanced language assistant fluent in input language, specializing in sentence compression and clarity improvement.
Your task is to shorten a given sentence while maintaining its core meaning, removing redundancy, and improving readability.


### Tasks
  - Reduce the sentence length while keeping its meaning intact.
  - Remove unnecessary words, redundancy, and filler phrases.
  - Maintain grammatical correctness and logical flow.


### Rules
  - Do not alter the intended meaning of the sentence.
  - Ensure the shortened sentence is grammatically correct and natural.
  - Avoid removing key details necessary for comprehension.
  - Do not make the sentence too abrupt or lose important context.


### Input Format (JSON)
{
  "sentence": "The original sentence that needs to be shortened.",
  "language": "Language of the provided sentence (e.g., English, Spanish, French)"
}


### Output Format (JSON)
{
  "original_sentence": "The original sentence provided by the user.",
  "shortened_sentence": "The revised, more concise version of the sentence."
}


### Additional Notes (optional)
  - Ensure the shortened sentence retains key meaning and context.
  - Avoid making the response too abrupt or overly simplified.
  - Keep readability high while ensuring effective compression.

`

// ShortenInstruction Make sentences more concise.
type ShortenInstruction struct {
	temperature       float32
	topP              float32
	topK              int32
	maxOutputToken    int32
	frequencyPenalty  float32
	presencePenalty   float32
	repetitionPenalty float32
}

func NewShortenInstruction() types.Instructor {
	return &ShortenInstruction{
		temperature:       0.4,
		topP:              0.75,
		topK:              30,
		maxOutputToken:    128,
		frequencyPenalty:  0.3,
		presencePenalty:   0,
		repetitionPenalty: 1.1,
	}
}

func (*ShortenInstruction) Title() string                { return "Shorten" }
func (*ShortenInstruction) Info() string                 { return "Make sentences more concise." }
func (*ShortenInstruction) Version() string              { return "v1.0.0" }
func (x *ShortenInstruction) Temperature() float32       { return x.temperature }
func (x *ShortenInstruction) TopP() float32              { return x.topP }
func (x *ShortenInstruction) TopK() int32                { return x.topK }
func (x *ShortenInstruction) MaxOutputToken() int32      { return x.maxOutputToken }
func (x *ShortenInstruction) FrequencyPenalty() float32  { return x.frequencyPenalty }
func (x *ShortenInstruction) PresencePenalty() float32   { return x.presencePenalty }
func (x *ShortenInstruction) RepetitionPenalty() float32 { return x.repetitionPenalty }
func (x *ShortenInstruction) Route() string              { return "/shorten" }
func (*ShortenInstruction) Instruction() string          { return rawShortenInstruction }

func (x *ShortenInstruction) AIProviders() map[types.AIProvider]map[string]struct{} {
	return map[types.AIProvider]map[string]struct{}{
		"gemini": {
			"gemini-2.0-flash":      {},
			"gemini-1.5-flash":      {},
			"gemini-1.5-pro":        {},
			"gemini-2.0-flash-lite": {},
		},
	}
}

func (x *ShortenInstruction) ValidateProvider(provider, model string) error {
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
