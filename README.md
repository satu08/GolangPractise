
**why go?**

efficient compilation

efficient execution

ease of programming

**ASCII, unicode and utf-8**

ascii - 2^8= 256 (1 byte)

unicode - 2^32=(4 byte)

utf-8 -  up to 4 bytes stores unicodes as binary

short declaration operator===    :=

**default values of variables in golang**

int =0

float=0.0

sting=””

boolean= false

pointers,functions,interfaces,slices,channels,maps=nil

golang is statically typed language

(:=) short declaration operator can only be used inside the functions not outside of function  
statically typed: if we declare a variable of one type then it can only hold values of that type

**variable types**

- Boolean, Numeric, String, Array, Slices, Struct, pointer, Function, Interfaces, Map, Channel
- aggregate types-brings many values together
- struct-includes values of different type

the act of constructing a Struct which is composite type is known as composition

**function declaration**

- func add(x, y int) int{

return x+y

}

**terminal commands**  
cd -change directory  
cd .. - go to parent directory  
mkdir - make directory  
touch - create file  
nano - write content of file  
cat - write down the  content of file on terminal  
clear-clear  
chmod - changes permissions  
rm file_name - remove particular file  
rm -rf temp/ - remove folder and all content of that folder  
pwd - print working directory  
ls -la - list all files  
chmod 777 text.txt (we are giving permissions to user, group and world for read(4) , write(2) and execute(1)
vi filename- open file and after opening if you want to insert something press I and insert anything  

**SHA checksum**  
consider we have a .msi file downloaded and when we run that file it contains some 0's and 1'sso when we apply
sha checksum algorithm on those 0's and 1's we get the hashcode
we can do this in terminal by command
- shasum -a 256 filename

**setup**  
-go mod init  
-go build- to build the package basically to create the executables  
now suppose if you want to create build for other os  
GOOS=linux go build -buildvcs=false  
GOOS=windows go build  
GOOS=darwin go build 


**dependency management**  
spaghetti code - write code in one file and using go to statements  
modular code - code distributed in many files using modules , packages  
direct dependency- a dependency your code directly imports  
indirect dependency-dependency used by your direct dependency  
go mod-configures our workspace  
like go mod init mymodule or go mod init github.com/xyzw/abc  
go get-allows to get a third party packages in your code  
if any variable or function is capital letter on first letter of word then it will be visible outside   
that means it is exported otherwise not



**tags git commit and versions**

Vmajor.minor.patch

major-major changes and not backward compatible  
minor-minor changes and backward compatible  
patch-bug fixes and backward compatible


init() function-executes before main as well 

concurrency= way to writing code in concurrent design pattern  
parallelism= executing that concurrent code on multiple cores or  
ability of program to executes multiple tasks simultaneously on multiple cores


checksum to a file
shasum -a 256 ./filepath

**array**  
    a numbered sequence of elements of the same type  
    does not change in size  
    used for go internals  

**slice**  
    build on top of array  
    holds value of same type  
    changes in size  
    has length and capacity  

**map**  
    key value storage  
    an unordered group of elements of one type called element type  
indexed by set of unique keys of another type

**struct**  
    a data structure  
    a composite type  
    allow us to collect values of different types together  


x1: = make([]int,50) -- crate a slice and initialize it with 50–0 values len here is 50 
x1: = make([]int,0,50) -- crate a slice, and it will not initialize it with 0 values so len here is 0

**functions**  
func (receiver) identifier(parameters) (returns) {code}  
parameters-we define a function with parameter  
arguments — we call a function and pass the arguments  

everything in go is pass by value

**interfaces and polymorphism**  
interfaces in go defines a set of method signatures  
polymorphism is ability of value of a certain type to also be of another type  
in go values can be of more than one type

anonymous functions  
self executing functions
defer functions are LIFO

go doc 
go doc -cmd -u

wrapper function wraps or modifies another function behavior while
callback function passed as an argument to be executed at a specific point or event


**pointers**
pointer refer to a variable that holds memory address  
value semantics-values are going to be on the stack  
pointer semantic — values are going to be on the heap  
all the methods of value semantics can be used by pointer semantics as well  
and all the methods of pointer semantics can be used by value semantics  

if you have value of a pointer, you can use pointer receiver or value receiver to implement interfaces  
if you have value of just a value not pointer, you can use only value receivers to implement interfaces


**json**  
Encode- writer  
decode- reader  

**concurrency**  
main function is first go routine  

**Fan-in and Fan-out**  
fan-in == taking values from many channels and putting those values onto one channel  
fan-out == taking some work and putting the chunks of work onto many goroutines


**Error**   
in go is just an interface  
any other type that has this method Error() is also a type Error in Go  

**Testing**  
is in the file that ends with "_test.go"  
put the file in the same package as he one being tested  
be in func with a signature "Func TestXxx(*testing.T)"  
run with - go test  
gofmt- formats the code  
golint - give sugggestion about the code  
go vet  
go test -bench . -- benchmark all the test benchmarks  

go test -cover = shows how  much coverage of code in testing  
go test -coverprofile c.out 
go tool cover -html c.out  -- renders ui on localhost showing coverage   


**MicroServices**  
GET    curl localhost:9091/path
POST   curl localhost:9091/path  -d  'json data'
PUT    curl -X PUT localhost:9091/path -d 'json data to update'


