# Go Programming Guide: From Basics to Advanced
## Table of Contents
- [Maps and Concurrency](#maps-and-concurrency)
- [Memory Management: Stack vs Heap](#memory-management-stack-vs-heap)
- [Interfaces](#interfaces)
- [Architecture Patterns](#architecture-patterns)
- [Concurrency](#concurrency)
- [Channels](#channels)
- [Algorithms](#algorithms)
- [Type System](#type-system)
- [Language Design](#language-design)
---
## Maps and Concurrency
### Map Safety Guidelines
| Scenario          | Solution    |
| ----------------- | ----------- |
| Single goroutine  | map         |
| Concurrent writes | map + mutex |
| Read-heavy cache  | sync.Map    |
| Ordered data      | slice + map |
> **Note:** Go maps are hash tables backed by buckets and are not safe for concurrent writes or read-write access without synchronization.
---
## Memory Management: Stack vs Heap
### Stack vs Heap Comparison
| Aspect         | Stack          | Heap         |
| -------------- | -------------- | ------------ |
| Speed          | Very fast      | Slower       |
| Lifetime       | Function scope | Until GC     |
| Managed by     | Compiler       | Runtime + GC |
| Sharing        | No             | Yes          |
| GC involvement | ❌              | ✅            |
### Escape Analysis
**Key Question:** "Does this variable need to live beyond this function call?"
- **If YES** → heap
- **If NO** → stack
```bash
go build -gcflags="-m"
```
> Stack memory is fast, function-scoped, and GC-free, while heap memory supports shared and long-lived data and is managed by Go's garbage collector through escape analysis.
**Definition:** Escape analysis is a compile-time analysis where the Go compiler determines whether a variable can be safely allocated on the stack or must be allocated on the heap.
**Key Points:**
- Variables always escape via goroutine → heap allocation
- Escape analysis is a compile-time process where Go determines variables must be heap allocated based on their lifetime and usage
---
## Interfaces
### Runtime Representation
At runtime, a Go interface is a pair of type information and a value pointer. Method calls use dynamic dispatch via an interface table. Interface conversions may cause heap allocation and require careful nil handling.
### Nil Interface Behavior
```go
var p *int = nil
var s []int = nil
var m map[string]int = nil
var ch chan int = nil
var f func() = nil
var i interface{} = nil
```
| Scenario                   | type  | data | i == nil |
| -------------------------- | ----- | ---- | -------- |
| `var i interface{} = nil`  | nil   | nil  | true     |
| `var p *int = nil; i = p`  | *int  | nil  | false    |
| `var s []int = nil; i = s` | []int | nil  | false    |
### Why Use Interfaces?
**Problems without interfaces:**
- Tight coupling
- Hard to test
- Hard to change implementations
- Violates SOLID principles
**Goals:**
- Depend on behavior, not implementation
- Enable mocking
- Allow multiple implementations
- Keep business logic clean
---
## Architecture Patterns
### Layered Architecture
```
API (HTTP)
↓
Service (Business Logic)
↓
Repository (Interfaces)
↓
Infrastructure (DB, API, Cache)
```
### Onion Architecture
```
        ┌────────────────────────┐
        │   Infrastructure       │
        │ (DB, HTTP, Kafka)      │
        └──────────▲─────────────┘
                   │
        ┌──────────┴─────────────┐
        │   Application           │
        │ (Use cases)             │
        └──────────▲─────────────┘
                   │
        ┌──────────┴─────────────┐
        │   Domain                │
        │ (Entities, Rules)       │
        └────────────────────────┘
```
**Dependency Rule:** All arrows point inwards
| Layer          | Depends On     |
| -------------- | -------------- |
| Domain         | Nothing        |
| Application    | Domain         |
| Infrastructure | Application    |
| Frameworks     | Infrastructure |
### Project Structure
```
/internal
  /domain
    user.go
  /application
    rename_user.go
  /infrastructure
    postgres_user_repo.go
  /interfaces
    /http
      user_handler.go
/cmd
  main.go
```
---
## Concurrency
### Goroutines
A **goroutine** is a lightweight user-space thread managed by the Go runtime.
#### Goroutine vs OS Thread
| Feature        | Goroutine             | OS Thread   |
| -------------- | --------------------- | ----------- |
| Creation cost  | Very cheap            | Expensive   |
| Stack size     | Dynamic (2KB → grows) | Fixed (MBs) |
| Scheduling     | Go runtime            | OS kernel   |
| Context switch | Fast                  | Slow        |
### G-P-M Scheduler Model
```
G (goroutines)  --> holds pointer of goroutine execution status, stack, Goroutine object
↓
P (run queue, context)  --> local run queue of goroutines
↓
M (OS thread) ---> actual OS thread, executes goroutine, can block in syscalls
↓
CPU
```
### Goroutine States
| State    | Meaning                           |
| -------- | --------------------------------- |
| Runnable | Ready to run                      |
| Running  | Currently executing               |
| Waiting  | Blocked (channel, mutex, syscall) |
| Dead     | Finished                          |
### Work Stealing Scheduler
Go uses work stealing M:N scheduler to efficiently multiplex goroutines onto OS threads:
```
Many goroutines (M)
↓
Few OS threads (N)
```
---
## Go Runtime
### Overview
```
Go runtime := linked into every binary, not a separate process, runs inside your program
binary = code + go runtime
```
### Core Responsibilities
| Area               | Runtime Handles          |
| ------------------ | ------------------------ |
| Goroutines         | Creation, scheduling     |
| Scheduler          | M:N scheduling           |
| Memory             | Stack/heap allocation    |
| Garbage Collection | Mark & sweep             |
| Channels           | Blocking, wakeups        |
| Maps               | Hashing, resizing        |
| Syscalls           | Non-blocking integration |
| Panic/Recover      | Stack unwinding          |
| Reflection         | Type metadata            |
### OS Kernel vs Go Runtime
| OS Kernel            | Go Runtime           |
| -------------------- | -------------------- |
| Schedules threads    | Schedules goroutines |
| Manages memory pages | Manages heap objects |
| Handles syscalls     | Wraps syscalls       |
| Interrupts           | Safe preemption      |
> The Go runtime is a core part of every Go program that manages goroutines, scheduling, memory allocation, garbage collection, and low-level system interactions, allowing Go to provide efficient concurrency and safety without exposing OS-level complexity to the developer.
---
## Channels
### Channel Types
```go
unbuffered channel = blocking
buffered channel = sender blocks only when buffer is full
                   receiver blocks only when buffer is empty
```
### Channel Close Behavior
```go
close(ch)
```
| Scenario                    | Result            |
| --------------------------- | ----------------- |
| Receive from closed channel | Zero value        |
| Send to closed channel      | Panic             |
| Range over channel          | Drains then exits |
### Unbuffered vs Buffered Channels
| Aspect        | Unbuffered | Buffered  |
| ------------- | ---------- | --------- |
| Communication | Sync       | Async-ish |
| Storage       | No         | Yes       |
| Safety        | High       | Medium    |
| Complexity    | Low        | Higher    |
---
## Synchronization Primitives
### sync.WaitGroup
`sync.WaitGroup` coordinates goroutines by counting active tasks and blocking until all signal completion, with strong memory ordering guarantees.
### context.Context
`context.Context` provides cancellation, deadlines, and request-scoped values, propagating lifecycle control across goroutines and service boundaries.
### Graceful Shutdown
Graceful shutdown ensures a service stops safely by:
- Handling OS signals
- Canceling contexts
- Completing in-flight work
- Releasing resources before exit
---
## Algorithms
### Prefix Sum
```c
prefixSum[0] = arr[0]; 
for (int i = 1; i < n; i++) 
    prefixSum[i] = prefixSum[i - 1] + arr[i]; 
return prefixSum;
```
### Sliding Window
```c
// Compute sum of first window of size k
int max_sum = 0;
for (int i = 0; i < k; i++)
    max_sum += arr[i];
// Compute sums of remaining windows by removing first element 
// of previous window and adding last element of current window
int window_sum = max_sum;
for (int i = k; i < n; i++) {
    window_sum += arr[i] - arr[i - k];
    max_sum = max(max_sum, window_sum);
}
return max_sum;
```
---
## Type System
### Zero Values
Go uses zero values to guarantee safety, simplicity, and predictability. APIs are designed so that the zero value of a type is immediately usable, eliminating the need for constructors and reducing bugs caused by uninitialized state.
### Value vs Reference Semantics
| Type      | Semantics    | Notes                   |
| --------- | ------------ | ----------------------- |
| int, bool | Value        | Full copy               |
| struct    | Value        | Full copy               |
| array     | Value        | Full copy               |
| pointer   | Reference    | Shared                  |
| slice     | Value header | Shared backing array    |
| map       | Value header | Shared hashmap          |
| channel   | Value header | Shared runtime object   |
| string    | Value header | Shared immutable data   |
| interface | Value        | Points to dynamic value |
> **Key Insight:** In Go, assignment copies values. Sharing happens only through pointers inside values.
### Slices: nil vs empty
> "A nil slice represents absence, an empty slice represents presence with no elements. They behave the same operationally but differ semantically and at system boundaries like JSON and databases."
### Slice Append Behavior
> "append checks available capacity; if sufficient it grows the slice in-place, otherwise it allocates a new backing array, copies data, and returns a new slice header. Reallocation happens only when len exceeds cap."
---
## Memory Allocation: new vs make
### What new Cannot Initialize
- Map hash tables
- Channel queues
- Slice backing arrays
> **Summary:** `new` allocates zeroed memory and returns a pointer, while `make` initializes and returns a usable slice, map, or channel by setting up required runtime structures.
---
## Composition vs Inheritance
### "Has-A" vs "Is-A" (Important Interview Point)
| Concept      | Inheritance | Go Embedding |
| ------------ | ----------- | ------------ |
| Relationship | Is-a        | Has-a        |
| Coupling     | Tight       | Loose        |
| Override     | Implicit    | Explicit     |
| Hierarchy    | Deep        | Flat         |
> "Go replaces inheritance with struct embedding and interfaces: embedding enables code reuse via method promotion, while interfaces provide polymorphism without tight coupling."
> "Go avoids inheritance because it creates hidden coupling and fragile hierarchies. Instead, Go uses composition and interfaces to achieve reuse and polymorphism with explicit, maintainable behavior."
---
## Language Design
### Type Systems 
| Concept              | Definition                                      |
| -------------------- | ----------------------------------------------- |
| Statically typed     | Types checked at compile time                   |
| Dynamically typed    | Types checked at runtime                        |
| Compiled language    | Code converted to machine code before execution |
| Interpreted language | Code executed by a runtime engine               |
### String Comparison
String comparison in Go is: **length check → pointer check → byte comparison**.
### Variable Shadowing
Shadowing in Go occurs when a variable is redeclared in an inner scope using `:=`, hiding the outer variable. It's legal but can cause subtle bugs, especially with `err`.
### Function Overloading
> "Go avoids function overloading to keep the language simple and explicit—each function name has one clear meaning, avoiding ambiguity and improving readability and maintainability."
---
## Summary
This guide covers essential Go concepts from basic memory management to advanced architectural patterns. Key takeaways:
1. **Memory:** Understand stack vs heap and escape analysis
2. **Concurrency:** Master goroutines, channels, and the G-P-M scheduler
3. **Design:** Use composition over inheritance, interfaces for abstraction
4. **Safety:** Leverage zero values and proper synchronization primitives
5. **Architecture:** Apply [httpserver](../httpserver)layered or onion architecture for clean, maintainable code

In Go, stack vs heap allocation is decided by the compiler using escape analysis, based on whether a variable can outlive the function in which it was created.
Escape analysis is a compile-time analysis used by the Go compiler to determine whether a value can be safely allocated on the stack. It is conservative because when the compiler cannot prove that a value won’t outlive its function, it chooses heap allocation to guarantee memory safety.

Returning a pointer usually forces heap allocation because the pointed-to value must outlive the function’s stack frame. Escape analysis detects this and moves the value to the heap to prevent use-after-free.

Goroutines cause escape analysis failures because they execute asynchronously. The compiler cannot guarantee that the goroutine will finish before the enclosing function returns, so any captured variables must be heap-allocated to ensure memory safety.

Slices and maps escape to the heap when their backing storage must outlive the creating function. Slices may stay on the stack if they don’t escape, but maps are runtime-managed and typically heap-allocated when they escape.

Heap allocation is slower than stack allocation because it requires allocator bookkeeping, possible synchronization, and garbage collection, while stack allocation is just a pointer adjustment and is freed automatically.

“Go’s GC tracks heap objects by starting from GC roots, precisely scanning pointers using compiler-generated metadata, marking all reachable objects with a tri-color algorithm, and sweeping unreachable objects. It runs mostly concurrently with the program using write barriers for safety.


Each goroutine starts with a small stack that grows and shrinks automatically. When a function needs more space, the runtime allocates a larger stack, copies the old one, updates pointers, and continues execution. Because stacks can move, variables that outlive a function must be heap-allocated.

Interface boxing causes allocations because assigning a concrete value to an interface hides its concrete type and lifetime. The interface stores a pointer to the value, and since the compiler cannot prove the value won’t outlive the stack frame, it conservatively allocates the value on the heap.


Heap allocation should be preferred when data must outlive a function, be shared across goroutines, preserve identity, avoid large copies, or represent long-lived state. Stack allocation is an optimization; heap allocation is a correctness and design tool.”

A Go interface is internally a pair of pointers: one to type/method metadata and one to the actual data.

Use interfaces at boundaries, not in hot loops.

An interface in Go is nil only if both its dynamic type and value are nil.
If an interface holds a typed nil value, the interface itself is non-nil because it still carries type information


Type assertion is a runtime check on an interface’s dynamic type, implemented via pointer comparisons and cached metadata.


Rule 1

A value type has only value-receiver methods.

Rule 2

A pointer type has both value and pointer-receiver methods.

Rule 3

A type implements an interface if its method set is a superset of the interface’s method set.


In Go, a type implements an interface automatically if it has all the required methods.


The package that uses the behavior should define the interface — not the package that provides the concrete type.


Interfaces in Go should be consumer-defined because the consumer knows the minimal behavior it requires. This minimizes coupling, keeps interfaces small and stable, enables retrofitting, and makes testing easier.


Interfaces should be avoided when there is only one implementation, no real abstraction need, performance-critical code paths, or when they obscure concrete behavior and invariants. In Go, interfaces are best introduced only when required by usage, not by design speculation.

Reflection in Go operates on interface values by inspecting their dynamic type and data pointer. It exposes the same runtime information used by interfaces, which is why reflection requires interface values and why pointers are necessary to modify data.


| Aspect                    | Concurrency             | Parallelism              |
| ------------------------- | ----------------------- | ------------------------ |
| Meaning                   | Managing multiple tasks | Executing multiple tasks |
| Focus                     | Program structure       | Hardware utilization     |
| CPU cores needed          | 1 or more               | 2 or more                |
| Execution                 | Interleaved             | Simultaneous             |
| Exists without the other? | Yes                     | No                       |


Concurrency is about dealing with many things at once.
Parallelism is about doing many things at the same time.

Go programs are written concurrently; they become parallel if hardware allows.

You can have concurrent code that is not parallel
You cannot have parallel execution without concurrency

Go uses M:N scheduling to multiplex many goroutines onto fewer OS threads.

The Go runtime, not the OS, controls scheduling.

Blocking a goroutine does not block an OS thread.


G is the task, M is the worker, P is the permit.

Ps limit parallelism; Ms execute; Gs do the work.

GOMAXPROCS controls Ps, not goroutines.

Go handles blocking system calls by detaching the OS thread from its scheduler context and using another thread to continue executing goroutines. For network I/O, Go uses non-blocking syscalls with a poller to avoid blocking threads altogether.

Go detects deadlocks by checking whether any goroutine can ever become runnable again.

Deadlock is detected when all goroutines are blocked and no timers or I/O events exist.

The Go runtime performs deadlock detection at the scheduler level, not via timeouts.

A goroutine leak occurs when a goroutine blocks forever and can never exit.

Go can garbage-collect memory, but not goroutines.

Most goroutine leaks are caused by missing cancellation or channel misuse.


