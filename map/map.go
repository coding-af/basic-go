package Map

import "fmt"

func MapLearn() {
	var animal = map[string]string{
		"animal1": "hourse",
		"animal2": "mongkey",
	}
	for _, val := range animal {
		fmt.Println(val)
	}

	//combain map slice
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken := range chickens {
		fmt.Println("name:", chicken["name"], "\t", "gender:", chicken["gender"])
	}

}
