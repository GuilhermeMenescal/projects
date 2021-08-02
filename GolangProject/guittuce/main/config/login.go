package config
import (
	"log"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

func UserLogin(w http.ResponseWriter, r *http.Request){
	
	var user UserLD 
	var u User

	err := DecodeJson(r.Body, &user)

	c := OpenSession("users")

	err = c.Find(bson.M{"email": user.Email}).One(&u)
	
	if err != nil {
		log.Print("[ERROR] failed to login: ", err)
		w.Write([]byte("This user doesn't exists. Register first!"))
		return 
	}

	if user.Password != u.Password {
		log.Print("Wrong password!")
		w.Write([]byte("Wrong infos!"))
		return 
	}
	if u.Active == false{
		ua:= u
		ua.Active = true
		UpdateUser(ua, u)
	}
	
	//u.Password = ""
	
	log.Print("Logged")
	w.Write([]byte("Logged"))
	return 
}

func DeliverymanLogin(w http.ResponseWriter, r *http.Request){
	
	var delivery DeliverymanLD 
	var d Deliveryman

	err := DecodeJson(r.Body, &delivery)

	c := OpenSession("deliverymen")

	err = c.Find(bson.M{"email": delivery.Email}).One(&d)
	
	if err != nil {
		log.Print("[ERROR] failed to login: ", err)
		w.Write([]byte("This user doesn't exists. Register first!"))
		return 
	}

	if delivery.Password != d.Password {
		log.Print("Wrong password!")
		w.Write([]byte("Wrong infos!"))
		return 
	}
	if d.Active == false{
		da:= d
		da.Active = true
		UpdateDeliveryman(da, d)
	}

	//u.Password = ""
	log.Print("Logged")
	w.Write([]byte("Logged"))
	return 
}

