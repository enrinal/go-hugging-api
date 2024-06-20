package hugging

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/hupe1980/go-huggingface"
	"github.com/redis/go-redis/v9"
	"time"
)

func (h *Hugging) TextClassification(ctx context.Context, req huggingface.TextClassificationRequest) (*huggingface.TextClassificationResponse, error) {
	ic := huggingface.NewInferenceClient(h.Token)

	hash := sha1.New()
	hash.Write([]byte(req.Inputs))
	key := hex.EncodeToString(hash.Sum(nil))

	cacheData, err := h.Rdb.Get(ctx, key).Result()

	if err == redis.Nil {
		res, err := ic.TextClassification(ctx, &huggingface.TextClassificationRequest{
			Inputs: req.Inputs,
			Model:  "SamLowe/roberta-base-go_emotions",
		})
		if err != nil {
			return nil, err
		}
		go func() {
			jsonData, _ := json.Marshal(res)
			h.Rdb.Set(context.Background(), key, jsonData, 60*time.Minute)
		}()
		return &res, nil
	} else if err != nil {
		return nil, err
	} else {
		var res huggingface.TextClassificationResponse
		json.Unmarshal([]byte(cacheData), &res)
		return &res, nil
	}
}
