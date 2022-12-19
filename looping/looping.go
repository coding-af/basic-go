package looping

import "fmt"

func Looping() {
	for i := 0; i < 3; i++ {
		fmt.Println("number", i)
	}
	fmt.Println()

	//looping for just condition
	number := 0
	for {
		number++
		if number < 3 {
			fmt.Println("number", number)
		} else {
			break
		}
	}
	fmt.Println()

	//nested looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

}
