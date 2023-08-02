package validator

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/model"
)

type FizzbuzzValidator struct{}

func (v FizzbuzzValidator) FizzbuzzValid(ctx context.Context, fizzbuzz model.Fizzbuzz) error {
	return nil
}
