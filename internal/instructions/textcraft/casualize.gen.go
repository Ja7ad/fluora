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

const rawCasualizeInstruction = `
### Description
  You are an advanced language assistant fluent in input language, specializing in transforming formal or rigid sentences into casual, friendly, and conversational language.
Your task is to rewrite the provided sentence in a more relaxed and approachable tone while keeping the original meaning intact.


### Tasks
  - Convert formal or structured sentences into relaxed, conversational language.
  - Use natural phrasing that sounds more like everyday speech.
  - Maintain clarity, correctness, and friendliness while making it sound informal.


### Rules
  - Do not remove key details or alter the core meaning of the sentence.
  - Keep the rewritten version engaging, natural, and easy to understand.
  - Use contractions and everyday expressions where appropriate.
  - Avoid unnecessary formality, but do not make the sentence overly slang-heavy unless needed.


### Input Format (JSON)
{
  "sentence": "The original formal sentence that needs to be casualized.",
  "language": "Language of the provided sentence (e.g., English, Spanish, French)"
}


### Output Format (JSON)
{
  "original_sentence": "The original formal sentence provided by the user.",
  "casualized_sentence": "The relaxed, more conversational version of the sentence."
}


### Additional Notes (optional)
  - Keep the response engaging and friendly while maintaining clarity.
  - Ensure that the tone is appropriate for informal but respectful communication.
  - Adapt the level of casualness based on context (e.g., social, business-casual, friendly conversation).

`

// CasualizeInstruction Make text sound more friendly and conversational.
type CasualizeInstruction struct {
	temperature       float32
	topP              float32
	topK              int32
	maxOutputToken    int32
	frequencyPenalty  float32
	presencePenalty   float32
	repetitionPenalty float32
}

func NewCasualizeInstruction() types.Instructor {
	return &CasualizeInstruction{
		temperature:       0.6,
		topP:              0.85,
		topK:              50,
		maxOutputToken:    256,
		frequencyPenalty:  0.2,
		presencePenalty:   0.1,
		repetitionPenalty: 1,
	}
}

func (*CasualizeInstruction) Title() string { return "Casualize" }
func (*CasualizeInstruction) Info() string {
	return "Make text sound more friendly and conversational."
}
func (*CasualizeInstruction) Version() string              { return "v1.0.0" }
func (x *CasualizeInstruction) Temperature() float32       { return x.temperature }
func (x *CasualizeInstruction) TopP() float32              { return x.topP }
func (x *CasualizeInstruction) TopK() int32                { return x.topK }
func (x *CasualizeInstruction) MaxOutputToken() int32      { return x.maxOutputToken }
func (x *CasualizeInstruction) FrequencyPenalty() float32  { return x.frequencyPenalty }
func (x *CasualizeInstruction) PresencePenalty() float32   { return x.presencePenalty }
func (x *CasualizeInstruction) RepetitionPenalty() float32 { return x.repetitionPenalty }
func (x *CasualizeInstruction) Route() string              { return "/casualize" }
func (*CasualizeInstruction) Instruction() string          { return rawCasualizeInstruction }

func (x *CasualizeInstruction) AIProviders() map[types.AIProvider]map[string]struct{} {
	return map[types.AIProvider]map[string]struct{}{
		"gemini": {
			"gemini-2.0-flash":      {},
			"gemini-1.5-flash":      {},
			"gemini-1.5-pro":        {},
			"gemini-2.0-flash-lite": {},
		},
	}
}

func (x *CasualizeInstruction) ValidateProvider(provider, model string) error {
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
