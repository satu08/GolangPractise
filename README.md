Here is your fully formatted README.md content. You can copy and paste this into your file:

---

# Go Language Notes

## Why Go?

- Efficient compilation
- Efficient execution
- Ease of programming

---

## ASCII, Unicode, and UTF-8

- **ASCII**: 2^8 = 256 (1 byte)
- **Unicode**: 2^32 = 4 bytes
- **UTF-8**: Up to 4 bytes, stores Unicode as binary

**Short declaration operator:** `:=`

---

## Default Values of Variables in Go

| Type                                   | Default Value |
|----------------------------------------|--------------|
| int                                   | 0            |
| float                                 | 0.0          |
| string                                | ""           |
| boolean                               | false        |
| pointers, functions, interfaces,       |
| slices, channels, maps                 | nil          |

- Go is a statically typed language.
- `:=` (short declaration operator) can only be used inside functions, not outside.
- Statically typed: if we declare a variable of one type, it can only hold values of that type.

---

## Variable Types

- Boolean, Numeric, String, Array, Slices, Struct, Pointer, Function, Interfaces, Map, Channel
- Aggregate types: bring many values together
- Struct: includes values of different types
- The act of constructing a struct (a composite type) is known as **composition**

---

## Function Declaration

```go
func add(x, y int) int {
    return x + y
}
```

---

## Terminal Commands

```sh
cd            # change directory
cd ..         # go to parent directory
mkdir         # make directory
touch         # create file
nano          # write content of file
cat           # write down the content of file on terminal
clear         # clear terminal
chmod         # changes permissions
rm file_name  # remove particular file
rm -rf temp/  # remove folder and all content of that folder
pwd           # print working directory
ls -la        # list all files
chmod 777 text.txt # give permissions to user, group, and world for read(4), write(2), execute(1)
vi filename   # open file; after opening, press I to insert
```

---

## SHA Checksum

- Consider we have a `.msi` file downloaded. When we run that file, it contains some 0's and 1's. When we apply the SHA checksum algorithm on those 0's and 1's, we get the hashcode.
- Command:

```sh
shasum -a 256 filename
```

---

## Setup

```sh
go mod init
go build # to build the package (create executables)
# To create build for other OS:
GOOS=linux go build -buildvcs=false
GOOS=windows go build
GOOS=darwin go build
```

---

## Dependency Management

- **Spaghetti code**: code in one file, using goto statements
- **Modular code**: code distributed in many files using modules, packages
- **Direct dependency**: a dependency your code directly imports
- **Indirect dependency**: dependency used by your direct dependency
- `go mod` configures our workspace
- Example: `go mod init mymodule` or `go mod init github.com/xyzw/abc`
- `go get` allows you to get third-party packages
- If any variable or function starts with a capital letter, it is exported (visible outside); otherwise, it is not

---

## Tags, Git Commit, and Versions

- Format: `Vmajor.minor.patch`
    - **major**: major changes, not backward compatible
    - **minor**: minor changes, backward compatible
    - **patch**: bug fixes, backward compatible

---

## Special Functions and Concepts

- `init()` function executes before `main()`
- **Concurrency**: way of writing code in concurrent design pattern
- **Parallelism**: executing concurrent code on multiple cores; ability to execute multiple tasks simultaneously on multiple cores

---

## Checksum to a File

```sh
shasum -a 256 ./filepath
```

---

## Array

- A numbered sequence of elements of the same type
- Does not change in size
- Used for Go internals

---

## Slice

- Built on top of array
- Holds values of the same type
- Changes in size
- Has length and capacity

---

## Map

- Key-value storage
- An unordered group of elements of one type (element type)
- Indexed by a set of unique keys of another type

---

## Struct

- A data structure
- A composite type
- Allows us to collect values of different types together

---

## Slice Creation Examples

```go
x1 := make([]int, 50)    // create a slice and initialize it with 50 zeros; len is 50
x1 := make([]int, 0, 50) // create a slice, len is 0, cap is 50
```

---

## Functions

```go
func (receiver) identifier(parameters) (returns) {
    // code
}
```
- Parameters: we define a function with parameters
- Arguments: we call a function and pass the arguments
- Everything in Go is pass by value

---

## Interfaces and Polymorphism

- Interfaces in Go define a set of method signatures
- Polymorphism: ability of a value of a certain type to also be of another type
- In Go, values can be of more than one type
- Anonymous functions
- Self-executing functions
- `defer` functions are LIFO

---

## Go Documentation

```sh
go doc
go doc -cmd -u
```

---

## Wrapper and Callback Functions

- Wrapper function: wraps or modifies another function's behavior
- Callback function: passed as an argument to be executed at a specific point or event

---

## Pointers

- Pointer refers to a variable that holds a memory address
- **Value semantics**: values are on the stack
- **Pointer semantics**: values are on the heap
- All methods of value semantics can be used by pointer semantics as well, and vice versa
- If you have a value of a pointer, you can use pointer receiver or value receiver to implement interfaces
- If you have a value (not pointer), you can use only value receivers to implement interfaces

---

## JSON

- Encode: writer
- Decode: reader

---

## Concurrency

- `main` function is the first goroutine

---

## Fan-in and Fan-out

- **Fan-in**: taking values from many channels and putting those values onto one channel
- **Fan-out**: taking some work and putting the chunks of work onto many goroutines

---

## Error

- In Go, `error` is just an interface
- Any other type that has the method `Error()` is also a type `Error` in Go

---

## Testing

- Test files end with `_test.go`
- Put the file in the same package as the one being tested
- Function signature: `func TestXxx(*testing.T)`
- Run with: `go test`
- `gofmt`: formats the code
- `golint`: gives suggestions about the code
- `go vet`: static analysis
- `go test -bench .`: benchmark all the test benchmarks
- `go test -cover`: shows how much code is covered by tests
- `go test -coverprofile c.out`
- `go tool cover -html c.out`: renders UI on localhost showing coverage

---

## Microservices (Curl Examples)

```sh
GET    curl localhost:9091/path
POST   curl localhost:9091/path  -d  'json data'
PUT    curl -X PUT localhost:9091/path -d 'json data to update'
```