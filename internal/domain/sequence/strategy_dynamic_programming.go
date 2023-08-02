package sequence

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/model"
)

type DynamicProgramming struct{}

func (d DynamicProgramming) Compute(ctx context.Context, fizzbuzz model.Fizzbuzz) Sequence {
	return Sequence{}
}
