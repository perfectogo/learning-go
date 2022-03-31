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
#####Rules of naming Variables
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
    ________________________________________________
    |           Operators          |     Example   |
    |------------------------------+----------------
    | == (equal to)                |     a == b    | 
    | != (not equal to)            |     a != b    |
    | > (greater than)             |     a > b     |
    | < (less than)                |     a < b     |
    |>= (greater than or equal to) |     a >= b    |
    |<= (less than or equal to)    |     a <= b    |
    ------------------------------------------------
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