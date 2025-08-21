package handlers

import (
	"net/http"

	"github.com/zeremyarby/go-porto/pkg/config"
	"github.com/zeremyarby/go-porto/pkg/renders"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renders.JustHtml(w, "about.html")
}

func MyWork(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "mywork.page.tmpl")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "contact.page.tmpl")
}
