package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"sync"
	"time"
)

func demoT() {
	client := &http.Client{
		Timeout: 15 * time.Minute,
	}
	for i := 0; i < 5; i++ {
		resp, err := client.Get("http://webcode.me")
		if err != nil {
			log.Fatal(err)
		}
		//defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(body))
		fmt.Println("Response received \n", string(body))

	}

}

func main() {

	//demoT()
	//demoWrite()
	// demoPost()
	// MakeRequest()

	// resp, err := http.Get("http://webcode.me")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(body))

	// demo()
	demoRegex()

}

func MakeRequest() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}

func demoWrite() {
	// r, err := http.Get("http://webcode.me")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer r.Body.Close()

	// f, err := os.Create("index.html")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// _, err = f.ReadFrom(r.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	r, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)
	}

	defer r.Body.Close()

	f, err := os.Create("index-f-go.html")

	f, err = os.Create("index-f-go1.html")

	h, err := os.Create("index-f-go2.html")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	defer h.Close()

	_, err = f.ReadFrom(r.Body)

	_, err = h.ReadFrom(r.Body)

	if err != nil {
		log.Fatal(err)
	}

}

func demoPost() {
	data := url.Values{
		"name":       {"John Doe"},
		"occupation": {"gardener"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", data)

	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])
}

func demo() {
	/* We make multiple asynchronous HTTP requests.
	   We get the contents of the title tag of each of the web pages. */
	/* WaitGroups are used to manage goroutines.
	   It waits for a collection of goroutines to finish */
	urls := []string{
		"http://webcode.me",
		"https://example.com",
		"http://httpbin.org",
		"https://www.perl.org",
		"https://www.php.net",
		"https://www.python.org",
		"https://code.visualstudio.com",
		"https://clojure.org",
	}
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			content := doReq(url)
			title := getTitle(content)
			fmt.Println(title)
		}(u)
	}
	wg.Wait()
}

func doReq(url string) (content string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func getTitle(content string) (title string) {
	re := regexp.MustCompile("<title>(.*)</title>")
	parts := re.FindStringSubmatch(content)
	if len(parts) > 0 {
		return parts[1]
	} else {
		return "no title"
	}
}

func demoRegex() {
	re := regexp.MustCompile(`a(x*)b(y|z)c`)
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))

}
