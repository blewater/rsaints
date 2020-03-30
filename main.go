package main

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/comp530/lib"
)

func assertBool(funcName string, res, expected bool) {

	if res != expected {
		log.Fatalf("%s resulted in %t, expected %t\n", funcName, res, expected)
	}

}

func assertEqInt64(funcName string, res, expected int64) {

	if res != expected {
		log.Fatalf("%s resulted in %d, expected %d\n", funcName, res, expected)
	}

}

func assertEqFactors(funcName string, res, expected lib.Factors) {

	if !reflect.DeepEqual(res, expected) {
		log.Fatalf("%s resulted in %v, expected %v\n", funcName, res, expected)
	}

}

func main() {

	// Q1 Tests
	assertBool("IsPrimeTrialByError",
		lib.IsPrimeTrialByError(5), true)
	assertBool("IsPrimeTrialByError",
		lib.IsPrimeTrialByError(23), true)
	assertBool("IsPrimeTrialByError",
		lib.IsPrimeTrialByError(81), false)

	// Q2 Tests
	assertBool("IsPrimeOptimized",
		lib.IsPrimeOptimized(5), true)
	assertBool("IsPrimeOptimized",
		lib.IsPrimeOptimized(23), true)
	assertBool("IsPrimeOptimized",
		lib.IsPrimeOptimized(81), false)
	// non-prime number: 7 * 157 * 8365633
	assertBool("IsPrimeOptimized",
		lib.IsPrimeOptimized(9193830667), false)
	// known 10 digit prime number < 0.5 sec
	start := time.Now()
	res := lib.IsPrimeOptimized(9576890767)
	fmt.Printf("large prime number check took %v\n", time.Since(start))
	assertBool("IsPrimeOptimized", res, true)

	// Q3 Tests
	assertEqFactors("Factor",
		lib.Factor(23), lib.Factors{23})
	assertEqFactors("Factor",
		lib.Factor(81), lib.Factors{3, 3, 3, 3})
	assertEqFactors("Factor",
		lib.Factor(150), lib.Factors{2, 3, 5, 5})
	assertEqFactors("Factor",
		lib.Factor(147), lib.Factors{3, 7, 7})
	assertEqFactors("Factor",
		lib.Factor(150), lib.Factors{2, 3, 5, 5})
	assertEqFactors("Factor",
		lib.Factor(330), lib.Factors{2, 3, 5, 11})
	// non-prime number: 7 * 157 * 8365633
	assertEqFactors("Factor",
		lib.Factor(9193830667), lib.Factors{7, 157, 8365633})
	// known 10 digit prime number
	assertEqFactors("Factor", lib.Factor(9576890767),
		lib.Factors{9576890767})

	// Q4 Tests
	assertEqInt64("Euclid",
		lib.CalcEuclid(499017086208, 676126714752),
		93312)

	assertEqInt64("Euclid",
		lib.CalcEuclid(5988737349, 578354589),
		9)

	// Q5 Tests
	assertEqInt64("Mod Mult Inverse",
		lib.CalcModInvByEuclid(15, 26),
		7)

	assertEqInt64("Mod Mult Inverse",
		lib.CalcModInvByEuclid(15, 26),
		7)

	assertEqInt64("Mod Mult Inverse",
		lib.CalcModInvByEuclid(342952340, 4230493243),
		583739113)

}
