package calculator

import "fmt"

// Divide: Divide two numbers
func Divide(quotient, divisor int) (int, error) {
	if divisor == 0 {
	        return 0, fmt.Errorf("El divisor no puede ser: %d", divisor)
        }

        return quotient / divisor, nil
}

