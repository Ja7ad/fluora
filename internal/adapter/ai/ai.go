package ai

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/Ja7ad/fluora/internal/types"
	"github.com/Ja7ad/fluora/pkg/logger"

	"github.com/Ja7ad/fluora/config"
)

type AI struct {
	mu         sync.Mutex
	geminiPool []Worker
	index      int
	logger     *logger.SubLogger
}

type Worker interface {
	GenerateJson(ctx context.Context, params *types.GenerativeParams) ([]byte, error)
}

func NewAIFromConfig(ctx context.Context, providers []*config.AI) (*AI, error) {
	ai := &AI{
		geminiPool: make([]Worker, 0),
		index:      -1,
	}

	ai.logger = logger.NewSubLogger("_ai", ai)

	for _, p := range providers {
		switch p.Provider {
		case "gemini":
			g, err := newGemini(ctx, p.APIKey, p.RateLimit, ai.logger)
			if err != nil {
				return nil, err
			}
			ai.geminiPool = append(ai.geminiPool, g)
		default:
			return nil, fmt.Errorf("unknown provider %s", p.Provider)
		}
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	return ai, nil
}

func (a *AI) String() string {
	return "AI Adapter"
}

func (a *AI) Worker(provider types.AIProvider) Worker {
	switch provider {
	case "gemini":
		return a.Gemini()
	default:
		return nil
	}
}

func (a *AI) Gemini() Worker {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.geminiPool) == 1 {
		return a.geminiPool[0]
	}

	if a.index == -1 {
		a.index = rand.Intn(len(a.geminiPool))
	} else {
		a.index = (a.index + 1) % len(a.geminiPool)
	}

	return a.geminiPool[a.index]
}
