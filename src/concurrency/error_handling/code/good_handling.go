package code

import (
	"fmt"
	"net/http"
)

// 응답을 받는 고루틴에서 에러에 관련된 아무런 정보를 알 수 없다.
func GoodHandling() {
	type Result struct {
		Error    error
		Response *http.Response
	}

	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)

		go func() {
			defer close(results)
			for _, url := range urls {
				var result Result
				resp, err := http.Get(url)
				result = Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}

	done := make(chan interface{})

	urls := []string{"https://google.co.kr", "https://badhost.com"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("Error : %v\n", result.Error)
			continue
		} else {
			fmt.Printf("Response : %v\n", result.Response)
		}
	}

}
