package models

//ResponseSuccess representa o response caso retorne sucesso
type ResponseSuccess struct {
	Records interface{} `json:"record"`
}

//ResponseError representa o response
type ResponseError struct {
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
	ErrorCode        int    `json:"errorCode"`
	MoreInfo         string `json:"moreInfo"`
}
