package main

func subtractProductAndSum(n int) int {
	Sum := 0
	Product := 1

	for n != 0 {
		Sum += n % 10
		Product *= n % 10
		n /= 10
	}

	return Product - Sum
}

func main() {

}
