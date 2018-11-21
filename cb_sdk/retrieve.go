package main

// retrive is designed to take an employee name and provide a data feed from couchbase
// returning relevent infromation about the user. We use FTS to find the employee based ona partoal name search
// and then return the UID to complete a full search. The data is based on the PTO requirements for salaried team members

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/couchbase/gocb.v1"
	"gopkg.in/couchbase/gocb.v1/cbft"
)

// Users struct which contains an array of users
type Users struct {
	Users []User `json:"users"`
}

// struc definition for the user PTO detail
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

// bucket reference - reuse as bucket reference in the application
var bucket *gocb.Bucket

func getName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	return text
}

// Query by name
func queryStringMethod(b *gocb.Bucket, retValue User) string {
	indexName := "name_list"
	LocalName := retValue.Employee
	query := gocb.NewSearchQuery(indexName, cbft.NewQueryStringQuery(LocalName)).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Query String Query Error:", err.Error())
	}

	var rslt string
	for _, hit := range result.Hits() {
		fmt.Printf("%s\n", hit.Id)
		rslt = hit.Id
	}

	printResult("String Query", result)
	return rslt
}

//Simple text query design to return the Id of a specific employee
func simpleTextQuery(b *gocb.Bucket, retValue User) string {
	indexName := "name_list"
	LocalName := retValue.Employee
	fmt.Println(LocalName)
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery(LocalName)).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Simple Text Query Error:", err.Error())
	}
	var rslt string
	for _, hit := range result.Hits() {
		rslt = hit.Id
	}

	printResult("Simple Text Query", result)
	return rslt
}

//test prining for result checking

func printResult(label string, results gocb.SearchResults) {
	fmt.Println()
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println("= = = = = = = = = = = = = = = = = = = = = = =")
	fmt.Println()
	fmt.Println(label)
	fmt.Println()

	for _, row := range results.Hits() {
		jRow, err := json.Marshal(row)
		if err != nil {
			fmt.Println("Print Error:", err.Error())
		}
		fmt.Println(string(jRow))
	}
}

func main() {
	// Connect to Cluster
	cluster, err := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Clorox503",
	})
	// Open Bucket
	bucket, err = cluster.OpenBucket("mytest", "")
	if err != nil {
		fmt.Println("ERROR OPENING BUCKET:", err)
	}
	var retValue User
	retValue.Employee = getName()
	key := simpleTextQuery(bucket, retValue)
	fmt.Println("Employee ID", key)

	_, err = bucket.Get(key, &retValue)
	if err != nil {
		fmt.Println("ERROR RETURNING DOCUMENT:", err)
	}

	fmt.Println("Manager: " + retValue.Manager)
	fmt.Println("Employee: " + retValue.Employee)
	fmt.Println("Date of Hire: " + retValue.DOH)
	fmt.Println("Eligibility: " + retValue.Eligibility)
	fmt.Println("Five year " + retValue.FiveYears)
	fmt.Println("Ten Year: " + retValue.TenYears)
	fmt.Println("Carry", retValue.Carry)
	fmt.Println("Accrued", retValue.Accrued)
	fmt.Println("PTOUsed", retValue.PTOUsed)
	fmt.Println("CurrentBalance", retValue.CurrentBalance)
	fmt.Println("Multiplier", retValue.Mutiplier)

	fmt.Println("Example Successful - Exiting")
	key = queryStringMethod(bucket, retValue)
	_, err = bucket.Get(key, &retValue)
	if err != nil {
		fmt.Println("ERROR RETURNING DOCUMENT:", err)
	}

	fmt.Println("new key", key)
	fmt.Println("Manager: " + retValue.Manager)
	fmt.Println("Employee: " + retValue.Employee)
	fmt.Println("Date of Hire: " + retValue.DOH)
	fmt.Println("Eligibility: " + retValue.Eligibility)
	fmt.Println("Five year " + retValue.FiveYears)
	fmt.Println("Ten Year: " + retValue.TenYears)
	fmt.Println("Carry", retValue.Carry)
	fmt.Println("Accrued", retValue.Accrued)
	fmt.Println("PTOUsed", retValue.PTOUsed)
	fmt.Println("CurrentBalance", retValue.CurrentBalance)
	fmt.Println("Multiplier", retValue.Mutiplier)

}
