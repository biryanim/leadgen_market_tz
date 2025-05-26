package service

import (
	"context"
	"github.com/biryanim/leadgenmarket_tz/internal/model"
)

type EncodeService interface {
	EncodeText(ctx context.Context, text *model.NormalStr) (*model.EncodedText, error)
}
