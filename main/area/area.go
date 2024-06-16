package area

import "fmt"

const PI float32 = 3.1415

func Area(shape string, args int32) float32 {
	fmt.Println(shape)
	fmt.Println(args)
	if shape == "circle" {
		return PI * float32(args*args)
	}
	return float32(args)
}
