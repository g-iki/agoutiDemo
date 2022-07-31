package structure

//Setting 設定ファイル構造体
type Setting struct {
	Headless bool   `json:"headless"`
	LoginURL string `json:"loginURL"`
	Password string `json:"password"`
	Users    []User `json:"users"`
}

type User struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}
