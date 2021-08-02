package config

import("gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"regexp"
)

var validate *validator.Validate

func StartValidator() { 
	validate = validator.New()
	validate.RegisterValidation("firstname", VerifyFirstNameFormat)
	validate.RegisterValidation("lastname", VerifyLastNameFormat)
	validate.RegisterValidation("password", VerifyPasswordLength)
	validate.RegisterValidation("cellphone", VerifyPhoneNumberFormat)
}

//func ValidateAge(age validator.FieldLevel) bool {
//	return age.Field().Int() >= 18
//}


func VerifyUser(u string) (User, error){

	var user User

	collection := GetSession().DB("db_guittuce").C("users")

	err := collection.Find(bson.M{"email": u}).One(&user)

	return user, err
}

func VerifyUsername(u string) (User, error){

	var user User

	collection := GetSession().DB("db_guittuce").C("users")

	err := collection.Find(bson.M{"username": u}).One(&user)

	return user, err
}	

func VerifyDeliveryman(d string) (Deliveryman, error){

	var deliveryman Deliveryman

	collection := GetSession().DB("db_guittuce").C("deliverymen")

	err := collection.Find(bson.M{"email": d}).One(&deliveryman)

	return deliveryman, err
}

func VerifyDeliveryName(d string) (Deliveryman, error){

	var deliveryman Deliveryman

	collection := GetSession().DB("db_guittuce").C("deliverymen")

	err := collection.Find(bson.M{"username": d}).One(&deliveryman)

	return deliveryman, err
}



//-----------------------tags---------------------------//

func VerifyPasswordLength(password validator.FieldLevel) bool {

	return len(password.Field().String()) >= 6 && len(password.Field().String()) <= 14

}

func VerifyFirstNameFormat(firstname validator.FieldLevel) bool {

	var (
		pattern = `^[a-zA-Z]*$`
	)

	return regexp.MustCompile(pattern).MatchString(firstname.Field().String())

}

func VerifyLastNameFormat(lastname validator.FieldLevel) bool {

	var (
		pattern = `^[a-zA-Z]*$`
	)

	return regexp.MustCompile(pattern).MatchString(lastname.Field().String())

}

func VerifyPhoneNumberFormat(cellphone validator.FieldLevel) bool {
	
	const (
		pattern = `^\d{2}\(?\d{2}\)?\d{5}-?\d{4}$`
	)

	return regexp.MustCompile(pattern).MatchString(cellphone.Field().String())

}

