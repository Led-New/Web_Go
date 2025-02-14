package routes

import (
	"Site/controlles"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controlles.Index)
	http.HandleFunc("/new", controlles.New)
	http.HandleFunc("/insert", controlles.Insert)
	http.HandleFunc("/delete", controlles.Delete)
	http.HandleFunc("/edit", controlles.Edit)
	http.HandleFunc("/update", controlles.Update)
}
