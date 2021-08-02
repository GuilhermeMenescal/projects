package test

import("testing"
	   "net/http"
	   "main/config"
	   "bytes"
	   "log"
	   )


func TestSessionSuccess(t *testing.T) {

	client := &http.Client{}
	//var u config.User
	var d config.Deliveryman
    

	expected := `Cart deleted successfully`

	user := `{
				"email": "caio@cebola.com",
				"pass": "cebolasroxa"	
			}`
	
	r := bytes.NewReader([]byte(user))		
			
	_, _ = http.Post("http://"+config.SERVER_HOST+"/users/userlogin", "application/json", r)
	//log.Print(u.UserName)
    //buf := new(bytes.Buffer)
    //buf.ReadFrom(resp.Body)
    //response := buf.String()
    //if response == "Logged" {
    //	u,_ = config.VerifyUser("caio@cebola.com")
    //}

     
	cart :=`{ "products": "Pilão",
		  "description": "De flango", 
		  "limitprice": 10.00,
		  "credcard": "true",
		  "cash": "false",
		  "username": "Cebola",
		  "lat": "-15.8048206", 
		  "lon": "-47.939495"
		  
}`
		 
	 
	r = bytes.NewReader([]byte(cart))		
	
			
	_, _ = http.Post("http://"+config.SERVER_HOST+"/carts/insertcarts", "application/json", r)

	deliveryman := `{
				"email": "valesca@popozuda.com",
				"pass": "valescapop"	
			}`
	
	r = bytes.NewReader([]byte(deliveryman))		
			
	_, _ = http.Post("http://"+config.SERVER_HOST+"/deliverymen/deliverymanlogin", "application/json", r)
	


    	cart =`{ "products": "Pilão",
		  "description": "De flango", 
		  "limitprice": 10.00,
		  "credcard": "true",
		  "cash": "false",
		  "username": "Cebola",
		  "lat": "-15.8048206", 
		  "lon": "-47.939495",
		  "deliveryname": "Popozuda",
		  "fullprice": 3.00

		  
}`
	r = bytes.NewReader([]byte(cart))		
			
	robson, _ := http.NewRequest("PUT", "http://"+config.SERVER_HOST+"/carts/updatepurchase", r)
	req, _ := client.Do(robson)

	cart =`{ "products": "Pilão",
		  "description": "De flango", 
		  "limitprice": 10.00,
		  "credcard": "true",
		  "cash": "false",
		  "username": "Cebola",
		  "lat": "-15.8048206", 
		  "lon": "-47.939495",
		  "deliveryname": "Popozuda",
		  "fullprice": 3.00

		  
}`
	r = bytes.NewReader([]byte(cart))		
			
	robson, _ = http.NewRequest("DELETE", "http://"+config.SERVER_HOST+"/carts/deliverydeletecart", r)
	req, _ = client.Do(robson)
   
	buf := new(bytes.Buffer)
   	buf.ReadFrom(req.Body)
    response := buf.String()

    //log.Print(u.UserName)
    log.Print(d.UserName)

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

