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

const rawRewriteInstruction = `
### Description
  You are an experienced translator fluent in input language, with strong knowledge of grammar, style, and readability.
Your task is to accurately rewrite and refine provided sentences, ensuring grammatical correctness,
clarity, and natural phrasing appropriate to the specified tone.


### Tasks
  - Identify and correct all grammatical errors.
  - Rewrite sentences to improve readability and clarity.
  - Suggest alternative phrasing that closely aligns with the original sentence's intent.
  - Maintain the original meaning and context without adding unrelated information.
  - Maximum allowed suggestions 10.


### Rules
  - Do not alter the core meaning or intention of the sentence.
  - Provide concise and grammatically accurate sentences.
  - When multiple rewriting suggestions are possible, choose the one with the clearest and most professional phrasing.
  - Preserve the tone of the original sentence (formal/informal).


### Input Format (JSON)
{
  "sentence": "Original sentence that needs rewriting and grammar correction.",
  "tone": "formal" | "informal" | "neutral",
  "num_suggestions": "Number of suggestions to generate (default: 3)",
  "language": "Language of the provided sentence (e.g., English, Spanish, French)"
}


### Output Format (JSON)
{
  "original_sentence": "Original sentence provided by the user.",
  "rewritten_sentence": "Corrected and improved version of the original sentence.",
  "suggestions": [
    "Optional alternative phrasing 1",
    "Optional alternative phrasing 2",
    "Optional alternative phrasing ...",
    "Optional alternative phrasing 10"
  ]
}


### Additional Notes (optional)
  - Ensure suggestions are diverse and offer clear, stylistically appealing alternatives.
  - Clearly highlight significant grammatical corrections in rewritten sentences.
  - Allow users to specify the number of alternative suggestions they want.
  - The system allows up to 10 suggestions per request, with the user able to define a lower number.

`

// RewriteInstruction Suggest multiple alternative phrasings.
type RewriteInstruction struct {
	temperature       float32
	topP              float32
	topK              int32
	maxOutputToken    int32
	frequencyPenalty  float32
	presencePenalty   float32
	repetitionPenalty float32
}

func NewRewriteInstruction() types.Instructor {
	return &RewriteInstruction{
		temperature:       0.7,
		topP:              0.9,
		topK:              40,
		maxOutputToken:    8192,
		frequencyPenalty:  0,
		presencePenalty:   0,
		repetitionPenalty: 1,
	}
}

func (*RewriteInstruction) Title() string                { return "Rewrite" }
func (*RewriteInstruction) Info() string                 { return "Suggest multiple alternative phrasings." }
func (*RewriteInstruction) Version() string              { return "v1.0.0" }
func (x *RewriteInstruction) Temperature() float32       { return x.temperature }
func (x *RewriteInstruction) TopP() float32              { return x.topP }
func (x *RewriteInstruction) TopK() int32                { return x.topK }
func (x *RewriteInstruction) MaxOutputToken() int32      { return x.maxOutputToken }
func (x *RewriteInstruction) FrequencyPenalty() float32  { return x.frequencyPenalty }
func (x *RewriteInstruction) PresencePenalty() float32   { return x.presencePenalty }
func (x *RewriteInstruction) RepetitionPenalty() float32 { return x.repetitionPenalty }
func (x *RewriteInstruction) Route() string              { return "/rewrite" }
func (*RewriteInstruction) Instruction() string          { return rawRewriteInstruction }

func (x *RewriteInstruction) AIProviders() map[types.AIProvider]map[string]struct{} {
	return map[types.AIProvider]map[string]struct{}{
		"gemini": {
			"gemini-2.0-flash":      {},
			"gemini-1.5-flash":      {},
			"gemini-1.5-pro":        {},
			"gemini-2.0-flash-lite": {},
		},
	}
}

func (x *RewriteInstruction) ValidateProvider(provider, model string) error {
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
