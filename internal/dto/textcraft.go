package dto

import (
	"encoding/json"
)

type BaseTextParam struct {
	Sentence string `json:"sentence"`
	Language string `json:"language"`
}

func (b *BaseTextParam) Text() string {
	data, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(data)
}

type CasualizeParam struct {
	BaseTextParam
}

type CasualizeResp struct {
	*Header            `json:"general,inline"`
	OriginalSentence   string `json:"original_sentence"`
	CasualizedSentence string `json:"casualized_sentence"`
}

type ExpandParam struct {
	BaseTextParam
}

type ExpandResp struct {
	*Header          `json:"general,inline"`
	OriginalSentence string `json:"original_sentence"`
	ExpandSentence   string `json:"expand_sentence"`
}

type FormalizeParam struct {
	BaseTextParam
}

type FormalizeResp struct {
	*Header            `json:"general,inline"`
	OriginalSentence   string `json:"original_sentence"`
	FormalizedSentence string `json:"formalized_sentence"`
}

type GrammarParam struct {
	BaseTextParam
}

type GrammarResp struct {
	*Header         `json:"general,inline"`
	Original        string            `json:"original"`
	Corrected       string            `json:"corrected"`
	WordSuggestions map[string]string `json:"word_suggestions"`
}

type RewriteParam struct {
	BaseTextParam
	NumSuggestions uint8       `json:"num_suggestions"`
	Tone           RewriteTone `json:"tone"`
}

type RewriteTone string

const (
	Formal   RewriteTone = "formal"
	Informal RewriteTone = "informal"
	Neutral  RewriteTone = "neutral"
)

func (r RewriteTone) String() string {
	return string(r)
}

type RewriteResp struct {
	*Header           `json:"general,inline"`
	OriginalSentence  string   `json:"original_sentence"`
	RewrittenSentence string   `json:"rewritten_sentence"`
	Suggestions       []string `json:"suggestions"`
}

type ShortenParam struct {
	BaseTextParam
}

type ShortenResp struct {
	*Header           `json:"general,inline"`
	OriginalSentence  string `json:"original_sentence"`
	ShortenedSentence string `json:"shortened_sentence"`
}

type WordChoiceParam struct {
	BaseTextParam
	NumSuggestions uint8 `json:"num_suggestions"`
}

type WordChoiceResp struct {
	*Header          `json:"general,inline"`
	OriginalSentence string              `json:"original_sentence"`
	WordSuggestions  map[string][]string `json:"word_suggestions"`
}
