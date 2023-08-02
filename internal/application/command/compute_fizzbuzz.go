package command

import "github.com/TestardR/fizzbuzz_ddd/internal/domain/fizzbuzz/model"

type ComputeFizzbuzz struct {
	fizzbuzz model.Fizzbuzz
}

func NewChangeFizzbuzz(fizzbuzz model.Fizzbuzz) ComputeFizzbuzz {
	return ComputeFizzbuzz{fizzbuzz: fizzbuzz}
}

func (c ComputeFizzbuzz) FizzBuzz() model.Fizzbuzz {
	return c.fizzbuzz
}
