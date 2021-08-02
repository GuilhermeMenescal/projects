package config

import ("net/http"
		"log"
		"encoding/json"
		"github.com/gorilla/mux")


type User struct {
	
	FirstName string     `json:"firstname" validate:"required,firstname"`
	LastName string      `json:"lastname" validate:"required,lastname"`
	UserName string		 `json:"username" validate:"required"`
	Email string         `json:"email" validate:"required,email"`
	Password string      `json:"pass" validate:"required,password"`
	Cellphone string     `json:"cellphone" validate:"required,cellphone"`
	Active bool			 `json:"active,omitempty"`
	

}

type UserLD struct {

	Email string         `json:"email" validate:"required,email"`
	Password string      `json:"pass" validate:"required,password"`

}

func InsertUser(w http.ResponseWriter, r *http.Request) {

	var u User
	
	err := DecodeJson(r.Body, &u)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	err = validate.Struct(u)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	user,_:=VerifyUser(u.Email)

	if user.Email == u.Email{
		log.Printf("This user email already exists")
		w.Write([]byte("This user email already exists"))
		return
	}

	user,_ = VerifyUsername(u.UserName)

	if user.UserName == u.UserName{
		log.Printf("This username already exists")
		w.Write([]byte("This username already exists"))
		return
	}

	u.Active = true

	err = RegisterUser(u)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["register-error"]))
		return	
	}

	w.Write([]byte(responses["register-success"]))
	return

}


func FindAllUsers(w http.ResponseWriter, r *http.Request){

	users, err := FindUsers()

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	usersJson, err := json.Marshal(users) 

	if err != nil{
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	//w.Write([]byte(responses["find-success"]))
	log.Printf("Success")
	w.Write([]byte(usersJson))
	
	return 
}

func FindUsersByName(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
    name := vars["fullname"]
    rs, err := FindByName(name)

    if err != nil {
    	log.Printf("[ERROR]: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

    bs, err := json.Marshal(rs)

    if err != nil {
    	log.Printf("[ERROR]: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	if rs == nil{
		log.Printf("User not found %v")
		w.Write([]byte("User not found"))
		return 
	}
	log.Printf("Success")
    w.Write(bs)

	}


func RemoveUser(w http.ResponseWriter, r *http.Request){
	var u UserLD
	var user User
	
	err := DecodeJson(r.Body, &u)

	if err != nil {
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	err = validate.Struct(u)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	user, err = VerifyUser(u.Email)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte("Erro ao deletar"))
		return	
	}

	if u.Password != user.Password{
		log.Printf("Senha incorreta", err)
		w.Write([]byte("Wrong password"))
		return	
	}

	err = DeleteUser(u)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte("Erro ao deletar"))
		return	
	}


	w.Write([]byte(responses["delete-success"]))
	return
}

func UpdateOneUser(w http.ResponseWriter, r *http.Request){

	var u User
	var user User


	err := DecodeJson(r.Body, &u)

	u.Active=true

	if err != nil {
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	err = validate.Struct(u)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	user, err = VerifyUser(u.Email)

	if err != nil {
		log.Printf("Email de usuário não encontrado", err)
		w.Write([]byte("We couldn't find you. You can't change your email"))
		return	
	}


	if u.Email != user.Email{
		log.Printf("Tentativa de alteração de chave primária", err)
		w.Write([]byte("You can't change your email"))
		return	
	}

	if u.UserName != user.UserName{
		log.Printf("Tentativa de alteração de chave primária", err)
		w.Write([]byte("You can't change your username"))
		return	
	}


	if u == user{
		log.Printf("Nenhuma alteração encontrada", err)
		w.Write([]byte("Nothing to update"))
		return	
	}

		if user.Active == false{
		log.Printf("Usuário está desativado", err)
		w.Write([]byte("This user is already deactivated, login to activate again"))
		return	
	}

	err = UpdateUser(u, user)

	if err != nil {
		log.Printf("Erro ao atualizar", err)
		w.Write([]byte("Something went wrong, sorry we couldn't update your account"))
		return	
	}

	log.Printf("Usuário atualizado")
	w.Write([]byte(responses["update-success"]))
	return

}

func DeactivateUser(w http.ResponseWriter, r *http.Request){
	var u UserLD
	var user User
	var deadUser User
	
	err := DecodeJson(r.Body, &u)

	if err != nil {
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	err = validate.Struct(u)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	user, err = VerifyUser(u.Email)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		w.Write([]byte("User email not found"))
		return	
	}

	if u.Password != user.Password{
		log.Printf("Wrong password", err)
		w.Write([]byte("Wrong password"))
		return	
	}

	if user.Active == false{
		log.Printf("This user doesn't exist anymore", err)
		w.Write([]byte("This user doesn't exist anymore"))
		return	
	}

	deadUser = user
	deadUser.Active = false

	err = UpdateUser(deadUser, user)

	if err != nil {
		log.Printf("Erro ao desativar: %v", err)
		w.Write([]byte("Erro ao desativar"))
		return	
	}

	log.Printf("Usuário desativado")
	w.Write([]byte(responses["delete-success"]))
	return
}


