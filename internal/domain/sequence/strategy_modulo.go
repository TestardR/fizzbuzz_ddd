package sequence

import (
	"context"
	"github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/model"
	"strconv"
)

type Modulo struct{}

func (m Modulo) Compute(ctx context.Context, fizzbuzz model.Fizzbuzz) Sequence {
	limit := fizzbuzz.UpperBound()
	if fizzbuzz.UpperBound() > fizzbuzz.Limit() {
		limit = fizzbuzz.Limit()
	}

	sequence := make([]string, 0, limit)

	for i := 1; i <= limit; i++ {
		mod1 := fizzbuzz.Int1()
		mod2 := fizzbuzz.Int2()

		switch {
		case i%mod1 == 0 && i%mod2 == 0:
			sequence = append(sequence, fizzbuzz.Str1()+fizzbuzz.Str2())
		case i%mod1 == 0:
			sequence = append(sequence, fizzbuzz.Str1())
		case i%mod2 == 0:
			sequence = append(sequence, fizzbuzz.Str2())
		default:
			sequence = append(sequence, strconv.Itoa(i))
		}
	}

	return sequence
}
