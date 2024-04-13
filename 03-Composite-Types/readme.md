# Composite Types

---

- [Arrays](#arrays)
  - [Array Declarations](#array-declarations)
  - [Comparing Arrays](#comparing-arrays)
  - [Multidimensional Array](#multidimensional-array)
  - [Array Access](#array-access)
  - [Getting Array Length](#getting-array-length)
  - [Limitations: Why Arrays Are Rarely Used](#limitations-why-arrays-are-rarely-used)
  - [Purpose: Why Arrays Exist If They Are So Limited](#purpose-why-arrays-exist-if-they-are-so-limited)
- [Slices](#slices)
  - [Working With Slices](#working-with-slices)
  - [Multidimensional Slice](#multidimensional-slice)
  - [Slice Access](#slice-access)
  - [Differences With Arrays](#differences-with-arrays)
    - [The Zero-Value for slice is `nil`](#the-zero-value-for-slice-is-nil)
    - [Slices Are Not Comparable With `==`](#slices-are-not-comparable-with-)
  - [Slices Helper Functions](#slices-helper-functions)
  - [`len()` Function](#len-function)
  - [`append()` Function](#append-function)
  - [Capacity](#capacity)
  - [NOTE: About Go Runtime](#note-about-go-runtime)
  - [`make()`](#make)
  - [Emptying A Slice](#emptying-a-slice)
  - [Which Slice Declarations To Use](#which-slice-declarations-to-use)
  - [Slicing Slices](#slicing-slices)
  - [`copy()`](#copy)
- [Converting Array To Slice](#converting-array-to-slice)
- [Converting Slice To Array](#converting-slice-to-array)

---

## Arrays

- Arrays are rarely used in Go
- **All elements in an array must be of the type of the array**

### Array Declarations

- There are different declaration styles

```go
// Basic array declaration
var arrA [3]int
```

- **If no value is assigned, all the elements are initialized with the Zero-Value of the type**
- But we can also specify the values with an *Array Literal*

```go
// Array declaration with Array literal
var arrB = [3]int{10, 20, 30}
```

- We can specify a *Sparse Array* as well
  - Only specify the Non-Zero elements using indices

```go
// Array declaration with specific non-zero elements
var arrC = [12]int{1, 5: 4, 6, 10, 100, 15}
```

- When using array literal, we can replace the length of the array to be dynamic with `...`
- However, this can be harder to read

```go
// Array declaration with ... length
var arrD = [...]int{100, 200, 300}
```

- With can combine any of the above with the `:=` operator as well

```go
// Array declaration with :=
arrE := [...]int{100, 200, 300}
```

### Comparing Arrays

- We can use `==` and `!=` to compare arrays
- **Arrays are equal if *same length and same values***

```go
// Comparing arrays
arrF := [...]int{1,2,3,4,5}
var arrG = [5]int{1,2,3,4,5}

fmt.Println("arrF == arrG:", arrF == arrG)
```

### Multidimensional Array

- **Go does not support multidimensional arrays**
- However, we can simulate them with *Array of Array*

```go
// Array of array of integers: 5 x 5
arrH = [5][5]int
```

### Array Access

- Arrays in Go are read and written via the `Array[index]` syntax

```go
// Array Access
arrI := [3]int{10, 20, 30}

// Accessing the values
fmt.Println("arrI[0] =", arrI[0])
fmt.Println("arrI[1] =", arrI[1])
fmt.Println("arrI[2] =", arrI[2])

// Changing the values
arrI[0] = 15
arrI[1] = 25
arrI[2] = 35

fmt.Println("arrI =", arrI)
```

- We cannot access beyond the length of the array: Compile-time error
- **We cannot access with negative index**: Compile-time error
- **Out-of-bonds access will *panic* at runtime**

### Getting Array Length

- Use the built-in function `len()` to  get the length of an array

```go
// Getting the length of an array
arrJ := [3]int{10, 20, 30}
arrK := [...]int{15, 25, 35}

fmt.Println("arrJ =", arrJ, "length =", len(arrJ))
fmt.Println("arrK =", arrK, "length =", len(arrK))
```

### Limitations: Why Arrays Are Rarely Used

- **Go considers the size of the array to be part of the type of the array**
  - `[3]int` is different from `[4]int`
- **We cannot use variable to specify the size**
  - Types must be resolved at compile-time
- **We cannot use type-conversion to convert arrays of different size to the same size**
  - We cannot write functions to work with arrays of any size
  - We cannot assign arrays of any sizes to the same variable
- ***Do not use array unless you know exactly the needed length ahead of time***

### Purpose: Why Arrays Exist If They Are So Limited

- The main purpose of arrays is to be used behind-the-scenes for *Slices*

## Slices

- Most of the time, *Slices* are what we should use for a *sequence* data structure
- **Slice length is dynamic: Can grow and shrink**
- **The slice length is NOT part of its type**
  - Removes the limitation from arrays
  - Allows to write functions to process slices of any size

### Working With Slices

- Very similar to working with Arrays
- With some subtle differences
- **Declaration With Literal: Do not specify the length**

```go
// Slice declaration with Slice literal
sliceA := []int{10, 20, 30, 40}
```

- We can specify a *Sparse Slice* as well
  - Only specify the Non-Zero elements using indices

```go
// Slice declaration with specific non-zero elements
sliceB := []int{1, 5: 4, 6, 10, 100, 15}
```

### Multidimensional Slice

- **Go does not support multidimensional slices**
- However, we can simulate them with *Slice of Slice*

```go
// Slice of slice of integers
var SliceC [][]int
```

### Slice Access

- Same rules as with *Arrays*
  - Read and write via the `Slice[index]` syntax
  - We cannot access beyond the length of the slice: Compile-time error
  - **We cannot access with negative index**: Compile-time error
  - **Out-of-bonds access will *panic* at runtime**

```go
// Slice Access
SliceD := []int{10, 20, 30}

// Accessing the values
fmt.Println("SliceD[0] =", SliceD[0])
fmt.Println("SliceD[1] =", SliceD[1])
fmt.Println("SliceD[2] =", SliceD[2])

// Changing the values
SliceD[0] = 15
SliceD[1] = 25
SliceD[2] = 35

fmt.Println("SliceD =", SliceD)
```

### Differences With Arrays

#### The Zero-Value for slice is `nil`

```go
// Declaring a slice with zero-value: nil
var sliceNil []int
```

- **`nil` is slightly different than `null` in other languages**
  - `nil` is an identifier that represent lack of value for some types
  - **`nil` is untyped**: Can be used for different types
- ***A `nil` slice contains nothing***

#### Slices Are Not Comparable With `==`

- **Using `==` and `!=` on 2 slices result in compile-time error**
- We can only use `slice == nil`

```go
// Checking if slice is nil
var sliceE []int

fmt.Println("sliceE is nil?", sliceE == nil)
```

### Slices Helper Functions

- `slices` package includes helper functions for dealing with slices
- Also has functions for comparing slices

Function|Description
:-|:-
`slices.Equal()`|Takes 2 slices *with comparable elements* and return `true` if the slices are same length with all equal elements. Else, return `false`. *The slice elements must be comparable*
`slices.EqualFunc()`|Allows to pass a function argument to customize conditions for equality. *The slice elements do not have to be comparable*

```go
import (
    "fmt"
    "slices"
)

sliceX := []int{1, 2, 3, 4, 5}
sliceY := []int{1, 2, 3, 4, 5}
sliceZ := []int{1, 2, 3, 4, 5, 6}
sliceS := []string{"a", "b", "c"}

fmt.Println("slices.Equal(sliceX, sliceY) ?", slices.Equal(sliceX, sliceY))
fmt.Println("slices.Equal(sliceX, sliceZ) ?", slices.Equal(sliceX, sliceZ))
// This one will not compile: Different types for each elements
fmt.Println("slices.Equal(sliceX, sliceS) ?", slices.Equal(sliceX, sliceS))
```

- **NOTE: `reflect` package contains a `DeepEqual` function that can compare almost anything**
  - Legacy function, primarily for testing
  - Was often used to compare slices before `slices.Equal` and `slice.EqualFunc`
  - **Do not use `reflect.DeepEqual` in new codes**
  - It is slower and less safe

### `len()` Function

- Built-in function for getting length of collections
  - Can be used on both **Arrays** and **Slices**
  - Also works on **Strings**, **Maps**, **Channels**
  - **However, any other type will result in compile-time error**
- Passing a `nil` slice returns `0`
- Parameter can be any type of array or slice

```go
// Getting the length of a slice
sliceF := []int{10, 20, 30, 40, 50}
fmt.Println("sliceF =", sliceF, "length =", len(sliceF))
```

### `append()` Function

- Used to grow a slice by appending new elements to its end
- **Takes 2 parameters**
  - *Slice of any type*
  - *Value of that type to append to the slice*
- Returns a new slice of the same type
  - Assign it to a variable to capture the result
  - We typically re-assign to the original variable
  - **NOTE: Not assigning the returned value to a variable is a compile-time error**
    - *Go is a call-by-value language*: All passed parameters are a copy of the value passed in
    - The function appends unto the copy of the slice
    - The copy of the slice is returned with appended values
    - *The copy needs to be re-assigned to make the change permanent*
- **`append()` can take more than 1 value at once to append**

```go
// Appending to a nil slice
var sliceG []int
sliceG = append(sliceG, 100)

// Appending to an existing slice
sliceH := []int{1, 2, 3, 4, 5}
// We can append more than 1 value at once
sliceH = append(sliceH, 6, 7, 8, 9)
```

- We can also append one slice unto another
  - Using the `...` operator
  - **This has the effect of *extending* the slice**

```go
// Extending a slice with another slice
sliceH = append(sliceH, sliceG...)
```

### Capacity

- A slice is a sequence of values
  - Each element assigned to consecutive memory locations
  - Quick read and write
- *Length of slice = Number of consecutive memory locations that have been assigned a value*
  - Each value added via `append()` increases the length
- **Capacity of slice = Number of consecutive memory locations that have been reserved**
  - Can be larger than *length*
  - When `length == capacity`, there are no more rooms to append additional values
  - If so:
    - The runtime allocate a new backing array for larger capacity
    - *The values in the original backing array are copied into the new backing array*
    - The new value is appended at the end of the new array
    - The slice is updated to refer to the new backing array
    - The updated slice is returned
- **For this reason, growing a slice can take some time**
  - Allocate new memory
  - Copy contents from old memory to new memory
  - Garbage collection
- **Go usually increase capacity by more than one each time it runs out of capacity**
  - Double when the current capacity is less than 256
  - Bigger slice: Increase by `(currCap + 768) / 4`
  - Slowly converge at 25% growth
- **Built-in `cap()` function returns the capacity of a slice**
  - Less used than `len()`
  - Mostly used to check if a slice is large enough to hold a data
  - If not, a call to `make()` is needed to create a new slice
  - We can also pass an array to `cap()`
    - But always `len(arr) == cap(arr)`
    - Not useful to use in code

```go
// Appending and Slice Capacity
var xSlice []int
fmt.Println(xSlice, len(xSlice), cap(xSlice))

xSlice = append(xSlice, 10)
fmt.Println(xSlice, len(xSlice), cap(xSlice))

xSlice = append(xSlice, 20)
fmt.Println(xSlice, len(xSlice), cap(xSlice))

xSlice = append(xSlice, 30)
fmt.Println(xSlice, len(xSlice), cap(xSlice))

// We reach the current slice cap here
xSlice = append(xSlice, 40)
fmt.Println(xSlice, len(xSlice), cap(xSlice))

// This one increases the cap
// Copy into a new slice
xSlice = append(xSlice, 50)
fmt.Println(xSlice, len(xSlice), cap(xSlice))
```

- **While it is nice that slices can grow dynamically, it is more efficient to size them once**
  - If the size can be determined at the beginning
  - We can use the `make()` function then

### NOTE: About Go Runtime

- **Provides different services**
  - Set of libraries that enables programs written in Go to run
  - Memory Allocation
  - Garbage Collection
  - Concurrency Support
  - Networking
  - Built-in types and functions
- **Compiled/Bundled into every Go programs**
  - No VM involved
  - Allows easier distribution of compiled Go programs
  - Avoids compatibility issues between runtime and program
  - But can increase the size of the compiled program slightly (Minimum 2MB, even for simplest programs)

### `make()`

- The way we currently declare slices does not allow to create an empty slice with length and capacity specified
- For that, we use `make()`
- Allows to specify the type, length, and capacity

```go
// Declaring a slice using make()
// Length: 5, Capacity: 5
sliceI := make([]int, 5)
```

- Using `make()`, all elements are initialized to `0`
- **NOTE: Since all elements are initialized, we should not use `append()` to the slice**
  - The slice is not empty
  - Using `append()` will increase the length and double the capacity of the slice
- We can also specify a different initial capacity

```go
// Declaring a slice using make()
// Length: 5, Capacity: 10
sliceJ := make([]int, 5, 10)
```

- We can also make a slice of `0` length but with capacity
  - **Using this approach allows to `append()` to the slice**

```go
// Declaring a slice using make()
// Length: 0, Capacity: 10
sliceK := make([]int, 0, 10)
sliceK = append(sliceK, 1, 2, 3, 4, 5)
```

- **NOTE: Never specify a capacity less than the length**
  - If done via constant, it is a compile-time error
  - If done via variables, it is a runtime error

### Emptying A Slice

- `clear()` allows to reset all elements in a slice back to the zero-value
- **NOTE: The length of the slice remains unchanged**

```go
// Resetting slice using clear()
sliceL := []string{"first", "second", "third"}
fmt.Println("Before:", sliceL, len(sliceL))
clear(sliceL)
fmt.Println("After:", sliceL, len(sliceL))
```

### Which Slice Declarations To Use

- **Primary goal is to try to minimize the number of times the slice has to grow**
- If the slice will not need to grow at all, create a `nil` slice

```go
// Slice with nil value
var data []int
```

- **NOTE: A slice with zero-length and zero-capacity is different from a `nil` slice**
  - Comparing the 2 will returns `false`
  - Comparing a `nil` to a `nil` returns `true`
  - *For simplicity, favor `nil` slice*
  - **A zero-length slice is only useful when converting a slice to JSON**

```go
// Slice with zero length and zero capacity
var data = []int{}
```

- With starting values or if slice's values will not change, use slice literal

```go
// Slice literal
data := []int{1, 2, 3, 4, 5}
```

- If you do not know the slice values but have an idea of how large the slice with be, use `make()`
  - If using slice as a buffer: *Specify a non-zero length*
  - If you know the exact size needed: *Specify the length and index into the slice to set the values*
  - In other situations: *Zero length and specify the capacity*
- **The choice is usually between the 2nd and the 3rd option**
  - Option 2 might be slower but introduces less bugs
  - `append()` always increases the length
  - *Specifying the length with `make()` must be done carefully*

### Slicing Slices

- Slice expression creates a slice from a slice
- Written inside brackets
- Consists of *starting* and *ending* offsets separated by colon
- Ending offset is *up-to-but-not-including*
- Default starting offset is `0`
- Default ending offset is the end of the slice

```go
// Slicing slice
// -------------
sliceM := []string{"a","b","c","d"}
subM1 := sliceM[:2]
subM2 := sliceM[1:]
subM3 := sliceM[1:3]
subM4 := sliceM[:]

fmt.Println("sliceM =", sliceM)
fmt.Println("sliceM[:2] =", subM1)
fmt.Println("sliceM[1:] =", subM2)
fmt.Println("sliceM[1:3] =", subM3)
fmt.Println("sliceM[:] =", subM4)
```

- ***WARNING: Slicing a slice does not make a copy***
  - Assigning the slice of the slice to a variable create an alias to the original variable
  - **Changing an element in one variable affects all other aliases**

```go
// Modifying one slice affects all others
subM1[1] = "y"
subM2[0] = "x"
subM3[1] = "z"

fmt.Println("sliceM =", sliceM)
fmt.Println("subM1 =", subM1)
fmt.Println("subM2 =", subM2)
fmt.Println("subM3 =", subM3)
fmt.Println("subM4 =", subM4)
```

- **This behavior can get extra-confusing when combined with `append()`**
  - `cap(subSlice) = cap(origSlice) - startingOffsetSubSlice`
  - Elements of the original slice beyond the end of the subslice are shared between them
  - ***This can cause slices and subslices to append and overwrite each other's data***

```go
// Slice of Slice with append()
fmt.Println("Slice of slice with append()")
fmt.Println("BEFORE:", "sliceM", cap(sliceM), sliceM)
fmt.Println("BEFORE:", "subM1", cap(subM1), subM1)

subM1 = append(subM1, "zz")

fmt.Println("AFTER:", "sliceM", cap(sliceM), sliceM)
fmt.Println("AFTER:", "subM1", cap(subM1), subM1)
```

- To avoid this issue, either:
  - **never use `append()` with subslices**
  - **always use *full slice expression* with `append()`**
    - Includes a 3rd part
    - *Indicates last position in the capacity of the original slice that is available for the subslice*
    - *Subtract the starting offset from this number to get the subslice capacity*
- Makes clear how much memory is shared between original slice and subslice

```go
// Full Slice Expressions
sliceN := make([]string, 0, 5)
sliceN = append(sliceN, "a", "b", "c", "d", "e")
subN1 := sliceN[:2:2]
subN2 := sliceN[2:4:4]

fmt.Println("Capacities:", cap(sliceN), cap(subN1), cap(subN2))

subN1 = append(subN1, "i", "j", "k")
sliceN = append(sliceN, "x", "y")
subN2 = append(subN2, "z", "zz")

fmt.Println("sliceN =", sliceN)
fmt.Println("subN1 =", subN1)
fmt.Println("subN2 =", subN2)
```

- Both `subN1` and `subN2` have a capacity of `2`
- **After limiting the capacities of the subslices to their lengths, appending to the subslices would now force to create new slices for them, which breaks the aliases with the other subslices**
- ***NOTE: If possible, avoid modifying slices after they have been sliced or produced by slicing***
  - **Else, use the *Full-slice expression* to prevent `append()` from sharing capacity**

### `copy()`

- Allows to create a slice independant of the original slice
- Takes 2 parameters:
  - `destination` slice
  - `source` slice
- Copies as many values as possible from source to destination
  - Limited by whichever slice is smaller
  - Capacities do not matter
  - Only the lengths
- Returns the number of elements copied

```go
// Slice Copy
// ----------
fmt.Println("Slicing copy:")
sliceO := []int{1, 2, 3, 4}
sliceP := make([]int, 4)

fmt.Println("Before Copying:", sliceO, sliceP)
num := copy(sliceP, sliceO)
fmt.Println("After Copying:", sliceO, sliceP, "copied", num, "elements")
```

- Can also copy a subset of a slice

```go
// Copying a subset
// ----------------
fmt.Println("Copying subset of slice:")
sliceQ := make([]int, 2)
num = copy(sliceQ, sliceO)
fmt.Println("After Copying:", sliceO, sliceQ, "copied", num, "elements")
```

- Can also copy from the middle of the source slice
- **NOTE: We do not have to assign the return of `copy()` if we do not need it**

```go
// Copying from middle of slice
// ----------------------------
fmt.Println("Copying from the middle of slice:")
sliceR := make([]int, 2)
copy(sliceR, sliceO[2:])
fmt.Println("After Copying:", sliceO, sliceR, "copied", num, "elements")
```

- We can copy between 2 slices that cover overlapping sections of an underlying slice

```go
// Overlapping copies
// ------------------
fmt.Println("Overlapping copies")
num = copy(sliceO[:3], sliceO[1:])
fmt.Println("After Copying:", sliceO, "copied", num, "elements")
```

- **NOTE: We can use `copy()` on arrays as well**
  - Can be source or destination

```go
// Using copy() with arrays
// ------------------------
fmt.Println("Using copy() with arrays")
sliceS := []int{1,2,3,4} // Slice
arrL := [4]int{5,6,7,8} //  Array
sliceT := make([]int, 2)

fmt.Println("Before copying:", sliceS, arrL, sliceT)
copy(sliceT, arrL[:])
fmt.Println("After copying from array:", sliceT, arrL)
copy(arrL[:], sliceS)
fmt.Println("After copying into array:", arrL, sliceS)
```

## Converting Array To Slice

- We can take a slice from an array
- Using slice expression
- Allows to use an array as a slice
- **To convert an entire array into a slice, use `[:]`**
- We can also convert only a subset with `[start:]` or `[:end]`
- **WARNING: This results in the same memory-sharing issue as a slice of slice**

```go
// Converting Array To Slice
bArray := [4]int{100, 200, 300, 400}
bSlice := bArray[:]
```

## Converting Slice To Array

- For *slice* to *array*, we use a type-conversion
  - Can convert an entire slice to an array of the same type
  - Can create an array from a subset of a slice
  - **The data is *copied* into a new memory**
  - There is no memory-sharing issue