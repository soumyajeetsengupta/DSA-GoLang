package main;

import (
	"fmt";
)

func main() {
	var num1, num2 int;
	fmt.Print("Num1: ");
	fmt.Scanf("%d\n", &num1);
	fmt.Print("Num2: ");
	fmt.Scanf("%d\n", &num2);

	for ;num1 % num2 != 0; {
		var temp int = num1;
		num1 = num2;
		num2 = temp % num2;
	}

	fmt.Println("The GCD is: ", num2);
}