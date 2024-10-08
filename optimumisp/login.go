package optimumisp

import (
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
)

// processLogin sends a POST request to log in and retrieves a session token
func (c *Client) ProcessLogin(username, password string) {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()
	page := stealth.MustPage(browser)
	var e proto.NetworkResponseReceived
	wait := page.WaitEvent(&e)
	page.MustNavigate("https://www.optimum.net/login")
	page.MustElement(`input[id="loginPageUsername"]`).MustInput(username)
	page.MustElement(`input[id="loginPagePassword"]`).MustInput(password)
	fmt.Println("Logging in...")
	page.MustElement(`#target`).MustClick() // Wait for the page to finish loading after login
	wait()
	fmt.Printf("Headers: %#v\n", e.Response)
	page.MustWaitNavigation()
	// Get all cookies

	var cookies []*proto.NetworkCookie
	var timeout = 20
	for i := 0; i < timeout; i++ {
		fmt.Println("Waiting for cookies, can take up to 15 seconds")
		cks := page.MustCookies()
		for _, cookie := range cks {
			if cookie.Name == "user-jwt" {
				fmt.Printf("Got %d cookies, including user-jwt\n", len(cks))
				cookies = cks
				break
			}
		}
		if len(cookies) > 0 {
			break
		}
		time.Sleep(time.Second)
	}

	if len(cookies) == 0 {
		log.Fatalf("Get cookies timed out")
	}

	c.cookies = cookies
}
