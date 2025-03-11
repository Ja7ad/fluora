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

const rawFormalizeInstruction = `
### Description
  You are an advanced language assistant fluent in input language, specializing in transforming informal or casual sentences into formal, professional, and polished language. 
Your task is to rewrite the provided sentence in a more sophisticated and appropriate tone for professional, academic, or official communication.


### Tasks
  - Transform informal or casual sentences into formal and professional language.
  - Improve sentence structure, grammar, and word choice to enhance professionalism.
  - Ensure the rewritten sentence is clear, polite, and appropriate for official communication.


### Rules
  - Do not change the core meaning of the sentence.
  - Maintain grammatical correctness and logical coherence.
  - Use formal vocabulary and sentence structures while ensuring readability.
  - Avoid slang, contractions, and overly casual expressions.


### Input Format (JSON)
{
  "sentence": "The original informal sentence that needs to be formalized.",
  "language": "Language of the provided sentence (e.g., English, Spanish, French)"
}


### Output Format (JSON)
{
  "original_sentence": "The original informal sentence provided by the user.",
  "formalized_sentence": "The rewritten, more formal version of the sentence."
}


### Additional Notes (optional)
  - Ensure the output aligns with a professional, business, or academic tone.
  - Preserve the original intent while improving clarity and formality.
  - Avoid making the response overly complex or unnatural.

`

// FormalizeInstruction Convert informal text into a professional tone.
type FormalizeInstruction struct {
	temperature       float32
	topP              float32
	topK              int32
	maxOutputToken    int32
	frequencyPenalty  float32
	presencePenalty   float32
	repetitionPenalty float32
}

func NewFormalizeInstruction() types.Instructor {
	return &FormalizeInstruction{
		temperature:       0.4,
		topP:              0.8,
		topK:              40,
		maxOutputToken:    256,
		frequencyPenalty:  0.3,
		presencePenalty:   0,
		repetitionPenalty: 1.1,
	}
}

func (*FormalizeInstruction) Title() string                { return "Formalize" }
func (*FormalizeInstruction) Info() string                 { return "Convert informal text into a professional tone." }
func (*FormalizeInstruction) Version() string              { return "v1.0.0" }
func (x *FormalizeInstruction) Temperature() float32       { return x.temperature }
func (x *FormalizeInstruction) TopP() float32              { return x.topP }
func (x *FormalizeInstruction) TopK() int32                { return x.topK }
func (x *FormalizeInstruction) MaxOutputToken() int32      { return x.maxOutputToken }
func (x *FormalizeInstruction) FrequencyPenalty() float32  { return x.frequencyPenalty }
func (x *FormalizeInstruction) PresencePenalty() float32   { return x.presencePenalty }
func (x *FormalizeInstruction) RepetitionPenalty() float32 { return x.repetitionPenalty }
func (x *FormalizeInstruction) Route() string              { return "/formalize" }
func (*FormalizeInstruction) Instruction() string          { return rawFormalizeInstruction }

func (x *FormalizeInstruction) AIProviders() map[types.AIProvider]map[string]struct{} {
	return map[types.AIProvider]map[string]struct{}{
		"gemini": {
			"gemini-2.0-flash":      {},
			"gemini-1.5-flash":      {},
			"gemini-1.5-pro":        {},
			"gemini-2.0-flash-lite": {},
		},
	}
}

func (x *FormalizeInstruction) ValidateProvider(provider, model string) error {
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
