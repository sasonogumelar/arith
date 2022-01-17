package arith

type Init struct {
	A, B float64
}

func (e *Init) Add() float64 {
	return e.A + e.B
}
