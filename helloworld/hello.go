package hello

import "fmt"

var greeting = map[string]string{
	"English": "Hello",
	"Spanish": "Hola",
	"French":  "Bonjour",
}

func Hello(name, lang string) string {
	greet, ok := greeting[lang]
	if !ok {
		greet = greeting["English"]
	}
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("%s, %s.", greet, name)
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
