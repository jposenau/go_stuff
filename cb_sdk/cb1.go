package main

//Cb1 takes a JSON file and converts to couchbase output correcting date stings in the process
import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/couchbase/gocb.v1"
)

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

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func fixDate(orig string) string {
	//v := strings.Replace(orig, "/", "-", 3)
	stringSlice := strings.Split(orig, "/")
	v := stringSlice[2] + "-" + stringSlice[0] + "-" + stringSlice[1]
	return v
}
func getfile() {
	var bucket *gocb.Bucket
	cluster, err := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Clorox503",
	})
	// Open Bucket
	bucket, err = cluster.OpenBucket("mytest1", "")
	if err != nil {
		fmt.Println("ERROR OPENING BUCKET:", err)
	}
	var users Users

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

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal([]byte(byteValue), &users)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	fmt.Println("User Length", len(users.Users))
	//fmt.Println(users.Users)
	var id string
	for i := 0; i < len(users.Users); i++ {
		id = fmt.Sprintf("%d", i+1)
		fmt.Println("Manager: " + users.Users[i].Manager)
		fmt.Println("Employee: " + users.Users[i].Employee)
		users.Users[i].DOH = fixDate(users.Users[i].DOH)
		fmt.Println("Date of Hire: " + users.Users[i].DOH)
		users.Users[i].Eligibility = fixDate(users.Users[i].Eligibility)
		fmt.Println("Eligibility: " + users.Users[i].Eligibility)
		users.Users[i].FiveYears = fixDate(users.Users[i].FiveYears)
		fmt.Println("Five year " + users.Users[i].FiveYears)
		users.Users[i].TenYears = fixDate(users.Users[i].TenYears)
		fmt.Println("Ten Year: " + users.Users[i].TenYears)
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
	// Use query
	//query := gocb.NewN1qlQuery("SELECT * FROM mytest1")
	//rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"employee"})
	//var row interface{}
	//for rows.Next(&row) {
	//	fmt.Printf("Row: %v \n", row)
	//}
}

func readFile() {
	// Open file for reading.
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Read file, line by line
	var text = make([]byte, 1024)
	for {
		_, err = file.Read(text)

		// Break if finally arrived at end of file
		if err == io.EOF {
			break
		}

		// Break if error occured
		if err != nil && err != io.EOF {
			isError(err)
			break
		}
	}

}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	if isError(err) {
		return
	}

	fmt.Println("File Deleted")
}
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func main() {

	getfile()

}
