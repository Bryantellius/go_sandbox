package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println(math.Pi)

	float := 1.56

	fmt.Printf("Ceil: %f, Floor: %f, Round: %f\n", math.Ceil(float), math.Floor(float), math.Round(float))

	x := 10.0

	fmt.Printf("Abs: %v\n", math.Abs(x))
	fmt.Printf("Pow: %v\n", math.Pow(x, 3))
	fmt.Printf("Exp: %v\n", math.Exp(x))
	fmt.Printf("Sqrt: %v\n", math.Sqrt(x))
	fmt.Printf("Cbrt: %v\n", math.Cbrt(x))

	rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf("Random Integer: %v\n", rand.Int())
	fmt.Printf("Random Integer with range: %v\n", rand.Intn(10))
	fmt.Printf("Random Float: %v\n", rand.Float64())

	scores := []int{1, 2, 3, 4, 5}
	indices := rand.Perm(len(scores))

	for _, i := range indices {
		fmt.Println(scores[i])
	}

	lowerBound, upperBound := 12, 144

	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", rand.Intn(upperBound-lowerBound+1)+lowerBound)
	}

	a, e := 'a', 'e'

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", string(rune(rand.Intn(int(e)-int(a)+1))+a))
	}

	words := strings.Fields("Hello my dear friend")

	rand.Shuffle(len(words), func (i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	fmt.Println(words)

	var sb strings.Builder

	length := 12

	for i := 0; i < length; i++ {
		sb.WriteRune(rune(rand.Intn(126 - 33 + 1) + 33));
	}

	fmt.Printf("Suggested Password: %s\n", sb.String())
}
