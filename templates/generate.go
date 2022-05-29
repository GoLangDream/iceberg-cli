package templates

import _ "embed"

//go:embed generate/controller.go.tmpl
var controllerTmpl string

//go:embed generate/view.html.tmpl
var viewTmpl string

func ControllerTmpl() string {
	return controllerTmpl
}

func ViewTmpl() string {
	return viewTmpl
}
