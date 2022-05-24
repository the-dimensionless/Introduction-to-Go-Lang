package sayings

import (
	"fmt"

	"rsc.io/quote"
)

func Pithy() {
	fmt.Println("Phrase: ", quote.Glass())
	fmt.Println("Proverb: ", quote.Go())
	fmt.Println("Greeting: ", quote.Hello())
	fmt.Println("Optimzation Truth: ", quote.Opt())
}
