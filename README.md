# Data Structures + Algorithms

## Big O

Keep this is mind, Big O doesn't measure things in seconds. Instead, we're focusing on **how quickly our runtime grows**.

<img src="./bigo.png">

 - **Constants - O(1)**: no loops.
 - **Logarithmic - O(log N)**: usually searching algorithms have log N if they are sorted (Binary Search).

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

 - **Log Linear - O(N log(N))**: usually sorting operations.
 - **Quadratic - O(N ^ 2)**: every element in a collection needs to be compared to every other element (Two nested loops).
 - **Exponential - O(2 ^ N)**: recursive algorithms that solves a problem of size N.
 - **Factorial - O(N!)**: you are adding a loop for every element.

In progress... 🏗