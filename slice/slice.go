package slice

import "fmt"

func Slicelearn() {
	employee := []string{"rifai", "erni", "andi", "wirahadi"}
	fmt.Println(employee[0:])
	newEmpolyee := employee[0:]
	fmt.Println(newEmpolyee)
	newEmpolyee[1] = "ali"
	fmt.Println(newEmpolyee[0:])
}
