package main

import (
	MarkdownToHTML "Markdown_Processor/pkg/md2html"
	"fmt"
	"testing"
)

const GeneralTest string = `# User interface
## The main page
#### The functionality of the main page
1. Book Recommendation 
2. Collections of books

#### What will be located on the main page
1. the first user will be greeted by a book recommendation. He can click on it and go to the page with the book
2. Scrolling through below, the user will see a selection of books
3. In the upper left corner there will be a circle with an avatar, when clicking on which the user will be able to go to the profile
4. In the center of the top there will be a search bar where you can find a specific book

===
+ Book Recomm **endation**
+ Collections of *books skoob*

Lorem ipsum dolor sit amet consectetur adipisicing elit. Officiis dolor assumenda ut consectetur cum maiores explicabo delectus soluta veritatis maxime, repellat earum autem! Libero ut, ad odio cupiditate porro pariatur?

@Pepsi_King %
`

type TestedToken struct {
	TestTokenType string
	TestToken     string
}

func TMain(TT TestedToken) error {
	MarkdownToHTML.Convert(GeneralTest)
	return nil
}

func BenchmarkMain(b *testing.B) {
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		_, err := MarkdownToHTML.Convert(GeneralTest)
		if err != nil {
			errors += 1
		}
	}
	fmt.Printf("\tTotat op: %d, \t errors %d\n", TotalOperations, errors)
}

func BenchmarkHEADING(b *testing.B) {
	TT := TestedToken{TestTokenType: "HEADING", TestToken: "#"}
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		if err := TMain(TT); err != nil {
			errors += 1
		}
	}
	fmt.Printf("\tTotat op: %d, \t errors %d\n", TotalOperations, errors)
}

func BenchmarkWORD(b *testing.B) {
	TT := TestedToken{TestTokenType: "WORD", TestToken: "Incomprehensibilities"}
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		if err := TMain(TT); err != nil {
			errors += 1
		}
	}
	fmt.Printf("\tTotat op: %d, \t errors %d\n", TotalOperations, errors)
}

func BenchmarLINE(b *testing.B) {
	TT := TestedToken{TestTokenType: "LINE", TestToken: "==="}
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		if err := TMain(TT); err != nil {
			errors += 1
		}
	}
	fmt.Printf("\tTotat op: %d, \t errors %d\n", TotalOperations, errors)
}

func BenchmarkLIST(b *testing.B) {
	TT := TestedToken{TestTokenType: "LIST", TestToken: `- Book Recommendation
- Collections of books`}
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		if err := TMain(TT); err != nil {
			errors += 1
		}
	}
	fmt.Printf("\tTotat op: %d, \t errors %d\n", TotalOperations, errors)
}

func BenchmarkNUMBEREDLIST(b *testing.B) {
	TT := TestedToken{TestTokenType: "NUMBEREDLIST", TestToken: `1. Book Recommendation
2. Collections of books`}
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		if err := TMain(TT); err != nil {
			errors += 1
		}
	}
	fmt.Printf("\tTotat op: %d, \t errors %d\n", TotalOperations, errors)
}

func BenchmarkBOLT(b *testing.B) {
	TT := TestedToken{TestTokenType: "BOLT", TestToken: " **test**"}
	errors := 0
	TotalOperations := 0
	for i := 0; i < b.N; i++ {
		TotalOperations += 1
		if err := TMain(TT); err != nil {
			errors += 1
		}
	}
	fmt.Printf("\tTotat op: %d, \t errors %d\n", TotalOperations, errors)
}
