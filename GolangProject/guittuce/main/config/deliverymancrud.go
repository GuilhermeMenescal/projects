package config
import ("gopkg.in/mgo.v2/bson")


func RegisterDeliveryman(d Deliveryman) (err error) {

	collection := GetSession().DB("db_guittuce").C("deliverymen")
	
	err = collection.Insert(d)

	return

}	

func FindDeliverymen() ([]Deliveryman, error) {

	var deliverymen []Deliveryman

	collection := GetSession().DB("db_guittuce").C("deliverymen")

	err := collection.Find(bson.M{"active": true}).All(&deliverymen)

	return deliverymen, err
}

func FindByNameD(name string) ([]Deliveryman, error){

	var deliverymen []Deliveryman

	collection := GetSession().DB("db_guittuce").C("deliverymen")

	err := collection.Find(bson.M{"firstname": name, "active": true}).All(&deliverymen)

	return deliverymen, err
}

		

func DeleteDeliveryman(d DeliverymanLD) (err error){

	collection := GetSession().DB("db_guittuce").C("deliverymen")

	err = collection.Remove(bson.M{"email": d.Email})

	return

}

func UpdateDeliveryman(d Deliveryman, deliveryman Deliveryman) (err error){

	collection := GetSession().DB("db_guittuce").C("deliverymen")

	err = collection.Update(deliveryman, &d)

	return
}

