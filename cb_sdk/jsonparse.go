package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}
type User struct {
	ID             string `json:"ID"`
	Manager        string `json:"manager"`
	Employee       string `json:"employee"`
	DOH            string `json:"doh"`
	Eligibility    string `json:"eligibility"`
	FiveYears      string `json:"5_years"`
	TenYears       string `json:"10_years"`
	Carry          string `json:"carry"`
	Accrued        string `json:"accrued"`
	PTOUsed        string `json:"pto_used"`
	CurrentBalance string `json:"current_balance"`
	Mutiplier      string `json:"multiplier"`
}

var path = "output4.json"
var fi = "output.json"

func fixDate(orig string) string {
	v := strings.Replace(orig, "/", "-", 3)
	return v
}
func main() {
	// Open our jsonFile
	jsonFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array
	var users Users
	//
	json.Unmarshal([]byte(byteValue), &users)
	fmt.Println("User Length", len(users.Users))
	var id string
	for i := 0; i < len(users.Users); i++ {
		id = fmt.Sprintf("%d", i+1)
		fmt.Println("ID", id)
		fmt.Println("Manager: " + users.Users[i].Manager)
		fmt.Println("Employee: " + users.Users[i].Employee)
		users.Users[i].DOH = fixDate(users.Users[i].DOH)
		fmt.Println("Date of Hire: ", users.Users[i].DOH)
		fmt.Println("Eligibility: " + users.Users[i].Eligibility)
		fmt.Println("Five year " + users.Users[i].FiveYears)
		fmt.Println("Ten Year: " + users.Users[i].TenYears)
		fmt.Println("Carry", users.Users[i].Carry)
		fmt.Println("Accrued", users.Users[i].Accrued)
		fmt.Println("PTOUsed", users.Users[i].PTOUsed)
		fmt.Println("CurrentBalance", users.Users[i].CurrentBalance)
		fmt.Println("Multiplier", users.Users[i].Mutiplier)

	}

}
