package repository

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/model"
)

type FizzbuzzPersister interface {
	Persist(ctx context.Context, sequence model.Sequence) error
}
