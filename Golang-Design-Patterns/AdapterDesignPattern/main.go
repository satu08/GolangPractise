package main

import "fmt"

/*
structural design pattern
it is used so that two unrelated objects can work together using adapter
it is also known as wrapper
the thing which join these two unrelated objects is called as adapter



four participants

Target Interface: this is the interface which will be used by clients
Adapter: this is the wrapper which implements the target interface and modifies the specific request available from the adaptee class
Adaptee : this is the object which is used by the adapter to reuse the existing functionality and modify them for desired use
Client : this will interact with adapter

when you dont want ti change the existing object or interface rather wants to add new functionality on top of what is existing
*/

// Target Interface
type mobile interface {
	chargeAppleMobile()
}

// Concrete Prototype Implementation
type apple struct{}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Charging APPLE mobile")
}

// Adaptee
type android struct{}

func (a *android) chargeAndroidMobile() {
	fmt.Println("Charging ANDROID mobile")
}

// Adapter
type androiddapter struct {
	android *android
}

func (ad *androiddapter) chargeAppleMobile() {
	ad.android.chargeAndroidMobile()
}

// Client
type client struct{}

func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleMobile()
}

func main() {
	// First/Initial Requirement
	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)

	// Extended requirement i.e. Charge Android Mobile.
	android := &android{}
	androiddapter := &androiddapter{
		android: android,
	}
	client.chargeMobile(androiddapter)
}
