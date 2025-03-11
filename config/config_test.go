package config

import (
	"testing"

	"github.com/Ja7ad/fluora/pkg/logger"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_Success(t *testing.T) {
	cfg, err := LoadConfig("./example.config.yml")

	assert.NoError(t, err, "Expected LoadConfig to succeed")
	assert.NotNil(t, cfg, "Config should not be nil")
	assert.NotNil(t, cfg.Rest, "Rest config should not be nil")
	assert.Equal(t, "127.0.0.1:8080", cfg.Rest.Address, "Incorrect REST Address")
	assert.NotEmpty(t, cfg.AI, "AI config should not be empty")
	assert.Equal(t, "gemini", cfg.AI[0].Provider, "Incorrect AI Provider")
	assert.NotNil(t, cfg.Logger, "Logger config should not be nil")
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	cfg, err := LoadConfig("non_existent_config.yml")

	assert.Error(t, err, "Expected error for missing config file")
	assert.Nil(t, cfg, "Config should be nil when file is not found")
}

func TestBasicCheck_Success(t *testing.T) {
	cfg := &Config{
		Rest: &Rest{
			Address: "127.0.0.1:8080",
			Origins: []string{"http://example.com"},
		},
		AI: []*AI{
			{
				Provider:  "gemini",
				Serve:     "api.gemini.com",
				APIKey:    "test-api-key",
				RateLimit: 5,
			},
		},
		Logger: logger.DefaultConfig(),
	}

	// ✅ **Assertions**
	assert.NoError(t, cfg.BasicCheck(), "Expected BasicCheck to pass with valid config")
}

func TestBasicCheck_MissingAI(t *testing.T) {
	cfg := &Config{
		Rest: &Rest{
			Address: "localhost:8080",
			Origins: []string{"*"},
		},
		Logger: logger.DefaultConfig(),
	}

	err := cfg.BasicCheck()

	assert.Error(t, err, "Expected error for missing AI config")
	assert.Equal(t, "no AI configuration found", err.Error(), "Incorrect error message")
}

func TestBasicCheck_MissingREST(t *testing.T) {
	cfg := &Config{
		AI: []*AI{
			{
				Provider:  "gemini",
				Serve:     "api.gemini.com",
				APIKey:    "test-api-key",
				RateLimit: 5,
			},
		},
		Logger: logger.DefaultConfig(),
	}

	err := cfg.BasicCheck()

	assert.Error(t, err, "Expected error for missing REST config")
	assert.Equal(t, "no REST configuration found", err.Error(), "Incorrect error message")
}

func TestBasicCheck_MissingRESTAddress(t *testing.T) {
	cfg := &Config{
		Rest: &Rest{
			Origins: []string{"*"},
		},
		AI: []*AI{
			{
				Provider:  "gemini",
				Serve:     "api.gemini.com",
				APIKey:    "test-api-key",
				RateLimit: 5,
			},
		},
		Logger: logger.DefaultConfig(),
	}

	err := cfg.BasicCheck()

	assert.Error(t, err, "Expected error for missing REST address")
	assert.Equal(t, "no REST configuration found", err.Error(), "Incorrect error message")
}
