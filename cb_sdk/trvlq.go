package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/couchbase/gocb.v1"
	"gopkg.in/couchbase/gocb.v1/cbft"
	//	"github.com/couchbase/gocb"
	//	"github.com/couchbase/gocb/cbft"
)

func simpleTextQuery(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery("swanky")).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Simple Text Query Error:", err.Error())
	}

	printResult("Simple Text Query", result)
}

func simpleTextQueryOnStoredField(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName,
		cbft.NewMatchQuery("MDG").Field("destinationairport")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Simple Text Query on Stored Field Error:", err.Error())
	}

	printResult("Simple Text Query on Stored Field", result)
}

func simpleTextQueryOnNonDefaultIndex(b *gocb.Bucket) {
	indexName := "travel-sample-index-hotel-description"
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery("swanky")).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Simple Text Query on Non-Default Index Error:", err.Error())
	}

	printResult("Simple Text Query on Non-Default Index", result)
}

func textQueryOnStoredFieldWithFacet(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewMatchQuery("La Rue Saint Denis!!").
		Field("reviews.content")).Limit(10).Highlight(gocb.DefaultHighlightStyle).
		AddFacet("Countries Referenced", cbft.NewTermFacet("country", 5))

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Match Query with Facet, Result by Row Error:", err.Error())
	}

	printResult("Match Query with Facet, Result by hits:", result)

	fmt.Println()
	fmt.Println("Match Query with Facet, Result by facet:")
	for _, row := range result.Facets() {
		jRow, err := json.Marshal(row)
		if err != nil {
			fmt.Println("Print Error:", err.Error())
		}
		fmt.Println(string(jRow))
	}
}

func docIdQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewDocIdQuery("hotel_26223", "hotel_28960"))

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("DocId Query Error:", err.Error())
	}

	printResult("DocId Query", result)
}

func unAnalyzedTermQuery(b *gocb.Bucket, fuzzinessLevel int) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewTermQuery("sushi").Field("reviews.content").
		Fuzziness(fuzzinessLevel)).Limit(50).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Printf("Unanalyzed Term Query with Fuzziness Level of %d Error: %s\n", fuzzinessLevel, err.Error())
	}

	printResult(fmt.Sprintf("Unanalyzed Term Query with Fuzziness Level of %d", fuzzinessLevel), result)
}

func matchPhraseQueryOnStoredField(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName,
		cbft.NewMatchPhraseQuery("Eiffel Tower").Field("description")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Match Phrase Query, using Analysis Error:", err.Error())
	}

	printResult("Match Phrase Query, using Analysis", result)
}

func unAnalyzedPhraseQuery(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName,
		cbft.NewPhraseQuery("dorm", "rooms").Field("description")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Phrase Query, without Analysis Error:", err.Error())
	}

	printResult("Phrase Query, without Analysis", result)
}

func conjunctionQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	firstQuery := cbft.NewMatchQuery("La Rue Saint Denis!!").Field("reviews.content")
	secondQuery := cbft.NewMatchQuery("boutique").Field("description")

	conjunctionQuery := cbft.NewConjunctionQuery(firstQuery, secondQuery)

	result, err := b.ExecuteSearchQuery(gocb.NewSearchQuery(indexName, conjunctionQuery).
		Limit(10).Highlight(gocb.DefaultHighlightStyle))
	if err != nil {
		fmt.Println()
		fmt.Println("Conjunction Query Error:", err.Error())
	}

	printResult("Conjunction Query", result)
}

func queryStringMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewQueryStringQuery("description: Imperial")).
		Limit(10)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Query String Query Error:", err.Error())
	}

	printResult("Query String Query", result)
}

func wildCardQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewWildcardQuery("bouti*ue").Field("description")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Wild Card Query Error:", err.Error())
	}

	printResult("Wild Card Query", result)
}

func numericRangeQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-unstored"
	query := gocb.NewSearchQuery(indexName, cbft.NewNumericRangeQuery().
		Min(10100, true).Max(10200, true).Field("id")).
		Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Wild Card Query Error:", err.Error())
	}

	printResult("Wild Card Query", result)
}

func regexpQueryMethod(b *gocb.Bucket) {
	indexName := "travel-sample-index-stored"
	query := gocb.NewSearchQuery(indexName, cbft.NewRegexpQuery("[a-z]").
		Field("description")).Limit(10).Highlight(gocb.DefaultHighlightStyle)

	result, err := b.ExecuteSearchQuery(query)
	if err != nil {
		fmt.Println()
		fmt.Println("Regexp Query Error:", err.Error())
	}

	printResult("Regexp Query", result)
}

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
	cluster, err := gocb.Connect("couchbase://localhost")
	if err != nil {
		panic("error conencting to cluster:" + err.Error())
	}
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Administrator",
		Password: "Clorox503",
	})
	bucket, err := cluster.OpenBucket("travel-sample", "")
	if err != nil {
		panic("error opening bucket:" + err.Error())
	}

	simpleTextQuery(bucket)
	simpleTextQueryOnStoredField(bucket)
	simpleTextQueryOnNonDefaultIndex(bucket)
	textQueryOnStoredFieldWithFacet(bucket)
	docIdQueryMethod(bucket)
	unAnalyzedTermQuery(bucket, 0)
	unAnalyzedTermQuery(bucket, 2)
	matchPhraseQueryOnStoredField(bucket)
	unAnalyzedPhraseQuery(bucket)
	conjunctionQueryMethod(bucket)
	queryStringMethod(bucket)
	wildCardQueryMethod(bucket)
	numericRangeQueryMethod(bucket)
	regexpQueryMethod(bucket)

	cluster.Close()
}
