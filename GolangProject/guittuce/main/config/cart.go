package config

import ("net/http"
		"log"
		"encoding/json"
		"github.com/gorilla/mux")


type Cart struct {
	
	Products string      `json:"products" validate:"required"`
	Description string   `json:"description" validate:"required"`
	LimitPrice float32   `json:"limitprice" validate:"required"`
	CredCard bool        `json:"credcard" validate:"required"`
	Cash bool            `json:"cash" validade:"required"`	 
	Username string	     `json:"username" validade:"required"`		
	Lat string	     	 `json:"lat" validade:"required"`
	Lon string	     	 `json:"lon" validade:"required"`
	Deliveryname string	 `json:"deliveryname"`
	FullPrice float32    `json:"fullprice"`
	DeliveryFee float32	 `json: "fee"`
	Open bool  	     	 `json:"open"`
}

type CartJson struct {
	
	Products string      `json:"products" validate:"required"`
	Description string   `json:"description" validate:"required"`
	LimitPrice float32   `json:"limitprice" validate:"required"`
	CredCard string      `json:"credcard" validate:"required"`
	Cash string          `json:"cash" validade:"required"`	 
	Username string	     `json:"username" validade:"required"`		
	Lat string	     	 `json:"lat" validade:"required"`
	Lon string	     	 `json:"lon" validade:"required"`
	Deliveryname string	 `json:"deliveryname"`
	FullPrice float32    `json:"fullprice"`
	DeliveryFee float32	 `json: "fee"`
	Open bool  	     	 `json:"open"`
}

type CartJsonUpdate struct {
	
	Products string      `json:"products" validate:"required"`
	Description string   `json:"description" validate:"required"`
	LimitPrice float32   `json:"limitprice" validate:"required"`
	CredCard string      `json:"credcard" validate:"required"`
	Cash string          `json:"cash" validade:"required"`	 
	Username string	     `json:"username" validade:"required"`		
	Lat string	     	 `json:"lat" validade:"required"`
	Lon string	     	 `json:"lon" validade:"required"`
	Deliveryname string	 `json:"deliveryname" validade:"required"`
	FullPrice float32    `json:"fullprice" validade:"required"`
	DeliveryFee float32	 `json: "fee"`
	Open bool  	     	 `json:"open"`
}





func InsertCart(w http.ResponseWriter, r *http.Request) {
	var cj CartJson
	var c Cart
	
	err := DecodeJson(r.Body, &cj)
	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["invalid-json"]))
		return
	}
		c.DeliveryFee = 6.90
		c.Products = cj.Products
		c.Description = cj.Description
		c.LimitPrice = cj.LimitPrice + c.DeliveryFee
		c.CredCard = ParseCredCard(cj)
		c.Cash = ParseCash(cj)
		c.Username = cj.Username	
		c.Lat = cj.Lat
		c.Lon = cj.Lon	 
		c.Open = true

	err = validate.Struct(c)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err, ParseCredCard(cj))
		w.Write([]byte(responses["missing-fields"]))
		return	
	}

	carts, users,_:= VerifyOpenCarts(c.Username)

	if users == nil{
		log.Printf("User not registered")
		w.Write([]byte("User not registered"))
		return	
	}

	if carts != nil{
		log.Printf("Too many open carts")
		w.Write([]byte("Too many open carts"))
		return	
	}

	err = RegisterCart(c)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["register-error"]))
		return	
	}

	w.Write([]byte(responses["register-success"]))
	return

}

func FindAllActiveCarts(w http.ResponseWriter, r *http.Request){

	carts, err := FindActiveCarts()

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	cartsJson, err := json.Marshal(carts) 

	if err != nil{
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}


	w.Write([]byte(cartsJson))
	return 
}

func UpdatePurchase(w http.ResponseWriter, r *http.Request){
	var cartJson CartJsonUpdate
	var c Cart
	var cart Cart
	

	err := DecodeJson(r.Body, &cartJson)

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["invalid-json"]))
		return
	}

		c.Products = cartJson.Products
		c.Description = cartJson.Description
		c.LimitPrice = cartJson.LimitPrice
		c.CredCard = ParseCredCardUpdate(cartJson)
		c.Cash = ParseCashUpdate(cartJson)
		c.Username = cartJson.Username	
		c.Lat = cartJson.Lat
		c.Lon = cartJson.Lon
		c.Deliveryname = cartJson.Deliveryname
		c.FullPrice = cartJson.FullPrice 
		c.DeliveryFee = cartJson.DeliveryFee
		c.Open = true

		


		cart,err = CartUpdateByUser(c.Username)

		if err != nil {
			log.Printf("Cart not found")
			w.Write([]byte("Cart not found"))
		return
	}

		c.FullPrice += cart.DeliveryFee
		c.LimitPrice = cart.LimitPrice
		c.DeliveryFee = cart.DeliveryFee

		cart.Deliveryname = c.Deliveryname
		cart.FullPrice = c.FullPrice


		
		

		if c!=cart{
			log.Printf("You've attempted to change an user field")
			w.Write([]byte("You've attempted to change an user field"))
		return
		}

		cart.Deliveryname = ""
		cart.FullPrice = 0

		if c.FullPrice > cart.LimitPrice{
			log.Printf("The full price exceed the limit price")
			w.Write([]byte("The full price exceed the limit price"))
		return
		}

		err = UpdateCart(cart, c)

		if err != nil {
		log.Printf("Update purchase fail")
		w.Write([]byte("Update purchase fail"))
		return
	}

		log.Printf("Purchase updated")
		w.Write([]byte("This cart is now yours"))


}

func UserDeactivateCart (w http.ResponseWriter, r *http.Request) {

	var cartJson CartJson
	var c Cart
	var cart Cart

	err := DecodeJson(r.Body, &cartJson)

		c.Products = cartJson.Products
		c.Description = cartJson.Description
		c.LimitPrice = cartJson.LimitPrice
		c.CredCard = ParseCredCard(cartJson)
		c.Cash = ParseCash(cartJson)
		c.Username = cartJson.Username	
		c.Lat = cartJson.Lat
		c.Lon = cartJson.Lon
		c.Deliveryname = cartJson.Deliveryname
		c.FullPrice = cartJson.FullPrice
		c.DeliveryFee = 6.90
		c.Open = true

		c.LimitPrice += c.DeliveryFee

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	if c.Deliveryname != ""{
		log.Printf("Couldn't delete a cart with a deliveryman")
		w.Write([]byte("You can't delete a cart that already have a deliveryman"))
		return 
	}

	cart = c
	cart.Open = false

	err = UpdateCart(c, cart)

	if err != nil{
		log.Printf("Couldn't delete the cart")
		w.Write([]byte("You can't delete this cart now"))
		return 
	}

	w.Write([]byte("Cart deleted successfully"))
	return 
}

func DeliveryDeactivateCart (w http.ResponseWriter, r *http.Request) {

	var cartJson CartJsonUpdate
	var c Cart
	var cart Cart

	err := DecodeJson(r.Body, &cartJson)

		c.Products = cartJson.Products
		c.Description = cartJson.Description
		c.LimitPrice = cartJson.LimitPrice
		c.CredCard = ParseCredCardUpdate(cartJson)
		c.Cash = ParseCashUpdate(cartJson)
		c.Username = cartJson.Username	
		c.Lat = cartJson.Lat
		c.Lon = cartJson.Lon
		c.Deliveryname = cartJson.Deliveryname
		c.FullPrice = cartJson.FullPrice
		c.DeliveryFee = cartJson.DeliveryFee
		c.Open = true

	if err != nil {
		log.Printf("[ERROR] Wrong data: %v", err)
		w.Write([]byte(responses["invalid-json"]))
		return
	}

	cart,err = CartByDeliverymanDelete(c.Username, c.Deliveryname)

	if err != nil {
		log.Printf("Cart not found")
		w.Write([]byte("Cart not found"))
		return
	}

	c = cart
	cart.Open = false

	err = UpdateCart(c, cart)

	if err != nil{
		log.Printf("Couldn't delete the cart")
		w.Write([]byte("You can't delete this cart now"))
		return 
	}

	w.Write([]byte("Cart deleted successfully"))
	return 
}

func FindCartsByProducts(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
    product := vars["product"]

    c,_ := FindByProduct(product)

    cartsJson, err := json.Marshal(c)
    if err != nil{
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

    w.Write([]byte(cartsJson))

}


func FindUserHistoryCarts(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
    name := vars["username"]

	carts, err := CartByUser(name)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	cartsJson, err := json.Marshal(carts) 

	if err != nil{
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	w.Write([]byte(cartsJson))
	return 

}

func FindDeliveryHistoryCarts(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
    name := vars["deliveryname"]

	carts, err := CartByDeliveryman(name)

	if err != nil {
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}

	cartsJson, err := json.Marshal(carts) 

	if err != nil{
		log.Printf("[ERROR] Problem with database: %v", err)
		w.Write([]byte(responses["find-failed"]))
		return 
	}


	w.Write([]byte(cartsJson))
	return 
}



