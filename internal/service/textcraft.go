package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Ja7ad/fluora/internal/errors"
	"github.com/Ja7ad/fluora/pkg/logger"
	"github.com/gofiber/fiber/v2"

	"github.com/Ja7ad/fluora/internal/adapter/ai"
	"github.com/Ja7ad/fluora/internal/dto"
	"github.com/Ja7ad/fluora/internal/instructions/textcraft"
	"github.com/Ja7ad/fluora/internal/types"
)

type TextCraft struct {
	services map[string]types.Instructor
	ai       *ai.AI
	logger   *logger.SubLogger
}

func NewTextCraft(ai *ai.AI) *TextCraft {
	services := map[string]types.Instructor{
		"rewrite":    textcraft.NewRewriteInstruction(),
		"casualize":  textcraft.NewCasualizeInstruction(),
		"expand":     textcraft.NewExpandInstruction(),
		"wordChoice": textcraft.NewWordChoiceSuggestionsInstruction(),
		"formalize":  textcraft.NewFormalizeInstruction(),
		"grammar":    textcraft.NewGrammarInstruction(),
		"shorten":    textcraft.NewShortenInstruction(),
	}

	svc := &TextCraft{
		services: services,
		ai:       ai,
	}

	svc.logger = logger.NewSubLogger("_service", svc)

	return svc
}

func (t *TextCraft) ServiceList() []*dto.ServiceListResp {
	list := make([]*dto.ServiceListResp, 0, len(t.services))

	for _, service := range t.services {
		list = append(list, &dto.ServiceListResp{
			Title:       service.Title(),
			Description: service.Info(),
			Version:     service.Version(),
			Route:       service.Route(),
			Providers:   service.AIProviders(),
		})
	}

	return list
}

func (t *TextCraft) String() string {
	return "TextCraft Service"
}

func (t *TextCraft) Service(key string) types.Instructor {
	return t.services[key]
}

func (t *TextCraft) Rewriter(ctx context.Context, provider, model string,
	req *dto.RewriteParam,
) (*dto.RewriteResp, error) {
	resp, err := t.generateText(ctx, "rewrite", provider, model, req.Text())
	if err != nil {
		return nil, err
	}

	var rewriteResp dto.RewriteResp
	if err := json.Unmarshal(resp, &rewriteResp); err != nil {
		t.logger.Error("Failed to unmarshal rewrite response", "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	rewriteResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &rewriteResp, nil
}

func (t *TextCraft) Casualize(ctx context.Context, provider, model string,
	req *dto.CasualizeParam,
) (*dto.CasualizeResp, error) {
	resp, err := t.generateText(ctx, "casualize", provider, model, req.Text())
	if err != nil {
		return nil, err
	}

	var casualizeResp dto.CasualizeResp
	if err := json.Unmarshal(resp, &casualizeResp); err != nil {
		t.logger.Error("Failed to unmarshal casualize response", "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	casualizeResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &casualizeResp, nil
}

func (t *TextCraft) Expand(ctx context.Context, provider, model string,
	req *dto.ExpandParam,
) (*dto.ExpandResp, error) {
	resp, err := t.generateText(ctx, "expand", provider, model, req.Text())
	if err != nil {
		return nil, err
	}

	var expandResp dto.ExpandResp
	if err := json.Unmarshal(resp, &expandResp); err != nil {
		t.logger.Error("Failed to unmarshal expand response", "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	expandResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &expandResp, nil
}

func (t *TextCraft) Formalize(ctx context.Context, provider, model string,
	req *dto.FormalizeParam,
) (*dto.FormalizeResp, error) {
	resp, err := t.generateText(ctx, "formalize", provider, model, req.Text())
	if err != nil {
		return nil, err
	}

	var formalizeResp dto.FormalizeResp
	if err := json.Unmarshal(resp, &formalizeResp); err != nil {
		t.logger.Error("Failed to unmarshal expand response", "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	formalizeResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &formalizeResp, nil
}

func (t *TextCraft) Grammar(ctx context.Context, provider, model string,
	req *dto.GrammarParam,
) (*dto.GrammarResp, error) {
	resp, err := t.generateText(ctx, "grammar", provider, model, req.Text())
	if err != nil {
		return nil, err
	}

	var grammarResp dto.GrammarResp
	if err := json.Unmarshal(resp, &grammarResp); err != nil {
		t.logger.Error("Failed to unmarshal expand response", "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	grammarResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &grammarResp, nil
}

func (t *TextCraft) Shorten(ctx context.Context, provider, model string,
	req *dto.ShortenParam,
) (*dto.ShortenResp, error) {
	resp, err := t.generateText(ctx, "shorten", provider, model, req.Text())
	if err != nil {
		return nil, err
	}

	var shortenResp dto.ShortenResp
	if err := json.Unmarshal(resp, &shortenResp); err != nil {
		t.logger.Error("Failed to unmarshal expand response", "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	shortenResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &shortenResp, nil
}

func (t *TextCraft) WordChoice(ctx context.Context, provider, model string,
	req *dto.WordChoiceParam,
) (*dto.WordChoiceResp, error) {
	resp, err := t.generateText(ctx, "wordChoice", provider, model, req.Text())
	if err != nil {
		return nil, err
	}

	var wordChoiceResp dto.WordChoiceResp
	if err := json.Unmarshal(resp, &wordChoiceResp); err != nil {
		t.logger.Error("Failed to unmarshal expand response", "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	wordChoiceResp.Header = &dto.Header{
		Provider: provider,
		Model:    model,
	}

	return &wordChoiceResp, nil
}

func (t *TextCraft) generateText(ctx context.Context, key, provider, model string, text string) ([]byte, error) {
	service := t.Service(key)

	if err := service.ValidateProvider(provider, model); err != nil {
		return nil, errors.New(fiber.StatusBadRequest, "invalid choice provider or model", nil)
	}

	worker, err := t.selectAIWorker(provider)
	if err != nil {
		return nil, err
	}

	params := &types.GenerativeParams{
		Instruction: service.Instruction(),
		Text:        text,
		Model:       model,
	}

	switch types.AIProvider(provider) {
	case types.GEMINI:
		params.Gemini = types.Gemini{
			Temperature:     service.Temperature(),
			TopP:            service.TopP(),
			TopK:            service.TopK(),
			MaxOutputTokens: service.MaxOutputToken(),
		}
	default:
	}

	resp, err := worker.GenerateJson(ctx, params)
	if err != nil {
		t.logger.Error("AI generation failed", "provider", provider, "model", model, "error", err)
		return nil, errors.New(fiber.StatusInternalServerError, "internal server error occurred", nil)
	}

	return resp, nil
}

func (t *TextCraft) selectAIWorker(provider string) (ai.Worker, error) {
	w := t.ai.Worker(types.AIProvider(provider))
	if w == nil {
		return nil, errors.New(fiber.StatusBadRequest, fmt.Sprintf("unsupported provider: %s", provider),
			nil)
	}
	return w, nil
}
