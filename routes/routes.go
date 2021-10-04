package routes

import (
	"alura/controllers"
	"fmt"
	"net/http"
)

func CarregaRotas() {
	fmt.Println("Inicializando servidor na porta 8000")
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)
}
