package base

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"tensorflow-back/models"

	_ "github.com/go-sql-driver/mysql" // OSEF
)

// DatabaseConnection fonction qui créee la connexion à la BDD
var DatabaseConnection = func() *sql.DB {

	conn, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/tensorflow")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}

// EmailCount fonction qui compte le nb d'occurence de data dans email
var EmailCount = func(db *sql.DB, data string) bool {
	var count int

	rows, err := db.Query("SELECT COUNT(*) FROM users WHERE email = ?", data)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	if count == 0 {
		return false
	}
	return true
}

// InsertQuery fonction qui demande à la BDD la liste des utilisateurs
var InsertQuery = func(db *sql.DB, newuser models.User) string {
	query, err := db.Prepare("INSERT INTO users (nom, prenom, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	} else {
		_, err = query.Exec(newuser.Nom, newuser.Prenom, newuser.Email, newuser.Password)
		if err != nil {
			panic(err.Error())
		}
		result := "Insert OK"
		return result
	}
}

// GetUserByEmailQuery fonction qui rapatrie les données de l'utilisateur
var GetUserByEmailQuery = func(db *sql.DB, email string) models.Reponse {
	var user models.User
	var message string
	var reponse models.Reponse
	query := `SELECT id, nom, prenom FROM users WHERE email = ?`
	row := db.QueryRow(query, email)
	switch err := row.Scan(&user.ID, &user.Nom, &user.Prenom); err {
	case sql.ErrNoRows:
		message = "Email non trouvé dans la BDD"
		user.ID = ""
		user.Nom = ""
		user.Prenom = ""
		user.Email = ""
	case nil:
		user.Email = email
		message = "utilisateur trouvé"
	default:
		panic(err)
	}
	reponse.User = user
	reponse.Message = message
	return reponse
}

// UpdateQuery fonction qui modifie les données de l'utilisateur
var UpdateQuery = func(db *sql.DB, user models.User) bool {
	query, err := db.Prepare("UPDATE users SET nom = ?, prenom = ?, password = ? WHERE email = ?")
	if err != nil {
		panic(err.Error())
	} else {
		_, err = query.Exec(user.Nom, user.Prenom, user.Password, user.Email)
		if err != nil {
			panic(err.Error())
		}
		return true
	}
}

// DeleteQuery fonction qui modifie les données de l'utilisateur
var DeleteQuery = func(db *sql.DB, email string) bool {
	query, err := db.Prepare("DELETE users WHERE email = ?")
	if err != nil {
		panic(err.Error())
	} else {
		_, err = query.Exec(email)
		if err != nil {
			panic(err.Error())
		}
		return true
	}
}

// GetUsersByID fonction qui renvoie tous les utilisateurs selon un critère
var GetUsersByID = func(db *sql.DB, id string) []models.User {
	var users []models.User
	var user models.User
	var query = `SELECT * FROM users WHERE id = ?`
	rows, err := db.Query(query, id)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Nom, &user.Prenom, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

// GetUsersByFirstName fonction qui renvoie tous les utilisateurs selon un critère
var GetUsersByFirstName = func(db *sql.DB, nom string) []models.User {
	var users []models.User
	var user models.User
	var query = `SELECT * FROM users WHERE nom = ?`
	rows, err := db.Query(query, nom)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Nom, &user.Prenom, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

// GetUsersByLastName fonction qui recherche tous les utilisateurs selon leur prenom
var GetUsersByLastName = func(db *sql.DB, prenom string) []models.User {
	var users []models.User
	var user models.User
	var query = `SELECT * FROM users WHERE prenom = ?`
	rows, err := db.Query(query, prenom)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Nom, &user.Prenom, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}

// GetUsersByEmail fonction qui recherche tous les utilisateurs selon l'email
var GetUsersByEmail = func(db *sql.DB, email string) []models.User {
	var users []models.User
	var user models.User
	var query = `SELECT * FROM users WHERE email = ?`
	rows, err := db.Query(query, email)
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Nom, &user.Prenom, &user.Email)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users
}
