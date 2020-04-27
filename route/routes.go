package routes

import (
	. "API_go/user"
	"fmt"
	"net/http"
)


func WebRoutes()  {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "<h1>Server Go</h1>")
		
	})
	http.HandleFunc("/users", Users)
	http.HandleFunc("/user", User)

	servers()
}

func servers()  {
	fmt.Println("Go Web Server starting at http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}