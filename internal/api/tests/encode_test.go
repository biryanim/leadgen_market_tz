package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/biryanim/leadgenmarket_tz/internal/api"
	"github.com/biryanim/leadgenmarket_tz/internal/model"
	"github.com/biryanim/leadgenmarket_tz/internal/service"
	serviceMock "github.com/biryanim/leadgenmarket_tz/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEncode(t *testing.T) {
	gin.SetMode(gin.TestMode)
	type encodeServiceMockFunc func(mc *minimock.Controller) service.EncodeService

	type args struct {
		ctx context.Context
		req *http.Request
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		str         = "some_text"
		encodedStr  = "some_encoded_text"
		emptyString = ""
	)

	defer mc.Finish()

	tests := []struct {
		name              string
		args              args
		wantStatus        int
		wantResponse      interface{}
		encodeServiceMock encodeServiceMockFunc
	}{
		{
			name: "success_case",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					reqBody := fmt.Sprintf(`{"text":"%s"}`, str)
					req, _ := http.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer([]byte(reqBody)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus: http.StatusOK,
			wantResponse: api.EncodedTextResponse{
				MD5:    encodedStr,
				SHA256: encodedStr,
			},
			encodeServiceMock: func(mc *minimock.Controller) service.EncodeService {
				mock := serviceMock.NewEncodeServiceMock(mc)
				mock.EncodeTextMock.Expect(ctx, &model.NormalStr{
					Text: str,
				}).Return(&model.EncodedText{
					MD5:    encodedStr,
					SHA256: encodedStr,
				}, nil)

				return mock
			},
		},
		{
			name: "error_case",
			args: args{
				ctx: ctx,
				req: func() *http.Request {
					reqBody := fmt.Sprintf(`{"text":"%s"}`, emptyString)
					req, _ := http.NewRequest(http.MethodPost, "/encode", bytes.NewBuffer([]byte(reqBody)))
					req.Header.Set("Content-Type", "application/json")
					return req
				}(),
			},
			wantStatus:   http.StatusBadRequest,
			wantResponse: `{"error":"invalid request body"}`,
			encodeServiceMock: func(mc *minimock.Controller) service.EncodeService {
				return serviceMock.NewEncodeServiceMock(mc)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encodeServiceMock := tt.encodeServiceMock(mc)
			apiImpl := api.NewImplementation(encodeServiceMock)

			router := gin.New()
			router.POST("/encode", apiImpl.EncodeText)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, tt.args.req)

			require.Equal(t, tt.wantStatus, w.Code)
			if tt.wantStatus == http.StatusOK {
				var response api.EncodedTextResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				require.NoError(t, err)
				require.Equal(t, tt.wantResponse, response)
			} else {
				require.Equal(t, tt.wantResponse, w.Body.String())
			}
		})
	}
}
