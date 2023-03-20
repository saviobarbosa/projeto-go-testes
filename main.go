package main

import (
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/routes"
)

//Apenas um comentario por aqui

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
