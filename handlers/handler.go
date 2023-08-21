package handlers

import (
	"github.com/willpang/data-entry/datastores/services"
)

type Handler struct {
	UserDatastore services.UserDatastore
}
