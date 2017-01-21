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
	DefaultMedia(HelloMedia)
	Action("hello", func() {
		Routing(GET("/hello/:name"))
		Description("API calls your name")
		Params(func() {
			Param("name", String, "Name")
		})
		Response(OK)
	})
})

var HelloMedia = MediaType("application/vnd.goa.example.hello+json", func() {
	Description("Message from Demo App API")
	Attributes(func() {
		Attribute("msg", String, "")
		Required("msg")
	})
	View("default", func() {
		Attribute("msg")
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "swaggerui/dist")
})
