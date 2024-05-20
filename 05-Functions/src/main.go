package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Example of Call-By-Value
// ------------------------
type Person2 struct {
	age  int
	name string
}

// This is the main entry of the application.
func main() {
	// Headers
	fmt.Println()
	fmt.Println(strings.Repeat("-", 100))
	fmt.Println()

	// Example of calling a Function
	// -----------------------------
	fmt.Println("Example of calling a Function:")
	result := div(100, 20)
	fmt.Println("div(100, 20) =", result)
	fmt.Println()

	// Example of a Function With Optional and Named Parameters
	// --------------------------------------------------------
	fmt.Println("Example of a Function With Optional and Named Parameters:")

	// Call MyFunc() one way
	MyFunc(FuncParams{
		LastName: "Smith",
		Age:      50,
	})

	// Call MyFunc() another way
	MyFunc(FuncParams{
		FirstName: "Mary",
		LastName:  "Smith",
	})
	fmt.Println()

	// Example of a Variadic Function
	// ------------------------------
	fmt.Println("Example of a Variadic Function:")
	x := addNums(1, 2, 3, 4, 5, 6, 7, 8, 9)
	y := addNums(21, 43, 65)
	z := addNums()
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("addNums(1, 2, 3, 4, 5, 6, 7, 8, 9) =", x)
	fmt.Println("addNums(21, 43, 65) =", y)
	fmt.Println("addNums() =", z)
	fmt.Println("nums =", nums)
	fmt.Println("addNums(nums...) =", addNums(nums...))
	fmt.Println()

	// Example of Function With Multiple Return Values
	// -----------------------------------------------
	fmt.Println("Example of Function With Multiple Return Values:")
	resDiv, resMod, err := divmod(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("divmod(5, 2) =>", "ResDiv =", resDiv, "ResMod =", resMod)
	fmt.Println()

	// Example of Function With Named Return Values
	// --------------------------------------------
	fmt.Println("Example of Function With Named Return Values:")
	resX, modY, errZ := divmodNamed(5, 2)
	// Error-handling
	if errZ != nil {
		fmt.Println(errZ)
		os.Exit(1)
	}
	fmt.Println("divmodNamed(5, 2) =>", "resX =", resX, "modY =", modY)
	fmt.Println()

	// Example of Declaring a Function Variable
	// ----------------------------------------
	fmt.Println("Example of Declaring a Function Variable:")
	// A function variable
	var myFuncVar func(string) int

	// Using the first case
	myFuncVar = f1
	res := myFuncVar("Hello")
	fmt.Println("myFuncVar(\"Hello\") using f1:", res)

	// Using the second case
	myFuncVar = f2
	res = myFuncVar("Hello")
	fmt.Println("myFuncVar(\"Hello\") using f2:", res)
	fmt.Println()

	// Example of a Simple Calculator With Functions
	// ---------------------------------------------
	fmt.Println("Example of a Simple Calculator With Functions:")
	opMap := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": divs,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expr := range expressions {
		if len(expr) != 3 {
			fmt.Println("Invalid expression:", expr)
			continue
		}
		p1, err := strconv.Atoi(expr[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expr[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Unsupported Operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expr[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)
		fmt.Println(expr, "=", result)
	}
	fmt.Println()

	// Example of Anonymous Function
	// -----------------------------
	fmt.Println("Example of Anonymous Function:")
	for i := range 5 {
		func(j int) {
			fmt.Println("Printing", j, "from inside an anonymous function")
		}(i)
	}
	fmt.Println()

	// Example of Closure
	// ------------------
	fmt.Println("Example of Closure:")
	a := 20
	fmt.Println("Outside fa(), a =", a)
	/*fa =*/ func() {
		fmt.Println("Inside fa() before assignment, a =", a)
		a = 30 // The assignment modifies outside a
		fmt.Println("Inside fa() after assignment, a =", a)
	}()
	fmt.Println("Outside fa(), a =", a)
	fmt.Println()

	// Example of Closure With Shadow
	// ------------------------------
	fmt.Println("Example of Closure With Shadow:")
	b := 20
	fmt.Println("Outside fb(), b =", b)
	fb := func() {
		fmt.Println("Inside fb() before assignment, b =", b)
		b := 30 // This assignment shadows instead of modifying outside b
		fmt.Println("Inside fb() after assignment, b =", b)
	}
	fb()
	fmt.Println("Outside fb(), b =", b)
	fmt.Println()

	// Example of Sorting Slice
	// ------------------------
	fmt.Println("Example of Sorting Slice")
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"John", "Smith", 37},
		{"Jeremy", "Trye", 18},
		{"Jasmine", "Alter", 20},
	}
	fmt.Println("Before Sorting:", people)
	// Sorting the slice by last name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("After Sorting By Last Name:", people)
	// Sorting the slice by age
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("After Sorting By Age:", people)
	fmt.Println()

	// Example of Function That Returns a Closure
	// ------------------------------------------
	fmt.Println("Example of Function That Returns a Closure:")
	base2Mult := makeMult(2)
	base3Mult := makeMult(3)
	base5Mult := makeMult(5)
	fmt.Println("i\tbase2\tbase3\tbase5")
	for i := range 10 {
		fmt.Println(i, "\t", base2Mult(i), "\t", base3Mult(i), "\t", base5Mult(i))
	}
	fmt.Println()

	// Example of defer With a cat Command
	// -----------------------------------
	fmt.Println("Example of defer With a cat Command:")
	// Make sure a filename was pass as argument: make try ARGS="<filename>""
	// Args[0] is the name of the program
	if len(os.Args) < 2 {
		log.Fatal("Error: No file was specified")
	}
	// Open the file: Read-only
	fl, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	// Close the file after using it
	// This must be run no matter any errors in the program
	defer func() {
		fmt.Println("Defer in main() is called here")
		fmt.Println()
		fl.Close()
	}()
	// Read from the file
	data := make([]byte, 2048)
	for {
		count, err := fl.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	fmt.Println()

	// Example of Using defer
	// ----------------------
	fmt.Println("Example of Using defer:")
	deferExample()
	fmt.Println()

	// Example of Call-By-Value
	// ------------------------
	fmt.Println("Example of Call-By-Value:")
	i := 2
	s := "Hello"
	p := Person2{}
	// Modifying the passed parameters has no effect on the arguments
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
	fmt.Println()

	// Example of Map-Slice Modification Calls
	// ---------------------------------------
	fmt.Println("Example of Map-Slice Modification Calls:")
	mapMod := map[int]string{
		1: "first",
		2: "second",
	}
	fmt.Println("Before: mapMod =", mapMod)
	modifyMap(mapMod)
	fmt.Println("After: mapMod =", mapMod)

	slcMod := []int{1, 2, 3}
	fmt.Println("Before: slcMod =", slcMod)
	modifySlice(slcMod)
	fmt.Println("After: slcMod =", slcMod)
	fmt.Println()
}

// Example of a Function
// ---------------------
func div(num int, den int) int {
	if den == 0 {
		return 0
	}
	return num / den
}

// Example of a Function With Optional and Named Parameters
// --------------------------------------------------------
type FuncParams struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(params FuncParams) error {
	// Do something
	fmt.Println("Passed parameters:", params)
	return nil
}

// Example of a Variadic Function
// ------------------------------
func addNums(nums ...int) int {
	var res int
	for _, n := range nums {
		res += n
	}
	return res
}

// Example of Function With Multiple Return Values
// -----------------------------------------------
func divmod(num, den int) (int, int, error) {
	if den == 0 {
		return 0, 0, errors.New("cannot divide by 0")
	}
	return num / den, num % den, nil
}

// Example of Function With Named Return Values
// --------------------------------------------
func divmodNamed(num, den int) (res int, mod int, err error) {
	if den == 0 {
		err = errors.New("cannot divide by 0")
		// Returning multiple values
		return res, mod, err
	}
	res, mod = num/den, num%den
	return res, mod, nil
}

// Example of Declaring a Function Variable
// ----------------------------------------
func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	sum := 0
	for _, v := range a {
		sum += int(v)
	}
	return sum
}

// Example of a Simple Calculator With Functions
// ---------------------------------------------
func add(i int, j int) int  { return i + j }
func sub(i int, j int) int  { return i - j }
func mul(i int, j int) int  { return i * j }
func divs(i int, j int) int { return i / j }

// Example of Function That Returns a Closure
// ------------------------------------------
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

// Example of Using defer
// ----------------------
func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("First value:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("Second value:", val)
	}(a)
	a = 30
	fmt.Println("Exiting deferExample:", a)
	return a
}

// Example of Call-By-Value
// ------------------------
func modifyFails(i int, s string, p Person2) {
	// Attempting to modify the passed parameters
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

// Example of Map-Slice Modification Calls
// ---------------------------------------
func modifyMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modifySlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

// AVAILABLE COMMANDS
// ------------------
//  make            Default to `make try`
//  make fmt        Format all source files
//  make vet        Verify any possible errors
//  make build      Build module
//  make run        Build module then run
//  make try        Build module, run, then remove built binary
