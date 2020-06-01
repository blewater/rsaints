package main

/*
 * Quick validation
 * 
 */

import (
	"fmt"
	"log"
	"reflect"
	"time"

	"github.com/blewater/rsaints/lib"
)

func assertBool(funcName string, res, expected bool) {

	if res != expected {
		log.Fatalf("%s resulted in %t, expected %t\n", funcName, res, expected)
	}

}

func assertEqInt64(funcName string, res, expected lib.Integer) {

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

	assertEqFactors("Factor",
		lib.Factor(23), lib.Factors{23})
	assertEqFactors("Factor",
		lib.Factor(26), lib.Factors{2, 13})
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

	assertEqInt64("Euclid",
		lib.CalcEuclid(499017086208, 676126714752),
		93312)

	assertEqInt64("Euclid",
		lib.CalcEuclid(5988737349, 578354589),
		9)

	assertEqInt64("Mod Mult Inverse",
		lib.CalcModInvByEuclid(15, 26),
		7)

	assertEqInt64("Mod Mult Inverse",
		lib.CalcModInvByEuclid(342952340, 4230493243),
		583739113)

	assertBool("Validate RSA Encryption and Decryption",
		lib.CheckRSA(654321, 937513, 638471), true)

	assertBool("Validate RSA Encryption and Decryption",
		lib.CheckRSA(10000, 937513, 638471), true)

	assertBool("Validate RSA Encryption and Decryption",
		lib.CheckRSA(937512, 937513, 638471), true)

	assertBool("Validate RSA Encryption and Decryption",
		lib.CheckRSA(1, 937513, 638471), true)

	fmt.Println("Successful completion of all tests.")
}
