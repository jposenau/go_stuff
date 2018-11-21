package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}
type User struct {
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
	fmt.Println(len(byteValue))
	// we initialize our Users array
	var users Users
	//fmt.Println(byteValue)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	//var result map[string]interface{}
	//json.Unmarshal([]byte(byteValue), &result)

	//fmt.Println(result["users"])
	json.Unmarshal(byteValue, &users)
	for i := 0; i < len(users.Users); i++ {
		//id = fmt.Sprintf("%d", i+1)
		fmt.Println("Manager: " + users.Users[i].Manager)
		fmt.Println("Employee: " + users.Users[i].Employee)
		fmt.Println("Date of Hire: " + users.Users[i].DOH)
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
