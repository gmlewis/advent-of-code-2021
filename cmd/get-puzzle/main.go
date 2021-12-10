// get-puzzle gets the puzzle for the given day (e.g. 'day08').
//
// The environment variable `AOC_COOKIES` must be set before this
// command is called.
//
// Usage from the root directory (where `go.mod` lives):
//   $ go run cmd/get-puzzle/main.go day08
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

const (
	aocCookies = "AOC_COOKIES"
	urlFmt     = "https://adventofcode.com/%v/day/%v"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("must supply a day to process, e.g. day08")
	}

	Each(flag.Args(), process)

	log.Printf("Done.")
}

func process(dayStr string) {
	if !strings.HasPrefix(dayStr, "day") {
		log.Fatalf("expected e.g. 'day08' but got %q", dayStr)
	}
	year := time.Now().Year()
	day := must.Atoi(dayStr[3:])

	log.Printf("Processing %v ...", dayStr)

	cookiesStr := os.Getenv(aocCookies)
	if cookiesStr == "" {
		log.Fatalf("must set env var %q", aocCookies)
	}
	cookies := parseCookies(cookiesStr)

	url := fmt.Sprintf(urlFmt, year, day)
	getFile(url, dayStr, "example1.txt", cookies, findPreCode)
	getFile(url+"/input", dayStr, "input.txt", cookies, nil)

	buf, err := exec.Command("git", "add", dayStr).CombinedOutput()
	if err != nil {
		log.Printf("%s", buf)
		log.Fatal(err)
	}
	log.Printf("%s", buf)

	msg := fmt.Sprintf("Initial commit for %v", dayStr)
	buf, err = exec.Command("git", "commit", "-am", msg).CombinedOutput()
	if err != nil {
		log.Printf("%s", buf)
		log.Fatal(err)
	}
	log.Printf("%s", buf)
}

func getFile(url, dayStr, outFile string, cookies []*http.Cookie, findPreCode func(string) string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	Each(cookies, req.AddCookie)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if findPreCode != nil {
		b = []byte(findPreCode(string(b)))
		if len(b) == 0 {
			log.Fatal("no puzzle found")
		}

		src := fmt.Sprintf("%v`%s`\n", mainTestSource, b)
		fn1 := filepath.Join(dayStr, "part1", "main_test.go")
		fn2 := filepath.Join(dayStr, "part2", "main_test.go")
		if _, err := os.Stat(fn1); err == nil {
			log.Fatalf("%v already exists; aborting", fn1)
		}
		if _, err := os.Stat(fn2); err == nil {
			log.Fatalf("%v already exists; aborting", fn2)
		}

		if err := ioutil.WriteFile(fn1, []byte(src), 0644); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(fn2, []byte(src), 0644); err != nil {
			log.Fatal(err)
		}
	} else {
		fn1 := filepath.Join(dayStr, "part1", "main.go")
		fn2 := filepath.Join(dayStr, "part2", "main.go")
		if _, err := os.Stat(fn1); err == nil {
			log.Fatalf("%v already exists; aborting", fn1)
		}
		if _, err := os.Stat(fn2); err == nil {
			log.Fatalf("%v already exists; aborting", fn2)
		}

		if err := ioutil.WriteFile(fn1, []byte(mainSource), 0644); err != nil {
			log.Fatal(err)
		}
		if err := ioutil.WriteFile(fn2, []byte(mainSource), 0644); err != nil {
			log.Fatal(err)
		}
	}

	filename := filepath.Join(dayStr, outFile)
	if err := ioutil.WriteFile(filename, b, 0644); err != nil {
		log.Fatal(err)
	}
}

var preCodeRE = regexp.MustCompile(`^(?ms).*?<pre><code>`)

func findPreCode(s string) string {
	s = strings.ReplaceAll(s, "<em>", "")
	s = strings.ReplaceAll(s, "</em>", "")
	parts := strings.Split(s, "</code></pre>")
	parts = Map(parts[:len(parts)-1], func(p string) string { return preCodeRE.ReplaceAllString(p, "") })
	return strings.Join(parts, "\n\n")
}

func parseCookies(cookies string) (ret []*http.Cookie) {
	return Reduce(strings.Split(cookies, "; "), ret, func(c string, acc []*http.Cookie) []*http.Cookie {
		p := strings.Split(c, "=")
		return append(acc, &http.Cookie{Name: p[0], Value: p[1]})
	})
}

var mainSource = `// -*- compile-command: "go run main.go ../example1.txt ../input.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"

	. "github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/advent-of-code-2021/must"
)

var logf = log.Printf
var printf = fmt.Printf

func main() {
	flag.Parse()

	Each(flag.Args(), process)
}

func process(filename string) {
	logf("Processing %v ...", filename)
	buf := must.ReadFile(filename)

	printf("Solution: %v\n", len(buf))
}
`

var mainTestSource = `package main

import (
	"testing"

	"github.com/gmlewis/advent-of-code-2021/test"
)

func TestExample(t *testing.T) {
	want := "Solution: 0\n"
	test.Runner(t, example1, want, process, &printf)
}

func BenchmarkExample(b *testing.B) {
	test.Benchmark(b, "../example1.txt", process, &logf, &printf)
}

func BenchmarkInput(b *testing.B) {
	test.Benchmark(b, "../input.txt", process, &logf, &printf)
}

var example1 = `
