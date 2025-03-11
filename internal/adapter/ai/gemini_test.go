package ai

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Ja7ad/fluora/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGemini_GenerateJson_Integration(t *testing.T) {
	apiKey := os.Getenv("GEMINI_API_KEY")

	if apiKey == "" {
		t.Skip("Skipping test: GEMINI_API_KEY is not set")
	}

	ctx := context.Background()

	client, err := newGemini(ctx, apiKey, 5, nil)
	require.NoError(t, err)
	require.NotNil(t, client)

	params := &types.GenerativeParams{
		Model:       "gemini-1.5-pro",
		Instruction: "Rewrite the following sentence in a more formal tone.",
		Text:        "Hey, can you send me the files asap?",
		Gemini: types.Gemini{
			Temperature:     0.7,
			TopP:            0.9,
			TopK:            40,
			MaxOutputTokens: 256,
		},
	}

	resp, err := client.GenerateJson(ctx, params)

	assert.NoError(t, err, "Expected no error from GenerateJson()")
	assert.NotNil(t, resp, "Response should not be nil")
}
