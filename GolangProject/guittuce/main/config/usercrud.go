package config
import ("gopkg.in/mgo.v2/bson" )


func RegisterUser(u User) (err error) {

	collection := GetSession().DB("db_guittuce").C("users")
	
	err = collection.Insert(u)

	return

}	

func FindUsers() ([]User, error) {

	var users []User

	collection := GetSession().DB("db_guittuce").C("users")

	err := collection.Find(bson.M{"active": true}).All(&users)

	return users, err
}

func FindByName(name string) ([]User, error){

	var users []User

	collection := GetSession().DB("db_guittuce").C("users")

	err := collection.Find(bson.M{"firstname": name, "active": true}).All(&users)

	return users, err
}

	

func DeleteUser(u UserLD) (err error){

	collection := GetSession().DB("db_guittuce").C("users")

	err = collection.Remove(bson.M{"email": u.Email})

	return

}

func UpdateUser(u User, user User) (err error){

	collection := GetSession().DB("db_guittuce").C("users")

	err = collection.Update(user, &u)

	return
}

