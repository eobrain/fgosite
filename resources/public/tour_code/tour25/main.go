package main

epsilon := 0.0001

func Sqrt(x float64) float64 {

	loop (z = x / 2) {
		if Math::abs(z * z - x) < epsilon  {
			z
		} else {
			recur(z - (z * z - x) / (2 * x))
		}
	}

}

func main() {
	printf("%7.4f", Sqrt(100))
}
