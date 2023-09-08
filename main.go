package main

import (
	"github/phat/product/routes"
)


func main() {
	routes.Routes().Run(":3000");
}