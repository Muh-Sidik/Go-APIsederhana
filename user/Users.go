package user

import (
	"encoding/json"
	"net/http"
)

type UserData struct {
	Id string
	Name string
	Grade int
}

var data = []UserData{
	UserData{"001", "Muhammad Sidik", 21},
	UserData{"002", "Tri Apria Ningsih", 19},
	UserData{"003", "Siapa aja", 20},
}

func Users(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func User(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var id = r.FormValue("id")
		var result []byte

		var err error

		for _, each := range data {
			if each.Id == id{
				result, err = json.Marshal(each)
				
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}

		http.Error(w, "User Not Found", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}