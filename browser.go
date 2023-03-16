package main

import (
	_ "database/sql/driver"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type SeleniumBot struct {
	driver selenium.WebDriver
	worker Worker
}

func NewSeleniumBot(w Worker) SeleniumBot {
	service, _ := selenium.NewChromeDriverService("./chromedriver", 4444)
	fmt.Println("\nДрайвер ок", service)
	caps := selenium.Capabilities{}

	caps.AddChrome(chrome.Capabilities{
		Path: "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
		Args: []string{
			"start-maximized",
			"disable-infobars",
			"--disable-gpu",
			"--user-data-dir=C:\\Users\\nikit\\AppData\\Local\\Google\\Chrome\\User Data",
			"--profile-directory=Profile 6",
			"--disable-dev-shm-usage",
			"--no-sandbox",

			// "--app-id=doifjfpmnajihigeedhhicefgjdpdnio",
			// "--headless", // comment out this line to see the browser
		}})

	fmt.Println("\nПараметры ок")
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 5)
	}
	fmt.Println("\nЗапуск ок")
	return SeleniumBot{
		driver: driver,
		worker: w,
	}
}

func (b SeleniumBot) wbPoint() {
	b.driver.Get("https://point.wb.ru/orders/ready")
	fmt.Println("\nИдем по ссылке ок")
	time.Sleep(time.Second * 5)
	for {
		name, err := b.driver.FindElement(selenium.ByClassName, "pp-info-block")
		if err == nil {
			textList := fmt.Sprintln(name.Text())
			s := strings.Fields(textList + "\t \n")
			fmt.Println("\nЧто-то происходит?")
			user := b.getName(s[6], s[5])
			fmt.Println("selenium user: ", user)
			b.worker.addName(user)
			break
		} else {
			fmt.Println("\nЛогина нет, ждем 1 мин browser")
			time.Sleep(time.Second * 30)
		}
	}
}

func (b SeleniumBot) getName(s1 string, s2 string) string {
	if s1+s2 != "" && !regexp.MustCompile(`\d`).MatchString(s1+s2) {
		return s1 + " " + s2
	}
	return "какой-то"
}
