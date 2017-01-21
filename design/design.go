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

var _ = Resource("Hello", func() {
	// DefaultMedia(HelloMedia)
	Action("hello", func() {
		Routing(GET("/hello/:name"))
		Description("API calls your name")
		Params(func() {
			Param("name", String, "Name")
		})
		Response(OK)
	})
})
