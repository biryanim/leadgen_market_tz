package api

import "github.com/biryanim/leadgenmarket_tz/internal/model"

type TextRequest struct {
	Text string `json:"text" binding:"required"`
}

func (t *TextRequest) ToModel() *model.NormalStr {
	return &model.NormalStr{
		Text: t.Text,
	}
}

type EncodedTextResponse struct {
	MD5    string `json:"md5"`
	SHA256 string `json:"sha256"`
}

func (e *EncodedTextResponse) FromModel(model *model.EncodedText) {
	e.MD5 = model.MD5
	e.SHA256 = model.SHA256
}
