# Learning<br>Go

## 1. The *"go"* comand

## 2. Go Variables

    var number int
  * number - name of the variable
  * var - keyword used to declare variable
  * int - data type associated with the variable
##### There are 3 ways to assign values to a variable.  
```
    var number int = 10 
```
```
    var number = 10
```
* shorthand 
```
    number := 10
```
##### Rules of naming Variables
```
    ___________________________________________________
    | Illegal Variable | Bad Variable | Good Variable |
    ==================================================|
    | first name       | fn           | firstName     |
    | last name        | ln           | lastName      |
    | user role        | ul           | userRole      |
```
## 3. Go Constants

## 4. The Basic Data Types In Golang

```
  ____________________________________________________________________________________
  | Data Types |            Description                 | Examples
  |============|========================================|=============================
  | int        |  Integer numbers.                      | -21, 7, 345
  | float      |  Numbers with decimal points.          | -21, -23.3, 23.23, 233
  | compleks   |  Complex numbers.                      | 2+4i, -9.5+18.3i
  | string     |  Sequence of characters.               | "hello, world", "this is go"
  | bool       |  Either true or false.                 | true, false
  | byte       |  A byte (8 bits) of                    | 2, 115, 7
  |            | non-negative integers.                 |
  | rune       |  Used for characters.                  | '<', '}', '4'
  |            | Internally used as 32-bit integers.    |
```

## 5. Type Casting

## 6. Go Operators

* **Go programming provides a wide range of operators that are categorized into the following major categories:**
    * Arithmetic operators
    * Assignment operator
    * Relational operators
    * Logical operators

* Arithmetic operators
```
    __________________________________________
    |      Operators       |      Example    |
    |----------------------+-----------------|
    |  + (Addition)        |       a + b     |
    |  - (Subtraction)     |       a - b     |
    |  * (Multiplication)  |       a * b     |
    |  / (Division)        |       a / b     |
    |  % (Modulo Division) |       a % b     |
    ------------------------------------------
```
 * Assignment operator
 ```
    ______________________________________________________________
    |           Operators            |     Example   |  Same as  |
    |----------------------+---------+---------------------------|
    | += (addition assignment)       |     a =+ b    | a = a + b |
    | -= (subtraction assignment)    |     a =- b    | a = a - b |
    | *= (multiplication assignment) |     a =* b    | a = a * b |
    | /= (division assignment)       |     a =/ b    | a = a / b |
    | %= (modulo assignment)         |     a =% b    | a = a % b |
    --------------------------------------------------------------
```
* Relational operators
    * true if the comparison between two values is correct
    * false if the comparison is wrong

 ```
    _________________________________________________
    |           Operators           |     Example   |
    |-------------------------------+----------------
    | == (equal to)                 |     a == b    | 
    | != (not equal to)             |     a != b    |
    | > (greater than)              |     a > b     |
    | < (less than)                 |     a < b     |
    | >= (greater than or equal to) |     a >= b    |
    | <= (less than or equal to)    |     a <= b    |
    -------------------------------------------------
```
* Logical operators

We use the logical operators to perform logical operations. A logical operator returns either true or false depending upon the conditions.
```
    ___________________________________
    |      Operators   |   Example    |
    |------------------+--------------|
    | && (Logical AND) | exp1 && exp2 | 
    | || (Logical OR)  | exp1 || exp2 |
    | ! (Logical NOT)  | !exp         |
    -----------------------------------
```
## 7. Go if else

In computer programming, we use the if statement to run a block code only when a certain condition is met.

* Go if statement
```
    if test_condition {
        // code
    }
```

* Go if...else statement
```
    if test_condition {
        // run code if test_condition is true
    } else {
        // run code if test_condition is false
    }
```

* Go if...else if ladder
```
    if test_condition1 {
        // code block 1
    } else if test_condition2 {
        // code block 2
    }.
    .
    .
    } else {
        // code block 3
    }
```
***
## 8. Go switch

In Go, the switch statement allows us to execute one code block among many alternatives.
```
    switch expression {
    case 1:
        // code block 1

    case 2:
        // code block 2

    case 3:
        // code block 3
    ...
    ...
        
    default:
        // default code block
    }
```
***
## 9. Go Loops
* increment operator
* decrement operator

### Go for Loop
```
    for initialization; condition; update {
       statement(s)
    }
```
ex:
```
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
```
### Go while Loop
```
    for condition {
        // code block
    }
```
ex:
```
    for number <= 5 {
        fmt.Println(number)
        number++
    }
```
### Go for range
```
    var numbers [7]int = [7]int{1, 2, 3, 4, 5, 6, 7}
    fmt.Println(numbers)

    for _, number := range numbers {
            fmt.Println(number)
	}
```

#### Go break and contin
* In Go, the break statement terminates the loop when it is encountered. For example,
```
    for initialization; condition; update {
         break;
    }
```
* In Go, the continue statement skips the current iteration of the loop. It passes the control flow of the program to the next iteration. For example,
```
    for initialization; condition; update {
        if condition {
            continue
        }
    }
```
***
## 10. Go Functions
### 10.1 normal functions
* **Create a Go Function**
```
    func greet(){
        // code
    }
```
* Function Call
```
    greet()
```
* **Define Function with Parameters**
```
    func addNumbers(n1 int, n2 int) {
        // code
    }
```
* Function Call
```
    // function call
    addNumbers(21, 13)
```
* **Return Value from Go Function**
```
    func addNumbers(n1 int, n2 int) int {
        // code
        return sum
    }
```
* Function Call
```
    result := addNumbers(21, 13)
```
* **Return Multiple Values from Go Function**
```
    func calculate(n1 int, n2 int) (int, int) {
        //code
        return sum, difference
    }
```
* Function Call
```
    sum, difference = calculate(21, 13)
```

### 10.2 Go Anonymous Function
* In Go, we can create a function without the function name, known as an anonymous function. For example,
```
    func () {
        fmt.Println("Function without name")
    }
```
* what we do is assign the anonymous function to a variable and then use the variable name to call the function. For example,
```
    //anonymous function
    var greet = func (){
        // code
    }

    // function call
    greet()
```
### 10.3 Return an Anonymous Function from a Function
```
    func name() func() int {
        num := 7
        return func() int {
            return num
        }
    }
```
* call
```
    fn := name()
    fmt.Println(fn())
```
### 10.4 Nested function in Golang
* In Go, we can create a function inside another function. This is known as a nested function. For example,
```
    func greet(name string) {

        // inner function
        var displayName = func() {
            fmt.Println("Hi", name)
        }

        // calling inner function
        displayName()
    }
```
### 10.5 Go Closure

Go closure is a nested function that helps us access the outer function's variables even after the outer function is closed. Let's see an example.
```

    func displayNumbers() func() int {
        number := 0

        // inner function
        return func() int {
            number++
            return number
        }
    }

    func main() {

        // returns a closure 
        num1 := displayNumbers()

        fmt.Println(num1())
        fmt.Println(num1())
        fmt.Println(num1())

        // returns a new closure
        num2 := displayNumbers()
        fmt.Println(num2())
        fmt.Println(num2())

    }
```
### 10.6 Go Recursion
* In computer programming, a recursive function calls itself. For example,
```
    func recurse() {
        … …
        … …
        recurse()
    }
```
### 11. Go Pointers
#### 11.1 Memory Address
When we create a variable, a memory address is allocated for the variable to store the value.
In Go, we can access the memory address using the & operator. For example,
```
    var num int = 5
    // prints the value stored in variable
    fmt.Println("Variable Value:", num)

    // prints the address of the variable
    fmt.Println("Memory Address:", &num)
```
Here, &num accesses the memory address where the num variable is stored.

#### 11.2 Go Pointer Variables
* In Go, we use the pointer variables to store the memory address. For example,
```
    var number int = 7

    // create the pointer variable
    var ptr *int = &number
```
* We can also create pointer variables of other types. For example,
```
    // pointer variable of string type
    var ptr1 *string

    // pointer variable of double type
    var ptr2 * float32
```
* Get Value Pointed by Pointers in Golang
We use the * operator to access the value present in the memory address pointed by the pointer. For example,
```
    var number int = 7

    // create the pointer variable
    var ptr *int = &number

    // * to get the value pointed by ptr
    fmt.Println(*ptr) // 7
```
* nil pointer
```
    var ptr *int
    
    fmt.Println("Value of pointer:", ptr) 
    
    // out:
    // Value of pointer: <nil>
```
* Create Pointers using Golang new()
```
    var ptr = new(int)
  
    *ptr = 20  
    fmt.Println(ptr)  // 0xc000016058

    // create pointer without *
    var ptr = &name
```

## 12. Go Pointers and Functions
Go pointers store the memory addresses of variables. And just like regular variables, we can also pass pointers to functions.

#### 12.1 Pointer as Function Argument in Golang
In Go, we can pass pointers as arguments to a function. For example,
```
    func abs(number *int) {
        if *number < 0 {
            *number = -(*number)
        }
    }

    func main() {
        number := -7
        abs(&number)
        fmt.Println(number)
    }

```
#### 12.2 Return Pointers From Function
Just like regular variables, we can also return pointers from a function. For example,
```
    func display() *string {
        // code
        return &message
    } 
```

#### 12.3 Call By Reference
```
    // Program to illustrate call by reference

    // call by value
    func callByValue(num int) {
        num = 30
        fmt.Println( num) // 30
    } 

    // call by reference
    func callByReference(num *int) {
        *num = 10
        fmt.Println(*num) // 10
    } 

    func main() {
        var number int

        // passing value
        callByValue(number)

        // passing a reference (address)
        callByReference(&number)
    }
```

## 13. Go Data Structures

### 13.1 Go Array
An array is a collection of similar types of data. For example,
#### 13.1.1 Creating an array in Go

Here is a syntax to declare an array in Golang.
```
    var array_variable = [size]datatype{elements of array}
```    
#### 13.1.2 Accessing array elements in Golang

In Go, each element in an array is associated with a number. The number is known as an array index.
We can access elements of an array using the index number (0, 1, 2 …). For example,
```
    progeammingLanguages := [3]string{"C++", "Go", "Java"}

    // access element at index 0
    fmt.Println(progeammingLanguages[0]) // C++

    // access element at index 2
    fmt.Println(progeammingLanguages[2]) // Java
```
#### 13.1.3 Initialize an Array in Golang

We can also use index numbers to initialize an array. For example,
```
    numbera[3] = 7
```
* ***Initialize specific elements of an array***

In Golang, we can also initialize the specific element of the array during the declaration. For example,
```
    // to initialize the elements of index 0 and 3 only
    numbers := [5]int{0: 7, 3: 9}

    fmt.Println(numbers)
```
#### 13.1.4 Find the length of an Array in Go

In Golang, we can use the len() function to find the number of elements present inside the array. For example,

#### 13.1.5 Multidimensional array in Golang

Arrays we have mentioned till now are called one-dimensional arrays. However, we can declare multidimensional arrays in Golang.

A multidimensional array is an array of arrays. That is, each element of a multidimensional array is an array itself. For example,

```
    arrayInteger := [2][2]int{{1, 2}, {3, 4}}

    // accessing the values of 2d array
    for i := 0; i < 2; i++ {
        for j := 0; j < 2; j++ {
            fmt.Println(arrayInteger[i][j])
        }
    }
```
### 13.2. Go Slice
In Go, slice a collection of similar types of data, just like arrays.

However, unlike arrays, slice doesn't have a fixed size. We can add or remove elements from the array.

* Create slice
```
    numbers := []int{1, 2, 3, 4, 5, 6, 7}
```
* **Slice Functions**

    * **append()** - adds element to a slice
    * **copy()** - copy elements of one slice to another
    * **Equals()** - compares two slices
    * **len()** - find the length of a slice

### 13.3 Go String
A string is a sequence of characters. For example, "Golang" is a string that includes characters: G, o, l, a, n, g.

We use double quotes to represent strings in Go. For example,
```
    var word string = "salom"
```
or ce can use use backtick:
```
    var word string = `salom`
```
#### 13.3.1 Access Characters of String in Go

We know that a string is a sequence of characters. Hence, we can access individual characters of a string.

Just like the Go array, we use index numbers to access string characters. For example,

```
    var word string = "salom"

    fmt.Println(word[0]) // s
```
#### 13.3.2 Find the length of a string

In Go, we use the len() function to find the length of a string. For example,
```
    var word string = "hello, world"

    fmt.Println(len(word)) // 12
```

#### 13.3.4 Join Two Strings Together

In Go, we can use the + operator to join (concatenates) strings together. For example,

```
    var word1 string = "hello"
    var word2 string = "world!"

    sentences = word1 + ", " + word2

    fmt.Prinln(sentence)
```
#### 13.3.5 Golang String Methods

In Go, the strings package provides various methods that can be used to perform different operations on strings.

* **Functions  Descriptions**
    * **Compare()** - compares two strings
    * **Contains()** - checks if a substring is present inside a string
    * **Replaces()** - replaces a substring with another substring
    * **ToLower()** - converts a string to lowercase
    * **ToUpper()** - converts a string to uppercase
    * **Split()** - splits a string into multiple substrings

### 13.4 Go map
In Go, the map data structure stores elements in key/value pairs. Here, keys are unique identifiers that are associated with each value on a map.

Create a map in Golang

The syntax to create a Go map is:
