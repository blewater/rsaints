package lib

import "fmt"

// Y is the 3rd key component of pub key
func Y(g, x, p Integer) Integer {

	if x <= 1 || x >= (p-1) {
		fmt.Println("Cannot continue, x : ", x, " has to be > 1 and < p-1 : ", p-1)
		return 0
	}

	gIsRoot := false
	roots, n := getRoots(p)
	if n == 0 {
		fmt.Println("Warning no roots found for ", p, " is p Prime?")
	}
	for _, root := range roots {
		if root == g {
			gIsRoot = true
			break
		}
	}

	if !gIsRoot {
		fmt.Println("Warning g : ", g, " is not a root : ", roots)
	}

	y := GetModOfPow(g, x, p)

	fmt.Printf("Public key is (%d, %d, %d)\n", p, g, y)

	return y
}

func GetCyphers(y, g, k, p, m Integer) (Integer, Integer) {
	c1 := GetModOfPow(g, k, p)
	c2 := (GetModOfPow(y, k, p) * m) % p

	return c1, c2
}

func getC2DivByYk(c2, yk, p Integer) Integer {
	ykInv := CalcModInvByEuclid(yk, p)

	return (c2 * ykInv) % p
}

func GetMessageByK(k, y, c2, p Integer) Integer {
	yk := GetModOfPow(y, k, p)

	return getC2DivByYk(c2, yk, p)
}

func GetMessageByX(c1, c2, x, p Integer) Integer {
	// or GetModOfPow(g, k, p), GetModOfPow(gk, x, p)
	yk := GetModOfPow(c1, x, p)

	return getC2DivByYk(c2, yk, p)
}
