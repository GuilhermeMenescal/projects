package config

import ("encoding/json"
		"log"
		"io")

var responses = map[string]string {
//--------------------------------success--------------------------------------//
	"find-success": `{"code":200, "message": "Found"}`,
	"register-success": `{"code": 201, "message":"Created success"}`,
	"delete-success": `{"code":200, "message": "Deleted"}`,
	"update-success": `{"code":200, "message": "Updated"}`,
//--------------------------------error--------------------------------------//
	"invalid-json" : `{"code": 400, "message":"Missing fields"}`,
	"register-error" : `{"code": 500, "message":"Couldn't create, try again later."}`,
	"find-failed": `{"code":404, "message": "Content not found"}`,
	"missing-fields": `{"code":403, "message": "Required fields missing or there are invalid chars in your fields"}`,
//--------------------------------generic----------------------------------//
	"forbidden": `{"code":403, "message": "I'm sorry Dave, I'm afraid I can't let you do that"}`,
}		

func DecodeJson(body io.Reader, entity interface{}) (err error) {
	d   := json.NewDecoder(body)	
	err = d.Decode(entity)
	
	if err != nil {
		log.Print("[ERROR] wrong JSON")
		return 	
	}

	return 
}
