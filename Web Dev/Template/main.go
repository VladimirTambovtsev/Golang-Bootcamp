package main

import (
	"fmt"
)

func main() {

	name := "Man"

	tmp := `
	<!DOCTYPE html>
	<html>
	<head>
		<title></title>
	</head>
	<body>
		<p>Hey ` + name + `</p>
	</body>
	</html>
	`

	fmt.Println(tmp)
}
