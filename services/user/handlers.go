package user

import (
	"api/ent"
	"api/ent/user"
	"api/utils"
	"context"
)

func CheckUserAvailability(email string, client *ent.Client) bool {
	user, _ := client.User.Query().Where(user.Email(email)).Only(context.Background())
	return user != nil
}

func CreateUser(email string, password string, username string, client *ent.Client) (*ent.User, error) {
	tx, err := client.Tx(context.Background())
	if err != nil {
		return nil, err
	}
	newUser, err := tx.User.Create().
		SetEmail(email).
		SetPassword(password).SetUsername(username).Save(context.Background())

	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return newUser, err
}

func UpdateUsers(id uint, u User, client *ent.Client) (*ent.User, error) {

	dob, err := utils.ConvertJsonDate(*u.DOB)

	if err != nil {
		return nil, err
	}
	user, err := client.User.UpdateOneID(uint(id)).
		SetUsername(u.UserName).
		SetFirstName(*u.FirstName).
		SetLastName(*u.LastName).
		SetDob(dob).
		SetDescription(*u.Description).
		SetMobileNumber(int32(*u.MobileNumber)).
		SetAddressBlock(*u.AddresBlock).
		SetAddressStreet(*u.AddressStreet).
		SetRecidentialCountry(*u.ResidentialCountry).
		SetNationality(*u.Nationality).
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return user, nil

}

func FetchUserByID(client *ent.Client, id uint) (*ent.User, error) {
	return client.User.Get(context.Background(), id)
}
