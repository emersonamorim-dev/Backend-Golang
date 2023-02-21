package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Define a estrutura dos dados do usuário
type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// Cria um slice de usuários para simular um banco de dados
var users []User

// Endpoint para retornar todos os usuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

// Endpoint para retornar um usuário específico
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Pega os parâmetros da requisição
	for _, user := range users {
		if user.ID == params["id"] {
			json.NewEncoder(w).Encode(user)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

// Endpoint para criar um novo usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

// Endpoint para atualizar um usuário existente
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users[index].Name = params["name"]
			users[index].Email = params["email"]
			users[index].Password = params["password"]
			json.NewEncoder(w).Encode(users[index])
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

// Endpoint para deletar um usuário
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, user := range users {
		if user.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(users)
}

// Função principal
func main() {
	// Inicializa o router
	router := mux.NewRouter()

	// Cria alguns usuários de exemplo
	users = append(users, User{ID: "1", Name: "Emerson", Email: "emerson@example.com", Password: "password"})
	users = append(users, User{ID: "2", Name: "Amorim", Email: "amorim@example.com", Password: "password"})
	users = append(users, User{ID: "3", Name: "Luiz", Email: "luiz@example.com", Password: "password"})

	// Define as rotas para cada endpoint da API
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", GetUser).Methods("GET")
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Inicia o servidor
	log.Fatal(http.ListenAndServe(":8018", router))
}
