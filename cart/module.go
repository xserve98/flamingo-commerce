package cart

import "flamingo/framework/dingo"

type (
	// Module registers our profiler
	Module struct {
		//RouterRegistry *router.RouterRegistry `inject:""`
		//EventRouter    event.Router           `inject:""`
	}
)

// Configure module
func (m *Module) Configure(injector *dingo.Injector) {
	//m.RouterRegistry.Handle("cart.view", new(controller.CartViewController))
	//m.RouterRegistry.Route("/cart", "cart.view")
	//
	//m.RouterRegistry.Handle("cart.api", new(controller.CartApiController))
	//m.RouterRegistry.Route("/api/cart", "cart.view")
	//m.RouterRegistry.Handle("/api/cart/item/add", new(controller.CartItemAddApiController))
	//m.RouterRegistry.Route("/api/cart", "cart.item.add.api")
	//
	////	m.RouterRegistry.Handle("foo", controller.CartItemAddApiController.AddToBasketAction)
	//
	//m.RouterRegistry.Handle("logintest", new(controller.TestLoginController))
	//m.RouterRegistry.Route("/logintest", "logintest")
	//
	//m.EventRouter.AddSubscriber(new(application.EventOrchestration))
	//
	////	a := controller.CartItemAddApiController.AddToBasketAction
	////	a

	//m.RouterRegistry.Mount("/api/cart", new(controller.CartApiController))
	//m.RouterRegistry.Mount("/cart", new(controller.CartController))
}
