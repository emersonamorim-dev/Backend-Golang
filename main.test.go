package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateUser(t *testing.T) {
	users = []User{}

	// Cria um novo usuário
	newUser := User{ID: "1", Name: "Test User", Email: "testuser@example.com", Password: "testpassword"}
	jsonUser, _ := json.Marshal(newUser)
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}

	// Cria um response recorder para gravar o response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)

	// Chama a função CreateUser
	handler.ServeHTTP(rr, req)

	// Verifica se o status code retornado é 200 (OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("CreateUser retornou um status code inválido: esperado %v, mas obteve %v", http.StatusOK, status)
	}

	// Verifica se o usuário foi adicionado corretamente
	expectedUser := newUser
	expectedUserJSON, _ := json.Marshal(expectedUser)
	if rr.Body.String() != string(expectedUserJSON) {
		t.Errorf("CreateUser não adicionou o usuário corretamente. Esperado: %v, mas obteve %v", expectedUserJSON, rr.Body.String())
	}
}

func TestGetUser(t *testing.T) {
	// Cria um novo usuário
	newUser := User{ID: "1", Name: "Test User", Email: "testuser@example.com", Password: "testpassword"}
	users = []User{newUser}

	// Cria uma requisição GET para o usuário criado
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Cria um response recorder para gravar o response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)

	// Cria um router e adiciona a rota para a função GetUser
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", handler)

	// Chama a função GetUser
	router.ServeHTTP(rr, req)

	// Verifica se o status code retornado é 200 (OK)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("GetUser retornou um status code inválido: esperado %v, mas obteve %v", http.StatusOK, status)
	}

	// Verifica se o usuário retornado é o esperado
	expectedUser := newUser
	expectedUserJSON, _ := json.Marshal(expectedUser)
	var actualUser User
	json.Unmarshal(rr.Body.Bytes(), &actualUser)
	actualUserJSON, _ := json.Marshal(actualUser)
	if !reflect.DeepEqual(actualUser, expectedUser) {
		t.Errorf("GetUser não retornou o usuário correto. Esperado: %v, mas obteve %v", expectedUserJSON, actualUserJSON)
	}
}
