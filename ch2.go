package main
import "fmt";

var x int = 10;

//don't have to put type on left side if the expected type is the one on the right side
var y = 5;

//can delclare multiple variable at once and they can be of the same type or different

var fee, fam = 10, "piglets"

//can use declaration list to list multiple variables

func myfunction() {
	fmt.Println("x")
	//inside a function you can use := (not legal outside of the functions)
	var b = 1;
	d := 2;

}

//avoid declaring variables outside of functions. When outside a function it is difficult to detect changes made to it 
