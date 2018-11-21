package main

// an example program connecting to my vm using the bucket mine
import (
	"fmt"

	"gopkg.in/couchbase/gocb.v1"
)

type User struct {
	Id        string   `json:"uid"`
	Email     string   `json:"email"`
	Interests []string `json:"interests"`
}

func main() {
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Clorox503",
	})
	bucket, _ := cluster.OpenBucket("mytest", "")

	bucket.Manager("", "").CreatePrimaryIndex("", true, false)
	i := 0
	for i < 10 {
		i = i + 1
		ids := fmt.Sprintf("%s", i)
		bucket.Upsert("u:"+i,
			User{
				Id:        "hello",
				Email:     "fred@couchbase.com",
				Interests: []string{"Holy Hound"},
			}, 0)
	}
	// Get the value back
	var inUser User
	bucket.Get("u:kingarthur4", &inUser)
	fmt.Printf("User: %v\n", inUser)

	// Use query
	query := gocb.NewN1qlQuery("SELECT * FROM mytest WHERE $1 IN interests")
	rows, _ := bucket.ExecuteN1qlQuery(query, []interface{}{"Holy Hound"})
	var row interface{}
	for rows.Next(&row) {
		fmt.Println("Row: %v", row)
	}
}
