package handlers

import (
	"encoding/json"
	"net/http"
	"tensorflow-back/base"
	"tensorflow-back/models"
	"tensorflow-back/tensorflow"
	"tensorflow-back/verif"

	"golang.org/x/crypto/bcrypt"
)

var conn = base.DatabaseConnection()

// GetUsers fonction qui renvoie la liste des utilisateurs et leurs données
// Accessible uniquement par les administrateurs
var GetUsers = func(w http.ResponseWriter, r *http.Request) {
	tri := r.URL.Query().Get("tri")
	value := r.URL.Query().Get("value")
	var users []models.User
	switch tri {
	case "id":
		users = base.GetUsersByID(conn, value)
	case "firstName":
		users = base.GetUsersByFirstName(conn, value)
	case "lastName":
		users = base.GetUsersByLastName(conn, value)
	default:
		users = base.GetUsersByEmail(conn, value)
	}
	json.NewEncoder(w).Encode(users)
}

// CreateUser fonction qui créé une entrée user dans la BDD
var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	var newuser models.User
	_ = json.NewDecoder(r.Body).Decode(&newuser)
	err, ok := verif.RegisterVerif(conn, newuser)
	if ok == true {
		// Hash du password
		bytes, _ := bcrypt.GenerateFromPassword([]byte(newuser.Password), 14)
		hash := string(bytes)
		newuser.Password = hash
		result := base.InsertQuery(conn, newuser)
		json.NewEncoder(w).Encode(result)
	} else {
		json.NewEncoder(w).Encode(err)
	}
}

// LoginUser fonction qui gère le login
var LoginUser = func(w http.ResponseWriter, r *http.Request) {
	// Ajouter Token
	var sent models.User
	json.NewDecoder(r.Body).Decode(&sent)
	reponse := verif.LoginVerif(conn, sent)
	println(reponse.Message)
	if reponse.Message == "login OK" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(reponse.User)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(reponse.Message)
	}
	// json.NewEncoder(w).Encode(reponse)
}

// GetUser fonction qui affiche les données d'un utilisateur
var GetUser = func(w http.ResponseWriter, r *http.Request) {
	// Ajouter session
	email := r.URL.Query().Get("email")
	reponse := base.GetUserByEmailQuery(conn, email)
	json.NewEncoder(w).Encode(reponse)
}

// UpdateUser fonction qui modifie les données d'un utilisateur
var UpdateUser = func(w http.ResponseWriter, r *http.Request) {
	// Ajouter session
	var sent models.User
	json.NewDecoder(r.Body).Decode(&sent)
	result := base.UpdateQuery(conn, sent)
	json.NewEncoder(w).Encode(result)
}

// DeleteUser fonction qui supprime un utilisateur
var DeleteUser = func(w http.ResponseWriter, r *http.Request) {
	// Ajouter session
	var email string
	json.NewDecoder(r.Body).Decode(&email)
	result := base.DeleteQuery(conn, email)
	json.NewEncoder(w).Encode(result)
}

// Tensorflow fonction qui appelle Tensorflow
var Tensorflow = func(w http.ResponseWriter, r *http.Request) {
	imageLink := r.URL.Query().Get("imageLink")
	testLink := imageLink[len(imageLink)-4:]
	if testLink == ".jpg" || testLink == ".png" {
		result := tensorflow.TensorflowMain(imageLink)
		json.NewEncoder(w).Encode(result)
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode("Requête non valide")
}
