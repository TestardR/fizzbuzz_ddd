package fizzbuzz_service

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/application/command"
	"github.com/TestardR/fizzbuzz_ddd/internal/application/query"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/model"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/repository"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/validator"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/sequence"
)

type Service struct {
	fizzbuzzValidator validator.FizzbuzzValidator
	fizzBuzzPersister repository.FizzbuzzPersister
	sequenceComputer  sequence.StrategyExecutor
}

func (s *Service) HandleGetFizzBuzz(ctx context.Context, query query.GetFizzbuzz) (model.Fizzbuzz, error) {
	return model.Fizzbuzz{}, nil
}

func (s *Service) HandleComputeFizzBuzz(ctx context.Context, cmd command.ComputeFizzbuzz) error {
	change, err := model.ComputeFizzbuzz(
		ctx,
		cmd.FizzBuzz(),
		s.fizzbuzzValidator,
	)
	if err != nil {
		return err
	}

	fizzbuzzSequence := s.sequenceComputer.Compute(ctx, change.Fizzbuzz())
	err = s.fizzBuzzPersister.Persist(ctx, fizzbuzzSequence)
	if err != nil {
		return err
	}

	return nil
}
