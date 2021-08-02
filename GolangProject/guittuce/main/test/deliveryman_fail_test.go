package test

import("testing"
	   "net/http"
	   "main/config"
	   "bytes")

//--------------------------------------------------Informações requeridas------------------------------------------------------------------------------//
func TestInsertDeliverymanMissFirstName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{
				"lastname": "rob",
				"username": "Robson",
				"email": "robso@orr.com",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanMissLastName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{	"firstname": "roberta",
				"username": "Roberta",
				"email": "roberta@orr.com",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanMissUserName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{	"firstname": "Roberta",
				"lastname": "Berto",
				"email": "roberta@orr.com",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanMissEmail(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{	"firstname": "Roberta",
				"lastname": "Berto",
				"username": "Robertina",
				"pass": "acb4789",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanMissPassword(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{	"firstname": "Roberta",
				"lastname": "Berto",
				"username": "Robertina",
				"email": "roberta@orr.com",
				"cellphone": "55(61)98756-1234"

			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanMissPhone(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{
				"firstname": "robert",
				"lastname": "rob",
				"username": "Robson",
				"email": "robson@orr.com",
				"pass": "acb4789"

			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}


//-----------------------------------------------------Usuário já cadastrado --------------------------------------------------------------//

func TestInsertDeliverymanEmailExists(t *testing.T) {

	expected := `This user email already exists`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanUsernameExists(t *testing.T) {

	expected := `This username already exists`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popo.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//-------------------------------------------------------Informações invalidas-----------------------------------------------------------------------------//

func TestInsertDeliverymanInvalidFirstName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai045789",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanInvalidLastName(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Raissa",
				"lastname": "Galves12",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai045789",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanInvalidEmail(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalvesgmail",
				"pass": "rai045789",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanInvalidShortPassword(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai04",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanInvalidBigPassword(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai041547896321",
				"cellphone": "55(61)98756-1234"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestInsertDeliverymanInvalidPhone(t *testing.T) {

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Raissa12",
				"lastname": "Galves",
				"username": "Raissawn",
				"email": "rgalves@gmail.com",
				"pass": "rai0412",
				"cellphone": "98756-1234"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//---------------------------------------------------Informações faltando para atualizar--------------------------------------------------------------------------//
func TestUpdateDeliverymanMissFirstName(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanMissLastName(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Valesca",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanMissUserName(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"email": "valesca@popozuda.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanMissEmail(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanMissPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanMissPhone(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`
	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "valescapop"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//-----------------------------------------------------------Tentativa de alteração de atributos identificadores-----------------------------------------------------//

func TestUpdateDeliverymanEmail(t *testing.T) {
	client := &http.Client{}

	expected := `We couldn't find you. You can't change your email`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozudas.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanUsername(t *testing.T) {
	client := &http.Client{}

	expected := `You can't change your username`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Catra",
				"email": "valesca@popozuda.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}
//--------------------------------------------------------------------dando erro-------------------------------------------------------------------------------------------//
func TestUpdateDeliverymanDeactivated(t *testing.T) {
	client := &http.Client{}

	expected := `This user is already deactivated, login to activate again`

	deliveryman := `{
				"firstname": "Johnny",
				"lastname": "Smith",
				"username": "Nhoji",
				"email": "johnny@morto.com",
				"pass": "johnnymorto",
				"cellphone": "55(61)90876-5324"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}
//--------------------------------------------------------------------dando erro-------------------------------------------------------------------------------------------//
func TestUpdateDeliverymanNothingToUpdate(t *testing.T) {
	client := &http.Client{}

	expected := `Nothing to update`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "valescapop",
				"cellphone": "55(61)98164-8732"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanShortPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "vale",
				"cellphone": "55(61)98164-5421"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestUpdateDeliverymanBigPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "valescapopozuda",
				"cellphone": "55(61)98164-5421"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/deliverymen/updatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//----------------------------------------------------------------Informações faltando para desativar----------------------------------------------------------------------//

func TestDeactivateDeliverymanMissEmail(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`

	deliveryman := `{
				"pass": "valescapop"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/deliverymen/deactivatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateDeliverymanMissPassword(t *testing.T) {
	client := &http.Client{}

	expected := `{"code": 400, "message":"Missing fields"}`

	deliveryman := `{
				
				"valesca@popozuda.com"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/deliverymen/deactivatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateDeliverymanIncorrectEmail(t *testing.T) {
	client := &http.Client{}

	expected := `User email not found`

	deliveryman := `{	
				"email": "valesca@pop.com",
				"pass": "valescapop"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/deliverymen/deactivatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateDeliverymanIncorrectPassword(t *testing.T) {
	client := &http.Client{}

	expected := `Wrong password`

	deliveryman := `{	
				"email": "valesca@popozuda.com",
				"pass": "valescapopozuda"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/deliverymen/deactivatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateDeliverymanAlreadyDeactivated(t *testing.T) {
	client := &http.Client{}

	expected := `This user doesn't exist anymore`

	deliveryman := `{	
				"email": "johnny@morto.com",
				"pass": "johnnymorto"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/deliverymen/deactivatedeliveryman", r)
	req, _ := client.Do(resp)

	buf := new(bytes.Buffer)
   buf.ReadFrom(req.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

//----------------------------------------------------------Falhas de busca---------------------------------------------------------------------------//


func TestFindDeliverymanNotFound(t *testing.T) {

	expected := `Deliveryman not found`

	resp, _ := http.Get("http://"+config.SERVER_HOST+"/deliverymen/finddeliverymanbyname/Matilda")

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}


