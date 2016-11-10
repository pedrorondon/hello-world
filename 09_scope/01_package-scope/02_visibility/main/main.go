package main

import (
	"fmt"

	"github.com/pedrorondon/udemytraining/09_scope/01_package-scope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName, "a") //Myname is a variable
	vis.PrintVar()               //function as indicated by parens at the end
}
