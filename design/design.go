package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Demo App", func() {
	Title("The Hello App")
	Description("Say Hello")
	Scheme("http")
	Host("localhost:8080")
})
