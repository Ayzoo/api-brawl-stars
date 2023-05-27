package utils

import (
	"text/template"

	core "github.com/gofiber/template"
)

type Username struct {
	User     string
	Password string
	Database string
}

type Engine struct {
	core.Engine

	Templates *template.Template
}
