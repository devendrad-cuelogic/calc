package main

import (
	"fmt"
	"github.com/devendra-cuelogic/calculator"
	"strconv"
)

func main() {
	fmt.Printf("Enter two numbers along with the operation to performi\n");
	fmt.Printf("Choose any one of Sum, Sub, Multply, Divide\n");
	var number_1;
	fmt.Scanf("%s", &number_1);
	var number_2;
	fmt.Scanf("%s", &number_2);

	var num_1 := strconv.Atoi(number_1);
	var num_2 := strconv.Atoi(number_2);
	

}
