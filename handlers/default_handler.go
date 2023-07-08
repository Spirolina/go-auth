package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	dat, err := os.ReadFile("./static/html/index.html")

	tpl := string(dat)

	if err != nil {
		fmt.Println(err)
	}

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}
	err = t.Execute(w, data)
	check(err)

}
