package lib

import (
	"math"
)

// IsPrimeTrialByError answers
// Question 1
// Loop through all in range 3..n-1 and check mod == 0
func IsPrimeTrialByError(posIntegerN uint64) bool {

	// candidate prime integers are greater than 1
	if posIntegerN <= 1 {
		return false
	}

	if posIntegerN == 2 {
		return true
	}

	var u uint64
	for u = 3; u < posIntegerN; u++ {
		if posIntegerN%u == 0 {
			// found divider other 1,2 or self
			return false
		}
	}

	return true
}

// IsPrimeOptimized answers
// Question 2
// Check mod 2 once to eliminate remaining even iterations as they are multiples of 2.
// loop up to √n testing odd numbers as n = a * b requires that at least one of a, b is ≤ √n
// (if they divide the number we found an odd factor and the number is composite)
func IsPrimeOptimized(posIntegerN uint64) bool {

	// candidate prime integers are greater than 1
	if posIntegerN <= 1 {
		return false
	}

	if posIntegerN == 2 {
		return true
	}

	if posIntegerN%2 == 0 {
		return false
	}

	// find square root of input
	sqrt := uint64(
		math.Floor(
			math.Sqrt(
				float64(
					posIntegerN))))

	var u uint64
	for u = 3; u <= sqrt; u += 2 {
		if posIntegerN%u == 0 {
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
// Factor up to square root by testing candidate even factors (if they do they are prime)
func Factor(posIntegerN uint64) Factors {
	// 4 is initial capacity, it can grow dynamically
	factors := make(Factors, 0, 4)

	toPrimeTarget := posIntegerN

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
					posIntegerN))))

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

// CalcEuclid answers Q4
// of computing GCD(A,B) by employing the Euclid's algorithm.
// Euclid's algorithm performs successive a,b substitutions such a <- b, b <- a divModulo b
// till we reduce remainder to 0
// Because GCD(a,b) is same as GCD(b, remainder of a/b)
func CalcEuclid(posIntegerA, posIntegerB int64) int64 {

	for posIntegerB != 0 {
		posIntegerA, posIntegerB =
			getReducedTermsByEuclidFormula(posIntegerA, posIntegerB)
	}

	return posIntegerA
}

// Replace (a,b) with (b, remainder of a/b)
func getReducedTermsByEuclidFormula(posIntegerA, posIntegerB int64) (a int64, b int64) {
	return posIntegerB, posIntegerA % posIntegerB
}

// CalcModInvByEuclid answers Q5
// to calculate a's multiplicative inverse x mod n or 1/a mod n so that
// a * x mod n results in 1
// by employing Euclid's algorithm.
// We presume a, n are co-primes or that gcd(a, n) = 1 for the inverse to exist.
// To answer x, we solve
// Bezout identity formula ax + my = gcd(a, m) = 1 (one for modulo inverse to exist)
// to calculate x. Y's final value is ignored in modulo inverse.
func CalcModInvByEuclid(a, m int64) int64 {
	var x int64 = 1
	var y int64 = 0
	var quotient int64

	modBase := m

	if modBase == 1 {
		return 0
	}

	for a > 1 {
		quotient = a / modBase

		x, y = y, x-quotient*y

		// Use one substitution step of Euclid's algorithm
		a, modBase = getReducedTermsByEuclidFormula(a, modBase)
	}

	// If x turns negative we reinstate it positive in the modulo ring
	if x < 0 {
		x = x + m
	}

	return x
}
