package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"agoutiDemo/structure"

	"github.com/sclevine/agouti"
)

var setting structure.Setting

func main() {
	loadSetting()

	posi := "--window-position=0,0" // ブラウザのxy位置を指定
	size := "--window-size=1920,1080"    // ブラウザのサイズを指定
	var param []string
	param = append(param, size)
	param = append(param, posi)
	if setting.Headless {
		param = append(param, "--headless")
	}

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions(
			"args",
			param,
		),
	)
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	page, err := driver.NewPage()
	if err != nil {
		log.Printf("Failed to open page:%v", err)
		return
	}
	// ログインページに遷移
	if err := page.Navigate(setting.LoginURL); err != nil {
		log.Printf("Failed to navigate:%v", err)
		return
	}

	// 要素を探し出して入力
	page.FindByID("email").Fill(setting.UserID)
	page.FindByID("password").Fill(setting.Password)

	page.Screenshot("screenshot_before_login.png")
	page.FindByID("login-button").Click()
	page.Screenshot("screenshot_after_login.png")

	log.Printf("******* Close all ChromeTestBrowser when this Program Closed *******")
	// 無限ループ(Ctrl+Cで抜ける) // ここで無限ループさせておいて、Ctrl+Cで終了することでChromeDriverのプロセスが残らないようにする
	for {

	}
}

func loadSetting() {
	// read setting file
	r, err := ioutil.ReadFile("./setting.json")
	if err != nil {
		log.Println(err.Error())
		log.Println("ERROR cannot read setting file")
		os.Exit(1)
	}
	json.Unmarshal(r, &setting)
}
