package main;

import ("fmt"; "math");

func main() {
	var input, count int;
	fmt.Print("User input: ");
	fmt.Scan(&input);

	// BRUTEFORCE - LOOP
	for i := input; i != 0; i /= 10 {
		count += 1;
	}
	fmt.Printf("Number of digits in your input %d are %d.", input, count);

	// OPTIMAL - STRING
	// var stringInput string = fmt.Sprintf("%v", input);
	// fmt.Printf("Number of digits in your input %d are %d.", input, len(stringInput));

	// MOST OPTIMAL - LOG to Base 10 >> Because 10 to reach n 10 will have to be raised to closest power to that power we will just + 1.
	// fmt.Printf("Number of digits in your input %d are %d.", input, int(math.Log10(float64(input))) + 1);
}