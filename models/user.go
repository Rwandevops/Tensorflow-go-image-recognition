package models

// User encodage des utilisateurs pour la BDD
type User struct {
	ID        string `json:"id"`
	Nom       string `json:"nom"`
	Prenom    string `json:"prenom"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Password2 string `json:"password2"`
}

// Reponse réponse http contenant un user et un message
type Reponse struct {
	User    User   `json:"user"`
	Message string `json:"message"`
}
