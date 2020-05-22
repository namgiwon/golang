package code

import (
	"fmt"
	"net/http"
)

// 응답을 받는 고루틴에서 에러에 관련된 아무런 정보를 알 수 없다.
func GoodHandling() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan *http.Response {
		responses := make(chan *http.Response)

		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})

	urls := []string{"https://google.co.kr", "https://badhost.com"}

	for response := range checkStatus(done, urls...) {
		fmt.Printf("Response : %v\n", response)
	}

}
