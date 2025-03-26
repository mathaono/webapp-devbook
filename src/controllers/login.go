package controllers

import "net/http"

// Carrega a tela de login
func LoadLoginScreen(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TELA DE LOGIN"))
}
