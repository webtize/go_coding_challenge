# Go Coding Challenges

Welcome to your Go coding challenges! These exercises are designed to test your understanding of Go basics, data structures, structs, and standard libraries.

## Instructions

1.  Navigate to the `challenges` directory.
2.  Open each file (e.g., `basics.go`) and read the comments to understand the task.
3.  Implement the function marked with `TODO`.
4.  Run the tests to verify your solution.

## Running Tests

To run all tests:

```bash
go test ./challenges/...
```

To run tests for a specific challenge (e.g., basics):

```bash
go test ./challenges/basics.go ./challenges/basics_test.go
# OR simply inside the directory
go test -v -run TestProcessNumbers ./challenges
```

## Challenges Overview

1.  **Basics (`basics.go`)**: Work with slices, loops, and basic arithmetic.
2.  **Data Structures (`structures.go`)**: Use maps and string manipulation to count word frequencies.
3.  **Structs & Methods (`structs.go`)**: Build a simple inventory system using structs and methods.
4.  **Standard Libraries (`libraries.go`)**: Parse log entries using `time` and `strings` packages.

Good luck!
