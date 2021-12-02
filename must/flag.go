package must

import (
	"flag"
	"log"
)

// Process calls the provided func for every argument
// provided on the command line.
func Process(f func(string)) {
	for _, arg := range flag.Args() {
		log.Printf("Processing file %v ...", arg)
		f(arg)
	}
}
