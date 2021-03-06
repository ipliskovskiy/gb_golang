package web

import (
	"fmt"
	"net/http"
	"html/template"
	// "gb_golang/internal/tools/stringWorker"
)

type IndexPage struct{
	Title string
	Message string
	BottomInformation string
	// Player1 string
	// Player2 string
}

var (
	tmpl = template.Must(template.ParseFiles("internal/web/templates/index.html"))
	player1 string
	player2 string
)

func handlerIndex(w http.ResponseWriter, r *http.Request){
	fmt.Println("path /")
	fmt.Println(player1)
	data := IndexPage {
		Title: "TicTacToe",
		Message: "Добро пожаловать в увлекательную интерактивную игру крестики нолики.",
		BottomInformation: "Version: Experemental",
	}
		// Player1: r.FormValue("player1"),
		// Player2: r.FormValue("player2"),
	// }
	player1 = r.FormValue("player1")
	player2 = r.FormValue("player2")
	tmpl.Execute(w, data)
}

func testResult(w http.ResponseWriter, r *http.Request) {
	
}

func StartServer() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/testplay", testResult)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("internal/web/templates/css"))))
	http.HandleFunc("/hello", sayhello)
	fmt.Println("Start server...")
	http.ListenAndServe(":8080", nil)
}

func sayhello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Привет!")
}