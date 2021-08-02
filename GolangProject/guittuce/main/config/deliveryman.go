package config

import ("net/http"
		"log"
		"encoding/json"
		"github.com/gorilla/mux")


type Deliveryman struct {
	
	FirstName string     `json:"firstname" validate:"required,firstname"`
	LastName string      `json:"lastname" validate:"required,lastname"`
	UserName string		 `json:"username" validate:"required"`
	Email string         `json:"email" validate:"required,email"`
	Password string      `json:"pass" validate:"required,password"`
	Cellphone string     `json:"cellphone" validate:"required,cellphone"`
	Active bool			 `json:"active,omitempty"`

}

type DeliverymanLD struct {

	Email string         `json:"email" validate:"required,email"`
	Password string      `json:"pass" validate:"required"`

}


func InsertDeliveryman(w http.ResponseWriter, r *http.Request) {

	var d Deliveryman
	
	err := DecodeJson(r.Body, &d)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	err = validate.Struct(d)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	deliveryman,_:=VerifyDeliveryman(d.Email)

	if deliveryman.Email == d.Email{
		log.Printf("This user email already exists")
		w.Write([]byte("This user email already exists"))
		return
	}

	deliveryman,_ = VerifyDeliveryName(d.UserName)

	if deliveryman.UserName == d.UserName{
		log.Printf("This username already exists")
		w.Write([]byte("This username already exists"))
		return
	}

	d.Active = true

	err = RegisterDeliveryman(d)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["register-error"]))
		return	
	}

	w.Write([]byte(responses["register-success"]))
	return

}


func FindAllDeliverymen(w http.ResponseWriter, r *http.Request){

	deliverymen, err := FindDeliverymen()

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	deliverymenJson, err := json.Marshal(deliverymen) 

	if err != nil{
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	log.Printf("Success")
	w.Write([]byte(deliverymenJson))
	return 
}

func FindDeliverymenByName(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
    name := vars["fullname"]
    rs, err := FindByNameD(name)

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
		w.Write([]byte("Deliveryman not found"))
		return 
	}
	
	log.Printf("Success")
    w.Write(bs)

	}


func RemoveDeliveryman(w http.ResponseWriter, r *http.Request){
	var d DeliverymanLD
	var deliveryman Deliveryman
	
	err := DecodeJson(r.Body, &d)

	if err != nil {
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	err = validate.Struct(d)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	deliveryman, err = VerifyDeliveryman(d.Email)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte("Erro ao deletar"))
		return	
	}

	if d.Password != deliveryman.Password{
		log.Printf("Senha incorreta", err)
		w.Write([]byte("Wrong password"))
		return	
	}

	err = DeleteDeliveryman(d)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte("Erro ao deletar"))
		return	
	}


	w.Write([]byte(responses["delete-success"]))
	return
}

func UpdateOneDeliveryman(w http.ResponseWriter, r *http.Request){

	var d Deliveryman
	var deliveryman Deliveryman

	err := DecodeJson(r.Body, &d)

	d.Active = true

	if err != nil {
		w.Write([]byte(responses["miss-json"]))
		return
	}

	err = validate.Struct(d)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	deliveryman, err = VerifyDeliveryman(d.Email)


	if err != nil {
		log.Printf("Email de usuário não encontrado", err)
		w.Write([]byte("We couldn't find you. You can't change your email"))
		return	
	}

	if d.Email != deliveryman.Email{
		log.Printf("Tentativa de alteração de chave primária", err)
		w.Write([]byte("You can't change your email"))
		return	
	}

	if d.UserName != deliveryman.UserName{
		log.Printf("Tentativa de alteração de chave primária", err)
		w.Write([]byte("You can't change your username"))
		return	
	}

	if deliveryman.Active == false{
		log.Printf("Usuário está desativado", err)
		w.Write([]byte("This user is already deactivated, login to activate again"))
		return	
	}

	if d == deliveryman{
		log.Printf("Nenhuma alteração encontrada", err)
		w.Write([]byte("Nothing to update"))
		return	
	}

	

	err = UpdateDeliveryman(d, deliveryman)

	if err != nil {
		log.Printf("Erro ao atualizar", err)
		w.Write([]byte("Something went wrong, sorry we couldn't update your account"))
		return	
	}

	log.Printf("Usuário atualizado")
	w.Write([]byte(responses["update-success"]))
	return

}

func DeactivateDeliveryman(w http.ResponseWriter, r *http.Request){
	var d DeliverymanLD
	var deliveryman Deliveryman
	var deadDeliveryman Deliveryman
	
	err := DecodeJson(r.Body, &d)

	if err != nil {
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	err = validate.Struct(d)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	deliveryman, err = VerifyDeliveryman(d.Email)

	if err != nil {
		log.Printf("[ERROR]: %v", err)
		w.Write([]byte("User email not found"))
		return	
	}

	if d.Password != deliveryman.Password{
		log.Printf("Wrong password", err)
		w.Write([]byte("Wrong password"))
		return	
	}

	if deliveryman.Active == false{
		log.Printf("This user doesn't exist anymore", err)
		w.Write([]byte("This user doesn't exist anymore"))
		return	
	}
	deadDeliveryman = deliveryman
	deadDeliveryman.Active = false

	err = UpdateDeliveryman(deadDeliveryman,deliveryman)

	if err != nil {
		log.Printf("Erro ao deletar: %v", err)
		w.Write([]byte("Erro ao desativar"))
		return	
	}

	log.Printf("Usuário desativado")
	w.Write([]byte(responses["delete-success"]))
	return
}






