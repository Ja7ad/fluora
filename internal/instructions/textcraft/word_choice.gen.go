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

const rawWordChoiceSuggestionsInstruction = `
### Description
  You are an advanced language assistant fluent in input language, specializing in vocabulary refinement and word choice optimization.
Your task is to analyze a given sentence and suggest alternative words or phrases that enhance clarity, style, and readability,
while preserving the original meaning and tone. Additionally, ensure that the suggested alternatives match the formality of the original sentence.


### Tasks
  - Identify words that can be improved or replaced with better alternatives.
  - Suggest synonyms or rephrasings that enhance clarity, conciseness, and engagement.
  - Ensure that suggestions match the original sentence's tone (formal/informal/neutral).
  - Avoid altering the intended meaning while refining word choice.
  - Maximum allowed word_suggestions for each word is 20.


### Rules
  - Provide alternatives that improve readability and fit the context.
  - Ensure word suggestions preserve proper grammar and syntax.
  - If a sentence has multiple word choices, prioritize the most natural-sounding alternative.
  - Avoid overly complex or uncommon words unless explicitly required by context.


### Input Format (JSON)
{
  "sentence": "The original sentence where word choices need improvement.",
  "num_suggestions": "Number of word_suggestions for each word to generate (default: 3)",
  "language": "Language of the provided sentence (e.g., English, Spanish, French)"
}


### Output Format (JSON)
{
  "original_sentence": "The original sentence provided by the user.",
  "word_suggestions": {
    "original_word": [
      "suggested_alternative_1",
      "suggested_alternative_...",
      "suggested_alternative_20" 
    ]
  }
}


### Additional Notes (optional)
  - Provide natural, context-aware word suggestions that align with the formality of the sentence.
  - Ensure that suggestions do not introduce ambiguity or significantly alter meaning.
  - Adapt suggestions for professional, casual, or neutral tones as necessary.
  - The system allows up to 20 suggestions per request for word_suggestions, with the user able to define a lower number.

`

// WordChoiceSuggestionsInstruction Suggest alternative words with slight tone variations.
type WordChoiceSuggestionsInstruction struct {
	temperature       float32
	topP              float32
	topK              int32
	maxOutputToken    int32
	frequencyPenalty  float32
	presencePenalty   float32
	repetitionPenalty float32
}

func NewWordChoiceSuggestionsInstruction() types.Instructor {
	return &WordChoiceSuggestionsInstruction{
		temperature:       0.4,
		topP:              0.75,
		topK:              35,
		maxOutputToken:    8192,
		frequencyPenalty:  0.3,
		presencePenalty:   0,
		repetitionPenalty: 1.1,
	}
}

func (*WordChoiceSuggestionsInstruction) Title() string { return "Word Choice Suggestions" }
func (*WordChoiceSuggestionsInstruction) Info() string {
	return "Suggest alternative words with slight tone variations."
}
func (*WordChoiceSuggestionsInstruction) Version() string              { return "v1.0.0" }
func (x *WordChoiceSuggestionsInstruction) Temperature() float32       { return x.temperature }
func (x *WordChoiceSuggestionsInstruction) TopP() float32              { return x.topP }
func (x *WordChoiceSuggestionsInstruction) TopK() int32                { return x.topK }
func (x *WordChoiceSuggestionsInstruction) MaxOutputToken() int32      { return x.maxOutputToken }
func (x *WordChoiceSuggestionsInstruction) FrequencyPenalty() float32  { return x.frequencyPenalty }
func (x *WordChoiceSuggestionsInstruction) PresencePenalty() float32   { return x.presencePenalty }
func (x *WordChoiceSuggestionsInstruction) RepetitionPenalty() float32 { return x.repetitionPenalty }
func (x *WordChoiceSuggestionsInstruction) Route() string              { return "/word-choice" }
func (*WordChoiceSuggestionsInstruction) Instruction() string {
	return rawWordChoiceSuggestionsInstruction
}

func (x *WordChoiceSuggestionsInstruction) AIProviders() map[types.AIProvider]map[string]struct{} {
	return map[types.AIProvider]map[string]struct{}{
		"gemini": {
			"gemini-2.0-flash":      {},
			"gemini-1.5-flash":      {},
			"gemini-1.5-pro":        {},
			"gemini-2.0-flash-lite": {},
		},
	}
}

func (x *WordChoiceSuggestionsInstruction) ValidateProvider(provider, model string) error {
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
