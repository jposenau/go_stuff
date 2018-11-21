package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gopkg.in/couchbase/gocb.v1"
)

type Users struct {
	Users []User `json:"staff"`
}
type User struct {
	ID             string  `json:"ID"`
	Manager        string  `json:"manager"`
	Employee       string  `json:"employee"`
	DOH            string  `json:"doh"`
	Eligibility    string  `json:"eligible"`
	FiveYears      string  `json:"fiveyears"`
	TenYears       string  `json:"tenyears"`
	Carry          float64 `json:"carry"`
	Accrued        float64 `json:"accrued"`
	PTOUsed        float64 `json:"ptoused"`
	CurrentBalance float64 `json:"currentbalance"`
	Mutiplier      float64 `json:"multiplier"`
}

var path = "output.json"
var fi = "output.json"

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func getfile() {
	f, err := os.Open(fi)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Clorox503",
	})
	bucket, _ := cluster.OpenBucket("mytest", "")
	var users Users
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		fmt.Println(s)

		s, e = Readln(r)
		//		fmt.Println("Reading from file.")
		var users Users

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'users' which we defined above
		json.Unmarshal([]byte(s), &users)
		// we iterate through every user within our users array and
		// print out the user Type, their name, and their facebook url
		// as just an example
		//fmt.Println(users.Users)
		var id string
		for i := 0; i < len(users.Users); i++ {
			id = fmt.Sprintf("%d", i+1)
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
			fmt.Println("Ready to insert Manager", users.Users[i].Manager)

			// Use query
			query := gocb.NewN1qlQuery("SELECT * FROM mytest")
			rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"employee"})
			var row interface{}
			for rows.Next(&row) {
				fmt.Printf("Row: %v", row)
			}
		}

	}
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

	//bucket.Manager("", "").CreatePrimaryIndex("", true, false)
	getfile()

	//readFile()
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

}
