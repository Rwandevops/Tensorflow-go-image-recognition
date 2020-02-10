package verif

import (
	"database/sql"
	"log"
	"regexp"
	"tensorflow-back/base"
	"tensorflow-back/models"

	"golang.org/x/crypto/bcrypt"
)

// RegisterVerif fonction qui vérifie les données avant l'enregistrement d'un utilisateur
var RegisterVerif = func(db *sql.DB, user models.User) (string, bool) {
	// Champs tous remplis?
	if user.Nom == "" || user.Prenom == "" || user.Email == "" || user.Password == "" {
		err := "Remplissez tous les champs"
		return err, false
	}
	// RegExp Email
	var validEmail = regexp.MustCompile(`^[^\W][a-zA-Z0-9_]+(\.[a-zA-Z0-9_]+)*\@[a-zA-Z0-9_]+(\.[a-zA-Z0-9_]+)*\.[a-zA-Z]{2,4}$`)
	if !validEmail.MatchString(user.Email) {
		err := "Email non valide"
		return err, false
	}
	// Email déjà enregistré?
	count := base.EmailCount(db, user.Email)
	if count == true {
		err := "email déjà utilisé"
		return err, false
	}
	// Password matches?
	if user.Password != user.Password2 {
		err := "mots de passe différents"
		return err, false
	}
	err := "données OK"
	return err, true
}

// LoginVerif fonction qui vérifie les identifiants pour le login
var LoginVerif = func(db *sql.DB, send models.User) models.Reponse {
	reponse := base.GetUserByEmailQuery(db, send.Email)
	fromBase := reponse.User
	if reponse.Message == "utilisateur trouvé" {
		if CheckPasswordHash(send.Password, fromBase.Password) {
			reponse.Message = "login OK"
		} else {
			fromBase.Nom = ""
			fromBase.Prenom = ""
			reponse.Message = "identifiants incorrects"
		}
	} else {
		reponse.Message = "email non référencé"
	}
	return reponse
}

// CheckPasswordHash fonction qui compare les mots de passe bcrypt
func CheckPasswordHash(sentPassword, hashBDD string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashBDD), []byte(sentPassword))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
