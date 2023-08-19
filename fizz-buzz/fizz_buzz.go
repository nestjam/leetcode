package fizzbuzz

import "fmt"

const (
    fizz = "Fizz"
    buzz = "Buzz"
)

func fizzBuzz(n int) []string {
	var result []string = make([]string, n)

	for i := 0; i < n; i++ {  
		result[i] = print(i + 1)
	}

	return result
}

func print(n int) string {
	divisibleBy3 := n % 3 == 0
    divisibleBy5 := n % 5 == 0
    
    if divisibleBy5 && divisibleBy3 {
        return fizz + buzz
    } else if divisibleBy5 {
		return buzz
	} else if divisibleBy3 {
		return fizz
	} else {
		return fmt.Sprintf("%v", n)
	}
}
