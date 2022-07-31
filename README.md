# agoutiDemo


## loginDemo

- ブラウザを一つ開いてログイン処理を行います

setting.json サンプル

```json
{
	"headless": false, // ヘッドレス(GUI上の表示なしで)実行を行うかどうか
	"loginURL": "https://hotel.testplanisphere.dev/ja/login.html",  // ログインページのURL
	"userID": "ichiro@example.com", // ID
	"password": "password" // パスワード
}
```


## multiLoginDemo

- 複数ユーザー分ブラウザを開いてログイン処理を行います

setting.json サンプル

```
{
	"headless": false,
	"loginURL": "https://hotel.testplanisphere.dev/ja/login.html",
	"users": [
		{
			"userID": "ichiro@example.com",
			"password": "password"
		},
		{
			"userID": "sakura@example.com",
			"password": "pass1234"
		},
		{
			"userID": "jun@example.com",
			"password": "pa55w0rd!"
		},
		{
			"userID": "yoshiki@example.com",
			"password": "pass-pass"
		}
	]

}
```
