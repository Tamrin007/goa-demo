package main

import (
	"github.com/Tamrin007/goa-demo/app"
	"github.com/goadesign/goa"
)

// HelloController implements the Hello resource.
type HelloController struct {
	*goa.Controller
}

// NewHelloController creates a Hello controller.
func NewHelloController(service *goa.Service) *HelloController {
	return &HelloController{Controller: service.NewController("HelloController")}
}

// Hello runs the hello action.
func (c *HelloController) Hello(ctx *app.HelloHelloContext) error {
	// HelloController_Hello: start_implement

	// Put your logic here
	msg := "Hello, " + ctx.Name + "!"

	// HelloController_Hello: end_implement
	res := &app.GoaExampleHello{
		Msg: msg,
	}
	return ctx.OK(res)
}
