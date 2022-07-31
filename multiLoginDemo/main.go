package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"sync"

	"multiLoginDemo/structure"

	"github.com/sclevine/agouti"
)

var setting structure.Setting

func main() {
	loadSetting()

	wg := &sync.WaitGroup{}
	for i := 0; i < len(setting.Users); i++ {
		wg.Add(1)
		go multiAccess(i, setting.Users[i], wg)
	}
	wg.Wait()

	log.Printf("******* Close all ChromeTestBrowser when this Program Closed *******")
	// 無限ループ(Ctrl+Cで抜ける)
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

func multiAccess(idx int, user structure.User, wg *sync.WaitGroup) {
	defer wg.Done()

	// 適度にずらす
	left := strconv.Itoa((idx + 1) * 20)
	top := strconv.Itoa((idx + 1) * 20)

	posi := "--window-position=" + left + "," + top

	size := "--window-size=1920,1080"
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

	page.FindByID("email").Fill(user.UserID)
	page.FindByID("password").Fill(user.Password)

	page.FindByID("login-button").Click()

	fileName := fmt.Sprintf("screenshot_logined%02d.png", idx)
	page.Screenshot(fileName)
}
