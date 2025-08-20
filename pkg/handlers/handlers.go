package handlers

import (
	"net/http"

	"github.com/zeremyarby/go-porto/pkg/config"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

}
