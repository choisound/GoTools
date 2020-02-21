package main

const boilingF = 212.0

func FtoC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// func main() {
// 	const f = boilingF
// 	fmt.Printf("%f %f\n", f, FtoC(f))
// 	fmt.Printf("%+v %+v", f, FtoC(f))
// }
