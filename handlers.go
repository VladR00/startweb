package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func HandleRegistrationFunc(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "SQL succesfully recorded"}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response = map[string]string{"error": "Only POST method allowed"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var data struct {
		Login    string `json: "login"`
		Password string `json: "password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{"error": "SQL not recorded"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := SaveToSQL(data.Login, data.Password); err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: users.login") {
			response = map[string]string{"error": "That login already exist."}
		} else {
			response = map[string]string{"error": fmt.Sprintf("%w", err)}
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response = map[string]string{"message": "SQL recorded"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func HandleLoginFunc(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Succesfully logining"}
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		response = map[string]string{"error": "Only POST method allowed"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var data struct {
		Login    string `json: "login"`
		Password string `json: "password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = map[string]string{"error": "Error decoding json"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	user, err := Login(data.Login, data.Password)
	if err != nil {
		response = map[string]string{"error": "Your account didn't exist or\nYour password is incorrect"}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	response = map[string]string{"message": fmt.Sprintf("Hi, %s!\nPass:%s\nRegistration time: %s", user.Login, user.Password, time.Unix(user.Time, 0).Format("2006-01-02 15:04"))}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
