package join

import "testing"

var testData = []string{
	"lorem",
	"ipsum",
	"dolor",
	"But",
	"I",
	"must",
	"explain",
	"to",
	"you",
	"how",
	"all",
	"this",
	"mistaken",
	"idea",
	"of",
	"denouncing",
	"pleasure",
	"and",
	"praising",
	"pain",
	"was",
	"born",
	"and",
	"I",
	"will",
	"give",
	"you",
	"a",
	"complete",
	"account",
	"of",
	"the",
	"system,",
	"and",
	"expound",
	"the",
	"actual",
	"teachings",
	"of",
	"the",
	"great",
	"explorer",
	"of",
	"the",
	"truth.",
}

func BenchmarkStrJoinIterated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrJoinIterated(testData)
	}
}

func BenchmarkStrJoinRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StrJoinRange(testData)
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsJoin(testData, " ")
	}
}

func BenchmarkStrJoinBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringsJoinBuilder(testData)
	}
}
