package domain

import "fmt"

type Homer interface {
	FindTemplate(ip string) string
}

type Home struct {
}

func (m *Home) FindTemplate(ip string) string {
	fmt.Println("ip", ip)
	var tmpl string
	if ip == "83.51.41.164" {
		tmpl = "home_krls.jet"
	} else {
		tmpl = "home.jet"
	}
	return tmpl
}
