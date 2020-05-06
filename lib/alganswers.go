package lib

import (
	"math"
)

// Integer is an alias for signed int64
type Integer int64

// IsPrimeOptimized checks mod 2 once to eliminate even iterations as they are multiples of 2.
// Loop up to √n testing odd numbers as n = a * b requires that at least one of a, b is ≤ √n
// (if they divide the number we found an odd factor and the number is composite)
func IsPrimeOptimized(integerN Integer) bool {

	// candidate prime integers are greater than 1
	if integerN <= 1 {
		return false
	}

	if integerN == 2 {
		return true
	}

	if integerN%2 == 0 {
		return false
	}

	// find square root of input
	sqrt := Integer(
		math.Floor(
			math.Sqrt(
				float64(
					integerN))))

	var i Integer
	for i = 3; i <= sqrt; i += 2 {
		if integerN%i == 0 {
			return false
		}
	}

	return true
}

// Factors is an alias for an unsigned slice
type Factors []Integer

// Factor reduces by 2 factors.
// Factor up to square root by testing candidate even factors (if they do they are prime)
func Factor(integerN Integer) Factors {
	// 4 is initial capacity, it can grow dynamically
	factors := make(Factors, 0, 4)

	toPrimeTarget := integerN

	// function helper that reduces toPrimeTarget
	// by factorization and collects prime factors
	reduceByPrimeFactor := func(factor Integer) {
		for ; toPrimeTarget%factor == 0; toPrimeTarget /= factor {
			factors = append(factors, factor)
		}
	}

	// get all two factors
	reduceByPrimeFactor(2)

	// find square root of remaining composite number
	sqrt := Integer(
		math.Floor(
			math.Sqrt(
				float64(
					integerN))))

	var u Integer
	for u = 3; u <= sqrt && toPrimeTarget > 1; u += 2 {
		reduceByPrimeFactor(u)
	}

	// Remaining unfactored number is prime
	if toPrimeTarget > 1 {
		factors = append(factors, toPrimeTarget)
	}

	return factors
}

// CalcEuclid computes GCD(A,B) by employing the Euclid's algorithm.
// Euclid's algorithm performs successive a,b substitutions such a <- b, b <- a divModulo b
// till we reduce remainder to 0.
// The algorithm applies substitutions because GCD(a,b) is equivalent as GCD(b, remainder of a/b)
// because a = b + remainder a/b
func CalcEuclid(integerA, integerB Integer) Integer {

	for integerB != 0 {
		integerA, integerB =
			getReducedTermsByEuclidFormula(integerA, integerB)
	}

	return integerA
}

// getReducedTermsByEuclidFormula replaces (a,b) with (b, remainder of a/b)
func getReducedTermsByEuclidFormula(integerA, integerB Integer) (a Integer, b Integer) {
	return integerB, integerA % integerB
}

// CalcModInvByEuclid to calculate a's multiplicative inverse x mod m or 1/a mod m so that
// a * x mod m results in 1
// by employing Euclid's algorithm (extended or reverse direction).
// To answer x, we solve
// Bezout identity formula (theorem) ax + my = gcd(a, m) = 1 (one for modular inverse to exist)
// to calculate x.
// Y's final value is ignored in modular inverse because
// ax = 1 + m*-y (applying mod m) on both sides results ->
// ax mod m = 1 mod m or ax mod m = 1.
// To solve x, y in ax + my = 1 we express x,y as integer combination of a, m equal to GCD of 1
// by applying coefficients of the terms involving the remainders
// from the a = qm+r relationship => r = a -qm.
// Thus working back through the equations generated by the Euclidean algorithm
// gcd(a, m) 	= 1 =>
//		= rprev - qr...
// arriving at gcd(a,m) = xa + ym
// Ref: https://proofwiki.org/wiki/Euclidean_Algorithm
func CalcModInvByEuclid(a, m Integer) Integer {
	var x Integer = 1
	var y Integer = 0
	var quotient Integer

	modBase := m

	for a > 1 {
		quotient = a / modBase

		// Apply Euclid's substitution step
		a, modBase = getReducedTermsByEuclidFormula(a, modBase)
		x, y = y, x-quotient*y

	}

	// If x turns negative we reinstate it positive in the modulo ring
	if x < 0 {
		x = x + m
	}

	return x
}

// CheckRSA applies RSA Encryption and subsequent decryption
// and validating that original message m is same to decrypted message m2
// Steps:
// 1 Encrypt m^e mod n
// 2 Performs prime factorization of n.
// 3 Calculates phi(n)
// 4 Calculates d such that d*e mod phi(n) = 1
// 5 Decrypts c -> m2 by c^d mod n
// 6 if m == m2 then it's a successful reversal
func CheckRSA(m, n, e Integer) bool {
	c := getModOfPow(m, e, n)

	factors := Factor(n)

	phiOfN := getPhin(factors[0], factors[1])

	d := CalcModInvByEuclid(e, phiOfN)

	m2 := getModOfPow(c, d, n)

	return m == m2
}

// getPhi calculates the Euler totient value for n
func getPhin(p, q Integer) Integer {
	return (p - 1) * (q - 1)
}

// getModOfPow encrypts or decrypts a value for the RSA algorithm using a simple
// linear complexity Oexp modular exponeniation algorithm such that
// i^exp mod n => ((r1=1^i mod n) * (r2=r1^i mod n)* ... * (rexp^i mod n)) mod n
// if integer is a message m and pow is e then it returns c mod n the encrypted message per RSA
// if integer is an encrypted message c and pow is d then it returns m mod n the decrypted message per RSA
func getModOfPow(integer, exponent, n Integer) Integer {
	var res Integer = 1
	for i := res; i <= exponent; i++ {
		res = (res * integer) % n
	}
	return res
}
