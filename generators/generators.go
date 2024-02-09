package generators

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Titulo(urls ...string) <-chan string {
	fmt.Println("Gerando títulos")
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

// func main() {

// 	urls := []string{"https://www.cod3r.com.br", "https://www.google.com", "https://www.mentorcycle.com.br", "https://www.youtube.com"}

// 	t1 := Titulo(urls...)

// 	for i := 0; i < len(urls); i++ {
// 		fmt.Println(<-t1)
// 	}

// }
