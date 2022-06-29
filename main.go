package main

import (
	"fmt"

	"github.com/ruyoutor/go-api-rest/database"
	"github.com/ruyoutor/go-api-rest/models"
	"github.com/ruyoutor/go-api-rest/routes"
)

func main() {

	models.Personalities = []models.Personality{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor rest")
	database.DatabaseConnect()
	routes.HandleRequest()
}
