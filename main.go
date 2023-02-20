package main

import (
	"api/database"
	"api/routes"
	"fmt"
)




func main() {

	database.ConectaDB()

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
	
}