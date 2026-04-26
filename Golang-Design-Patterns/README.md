Classified into 3 Categories


1) Creational patterns

    a) Abstract Factory  
    b) Builder   -----> done
    c) Factory   ----> done
    d) prototype  ----> done
    e) singleton  ---> done


2) Structural Patterns

    a) adapter   ---> done
    b) bridge   
    c) Composite  
    d) decorater   -----> done
    e) facade    -------> done
    f) proxy  -----> done


3) Behavioural Patterns

    a) chain of Responsibility  -------> done
    b) Iterator  ---------------------> done
    c) Momento  
    d) State  --------------------------> done
    e) Template Method  
    f) Command  
    g) Mediator   
    h) Observer   -------------> done
    i) Strategy   --------------> done
    j) visitor  



Dependency Injection is a design technique where dependencies are provided to a component instead of being created inside it. In Go, it is typically implemented using constructor injection, often combined with interfaces for abstraction. DI improves testability, flexibility, and decoupling, and is preferred over global state like singletons.



type Name interface {
	GetName() string
}

type satya struct {
	name string
}

func (a *satya) GetName() string {
	return a.name
}

// Option 1
func NewSatya(name string) Name {
	return &satya{name: name}
}

// Option 2
func NewSatyaConcrete(name string) *satya {
	return &satya{name: name}
}


| Case                     | Return Interface    | Return Struct |
| ------------------------ | ------------------- | ------------- |
| Multiple implementations | ✅                   | ❌             |
| Need abstraction         | ✅                   | ❌             |
| Simple use case          | ❌                   | ✅             |
| Internal code            | ❌                   | ✅             |
| Public API               | Often ✅             | Sometimes     |
| Performance critical     | ❌ (avoid interface) | ✅             |


In Go, constructors typically return concrete struct pointers for simplicity and performance. Interfaces are returned only when abstraction, multiple implementations, or encapsulation is required. A common guideline is “accept interfaces, return structs.”




| Pattern | Purpose          |
| ------- | ---------------- |
| Facade  | Simplify system  |
| Proxy   | Control access   |
| Adapter | Change interface |
