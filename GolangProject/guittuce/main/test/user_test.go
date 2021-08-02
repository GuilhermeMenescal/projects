package test

import("testing"
	   "net/http"
	   "main/config"
	   "bytes")

func TestInsertUserSuccess(t *testing.T) {

	expected := `{"code": 201, "message":"Created success"}`

	user := `{
				"firstname": "Caio",
				"lastname": "Silva",
				"username": "Cebola",
				"email": "caio@cebola.com",
				"pass": "cebolaroxa",
				"cellphone": "55(61)95678-1234"
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)

	user = `{
				"firstname": "Batata",
				"lastname": "Capota",
				"username": "Batatinha",
				"email": "batata@morta.com",
				"pass": "batata",
				"cellphone": "55(61)96778-5678"
			}`

	r = bytes.NewReader([]byte(user))	

	resp, _ = http.Post("http://"+config.SERVER_HOST+"/users/insertusers", "application/json", r)


	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateUserSuccess(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":200, "message": "Deleted"}`

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



func TestUpdateUserSuccess(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":200, "message": "Updated"}`

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



func TestFindActiveUsersSuccess(t *testing.T) {

	expected := `[]User`
	
	r := `[]User`

	resp, _ := http.Get("http://"+config.SERVER_HOST+"/users/findusers")
	
	_ = config.DecodeJson(resp.Body, &r)
	
	if r != expected {
		t.Errorf("The http response was: %v Expected: %v", r, expected)
	}

}

func TestFindUserbyNameSuccess(t *testing.T) {

	expected := `User`
	
	r := `User`

	resp, _ := http.Get("http://"+config.SERVER_HOST+"/users/finduserbyname/Caio")
	
	_ = config.DecodeJson(resp.Body, &r)
	
	if r != expected {
		t.Errorf("The http response was: %v Expected: %v", r, expected)
	}

}


	
	



