package redis_cache

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/sequence"
	redisCache "github.com/TestardR/fizzbuzz_ddd/internal/infrastructure/shared/redis_cache"
)

const (
	maxKey  = "max_key"
	maxHits = "max_hits"
)

type sequenceStore struct {
	redis *redisCache.Client
}

func NewSequenceStore(
	redis *redisCache.Client,
) *sequenceStore {
	return &sequenceStore{
		redis: redis,
	}
}

func (s *sequenceStore) Persist(ctx context.Context, fizzbuzzSequence sequence.Sequence) error {
	return nil
}
