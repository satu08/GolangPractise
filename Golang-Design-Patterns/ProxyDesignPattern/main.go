package main

import "fmt"

//////////////////// SUBJECT INTERFACE ////////////////////

type DataSource interface {
	GetData()
}

//////////////////// REAL OBJECT ////////////////////

type RealDB struct{}

func (r RealDB) GetData() {
	fmt.Println("Fetching from real DB...")
}

//////////////////// PROXY ////////////////////

type ProxyDB struct {
	real  RealDB
	cache bool
}

func (p *ProxyDB) GetData() {
	if p.cache {
		fmt.Println("Returning cached data")
		return
	}

	fmt.Println("Proxy: checking access...")
	p.real.GetData()
	p.cache = true
}

//////////////////// MAIN ////////////////////

func main() {
	db := &ProxyDB{}

	db.GetData() // hits real DB
	db.GetData() // cached
}


// Proxy pattern provides a surrogate object that controls access to a real object.
//  It implements the same interface and can add functionality like access control, caching, or lazy initialization.
//  In Go, it is implemented using interfaces and struct composition and is commonly used in API gateways and caching layers.