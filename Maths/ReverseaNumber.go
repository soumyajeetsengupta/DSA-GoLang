package main;

import (
	"fmt";
);

func main() {
	fmt.Print("User input: ");
	var ip int; 
	fmt.Scan(&ip);
	
	var reverseNum int;

	for i := 1; ip != 0; i++ {
		if i == 1 {
			reverseNum = ip % 10;
		} else {
			reverseNum = (reverseNum * 10) + (ip % 10);
		}

		ip /= 10;
	}

	fmt.Printf("Reverse print of your num ip: %d", reverseNum);
}

/* REVERSING A STRING IN GOLANG
var stringIp string = fmt.Sprintf("%v", ip);
fmt.Println("Reverse print of your num ip: ", reverse(stringIp));

func reverse(s string) string { 
    rns := []rune(s) // convert to rune 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  
        // swap the letters of the string, 
        // like first with last and so on. 
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    // return the reversed string. 
    return string(rns) 
}
*/