package config
import ("gopkg.in/mgo.v2/bson"
		"strings")


func RegisterCart(c Cart) (err error) {

	collection := GetSession().DB("db_guittuce").C("carts")
	
	err = collection.Insert(c)

	return

}	

func FindActiveCarts() ([]Cart, error) {

	var carts []Cart

	collection := GetSession().DB("db_guittuce").C("carts")

	err := collection.Find(bson.M{}).All(&carts)

	err = collection.Find(bson.M{"open": true}).All(&carts)

	return carts, err
}

		
func CartByUser(username string) ([]Cart, error){

	var carts []Cart

	collection := GetSession().DB("db_guittuce").C("carts")

	err := collection.Find(bson.M{"username": username}).All(&carts)

	return carts, err
}

func CartUpdateByUser(username string) (Cart, error){

	var cart Cart

	collection := GetSession().DB("db_guittuce").C("carts")

	err := collection.Find(bson.M{"username": username, "open": true}).One(&cart)

	return cart, err
}

func CartByDeliveryman(deliveryname string) ([]Cart, error){

	var carts []Cart

	collection := GetSession().DB("db_guittuce").C("carts")

	err := collection.Find(bson.M{"deliveryname": deliveryname}).All(&carts)

	return carts, err
}

func CartByDeliverymanDelete(username string, deliveryname string) (Cart, error){

	var cart Cart

	collection := GetSession().DB("db_guittuce").C("carts")

	err := collection.Find(bson.M{"username": username,"deliveryname": deliveryname, "open": true}).One(&cart)

	return cart, err
}
func VerifyOpenCarts(username string) ([]Cart, []User, error){

	var carts []Cart
	var users []User

	collection := GetSession().DB("db_guittuce").C("carts")
	c := GetSession().DB("db_guittuce").C("users")

	err := c.Find(bson.M{"username": username}).All(&users)

	// = collection.Find(bson.M{"username": username}).All(&carts)

	err = collection.Find(bson.M{"username": username,"open": true}).All(&carts)

	return carts, users, err
}




func UpdateCart(c Cart, cart Cart) error{

	collection := GetSession().DB("db_guittuce").C("carts")

	err := collection.Update(c, &cart)

	return err
}


func ParseCredCard(cj CartJson) bool{
	 if cj.CredCard == "true"{
	 	return true
	 }
	 return false
}

func ParseCash(cj CartJson) bool{
	return cj.Cash == "true"
}

func ParseCredCardUpdate(cj CartJsonUpdate) bool{
	 if cj.CredCard == "true"{
	 	return true
	 }
	 return false
}

func ParseCashUpdate(cj CartJsonUpdate) bool{
	return cj.Cash == "true"
}

func FindByProduct(product string) ([]Cart,error){
	var carts []Cart
	var c []Cart

	collection := GetSession().DB("db_guittuce").C("carts")
	err:= collection.Find(bson.M{"open": true}).All(&carts)

	for i := 0; i < len(carts); i++ {
		if strings.Contains(carts[i].Products, product) {
		c = append (c, carts[i])}
	}

	return c, err
}
