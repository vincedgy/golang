package main

import "fmt"

/**

 */
func TestAddThemUp() {

	listNums := []float64{1, 2, 3, 4.5}

	fmt.Println(AddThemUp(listNums))

}

func AddThemUp(numbers []float64) float64 {

	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

func main() {

	var Message string = "Hello"

	var (
		var1      = 2
		var2      = 3
		var3 bool = true
	)

	fmt.Println(Message)

	randnum := 1

	randnum++

	fmt.Println("Number is ", (randnum*var1+var2)%2)

	fmt.Println("Lenght is ", len(Message), " super "+" \ncool", var3)

	fmt.Printf("Super number here %d yeah ?\n	", 100)

	fmt.Println("true && false = ", true && !false)

	i := 0
	for ; i < 10; i++ {
		fmt.Printf("%d, ", i)
	}

	if i >= 9 {
		fmt.Printf("%d", i)
	}

	fmt.Print("Hi")

	yourAge := 20

	// Switch
	switch yourAge {
	case 16:
		fmt.Print("Cool")
	default:
		fmt.Println("Too bad")
	}

	// Tab
	tab := [5]int{1, 2, 4}
	tab[0]++
	fmt.Println(tab)

	// Slices
	slice := []int{1, 2, 4, 5, 6, 7, 8}
	numSlice := slice[0:]
	fmt.Println(numSlice)
	fmt.Println(numSlice[:4])
	fmt.Println(numSlice[1:])

	numSlice3 := append(numSlice, 0, -1)
	fmt.Println(numSlice3)

	// Maps
	presAge := make(map[string]int)
	presAge["Paul"] = 2
	presAge["Nicole"] = 20
	presAge["Bertrand"] = 12

	fmt.Println(presAge)

	delete(presAge, "Paul")
	fmt.Println(presAge)

}
