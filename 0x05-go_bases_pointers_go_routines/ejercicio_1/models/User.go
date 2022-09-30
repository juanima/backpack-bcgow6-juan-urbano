package models 

import "fmt"

type User struct {
    Name     string `json:"ID"`
    LastName string `json:"price"`
    Age	     int    `json:"amount"`
    Email    string `json:"email"`
    Password string `json:"password`
}


func NewUser(name, lastName, email, password string, age int) (user *User, err error) {

    user = &User {
	Name: name, 
	LastName: lastName,
	Email: email, 
	Password: password, 
	Age: age,
    }

    return user, nil
}

func (u *User) PrintUser() (userStr string) {
    return fmt.Sprintf(
	"Info User: \n\tName: %s,\n\tLast name: %s, \n\tEmail: %s, \n\tAge: %d\n", 
	u.Name, u.LastName, u.Email, u.Age)    
}


func (u *User) SetName(newName string) {
    u.Name = newName
}


func (u *User) SetAge(newAge int) {
    u.Age = newAge
}


func (u *User) SetEmail(newEmail string) {
    u.Email = newEmail
}


func (u *User) SetPassword(newPassword string) {
    u.Password = newPassword
}

