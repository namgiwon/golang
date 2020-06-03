package model

// HelloWorldResponse represents a http response.
type HelloWorldResponse struct {
	// 출력 필드를 message로 바꿈
	Message string `json:"message"`

	// 필드를 출력하지 않음
	Author string `json:"-"`

	// 값이 비어 있으면 출력하지 않음
	Date string `json:", omitempty"`

	// 출력을 문자열로 변환하고 이름을 id로 바꿈
	Id int `json:"id, string"`
}
