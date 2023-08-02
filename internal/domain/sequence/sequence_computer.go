package sequence

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/model"
)

const (
	DynamicProgrammingFormula string = "dynamic_programming"
	ModuloFormula             string = "modulo"
)

type StrategyExecutor interface {
	Compute(ctx context.Context, fizzbuzz model.Fizzbuzz) Sequence
}

type Strategy string

type sequenceComputer struct {
	strategy   Strategy
	strategies map[Strategy]StrategyExecutor
}

func NewSequenceComputer(
	strategy Strategy,
	strategies map[Strategy]StrategyExecutor,
) sequenceComputer {
	return sequenceComputer{
		strategy:   strategy,
		strategies: strategies,
	}
}

func (c sequenceComputer) Compute(ctx context.Context, fizzbuzz model.Fizzbuzz) Sequence {
	return c.strategies[c.strategy].Compute(ctx, fizzbuzz)
}
