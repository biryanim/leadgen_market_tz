package converter

import (
	"github.com/biryanim/leadgenmarket_tz/internal/api"
	"github.com/biryanim/leadgenmarket_tz/internal/model"
)

func FromDTOReqToModel(dto *api.TextRequest) *model.NormalStr {
	return &model.NormalStr{
		Text: dto.Text,
	}
}

func FromEncodedToResp(model *model.EncodedText) *api.EncodedTextResponse {
	return &api.EncodedTextResponse{
		MD5:    model.MD5,
		SHA256: model.SHA256,
	}
}
