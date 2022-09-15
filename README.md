# Data Structures + Algorithms

## What is Big-O Notation?

This notations focuses on finding the best and worst case that our algorithm takes to execute.

> *We're always looking at the worst case scenario.*

We talk about Big-O when we want to understand:
- how long[^1] an algorithm takes to execute (Time Complexity);
- how much memory is used by an algorithm (Space Complexity);
- says which function, algorithm or code is best;
- describe how efficient we can run our code;

[^1]: The "running time" when using Big-O is not directly linked to the time we know (seconds, milliseconds, microseconds, etc.). Instead we can understand the **time** as the number of operations or steps neeeded to complete a problem of size *n*.

> *Keep this is mind, Big-O doesn't measure things in seconds. Instead, we're focusing on **how quickly our runtime grows**.*

Good to remember:
- What can cause time in a function?
  - operations (+, -, *, /)
  - comparisons (<, >, ==, !=)
  - looping (for, while)
  - outside Function call (functionOutOfScope())
- What causes Space complexity?
  - variables
  - data Structures
  - function Call
  - allocations

- Worst Case: only cares about the worst case.
	- O(N) > O(1)

- Remove Constants: drop the constants.
	- O(1 + N/2 + 100) => O(N)
	- O(2N) => O(N)

- Different terms for inputs
	- ```javascript
		function x(nums, nums2) { 
			// loop for nums

			// loop for nums2
		}

		// O(nums + nums2) => O(A + B)
	  ```
	- ```javascript
		function x(nums, nums2) { 
			// loop for nums
			// loop for nums2 inside nums
		}

		// O(nums * nums2) => O(A * B)
	  ```

- Drop Non Dominants: we care about the most important term
	- O(N + N^2) => O(N^2)
	- O(N^2 + 3N + 100 + N/2) => O(N^2)

<br>

<img src="./bigo.png">

<br>

### Constants - O(1)

no loops.

```go
package main

import "fmt"

func compressFirstAndSecondBox(boxes []string) {
	fmt.Println(arr[0]) // O(1)
	fmt.Println(arr[1]) // O(2)
}

func main() {
	boxes := []string{"a", "b", "c", "d", "e"}

	compressFirstBox(boxes) // O(2) - No matter how big this the number of boxes are, we're only doing one.
}
```

### Logarithmic - O(log N)

usually searching algorithms have log N if they are sorted (Binary Search).

### Linear - O(N)

for loops, while loops through N items.

```go
package main

import "fmt"

func printValues(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main() {
	arr := []string{"a", "b", "c", "d", "e"} // and continue...

	printValues(arr) // O(arr) - It takes linear time do print values
}
```

### Log Linear - O(N log(N))

usually sorting operations.

### Quadratic - O(N ^ 2)

every element in a collection needs to be compared to every other element (Two nested loops).

when you see loops that are nested, we use multiplication.

```go
package main

import "fmt"

func logAllPairsOfArray(arr []string) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			fmt.Println(arr[i], arr[j])
		}
	}
}

func main() {
	arr := []string{"a", "b", "c", "d", "e"}

	logAllPairsOfArray(arr)
}
```

### Exponential - O(2 ^ N)

recursive algorithms that solves a problem of size N.

### Factorial - O(N!)

you are adding a loop for every element.

<br>

### Big O Exercises

#### Exercise 1

```javascript
// What is the Big O of the below function?
// Answer: O(N)

function funChallenge(input) {
  let a = 10; // O(1)
  a = 50 + 3; // O(1)

  for (let i = 0; i < input.length; i++) { // O(N)
    anotherFunction(); // O(N)
    let stranger = true; // O(N)
    a++; // O(N)
  }

  return a; // O(1)
}

// Big O(3 + 4N) => Big O(N)
```

#### Exercise 2

```javascript
// What is the Big O of the below function?
// Answer: O(n)

function anotherFunChallenge(input) {
  let a = 5; // O(1)
  let b = 10; // O(1)
  let c = 50; // O(1)

  for (let i = 0; i < input; i++) { // O(N)
    let x = i + 1; // O(N)
    let y = i + 2; // O(N)
    let z = i + 3; // O(N)
  }

  for (let j = 0; j < input; j++) { // O(N)
    let p = j * 2; // O(N)
    let q = j * 2; // O(N)
  }

  let whoAmI = "I don't know"; // O(1)
}

// Big O(4 + 7N) => Big O(N)
```

## Simplifying Big O

### Drop the constants

Since Big-O is concerned with how quickly our runtime grows, the first rule that you want to remember is to **drop any constants** when you analyze an algorithm:

- $O(4 * N) = O(N)$
- $O(512 * N) = O(N)$
- $O(N / 3) = O((1/3) * N): O(N)$
- $O(5 * N * N) = O(N)$

### Different inputs should have different variables

If an algorithm receives different input parameters (e.g., $f(A, B)$), a different calculation is required for each type of situation:

- $O(f(A + B))$: `for` steps in order
- $O(f(A * B))$: `for` nested steps

In progress... 🏗