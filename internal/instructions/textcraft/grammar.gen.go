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

const rawGrammarInstruction = `
### Description
  You are an experienced linguist and grammar expert fluent in input language. 
Your task is to analyze a given sentence for grammatical correctness, suggest necessary corrections, 
and provide alternative word or token suggestions to enhance clarity, readability, and style while preserving the original meaning.


### Tasks
  - Identify and correct grammatical errors in the provided sentence.
  - Suggest alternative words or tokens that can improve readability and clarity.
  - Preserve the original meaning and tone of the sentence while making necessary refinements.


### Rules
  - Do not alter the intended meaning of the sentence.
  - Ensure grammatical correctness while maintaining fluency.
  - Provide suggestions that enhance readability without making unnecessary changes.
  - Suggestions should be clear, concise, and grammatically accurate.


### Input Format (JSON)
{
  "sentence": "The original sentence provided by the user.",
  "language": "Language of the provided sentence (e.g., English, Spanish, French)"
}


### Output Format (JSON)
{
  "original": "The original sentence provided by the user.",
  "corrected": "Grammatically corrected version of the sentence.",
  "word_suggestions": {
    "original_word": "suggested_alternative"
  }
}


### Additional Notes (optional)
  - Suggestions should provide clear improvements without overcomplicating the sentence.
  - Ensure that grammar corrections align with natural language usage and readability.
  - Highlight key grammar fixes and suggested word alternatives where applicable.

`

// GrammarInstruction Correct mistakes while keeping the style.
type GrammarInstruction struct {
	temperature       float32
	topP              float32
	topK              int32
	maxOutputToken    int32
	frequencyPenalty  float32
	presencePenalty   float32
	repetitionPenalty float32
}

func NewGrammarInstruction() types.Instructor {
	return &GrammarInstruction{
		temperature:       0.5,
		topP:              0.8,
		topK:              50,
		maxOutputToken:    256,
		frequencyPenalty:  0.2,
		presencePenalty:   0,
		repetitionPenalty: 1,
	}
}

func (*GrammarInstruction) Title() string                { return "Grammar" }
func (*GrammarInstruction) Info() string                 { return "Correct mistakes while keeping the style." }
func (*GrammarInstruction) Version() string              { return "v1.0.0" }
func (x *GrammarInstruction) Temperature() float32       { return x.temperature }
func (x *GrammarInstruction) TopP() float32              { return x.topP }
func (x *GrammarInstruction) TopK() int32                { return x.topK }
func (x *GrammarInstruction) MaxOutputToken() int32      { return x.maxOutputToken }
func (x *GrammarInstruction) FrequencyPenalty() float32  { return x.frequencyPenalty }
func (x *GrammarInstruction) PresencePenalty() float32   { return x.presencePenalty }
func (x *GrammarInstruction) RepetitionPenalty() float32 { return x.repetitionPenalty }
func (x *GrammarInstruction) Route() string              { return "/grammar" }
func (*GrammarInstruction) Instruction() string          { return rawGrammarInstruction }

func (x *GrammarInstruction) AIProviders() map[types.AIProvider]map[string]struct{} {
	return map[types.AIProvider]map[string]struct{}{
		"gemini": {
			"gemini-2.0-flash":      {},
			"gemini-1.5-flash":      {},
			"gemini-1.5-pro":        {},
			"gemini-2.0-flash-lite": {},
		},
	}
}

func (x *GrammarInstruction) ValidateProvider(provider, model string) error {
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
