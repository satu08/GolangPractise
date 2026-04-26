// Stategy Design Pattern in Go?
// 1. What ?
// 2. Why ?

//Strategy pattern defines a family of algorithms, encapsulates each one, and makes them interchangeable at runtime.
//  It replaces conditional logic with polymorphism.
//  In Go, it is implemented using interfaces and struct composition and is widely used for pluggable business logic.
// 3. How ? (Interfaces, Structs and Receiver Functions)
// 4. Pros and Cons

package main

import "fmt"

type IDBconnection interface { // struct can be assigned any type in runtime
	Connect()
}

type DBConnection struct {
	Db IDBconnection // Compatible to accept any type in runtime
}

func (con DBConnection) DBConnect() { // Receiver Function for struct DBConnection
	con.Db.Connect()
}

// Lets implement First behaviour to connect to MySql
type MySqlConnection struct {
	ConnectionString string
}

func (con MySqlConnection) Connect() { // Receiver Function for struct MySqlConnection
	fmt.Println(("MySql " + con.ConnectionString))
}

// Second Behaviour
type PostgressConnection struct {
	ConnectionString string
}

func (con PostgressConnection) Connect() {
	fmt.Println("Postgress " + con.ConnectionString)
}

// Third Behaviour
type MongoDbConnection struct {
	ConnectionString string
}

func (con MongoDbConnection) Connect() {
	fmt.Println("MongoDb " + con.ConnectionString)
}

func main() {

	MySqlConnection := MySqlConnection{ConnectionString: "MySql DB is connected"}
	con := DBConnection{Db: MySqlConnection}
	con.DBConnect()

	pgConnection := PostgressConnection{ConnectionString: "Postgress DB is connected"}
	con2 := DBConnection{Db: pgConnection}
	con2.DBConnect()

	mgConnection := MongoDbConnection{ConnectionString: "Mongo DB is connected"}
	con3 := DBConnection{Db: mgConnection}
	con3.DBConnect()

	gzip := Gzip{}
	service := Service{compressor: gzip}
	service.Process("Hello World")

	zip := Zip{}
	service2 := Service{compressor: zip}
	service2.Process("Hello World")

	service1 := &PaymentService{}

	service1.SetStrategy(CreditCard{})
	service1.Pay(100)

	service1.SetStrategy(UPI{})
	service1.Pay(200)

}

type Compressor interface {
	Compress(data string) string
}

type Gzip struct{}

func (g Gzip) Compress(data string) string {
	return "gzip:" + data
}

type Zip struct{}

func (z Zip) Compress(data string) string {
	return "zip:" + data
}

type Service struct {
	compressor Compressor
}

func (s Service) Process(data string) {
	println(s.compressor.Compress(data))
}

//////////////////// STRATEGY INTERFACE ////////////////////

type PaymentStrategy interface {
	Pay(amount int)
}

//////////////////// CONCRETE STRATEGIES ////////////////////

type CreditCard struct{}

func (c CreditCard) Pay(amount int) {
	fmt.Println("Paid via Credit Card:", amount)
}

type UPI struct{}

func (u UPI) Pay(amount int) {
	fmt.Println("Paid via UPI:", amount)
}

//////////////////// CONTEXT ////////////////////

type PaymentService struct {
	strategy PaymentStrategy
}

func (p *PaymentService) SetStrategy(s PaymentStrategy) {
	p.strategy = s
}

func (p *PaymentService) Pay(amount int) {
	p.strategy.Pay(amount)
}
