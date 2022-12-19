package array

import "fmt"

func Array() {
	fruits := [4]string{"apple", "orange", "manggo", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("fruits name %v: %s \n", i+1, fruits[i])
	}
}
