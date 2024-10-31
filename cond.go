package main

import (
	"fmt"
)

func DomainForLocale(domain, locale string) string {
	subdomain := ""
	if locale == "" {
		subdomain = "en"
	} else {
		subdomain = locale
	}
	return fmt.Sprintf("%s.%s", subdomain, domain)
}

func main() {
	locale1 := ""
	locale2 := "ru"
	domain := "site.com"
	fmt.Println(DomainForLocale(domain, locale1))
	fmt.Println(DomainForLocale(domain, locale2))
}
