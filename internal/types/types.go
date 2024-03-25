package types

type Item struct {
	Shelf     string `json:"shelf"`
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Amount    int    `json:"amount"`
	Extra     string `json:"extra"`
	ProductID int    `json:"productid"`
}

type Config struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}
