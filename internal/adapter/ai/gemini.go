package ai

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Ja7ad/fluora/pkg/logger"
	"golang.org/x/time/rate"

	"github.com/Ja7ad/fluora/internal/types"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Gemini struct {
	client      *genai.Client
	ratelimiter *rate.Limiter
	log         *logger.SubLogger
}

func newGemini(ctx context.Context, apiKey string, ratelimit int, slog *logger.SubLogger) (Worker, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	g := &Gemini{client: client, ratelimiter: rate.NewLimiter(rate.Limit(ratelimit), 1)}

	if slog == nil {
		slog = logger.NewSubLogger("_gemini", g)
	}

	g.log = slog

	g.log.Debug("Init new gemini client")

	return g, nil
}

func (g *Gemini) GenerateJson(ctx context.Context, params *types.GenerativeParams) ([]byte, error) {
	m := g.client.GenerativeModel(params.Model)
	m.SystemInstruction = genai.NewUserContent(genai.Text(params.Instruction))

	if params.Gemini.IsZero() {
		params.Gemini.Default()
	}

	m.SetTemperature(params.Temperature)
	m.SetTopP(params.TopP)
	m.SetTopK(params.TopK)
	m.SetMaxOutputTokens(params.MaxOutputTokens)
	m.ResponseMIMEType = "application/json"

	g.log.Debug("new generative json request",
		"model", params.Model,
		"temp", params.Temperature,
		"topP", params.TopP,
		"topK", params.TopK,
		"max_output_token", params.MaxOutputTokens,
	)

	resp, err := m.GenerateContent(ctx, genai.Text(params.Text))
	if err != nil {
		return nil, err
	}

	return g.responseToByte(resp)
}

func (g *Gemini) String() string {
	return "gemini"
}

func (g *Gemini) responseToByte(resp *genai.GenerateContentResponse) ([]byte, error) {
	parts := make([]byte, 0)
	for _, candidate := range resp.Candidates {
		for _, part := range candidate.Content.Parts {
			parts = append(parts, []byte(strings.TrimSpace(fmt.Sprint(part)))...)
		}
	}

	if len(parts) == 0 {
		return nil, errors.New("cannot find any parts for gemini response")
	}

	return parts, nil
}
