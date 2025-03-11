package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Ja7ad/fluora/internal/adapter/ai"
	"github.com/Ja7ad/fluora/internal/dto"
	"github.com/Ja7ad/fluora/internal/instructions/textcraft"
	"github.com/Ja7ad/fluora/internal/types"
)

type TextCraft struct {
	Rewrite    *textcraft.RewriteInstruction
	Casualize  *textcraft.CasualizeInstruction
	Expand     *textcraft.ExpandInstruction
	WordChoice *textcraft.WordChoiceSuggestionsInstruction
	Formalize  *textcraft.FormalizeInstruction
	Grammar    *textcraft.GrammarInstruction
	Shorten    *textcraft.ShortenInstruction

	gemini *ai.Gemini
}

func NewTextCraft(gemini *ai.Gemini) *TextCraft {
	return &TextCraft{
		Rewrite:    textcraft.NewRewriteInstruction(),
		Casualize:  textcraft.NewCasualizeInstruction(),
		Expand:     textcraft.NewExpandInstruction(),
		WordChoice: textcraft.NewWordChoiceSuggestionsInstruction(),
		Formalize:  textcraft.NewFormalizeInstruction(),
		Grammar:    textcraft.NewGrammarInstruction(),
		Shorten:    textcraft.NewShortenInstruction(),
		gemini:     gemini,
	}
}

func (t *TextCraft) ServiceList() []*dto.ServiceListResp {
	list := make([]*dto.ServiceListResp, 0)
	return append(list,
		&dto.ServiceListResp{
			Title:       t.Rewrite.Title(),
			Description: t.Rewrite.Info(),
			Version:     t.Rewrite.Version(),
			Route:       t.Rewrite.Route(),
			Providers:   t.Rewrite.AIProviders(),
		},
		&dto.ServiceListResp{
			Title:       t.Casualize.Title(),
			Description: t.Casualize.Info(),
			Version:     t.Casualize.Version(),
			Route:       t.Casualize.Route(),
			Providers:   t.Casualize.AIProviders(),
		},
		&dto.ServiceListResp{
			Title:       t.Expand.Title(),
			Description: t.Expand.Info(),
			Version:     t.Expand.Version(),
			Route:       t.Expand.Route(),
			Providers:   t.Expand.AIProviders(),
		},
		&dto.ServiceListResp{
			Title:       t.WordChoice.Title(),
			Description: t.WordChoice.Info(),
			Version:     t.WordChoice.Version(),
			Route:       t.WordChoice.Route(),
			Providers:   t.WordChoice.AIProviders(),
		},
		&dto.ServiceListResp{
			Title:       t.Formalize.Title(),
			Description: t.Formalize.Info(),
			Version:     t.Formalize.Version(),
			Route:       t.Formalize.Route(),
			Providers:   t.Formalize.AIProviders(),
		},
		&dto.ServiceListResp{
			Title:       t.Grammar.Title(),
			Description: t.Grammar.Info(),
			Version:     t.Grammar.Version(),
			Route:       t.Grammar.Route(),
			Providers:   t.Grammar.AIProviders(),
		},
		&dto.ServiceListResp{
			Title:       t.Shorten.Title(),
			Description: t.Shorten.Info(),
			Version:     t.Shorten.Version(),
			Route:       t.Shorten.Route(),
			Providers:   t.Shorten.AIProviders(),
		},
	)
}

func (t *TextCraft) Rewriter(ctx context.Context,
	provider, model, targetLanguage string,
	req *dto.RewriteParam,
) (*dto.RewriteResp, error) {
	if err := t.validateAIProvider(provider, model, t.Rewrite.AIProviders()); err != nil {
		return nil, err
	}

	resp, err := t.gemini.GenerateJson(ctx, &types.GeminiGenerateParams{
		Instruction:     t.Rewrite.Instruction(),
		Text:            req.Text(),
		Model:           model,
		Temperature:     t.Rewrite.Temperature(),
		TopP:            t.Rewrite.TopP(),
		TopK:            t.Rewrite.TopK(),
		MaxOutputTokens: t.Rewrite.MaxOutputToken(),
	})
	if err != nil {
		return nil, err
	}

	var rewriteResp dto.RewriteResp
	if err := json.Unmarshal(resp, &rewriteResp); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	rewriteResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &rewriteResp, nil
}

func (t *TextCraft) validateAIProvider(provider, model string, providers map[types.AIProviders][]string) error {
	models, ok := providers[types.AIProviders(provider)]
	if !ok {
		return fmt.Errorf("no such valid ai provider: %s", provider)
	}

	isValidModel := false

	for _, m := range models {
		if m == model {
			isValidModel = true
		}
	}

	if !isValidModel {
		return fmt.Errorf("no such valid model: %s", model)
	}

	return nil
}
