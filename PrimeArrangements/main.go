package main

import (
	"fmt"

)


func numPrimeArrangements(n int) int {
	nPrimeCount := 0
	for i := 1 ; i <= n ; i++{
		if isPrime(i){
			nPrimeCount++
		}
	}
	
	nResult := 1

	nNormal := n-nPrimeCount
	for nNormal > 0{
		nResult*= nNormal
		nResult = nResult % (1000000007)
		nNormal--
	}

	for nPrimeCount > 0{
		nResult*= nPrimeCount
		nResult = nResult % (1000000007)
		nPrimeCount--
	}

    return nResult
}

func isPrime(n int) bool {
    if n <=3 {
        return n > 1
    }
    
	if (n % 2 == 0) || (n % 3 == 0) {
		return false
	}
	
	for i := 5 ; (i*i) <= n ; i+=6 {
		if (n % i == 0) || (n % (i+2) == 0) {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(numPrimeArrangements(100))
}