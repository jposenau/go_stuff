package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"gopkg.in/couchbase/gocb.v1"
)

type Users struct {
	Users []User `json:"staff"`
}
type User struct {
	//	ID             string  `json:"ID"`
	Manager        string    `json:"manager"`
	Employee       string    `json:"employee"`
	DOH            time.Time `json:"doh"`
	Eligibility    string    `json:"eligible"`
	FiveYears      string    `json:"fiveyears"`
	TenYears       string    `json:"tenyears"`
	Carry          float64   `json:"carry"`
	Accrued        float64   `json:"accrued"`
	PTOUsed        float64   `json:"ptoused"`
	CurrentBalance float64   `json:"currentbalance"`
	Mutiplier      float64   `json:"multiplier"`
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
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Clorox503",
	})
	bucket, _ := cluster.OpenBucket("mytest", "")

	bucket.Manager("", "").CreatePrimaryIndex("", true, false)
	dt := time.Now()
	bucket.Upsert("u:SageTeam1",
		User{
			//cb_sd	ID:             " 100",
			Manager:        "John",
			Employee:       "Smith Suzie",
			DOH:            dt,
			Eligibility:    "11/25/2005",
			FiveYears:      "9/26/2010",
			TenYears:       "9/26/2015",
			Carry:          11.5,
			Accrued:        1.8,
			PTOUsed:        2.5,
			CurrentBalance: 4.25,
			Mutiplier:      1.75}, 0)

	// Use query
	query := gocb.NewN1qlQuery("SELECT * FROM mytest WHERE $1 IN DOH")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"doh"})
	var row interface{}
	for rows.Next(&row) {
		fmt.Printf("Row: %v", row)
	}
	//readFile()
	//getfile()
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above

}
