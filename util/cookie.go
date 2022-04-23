package util

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/SultanKs4/gi-daily/setting"
)

func loadCookies() ([]*http.Cookie, error) {
	var cookies []*http.Cookie
	cookieParse, err := parseCookieEnv()
	if err != nil {
		return nil, err
	}
	for k, v := range cookieParse {
		cookie := &http.Cookie{
			Name:  k,
			Value: v,
		}
		cookies = append(cookies, cookie)
	}
	return cookies, nil
}

func parseCookieEnv() (map[string]string, error) {
	cookie := map[string]string{}
	cookieRaw := setting.Environment.Cookie
	if cookieRaw == "" {
		return nil, fmt.Errorf("cookie error: cookie not found")
	}
	for _, v := range strings.Split(cookieRaw, "; ") {
		data := strings.Split(v, "=")
		cookie[data[0]] = data[1]
	}
	return cookie, nil
}

func JarCookies(rawurl string) (jar *cookiejar.Jar, err error) {
	jar, err = cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("create Jar error: %v", err)
	}
	u, err := url.Parse(rawurl)
	if err != nil {
		return nil, fmt.Errorf("parse URL error: %v", err)
	}
	cookie, err := loadCookies()
	if err != nil {
		return nil, err
	}
	jar.SetCookies(u, cookie)
	return jar, nil
}
