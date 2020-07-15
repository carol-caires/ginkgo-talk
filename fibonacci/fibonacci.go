package fibonacci

import "errors"

// F Fibonacci
func F(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("Numero deve ser maior do que zero")
	}
	current, prev := 0, 1
	for i := 0; i < n; i++ {
		current, prev = current+prev, current
	}
	returnErr()
	return current, nil
}

func returnErr() error {
	return nil
}
