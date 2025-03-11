package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRewriteReq_Validate(t *testing.T) {
	validReq := &RewriteReq{
		BaseTextCraftReq: BaseTextCraftReq{
			Provider: "gemini",
			Model:    "gemini-2.0",
			Language: "English",
			Sentence: "This is a test sentence.",
		},
		NumSuggestions: 5,
		Tone:           Formal,
	}

	err := validReq.Validate()
	assert.NoError(t, err, "Expected no error for valid RewriteReq")

	invalidToneReq := &RewriteReq{
		BaseTextCraftReq: validReq.BaseTextCraftReq,
		NumSuggestions:   5,
		Tone:             "wrong-tone",
	}

	err = invalidToneReq.Validate()
	assert.Error(t, err, "Expected error for invalid tone")
	assert.Contains(t, err.Error(), "Invalid tone",
		"Tone validation message mismatch")

	tooManySuggestionsReq := &RewriteReq{
		BaseTextCraftReq: validReq.BaseTextCraftReq,
		NumSuggestions:   15, // Exceeds max (10)
		Tone:             Formal,
	}

	err = tooManySuggestionsReq.Validate()
	assert.Error(t, err, "Expected error for too many suggestions")
	assert.Contains(t, err.Error(), "Invalid num_suggestions",
		"NumSuggestions validation message mismatch")
}

func TestWordChoiceReq_Validate(t *testing.T) {
	validReq := &WordChoiceReq{
		BaseTextCraftReq: BaseTextCraftReq{
			Provider: "gemini",
			Model:    "gemini-2.0",
			Language: "English",
			Sentence: "This is a sample sentence.",
		},
		NumSuggestions: 15, // ✅ Allowed (Max: 20)
	}

	err := validReq.Validate()
	assert.NoError(t, err, "Expected no error for valid WordChoiceReq")

	tooManySuggestionsReq := &WordChoiceReq{
		BaseTextCraftReq: validReq.BaseTextCraftReq,
		NumSuggestions:   25,
	}

	err = tooManySuggestionsReq.Validate()
	assert.Error(t, err, "Expected error for too many word choice suggestions")
	assert.Contains(t, err.Error(), "Invalid num_suggestions",
		"NumSuggestions validation message mismatch")
}

func TestTextcraftCommonValidate_MissingLanguage(t *testing.T) {
	err := textcraftCommonValidate("", "This is a test sentence.")
	assert.Error(t, err, "Expected error for missing language")
	assert.Contains(t, err.Error(), "Language field is required",
		"Missing language validation message mismatch")
}

func TestTextcraftCommonValidate_InvalidLanguage(t *testing.T) {
	err := textcraftCommonValidate("UnknownLanguage", "This is a test sentence.")
	assert.Error(t, err, "Expected error for invalid language")
	assert.Contains(t, err.Error(), "Invalid language",
		"Invalid language validation message mismatch")
}

func TestTextcraftCommonValidate_SentenceLength(t *testing.T) {
	longSentence := ""
	for i := 0; i < _MaxSentenceLen+1; i++ {
		longSentence += "a"
	}

	err := textcraftCommonValidate("English", longSentence)
	assert.Error(t, err, "Expected error for sentence exceeding max length")
	assert.Contains(t, err.Error(), "Sentence exceeds the maximum allowed length",
		"Sentence length validation message mismatch")
}

func TestTextcraftCommonValidate_Success(t *testing.T) {
	err := textcraftCommonValidate("English", "This is a valid sentence.")
	assert.NoError(t, err, "Expected no error for valid input")
}
