package dto

import (
	"fmt"
	"strconv"

	"github.com/Ja7ad/fluora/internal/errors"
	"github.com/Ja7ad/fluora/internal/types"
	"github.com/gofiber/fiber/v2"
)

const (
	_MaxSentenceLen           = 1024
	_MaxRewriteSuggestions    = 10
	_MaxWordChoiceSuggestions = 20
)

type BaseTextCraftReq struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	Language string `json:"language"`
	Sentence string `json:"sentence"`
}

type RewriteReq struct {
	BaseTextCraftReq
	NumSuggestions uint8       `json:"num_suggestions"`
	Tone           RewriteTone `json:"tone"`
}

func (r *RewriteReq) Validate() error {
	if err := textcraftCommonValidate(r.Language, r.Sentence); err != nil {
		return err
	}

	switch r.Tone {
	case Informal, Neutral, Formal:
	default:
		return errors.New(fiber.StatusBadRequest, "Invalid tone. Accepted values: "+
			"informal, neutral, formal.", map[string]string{
			"tone": r.Tone.String(),
		})
	}

	// ✅ Validate Number of Suggestions
	if r.NumSuggestions > _MaxRewriteSuggestions {
		return errors.New(fiber.StatusBadRequest, fmt.Sprintf("Invalid num_suggestions. "+
			"Maximum allowed: %d.", _MaxRewriteSuggestions), nil)
	}

	return nil
}

type CasualizeReq struct {
	BaseTextCraftReq
}

func (r *CasualizeReq) Validate() error {
	return textcraftCommonValidate(r.Language, r.Sentence)
}

type ExpandReq struct {
	BaseTextCraftReq
}

func (e *ExpandReq) Validate() error {
	return textcraftCommonValidate(e.Language, e.Sentence)
}

type FormalizeReq struct {
	BaseTextCraftReq
}

func (f *FormalizeReq) Validate() error {
	return textcraftCommonValidate(f.Language, f.Sentence)
}

type GrammarReq struct {
	BaseTextCraftReq
}

func (g *GrammarReq) Validate() error {
	return textcraftCommonValidate(g.Language, g.Sentence)
}

type ShortenReq struct {
	BaseTextCraftReq
}

func (s *ShortenReq) Validate() error {
	return textcraftCommonValidate(s.Language, s.Sentence)
}

type WordChoiceReq struct {
	BaseTextCraftReq
	NumSuggestions uint8 `json:"num_suggestions"`
}

func (w *WordChoiceReq) Validate() error {
	if err := textcraftCommonValidate(w.Language, w.Sentence); err != nil {
		return err
	}

	if w.NumSuggestions > _MaxWordChoiceSuggestions {
		return errors.New(fiber.StatusBadRequest, fmt.Sprintf("Invalid num_suggestions. "+
			"Maximum allowed: %d.", _MaxWordChoiceSuggestions), nil)
	}

	return nil
}

func textcraftCommonValidate(language, sentence string) error {
	if language == "" {
		return errors.New(fiber.StatusBadRequest, "Language field is required.", nil)
	}

	if _, ok := types.Language[language]; !ok {
		return errors.New(fiber.StatusBadRequest, fmt.Sprintf("Invalid language: %s. Accepted values: "+
			"English, Spanish, French, etc.", language), map[string]string{
			"language": language,
		})
	}

	sentenceLen := len(sentence)
	if sentenceLen > _MaxSentenceLen {
		return errors.New(fiber.StatusBadRequest, fmt.Sprintf("Sentence exceeds the maximum "+
			"allowed length of %d characters.", _MaxSentenceLen), map[string]string{
			"max_length":    strconv.Itoa(_MaxSentenceLen),
			"actual_length": strconv.Itoa(sentenceLen),
		})
	}

	return nil
}
