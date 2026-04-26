package main

//Facade pattern provides a simplified interface to a complex subsystem by wrapping multiple components behind a single entry point.
// In Go, it is implemented using a struct that orchestrates calls to underlying services, improving usability and reducing coupling

type AuthService struct{}

func (a AuthService) Login() {
	println("User authenticated")
}

type PaymentService struct{}

func (p PaymentService) Pay() {
	println("Payment processed")
}

type InventoryService struct{}

func (i InventoryService) Check() {
	println("Inventory checked")
}

type OrderFacade struct {
	auth      AuthService
	payment   PaymentService
	inventory InventoryService
}

func (o OrderFacade) PlaceOrder() {
	o.auth.Login()
	o.inventory.Check()
	o.payment.Pay()
	println("Order placed successfully")
}

func main() {
	facade := OrderFacade{
		auth:      AuthService{},
		payment:   PaymentService{},
		inventory: InventoryService{},
	}
	facade.PlaceOrder()
}
