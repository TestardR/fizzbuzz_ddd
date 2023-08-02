package model

type Fizzbuzz struct {
	int1       int
	str1       string
	int2       int
	str2       string
	limit      int
	upperBound int
}

func NewFizzbuzz(
	int1 int,
	str1 string,
	int2 int,
	str2 string,
	limit int,
) Fizzbuzz {
	return Fizzbuzz{
		int1:  int1,
		str1:  str1,
		int2:  int2,
		str2:  str2,
		limit: limit,
	}
}

func (f Fizzbuzz) Int1() int {
	return f.int1
}

func (f Fizzbuzz) Str1() string {
	return f.str1
}

func (f Fizzbuzz) Int2() int {
	return f.int2
}

func (f Fizzbuzz) Str2() string {
	return f.str2
}

func (f Fizzbuzz) Limit() int {
	return f.limit
}

func (f Fizzbuzz) UpperBound() int {
	return f.upperBound
}
