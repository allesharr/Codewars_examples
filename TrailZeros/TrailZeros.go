package trailzeros

import (
	"fmt"
	"math/big"
	"time"
)

/*Write a program that will calculate the number of trailing zeros in a factorial of a given number.

N! = 1 * 2 * 3 *  ... * N

Be careful 1000! has 2568 digits...

For more info, see: http://mathworld.wolfram.com/Factorial.html

Examples
zeros(6) = 1
# 6! = 1 * 2 * 3 * 4 * 5 * 6 = 720 --> 1 trailing zero

zeros(12) = 2
# 12! = 479001600 --> 2 trailing zeros
Hint: You're not meant to calculate the factorial. Find another way to find the number of zeros.
*/

func Zeros(n int) int {
	timer := time.Now()
	nn := big.NewInt(int64(n))
	number := factorial(nn)
	var count int = 0
	for {

		fmt.Println("number is", number)
		// key := number % 10
		key := new(big.Int).Mod(number, big.NewInt(10))
		fmt.Println("key is ", key)
		// number = number / 10
		number = new(big.Int).Div(number, big.NewInt(10))
		// fmt.Println("first ", first)

		fmt.Println("result is", number) //number = 0 always as a fisrt variable too
		zero := big.NewInt(0)
		fmt.Println(key.Cmp(zero) != 0)

		if key.Cmp(zero) != 0 {
			break
		} else {
			count++
		}
	}
	newTime := time.Now()
	deltaTime := newTime.Sub(timer)
	fmt.Println(deltaTime)
	return count
}

// function to find factorial
func factorial(n *big.Int) (result *big.Int) {
	//fmt.Println("n = ", n)
	b := big.NewInt(0)
	c := big.NewInt(1)

	if n.Cmp(b) == -1 {
		result = big.NewInt(1)
	}
	if n.Cmp(b) == 0 {
		result = big.NewInt(1)
	} else {
		// return n * factorial(n - 1);
		result = new(big.Int)
		result.Set(n)
		result.Mul(result, factorial(n.Sub(n, c)))
	}
	return
}
