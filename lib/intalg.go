package lib

import (
	"math"
)

// IsPrimeTrialByError answers
// Question 1
// Loop through all in range 3..n-1 and check mod == 0
func IsPrimeTrialByError(posInteger uint64) bool {

	// candidate prime integers are greater than 1
	if posInteger <= 1 {
		return false
	}

	var u uint64
	for u = 3; u < posInteger; u++ {
		if posInteger%u == 0 {
			// found divisor other 1 or self
			return false
		}
	}

	// includes 2
	return true
}

// IsPrimeOptimized answers
// Question 2
// check mod 2 once to eliminate even numbers iterations
// loop up to square root
func IsPrimeOptimized(posInteger uint64) bool {

	// candidate prime integers are greater than 1
	if posInteger <= 1 {
		return false
	}

	if posInteger == 2 {
		return true
	}

	if posInteger%2 == 0 {
		return false
	}

	// find square root of input
	sqrt := uint64(
		math.Floor(
			math.Sqrt(
				float64(
					posInteger))))

	var u uint64
	for u = 3; u <= sqrt; u += 2 {
		if posInteger%u == 0 {
			return false
		}
	}

	return true
}

// Factors is an alias for an unsigned slice
type Factors []uint64

// Factor answers q3
// By employing optimizations of q2 answer: IsPrimeOptimized
// Reduce by 2 factors
// Factor up to square root
func Factor(posInteger uint64) Factors {
	// 4 is initial capacity, it can grow dynamically
	factors := make(Factors, 0, 4)

	toPrimeTarget := posInteger

	// function helper that reduces toPrimeTarget
	// by factorization and collects prime factors
	reduceByPrimeFactor := func(factor uint64) {
		for ; toPrimeTarget%factor == 0; toPrimeTarget /= factor {
			factors = append(factors, factor)
		}
	}

	// get all two factors
	reduceByPrimeFactor(2)

	// find square root of remaining composite number
	sqrt := uint64(
		math.Floor(
			math.Sqrt(
				float64(
					posInteger))))

	var u uint64
	for u = 3; u <= sqrt && toPrimeTarget > 1; u += 2 {
		reduceByPrimeFactor(u)
	}

	// Remaining unfactored number is prime
	if toPrimeTarget > 1 {
		factors = append(factors, toPrimeTarget)
	}

	return factors
}
