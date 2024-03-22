package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User
var lastUserID uint = 0

func main() {
	NetHttp()
}

func NetHttp() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(users)
		case http.MethodPost:
			user := User{}
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			lastUserID++
			user.ID = lastUserID
			users = append(users, user)
			w.WriteHeader(http.StatusAccepted)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Path[len("/users/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			for _, user := range users {
				if int(user.ID) == id {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(user)
					return
				}
			}
			w.WriteHeader(http.StatusNotFound)
		case http.MethodPut:
			userTemp := User{}
			if err := json.NewDecoder(r.Body).Decode(&userTemp); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			for i, user := range users {
				if int(user.ID) == id {
					user.Username = userTemp.Username
					user.Email = userTemp.Email
					users[i] = user

					w.WriteHeader(http.StatusAccepted)
					json.NewEncoder(w).Encode(users[i])
					return
				}
			}

			w.WriteHeader(http.StatusBadRequest)
			res := errorResponseMap("user with id did not exist")
			json.NewEncoder(w).Encode(res)
		case http.MethodDelete:
			for i, user := range users {
				if int(user.ID) == id {
					users = append(users[:i], users[i+1:]...)
					w.WriteHeader(http.StatusAccepted)
					json.NewEncoder(w).Encode(users)
					return
				}
			}

			w.WriteHeader(http.StatusBadRequest)
			res := errorResponseMap("user with id did not exist")
			json.NewEncoder(w).Encode(res)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listen to port 8080")
	}
}

func errorResponseMap(errorText string) map[string]interface{} {
	response := map[string]interface{}{
		"success": false,
		"error":   errorText,
	}
	return response
}
