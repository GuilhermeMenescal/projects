package config

import ("github.com/gorilla/mux")

func SetupUsersRoutes(r *mux.Router) {

	//----------------------User------------------------------------//
	r.HandleFunc("/users/insertusers", InsertUser).Methods("POST")
	r.HandleFunc("/users/deleteuser", RemoveUser).Methods("DELETE")
	r.HandleFunc("/users/deactivateuser", DeactivateUser).Methods("DELETE")
	r.HandleFunc("/users/findusers", FindAllUsers).Methods("GET")
	r.HandleFunc("/users/finduserbyname/{fullname}", FindUsersByName).Methods("GET")
	r.HandleFunc("/users/updateuser", UpdateOneUser).Methods("PUT")


	//----------------------Deliveryman-------------------------------//
	r.HandleFunc("/deliverymen/insertdeliverymen", InsertDeliveryman).Methods("POST")
	r.HandleFunc("/deliverymen/deletedeliveryman", RemoveDeliveryman).Methods("DELETE")
	r.HandleFunc("/deliverymen/deactivatedeliveryman", DeactivateDeliveryman).Methods("DELETE")
	r.HandleFunc("/deliverymen/finddeliverymen", FindAllDeliverymen).Methods("GET")
	r.HandleFunc("/deliverymen/finddeliverymanbyname/{fullname}", FindDeliverymenByName).Methods("GET")
	r.HandleFunc("/deliverymen/updatedeliveryman", UpdateOneDeliveryman).Methods("PUT")


	//--------------------Cart----------------------------------------//
	r.HandleFunc("/carts/insertcarts", InsertCart).Methods("POST")
	r.HandleFunc("/carts/findactivecarts", FindAllActiveCarts).Methods("GET")
	r.HandleFunc("/carts/{username}/finduserhistorycarts", FindUserHistoryCarts).Methods("GET")
	r.HandleFunc("/carts/{deliveryname}/finddeliveryhistorycarts", FindDeliveryHistoryCarts).Methods("GET")
	r.HandleFunc("/carts/updatepurchase", UpdatePurchase).Methods("PUT")
	r.HandleFunc("/carts/findbyproduct/{product}", FindCartsByProducts).Methods("GET")
	r.HandleFunc("/carts/userdeletecart", UserDeactivateCart).Methods("DELETE")
	r.HandleFunc("/carts/deliverydeletecart", DeliveryDeactivateCart).Methods("DELETE")


	//----------------Login------------------------------------------------//
	r.HandleFunc("/users/userlogin", UserLogin).Methods("POST")
	r.HandleFunc("/deliverymen/deliverymanlogin", DeliverymanLogin).Methods("POST")

	
}
