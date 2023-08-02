package model

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/validator"
)

type ComputeFizzbuzzChange struct {
	fizzbuzz Fizzbuzz
}

func NewComputeFizzbuzzChange(fizzbuzz Fizzbuzz) ComputeFizzbuzzChange {
	return ComputeFizzbuzzChange{fizzbuzz: fizzbuzz}
}

func (c ComputeFizzbuzzChange) Fizzbuzz() Fizzbuzz {
	return c.fizzbuzz
}

func ComputeFizzbuzz(
	ctx context.Context,
	fizzbuzz Fizzbuzz,
	validator validator.FizzbuzzValidator,
) (ComputeFizzbuzzChange, error) {
	err := validator.FizzbuzzValid(ctx, fizzbuzz)
	if err != nil {
		return ComputeFizzbuzzChange{}, err
	}

	return NewComputeFizzbuzzChange(fizzbuzz), nil
}
