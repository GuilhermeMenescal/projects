package test

import("testing"
	   "net/http"
	   "main/config"
	   "bytes")


func TestInsertCartSuccess(t *testing.T) {

	expected := `{"code": 200, "message":"You are now able to get fat"}`
	
      

	cart := `{"products": "Arroz",
		  "description": "me da comida", 
		  "limitprice": 10.00,
		  "credcard": "true",
		  "cash": "true",
		  
}`
		 
	
	r := bytes.NewReader([]byte(cart))		
	
			
	resp, _ := http.Post("http://"+config.SERVER_HOST+"/insertcarts", "application/json", r)

	buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    response := buf.String()
	

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}

