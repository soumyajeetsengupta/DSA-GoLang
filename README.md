# DSA Journey in Go: Striver's A 2 Z Overview
This repository documents my learning journey through **Striver's A 2 Z DSA** overview, implemented in ***Go Lang***. The goal is to solidify my understanding of Data Structures and Algorithms through practical implementation and code analysis.

## Striver's A 2 Z DSA overview:

1. This free course provides a well-organized and structured approach to learning DSA fundamentals.
2. High-quality lectures offer comprehensive explanations.
3. While doubt support isn't directly provided, a vibrant community of over 250,000 individuals actively engage in discussions and support each other.
4. What you'll find here:
    1. **Code solutions:** Implementations of various DSA concepts in Go Lang, following Striver's guidance.
    2. **Notes and explanations:** Personal notes and annotations to enhance understanding and memory retention.
    3. **Challenges and practice:** Exercises and problems to test my knowledge and push my learning further.

### Getting started:

1. Clone this repository: `git clone git@github.com:soumyajeetsengupta/DSA-GoLang.git`
2. Open the directory in your preferred code editor.
3. Follow along with the Striver's A 2 Z DSA videos, implementing the concepts and exploring the corresponding code in this repository.

# Go Lang Basics and Syntaxes:

## Data Types in Go

Go provides a variety of built-in data types to represent different kinds of information:

### Basic Types:

* **rune:** Represents a single Unicode code point (character).
* **complex64/complex128:** Stores complex numbers with 32/64-bit floating-point precision.
* **uintptr:** An unsigned integer representing memory addresses.
* **byte:** Represents 8-bit unsigned integers (0-255).
* **int:** Represents signed integers within a specific size range (e.g., 16, 32, 64 bits).
* **uint:** Represents unsigned integers within a specific size range (e.g., 16, 32, 64 bits).
    * **(#unsigned int = uint):** Go implicitly treats `uint` as the equivalent of `unsigned int`.
* **float32/float64:** Represents single/double-precision floating-point numbers.
* **bool:** Represents boolean values (true or false).
* **string:** Represents sequences of Unicode characters.

### Derived Types:

* **Pointers:** Variables storing memory addresses of other variables. Enable reference passing to functions.
* **Arrays:** Fixed-size collections of elements of the same type. Accessed using indices.
* **Structures:** User-defined collections of named variables (fields) of different types.
* **Maps:** Unordered collections of key-value pairs.
* **Interfaces:** Define sets of methods that specific types can implement.

This is just a brief overview of data types in Go. Remember to explore further and dive deeper into each type and its functionalities based on your specific needs.
