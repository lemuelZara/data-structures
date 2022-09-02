# Data Structures + Algorithms

## Big O

Keep this is mind, Big O doesn't measure things in seconds. Instead, we're focusing on **how quickly our runtime grows**.

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

In progress... 🏗