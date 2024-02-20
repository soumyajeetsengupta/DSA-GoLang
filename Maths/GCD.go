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

	// Bruteforce
	var min int;
	if num1 < num2 {
		min = num1;
	} else {
		min = num2;
	}

	for i := 2; i <= min; i++ {
		if num1 % i == 0 && num2 % i == 0 {
			min = i;
			break;
		}
	}

	fmt.Println("The GCD is: ", min);

	/* Most Optmial
	for ;num1 % num2 != 0; {
		var temp int = num1;
		num1 = num2;
		num2 = temp % num2;
	}

	fmt.Println("The GCD is: ", num2);
	*/
}