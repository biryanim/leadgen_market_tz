package encoder

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/biryanim/leadgenmarket_tz/internal/model"
	"strings"
)

func (s *serv) EncodeText(ctx context.Context, text *model.NormalStr) (*model.EncodedText, error) {
	cached, err := s.redisClient.Get(ctx, text.Text)
	if err != nil {
		return nil, fmt.Errorf("redis get failed: %w", err)
	}

	if cached != nil {
		var cachedStr string

		switch v := cached.(type) {
		case []byte:
			cachedStr = string(v)
		case string:
			cachedStr = v
		default:
			return nil, fmt.Errorf("unexpected cache type: %T", cached)
		}
		parts := strings.Split(cachedStr, " ")
		if len(parts) == 2 {
			return &model.EncodedText{
				MD5:    parts[0],
				SHA256: parts[1],
			}, nil
		}
	}

	md5hash := md5.Sum([]byte(text.Text))
	sha256hash := sha256.Sum256([]byte(text.Text))

	result := &model.EncodedText{
		MD5:    hex.EncodeToString(md5hash[:]),
		SHA256: hex.EncodeToString(sha256hash[:]),
	}

	if err = s.redisClient.Set(ctx, text.Text, result, 0); err != nil {
		return nil, fmt.Errorf("redis set failed: %w", err)
	}

	return result, nil
}
