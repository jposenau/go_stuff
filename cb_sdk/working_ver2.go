package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	gocb "gopkg.in/couchbase/gocb.v1"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}
type User struct {
	Manager        string    `json:"manager"`
	Employee       string    `json:"employee"`
	DOH            time.Time `json:"doh"`
	Eligibility    time.Time `json:"eligibility"`
	FiveYears      time.Time `json:"5_years"`
	TenYears       time.Time `json:"10_years"`
	Carry          string    `json:"carry"`
	Accrued        string    `json:"accrued"`
	PTOUsed        string    `json:"pto_used"`
	CurrentBalance string    `json:"current_balance"`
	Mutiplier      string    `json:"multiplier"`
}

var path = "output4.json"

func fixDate(orig string) string {
	//v := strings.Replace(orig, "/", "-", 3)
	stringSlice := strings.Split(orig, "/")
	v := stringSlice[2] + "-" + stringSlice[0] + "-" + stringSlice[1]
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
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Clorox503",
	})
	bucket, _ := cluster.OpenBucket("mytest1", "")
	bucket.Manager("", "").CreatePrimaryIndex("", true, false)
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Println(len(byteValue))
	// we initialize our Users array
	var users Users
	// we unmarshal our byteArray which contains our

	json.Unmarshal(byteValue, &users)
	var id string
	for i := 0; i < len(users.Users); i++ {
		id = fmt.Sprintf("%d", i+1)
		fmt.Println("Manager: " + users.Users[i].Manager)
		fmt.Println("Employee: " + users.Users[i].Employee)
		fmt.Println("Date of Hire: ", users.Users[i].DOH)
		fmt.Println("Eligibility: ", users.Users[i].Eligibility)
		fmt.Println("Five year ", users.Users[i].FiveYears)
		fmt.Println("Ten Year: ", users.Users[i].TenYears)
		fmt.Println("Carry", users.Users[i].Carry)
		fmt.Println("Accrued", users.Users[i].Accrued)
		fmt.Println("PTOUsed", users.Users[i].PTOUsed)
		fmt.Println("CurrentBalance", users.Users[i].CurrentBalance)
		fmt.Println("Multiplier", users.Users[i].Mutiplier)
		bucket.Upsert("u:"+id,
			User{
				Manager:        users.Users[i].Manager,
				Employee:       users.Users[i].Employee,
				DOH:            users.Users[i].DOH,
				Eligibility:    users.Users[i].Eligibility,
				FiveYears:      users.Users[i].FiveYears,
				TenYears:       users.Users[i].TenYears,
				Carry:          users.Users[i].Carry,
				Accrued:        users.Users[i].Accrued,
				PTOUsed:        users.Users[i].PTOUsed,
				CurrentBalance: users.Users[i].CurrentBalance,
				Mutiplier:      users.Users[i].Mutiplier}, 0)

	}

}
