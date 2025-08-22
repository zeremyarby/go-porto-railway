package renders

import (
	"fmt"
	"net/http"
	"text/template"
)

func JustHtml(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./static/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}

}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./static/template/"+tmpl, "./static/template/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
