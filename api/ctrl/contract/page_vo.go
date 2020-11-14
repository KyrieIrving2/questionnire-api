package contract

type PageVo struct {
	Total   int         `json:"total"`
	Results interface{} `json:"results"`
}
