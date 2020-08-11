package model

type DemoOrder struct {
	Id       int64   `json:"id"`
	Orderno  string  `json:"orderno"`
	Username string  `json:"username"`
	Amount   float32 `json:"amount"`
	Status   string  `json:"status"`
	Fileurl  string  `json:"fileurl"`
	Time     string  `json:"time"`
}

type DemoOrderList struct {
	Matches []DemoOrder `json:"matches"`
}

type Student struct {
	Id      int64    `json:"id"`
	Number  string   `json:"number"`
}
