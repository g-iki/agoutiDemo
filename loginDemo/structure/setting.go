package structure

//Setting 設定ファイル構造体
type Setting struct {
	Headless     bool   `json:"headless"`
	LoginURL     string `json:"loginURL"`
	UserID       string `json:"userID"`
	Password     string `json:"password"`
}
