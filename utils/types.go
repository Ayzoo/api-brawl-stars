package utils

import (
	"text/template"
)

type Username struct {
	User     string
	Password string
	Database string
}

type Engine struct {
	Templates *template.Template
}
