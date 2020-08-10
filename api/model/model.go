package model

type DemoOrder struct {
	Id       int   `json:"id"`
	Orderno  string  `json:"orderno"`
	Username string  `json:"username"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	Fileurl  string  `json:"fileurl"`
	Time     string  `json:"time"`
}

type DemoOrderList struct {
	Matches []DemoOrder `json:"matches"`
}

type Student struct {
	Id      int    `json:"id"`
	Number  string   `json:"number"`
}
