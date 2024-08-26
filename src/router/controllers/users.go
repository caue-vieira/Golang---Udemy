package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuários"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuário por id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alterando usuário"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário"))
}
