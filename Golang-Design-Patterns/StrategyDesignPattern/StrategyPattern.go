// Stategy Design Pattern in Go?
// 1. What ?
// 2. Why ?
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
