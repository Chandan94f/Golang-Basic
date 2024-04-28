package main

/*

interfaces


let's assume that we have to suffle the a slices

now data inside the slices can be of different type like

int, float, string or deck - for all these , we need a separate function
but in all the function logics will be the same

To avoid writing same type of logics in different function - interfaces comes in picture

*/

/*

Concrete Type

map, struct, int, string, engEnglish(custimised)


Interface Type

bot




*/
/*

type bot interface {
	getGreeting(string, int) (string, error)
	getBotVersion() float64
	respondToUser(user) string

}
*/

// example

/*

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}

	sb := spanishBot{}

	printGreeting(eb)

	printGreeting(sb)
}

// type englishBot struct{}

// type spanishBot struct{}

func (eb englishBot) getGreeting() string {

	// custom logic
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Halo"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }
*/
