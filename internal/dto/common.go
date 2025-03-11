package dto

import (
	"github.com/Ja7ad/fluora/internal/types"
)

type Header struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
}

type ServiceListResp struct {
	Title       string                                   `json:"title"`
	Description string                                   `json:"description"`
	Version     string                                   `json:"version"`
	Route       string                                   `json:"route"`
	Providers   map[types.AIProvider]map[string]struct{} `json:"providers"`
}
