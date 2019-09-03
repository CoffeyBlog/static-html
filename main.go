package main

import (
	"fmt"
)

func main() {

	name := "Customer Name"


	tpl := `															// tpl = template 
	<DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF - 8">
	<title> Hello </title>  
	</head>
	<body>
	<h1> ` + name + ` </h1> 
	</body>
	</html>
	`
	fmt.Println(tpl)
}


