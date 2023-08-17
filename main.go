package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Point struct {
	x int
	y int
}

var some_global, another bool = true, false

const twoja_stara = ";)"

// Numeric constants are high-precision values.
// An untyped constant takes the type needed by its context.
const that_number = 69

func main() {
	defer fmt.Println("This has been: fun with go")
	defer fmt.Println("The current time is", time.Now())

	fmt.Println("This is a number:", rand.Intn(100))
	fmt.Printf("Whatever %g\n", math.Sqrt(2))

	a, b := x_plus_times_y(4, 4)
	fmt.Println("3 + 5", a, b)

	a, b = nekkid_return(34, 33)
	fmt.Println("The result of this nekkid return:", a, b)

	fmt.Println("Sigma 10 = ", suma_suma(10))
	fmt.Println("When did you say Saturday is?")
	when_is_saturday()

	point := Point{55, 99}
	pr := &point

	pr2 := &Point{y: -44}
	fmt.Println("first point.x =", pr.x, "& point.y =", pr.y)
	fmt.Println("second point.x =", pr2.x, "& point.y =", pr2.y)

	fibonacci := [8]int{1, 1, 2, 3, 5, 8, 13, 21}
	// for idx, val := range fibonacci {
	// 	fmt.Println("Number ", idx+1, "of fibonacci is", val)
	// }
	fmt.Println("fibonacci:", fibonacci)
	slice := fibonacci[:7]
	slice[6] = 666
	fmt.Println("fibonacci and slice after modifying slice:", fibonacci, slice)
	// it seems the right bound can be moved back and forward without consequence
	// but the left bound moving forward is irreversable
	slice = slice[2:] // this apparently drops the first two items?
	printSlice(slice)

	slice = append(slice, 34, 55)
	printSlice(slice)
	fmt.Println("fibonacci remains unchanged:", fibonacci)

	// five_zeroes := make([]int, 5)           // [0 0 0 0 0]
	five_hidden_zeroes := make([]int, 0, 5) // [] as a slice of the above
	five_revealed_zeroes := five_hidden_zeroes[:5]
	fmt.Println(five_hidden_zeroes, five_revealed_zeroes)
}

func printSlice(s []int) {
	fmt.Println("Len =", len(s), "| Cap =", cap(s), "|", s)
}

// func x_plus_y(x int, y int) int {
func x_plus_times_y(x, y int) (int, int) {
	return x + y, x * y
}

func nekkid_return(x, y int) (a, b int) {
	a = x + y + 5
	b = 99 - a
	return // disgusting
}

func suma_suma(a int) int {
	suma := 0
	// for i := 0 ; i <= a ; i++ {
	// suma += i
	if dingy_dongy := 0; a <= 0 {
		return dingy_dongy
	} // dingy_dongy is dropped
	for ; a > 0; a-- {
		suma += a
	}
	return suma
}

func saturday() time.Weekday {
	return time.Now().Weekday() + 0
}

func friday() time.Weekday {
	return time.Now().Weekday() + 1
}

func when_is_saturday() {
	// today := time.Now().Weekday()
	switch time.Saturday {
	case saturday():
		fmt.Println("En effect, c'est aujourd'hui :)")
	case friday():
		fmt.Println("O la la, c'est demain")
	default:
		fmt.Println("il faut toujours attendre :(")
	}
}

// Implement Pic. It should return a slice of length dy, each element of which is a slice of dx
// 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the
// integers as grayscale (well, bluescale) values.
func Pic(dx, dy int) [][]uint8 {
	out := make([][]uint8, dy)
	for y := range out {
		out[y] = make([]uint8, dx)
		for x := range out[y] {
			out[y][x] = uint8(x * y)
		}
	}
	return out
}
