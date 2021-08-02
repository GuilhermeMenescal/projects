package test

import("testing"
	   "net/http"
	   "main/config"
	   "bytes"
	   )

//--------------------------------------------------Informações requeridas------------------------------------------------------------------------------//
func TestInsertUserMissFirstName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{
				"lastname": "rob",
				"username": "Robson",
				"email": "robso@orr.com",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserMissLastName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{	"firstname": "roberta",
				"username": "Roberta",
				"email": "roberta@orr.com",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserMissUserName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{	"firstname": "Roberta",
				"lastname": "Berto",
				"email": "roberta@orr.com",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserMissEmail(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{	"firstname": "Roberta",
				"lastname": "Berto",
				"username": "Robertina",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserMissPassword(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{	"firstname": "Roberta",
				"lastname": "Berto",
				"username": "Robertina",
				"email": "roberta@orr.com",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserMissPhone(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{
				"firstname": "robert",
				"lastname": "rob",
				"username": "Robson",
				"email": "robson@orr.com",
				"pass": "acb4789"

			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}


//-----------------------------------------------------Usuário já cadastrado --------------------------------------------------------------//

func TestInsertUserEmailExists(t *testing.T) {

	expected := `This user email already exists`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebo",
				"email": "caio@cebola.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUsernameExists(t *testing.T) {

	expected := `This username already exists`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebolas.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//-------------------------------------------------------Informações invalidas-----------------------------------------------------------------------------//

func TestInsertUserInvalidFirstName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai045789",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserInvalidLastName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Raissa",
				"lastname": "Galves12",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai045789",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserInvalidEmail(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalvesgmail",
				"pass": "rai045789",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserInvalidShortPassword(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai04",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserInvalidBigPassword(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai041547896321",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertUserInvalidPhone(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai0412",
				"cellphone": "98756-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//---------------------------------------------------Informações faltando para atualizar--------------------------------------------------------------------------//
func TestUpdateUserMissFirstName(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserMissLastName(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Caio",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserMissUserName(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"email": "caio@cebola.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserMissEmail(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserMissPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserMissPhone(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"pass": "cebolasroxa"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//-----------------------------------------------------------Tentativa de alteração de atributos identificadores-----------------------------------------------------//

func TestUpdateUserEmail(t *testing.T) {
	client := &http.Client{}

	expected := `We couldn't find you. You can't change your email`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebolaroxa.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserUsername(t *testing.T) {
	client := &http.Client{}

	expected := `You can't change your username`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "CebolaRoxa",
				"email": "caio@cebola.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserDeactivated(t *testing.T) {
	client := &http.Client{}

	expected := `This user is already deactivated, login to activate again`

	user := `{
				"firstname": "Batata",
				"lastname": "Capota",
				"username": "Batatinha",
				"email": "batata@morta.com",
				"pass": "batatamorta",
				"cellphone": "55(61)96778-5678"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserNothingToUpdate(t *testing.T) {
	client := &http.Client{}

	expected := `Nothing to update`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"pass": "cebolasroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserShortPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"pass": "roxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateUserBigPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"pass": "cebolasroxaserosas",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/users/updateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//----------------------------------------------------------------Informações faltando para desativar----------------------------------------------------------------------//

func TestDeactivateUserMissEmail(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{
				"pass": "1234567"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/users/deactivateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateUserMissPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	user := `{
				
				"email": "caio@cebola.com"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/users/deactivateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateUserIncorrectEmail(t *testing.T) {
	client := &http.Client{}

	expected := `User email not found`

	user := `{	
				"email": "caio@cebolas.com",
				"pass": "cebolasroxa"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/users/deactivateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateUserIncorrectPassword(t *testing.T) {
	client := &http.Client{}

	expected := `Wrong password`

	user := `{	
				"email": "caio@cebola.com",
				"pass": "cebolaroxa"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/users/deactivateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateUserAlreadyDeactivated(t *testing.T) {
	client := &http.Client{}

	expected := `This user doesn't exist anymore`

	user := `{	
				"email": "batata@morta.com",
				"pass": "batata"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/users/deactivateuser", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//----------------------------------------------------------Falhas de busca---------------------------------------------------------------------------//


func TestFindUserNotFound(t *testing.T) {

	expected := `User not found`

	resp, _ := http.Get("http://"+config.SERVER_HOST+"/users/finduserbyname/Garoto")

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}




