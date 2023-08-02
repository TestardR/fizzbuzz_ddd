package statistics_service

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/application/query"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/statistics/model"
)

type Service struct{}

func (s *Service) HandleGetStatistics(ctx context.Context, query query.GetStatistics) (model.Statistics, error) {
	return model.Statistics{}, nil
}
