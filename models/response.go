package models

//ResponseSuccess represents the success response
type ResponseSuccess struct {
	Records interface{} `json:"record"`
}

//ResponseError represents the error response
type ResponseError struct {
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
	ErrorCode        int    `json:"errorCode"`
	MoreInfo         string `json:"moreInfo"`
}
