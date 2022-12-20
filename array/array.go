package array

import (
	"fmt"
	"strings"
)

func Array() {
	fruits := [4]string{"apple", "orange", "manggo", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

	temp, params := fillterString("fuck")
	if temp {
		fmt.Println("***")
	} else {
		fmt.Println(params)
	}

}

func fillterString(params string) (value bool, data string) {
	dataString := "dog animal dick fuck"
	var isExists = strings.Contains(dataString, params)

	return isExists, params
}
