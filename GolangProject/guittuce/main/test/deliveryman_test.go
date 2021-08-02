package test

import("testing"
	   "net/http"
	   "main/config"
	   "bytes")


func TestInsertDeliverymanSuccess(t *testing.T) {

	expected := `{"code": 201, "message":"Created success"}`

	deliveryman := `{
				"firstname": "Valesca",
				"lastname": "Catra",
				"username": "Popozuda",
				"email": "valesca@popozuda.com",
				"pass": "valescapopa",
				"cellphone": "55(61)98164-5421"
			}`
	
	r := bytes.NewReader([]byte(deliveryman))		
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)


	deliveryman = `{
				"firstname": "Johnny",
				"lastname": "Smith",
				"username": "Nhoji",
				"email": "johnny@morto.com",
				"pass": "johnnymorto",
				"cellphone": "55(61)90876-5324"
			}`
	
	r = bytes.NewReader([]byte(deliveryman))		
			
	resp, _ = http.Post("http://"+config.SERVER_HOST+"/deliverymen/insertdeliverymen", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

func TestDeactivateDeliverymanSuccess(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":200, "message": "Deleted"}`

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

func TestUpdateDeliverymanSuccess(t *testing.T) {
	client := &http.Client{}

	expected := `{"code":200, "message": "Updated"}`

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



func TestFindActiveDeliverymenSuccess(t *testing.T) {

	expected := `[]Deliveryman`
	
	r := `[]Deliveryman`

	resp, _ := http.Get("http://"+config.SERVER_HOST+"/deliverymen/finddeliverymen")
	
	_ = config.DecodeJson(resp.Body, &r)

	if r != expected {
		t.Errorf("The http response was: %v Expected: %v", r, expected)
	}

}

func TestFindDeliverymanbyNameSuccess(t *testing.T) {

	expected := `Deliveryman`
	
	r := `Deliveryman`

	resp, _ := http.Get("http://"+config.SERVER_HOST+"/deliverymen/finddeliverymanbyname/Valesca")
	
	_ = config.DecodeJson(resp.Body, &r)
		
	if r != expected {
		t.Errorf("The http response was: %v Expected: %v", r, expected)
	}

}





	
	



