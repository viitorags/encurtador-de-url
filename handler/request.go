package handler

import (
	"fmt"
)

func errParamRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type LinkRequest struct {
	OriginalURL string `json:"original_url"`
	ShortUrl    string `json:"short_url"`
}
