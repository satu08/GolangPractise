# 🧠 1. DSA (120 Questions)

## Arrays & Strings (25)

1. Two sum (all variations)
2. Three sum (duplicates handling)
3. Four sum optimized
4. Container with most water
5. Product of array except self (no division)
6. Longest substring without repeating chars
7. Longest substring with K distinct
8. Minimum window substring
9. Sliding window maximum
10. Rotate array (in-place)
11. Spiral matrix
12. Set matrix zeroes (O(1) space)
13. Trapping rain water
14. Candy problem
15. Gas station circular tour
16. Maximum product subarray
17. Kadane’s variations
18. Partition array (Dutch flag)
19. Rearrange positives/negatives
20. Subarray sum equals K
21. Count subarrays with XOR K
22. Find duplicate number (cycle detection)
23. Missing number (bit manipulation)
24. Merge intervals
25. Insert interval

---

## Linked List (15)

26. Reverse linked list (iterative + recursive)
27. Reverse in K groups
28. Detect cycle (Floyd)
29. Intersection of two lists
30. Remove nth from end
31. Flatten multilevel list
32. Copy list with random pointer
33. Merge K sorted lists
34. Palindrome linked list
35. Rotate list
36. Add two numbers (LL)
37. Split list into parts
38. Odd-even list
39. Delete node without head
40. Design browser history

---

## Stack & Queue (15)

41. Min stack
42. Max stack
43. Valid parentheses
44. Largest rectangle histogram
45. Daily temperatures
46. Next greater element
47. Sliding window max (deque)
48. Evaluate postfix
49. Design circular queue
50. Implement queue using stacks
51. Implement stack using queue
52. Decode string
53. Simplify path
54. Basic calculator (I, II, III)
55. Design hit counter

---

## Trees (25)

56. Binary tree traversal (all)
57. Level order zigzag
58. Diameter of tree
59. Maximum path sum
60. Lowest common ancestor
61. Serialize/deserialize tree
62. Validate BST
63. Kth smallest in BST
64. Convert BST to DLL
65. Flatten binary tree
66. Symmetric tree
67. Path sum variants
68. Build tree from traversals
69. Count complete tree nodes
70. Binary tree right view
71. Vertical order traversal
72. Boundary traversal
73. Distance between nodes
74. Burn tree problem
75. Recover BST
76. BST iterator
77. Merge two BSTs
78. Largest BST in BT
79. Sum root-to-leaf numbers
80. Check balanced tree

---

## Graphs (20)

81. BFS/DFS basics
82. Number of islands
83. Clone graph
84. Course schedule (topo sort)
85. Detect cycle (directed/undirected)
86. Dijkstra
87. Bellman-Ford
88. Floyd-Warshall
89. Minimum spanning tree (Kruskal/Prim)
90. Word ladder
91. Alien dictionary
92. Shortest path in grid
93. Rotten oranges
94. Pacific Atlantic water flow
95. Graph coloring
96. Bipartite graph
97. Reconstruct itinerary
98. Strongly connected components
99. Union-Find applications
100. Critical connections (Tarjan)

---

## DP (20)

101. Fibonacci variations
102. Climbing stairs
103. House robber (I, II)
104. Coin change (min ways)
105. LIS (O(n log n))
106. LCS
107. Edit distance
108. Knapsack
109. Partition equal subset
110. Word break
111. Palindrome partitioning
112. Matrix chain multiplication
113. Burst balloons
114. Egg dropping
115. Decode ways
116. Unique paths
117. Maximum square
118. Rod cutting
119. Paint house
120. Stock buy/sell (all variations)

---

# ⚙️ 2. Golang Deep Dive (120 Questions)

## Basics → Advanced

1. What is Go runtime?
2. How scheduler works?
3. What is GOMAXPROCS?
4. Goroutine vs thread
5. Channel internals
6. Buffered vs unbuffered
7. Select behavior
8. Deadlock detection
9. Mutex vs RWMutex
10. Atomic operations

---

## Memory & Performance (30)

11. Escape analysis
12. Stack growth
13. Heap allocation triggers
14. Garbage collector phases
15. STW pauses
16. Memory leaks in Go
17. Profiling using pprof
18. CPU vs memory bottleneck
19. Object pooling
20. sync.Pool usage
21. Cache locality
22. False sharing
23. Memory alignment
24. Struct optimization
25. Slice internals
26. Map internals
27. Map resizing
28. Nil vs empty slice
29. Pointer vs value
30. Copy vs reference

---

## Concurrency (40)

31. Goroutine leak scenarios
32. Context cancellation
33. Worker pool design
34. Rate limiter
35. Semaphore
36. Fan-in/fan-out
37. Pipeline design
38. Backpressure handling
39. Channel closing rules
40. Select starvation
41. Race condition debugging
42. Deadlock debugging
43. Concurrent map patterns
44. Lock contention
45. Distributed locking
46. Retry mechanism
47. Circuit breaker
48. Timeout handling
49. Graceful shutdown
50. Signal handling

---

## Advanced Go Design (50)

51. Interface design best practices
52. Composition vs inheritance
53. Dependency injection in Go
54. Functional options pattern
55. Middleware design
56. Plugin system
57. Reflection usage
58. Generics in Go
59. Error handling patterns
60. Custom error types
61. Logging architecture
62. Config management
63. API versioning
64. Code organization
65. Clean architecture in Go
66. Hexagonal architecture
67. Testing strategies
68. Mocking interfaces
69. Integration testing
70. Benchmarking
71. Writing reusable packages
72. Avoiding cyclic dependencies
73. Code smells in Go
74. Refactoring techniques
75. Writing thread-safe code
76. Performance tuning
77. Handling high concurrency
78. Building CLI tools
79. Plugin architecture
80. Feature flags
    81–120. (Variants: same topics applied to real scenarios)

---

# 🏗️ 3. System Design (120 Questions)

## Core Systems

1. URL shortener
2. Pastebin
3. Dropbox
4. Netflix
5. YouTube
6. WhatsApp
7. Twitter
8. Uber
9. Google Docs
10. Notification system

---

## Your Domain (CRITICAL)

11. Edge device system
12. IoT platform
13. Protocol gateway
14. Telemetry pipeline
15. Device sync system
16. Offline-first system
17. Firmware update system
18. Device authentication
19. Gateway clustering
20. Real-time monitoring

---

## Deep Concepts

21. CAP theorem
22. Consistency models
23. Idempotency
24. Retry strategies
25. Backpressure
26. Rate limiting
27. Caching strategies
28. CDN design
29. Load balancing
30. Database sharding

---

## Advanced (60 more variations)

* Scale to 1M → 100M users
* Multi-region design
* Failure handling
* Cost optimization
* Security design
* Observability

---

# 🌐 4. Distributed Systems (80 Questions)

1. What is distributed system?
2. CAP theorem real case
3. Eventual consistency
4. Strong consistency
5. Leader election
6. Raft vs Paxos
7. Gossip protocol
8. Vector clocks
9. Distributed locks
10. Exactly-once semantics
11. Kafka internals
12. Partitioning strategy
13. Consumer groups
14. Offset management
15. Message ordering
16. Retry + DLQ
17. Saga pattern
18. Circuit breaker
19. Bulkhead pattern
20. Failover strategies

(+ 60 scenario-based variations)

---

# 🧱 5. SOLID + Design Patterns (40 Questions)

1. Explain SOLID deeply
2. Violations examples
3. Strategy pattern in Go
4. Factory pattern
5. Builder pattern
6. Observer pattern
7. Adapter pattern
8. Decorator pattern
9. Chain of responsibility
10. State pattern
11. Singleton pitfalls
12. Dependency inversion
13. Interface segregation
14. Open/closed principle
15. Real-world refactoring
    16–40. Variations on real systems

---

# 🐳 6. DevOps / Docker / CI-CD (40 Questions)

1. Docker architecture
2. Container lifecycle
3. Multi-stage builds
4. Image optimization
5. Networking
6. Volumes
7. Kubernetes basics
8. CI/CD pipelines
9. GitLab runners
10. Rollbacks
11. Canary deploy
12. Blue-green deploy
13. Observability
14. Logging
15. Secrets management
    16–40. Real-world debugging

---

# 🌍 7. Networking & Protocols (30 Questions)

1. TCP vs UDP
2. HTTP/1.1 vs HTTP/2
3. gRPC internals
4. WebSockets
5. MQTT QoS
6. BACnet working
7. Modbus protocol
8. TLS handshake
9. Load balancing
10. Reverse proxy
    11–30. Failure + optimization scenarios

---

# 🧠 8. Behavioral + Senior (30 Questions)

1. Biggest system built
2. Production failure
3. Scaling decision
4. Conflict resolution
5. Mentoring juniors
6. Handling ambiguity
7. Tradeoffs taken
8. Ownership example
9. Design decision regret
10. Performance optimization story
    11–30. Deep probing follow-ups




# 🔥 PART 1: PUBMATIC-STYLE CODING SHEET (80 QUESTIONS)

---

## 🧩 A. ARRAY / TWO POINTER (10)

1. Two Sum (all variations)
2. 3Sum / 4Sum
3. Container With Most Water
4. Trapping Rain Water ⭐
5. Maximum Subarray (Kadane)
6. Product of Array Except Self
7. Merge Intervals ⭐
8. Insert Interval
9. Rotate Array
10. Find Duplicate Number (Floyd cycle)

---

## 🔤 B. STRING (10)

11. Longest Substring Without Repeating
12. Minimum Window Substring ⭐
13. Group Anagrams ⭐ (IMPORTANT - PubMatic style)
14. Valid Anagram
15. K-Anagrams ⭐ (ACTUAL asked)
16. String Compression
17. Decode String
18. Longest Palindromic Substring
19. Rabin-Karp pattern search
20. Check if one string is rotation of another

---

## 🔗 C. LINKED LIST (8)

21. Reverse Linked List ⭐
22. Detect Cycle (Floyd) ⭐
23. LRU Cache ⭐⭐⭐ (VERY IMPORTANT)
24. Merge Two Sorted Lists
25. Remove Nth Node From End
26. Flatten Linked List
27. Intersection of Two Lists
28. Copy List with Random Pointer

---

## 🌳 D. TREES (10)

29. Binary Tree Level Order
30. Lowest Common Ancestor ⭐
31. Diameter of Binary Tree
32. Validate BST
33. Serialize/Deserialize Tree ⭐
34. Path Sum
35. Balanced Binary Tree
36. Zigzag Traversal
37. Kth Smallest in BST
38. Max Path Sum (Hard)

---

## 🌐 E. GRAPH (10)

39. DFS / BFS ⭐ (ACTUAL asked)
40. Number of Islands ⭐
41. Clone Graph
42. Course Schedule ⭐
43. Detect Cycle in Graph ⭐
44. Topological Sort ⭐
45. Dijkstra Algorithm
46. Shortest Path in Grid
47. Union Find ⭐
48. Word Ladder (Hard)

---

## 📚 F. STACK / QUEUE (6)

49. Valid Parentheses
50. Min Stack
51. Next Greater Element ⭐
52. Largest Rectangle in Histogram
53. Sliding Window Maximum ⭐
54. Implement Queue using Stack

---

## 🔢 G. HEAP / GREEDY (6)

55. Kth Largest Element ⭐
56. Top K Frequent Elements ⭐
57. Merge K Sorted Lists ⭐
58. Task Scheduler
59. Meeting Rooms II
60. Huffman Coding

---

## 🧠 H. DYNAMIC PROGRAMMING (10)

61. Climbing Stairs
62. House Robber ⭐
63. Coin Change ⭐
64. Longest Increasing Subsequence ⭐
65. Edit Distance (Hard)
66. 0/1 Knapsack ⭐
67. Longest Common Subsequence ⭐
68. Partition Equal Subset Sum
69. Word Break ⭐
70. DP on Grid

---

## ⚙️ I. PUBMATIC-SPECIFIC (VERY IMPORTANT)

71. Design LRU Cache ⭐⭐⭐
72. Count number of triangles in array ⭐
73. Memoization function
74. Rate limiter (sliding window) ⭐
75. Concurrent queue (thread-safe) ⭐
76. Implement cache with TTL ⭐
77. Log processing system
78. Find top-K streams (real-time)
79. String grouping / hashing problems
80. Data deduplication system

---

# 🧠 PART 2: EXACT LLD / HLD QUESTIONS (WITH SOLUTIONS)

These are **REAL interview-style problems asked / expected**

---

# ⚙️ 1. DESIGN LRU CACHE (LLD)

### 🔥 Requirements

* O(1) get / put
* Capacity bound
* Eviction policy

---

### ✅ Solution (Go)

```go
type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.cache[key]; ok {
		l.moveToFront(node)
		return node.value
	}
	return -1
}

func (l *LRUCache) Put(key, value int) {
	if node, ok := l.cache[key]; ok {
		node.value = value
		l.moveToFront(node)
		return
	}

	if len(l.cache) == l.capacity {
		lru := l.tail.prev
		delete(l.cache, lru.key)
		l.remove(lru)
	}

	node := &Node{key: key, value: value}
	l.cache[key] = node
	l.add(node)
}
```

👉 Interview follow-ups:

* Thread-safe version?
* Distributed cache?

---

# 🌐 2. DESIGN RATE LIMITER (VERY IMPORTANT)

### 🔥 Asked in ad-tech companies

### Types:

* Token Bucket
* Leaky Bucket
* Sliding Window ⭐ (most asked)

---

### ✅ Sliding Window Approach

```go
type RateLimiter struct {
	requests map[string][]int64
	limit    int
	window   int64
}

func (r *RateLimiter) Allow(user string, now int64) bool {
	if _, ok := r.requests[user]; !ok {
		r.requests[user] = []int64{}
	}

	reqs := r.requests[user]
	cutoff := now - r.window

	// remove old requests
	newReqs := []int64{}
	for _, t := range reqs {
		if t > cutoff {
			newReqs = append(newReqs, t)
		}
	}

	if len(newReqs) >= r.limit {
		return false
	}

	newReqs = append(newReqs, now)
	r.requests[user] = newReqs
	return true
}
```

---

# 🏗️ 3. DESIGN URL SHORTENER (HLD)

### 🔥 Requirements

* Generate short URL
* Redirect fast
* Handle billions of requests

---

### ✅ Design

**Components:**

* API Service
* DB (NoSQL like Cassandra)
* Cache (Redis)
* Load Balancer

**Flow:**

```
User → API → Generate Hash → Store → Return short URL
```

👉 Key Points:

* Base62 encoding
* Cache hot URLs
* TTL support

---

# 📡 4. DESIGN REAL-TIME LOG PROCESSING SYSTEM

### 🔥 Very relevant to PubMatic

---

### ✅ Architecture

* Producers → Kafka → Consumers → DB
* Use:

  * **Kafka**
  * Stream processing (Flink / Spark)
  * Aggregation

---

### Key Questions:

* How to handle millions of logs/sec?
* Fault tolerance?
* Ordering?

---

# ⚡ 5. DESIGN TOP-K FREQUENT SYSTEM (REAL-TIME)

### Example:

“Top ads clicked in last 1 minute”

---

### ✅ Approach

* Min Heap + HashMap
* Streaming + windowing

---

# 🔐 6. DESIGN DISTRIBUTED CACHE

---

### Components:

* Cache nodes (Redis cluster)
* Consistent hashing ⭐
* Replication
* Eviction policies

---

# 🧵 7. DESIGN THREAD-SAFE QUEUE (Go specific)

---

```go
type SafeQueue struct {
	queue []int
	mu    sync.Mutex
}

func (q *SafeQueue) Enqueue(val int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, val)
}
```

👉 Follow-ups:

* Use channels?
* Blocking queue?

---

# 🎯 FINAL STRATEGY (VERY IMPORTANT)

For **PubMatic senior role**, focus:

### 💣 MUST MASTER

* LRU Cache
* Rate Limiter
* Graph (DFS/BFS)
* Heap problems
* System design (real-time systems)




---

# 🚀 PHASED INTERVIEW PREP PLAN (Highly Effective)

We’ll cover everything you listed, but in **layers**:

### ✅ Phase 1: Core Go + Concurrency

* Goroutines
* Channels
* Context
* Memory model
* Interfaces

### ✅ Phase 2: Backend Engineering

* REST, gRPC
* Middleware
* Auth (JWT)
* API design

### ✅ Phase 3: Distributed Systems

* Microservices
* Event-driven architecture
* Kafka/MQTT
* Caching

### ✅ Phase 4: Infra & DevOps

* Docker
* CI/CD
* Nginx/Kong
* Observability (Prometheus, Grafana)

### ✅ Phase 5: Specialized (Your Edge)

* BACnet, Modbus
* Edge computing
* CGO
* File + networking

---

# 🎯 HOW WE’LL DO 500 QUESTIONS (SMART WAY)

Instead of dumping:

* I’ll give **50–80 HIGH-QUALITY questions per phase**
* Each question:

  * 🔥 Interview-level
  * 💡 Concept-driven
  * 🧠 Includes traps & edge cases

---

# 🔥 LET’S START: PHASE 1 (GO CORE + CONCURRENCY)

## 🧠 SECTION 1: Goroutines & Scheduling (15 Questions)

1. What happens internally when you start a goroutine?
2. How does Go scheduler work (M:N model)?
3. What is GOMAXPROCS?
4. Difference between OS thread vs goroutine?
5. Can goroutines run in parallel always?
6. What is work stealing?
7. What happens if goroutine panics?
8. How do you recover from panic inside goroutine?
9. What is goroutine leak? Give real example.
10. How to debug goroutine leak?
11. What happens when main exits?
12. Are goroutines preemptive or cooperative?
13. Stack growth in goroutines?
14. How many goroutines can run in 1GB RAM?
15. Why goroutines are lightweight?

---

## ⚡ SECTION 2: Channels Deep Dive (20 Questions)

16. How channels work internally?
17. Buffered vs unbuffered channels (deep difference)?
18. What happens if you send to closed channel?
19. What happens if you receive from closed channel?
20. How to detect closed channel?
21. Can you close a nil channel?
22. Select statement internals?
23. How select chooses cases?
24. What is channel deadlock?
25. How to debug deadlocks?
26. Can multiple goroutines read/write same channel?
27. When should you close channel?
28. Who should close channel?
29. Fan-in / fan-out pattern?
30. Rate limiting using channels?
31. Worker pool implementation?
32. What happens if channel buffer is full?
33. Nil channel behavior?
34. Directional channels (why)?
35. Channel vs mutex (when to use what)?

---

## 🧩 SECTION 3: Context Package (15 Questions)

36. Why context is needed?
37. What happens if you don’t cancel context?
38. Context tree structure?
39. WithCancel vs WithTimeout vs WithDeadline?
40. How context propagates in microservices?
41. Context vs global variables?
42. Can context store large data?
43. What is context leak?
44. Best practices of context usage?
45. Why context first argument?
46. What happens when parent context cancels?
47. How to debug context cancellation?
48. Context in HTTP server?
49. Context in gRPC?
50. Can you create custom context?

---

## 🧱 SECTION 4: Interfaces & Design (10 Questions)

51. What is interface internally?
52. Nil interface vs nil pointer?
53. Type assertion vs type switch?
54. Empty interface vs generics?
55. Interface composition?
56. Implicit implementation (pros/cons)?
57. Interface pollution problem?
58. When NOT to use interface?
59. How Go achieves polymorphism?
60. Real-world interface design example?

---

## 💣 SECTION 5: Memory + Maps + Internals (15 Questions)

61. How maps work internally?
62. Are maps thread-safe?
63. What happens during map resize?
64. Why map iteration is random?
65. Pointer vs value receiver?
66. Escape analysis?
67. Stack vs heap allocation?
68. Garbage collector working?
69. Stop-the-world?
70. How to reduce GC pressure?
71. Memory leak in Go?
72. Slice internals?
73. Append behavior?
74. Copy vs reference in slices?
75. Struct alignment?

---



# 🚀 PHASE 2: BACKEND ENGINEERING (INTERVIEW MASTER SET)

---

## 🌐 SECTION 1: REST API (20 Questions)

1. What makes an API truly RESTful?
2. Difference between REST and HTTP?
3. Idempotency — explain with PUT vs POST.
4. Why PATCH exists if PUT is there?
5. Safe vs idempotent methods?
6. How do you version REST APIs?
7. URI vs Query Params vs Headers — when to use what?
8. What is HATEOAS? Is it used in real systems?
9. How do you design pagination?
10. Offset vs cursor pagination?
11. How do you handle filtering/sorting?
12. Proper error handling structure?
13. What HTTP status codes do you use for:

* validation error
* auth failure
* server crash

14. How do you handle partial failures?
15. What is rate limiting in REST?
16. How caching works in REST (ETag, Cache-Control)?
17. How do you make REST scalable?
18. How do you secure REST APIs?
19. What is statelessness? Is it always true?
20. REST anti-patterns?

---

## ⚡ SECTION 2: gRPC (15 Questions)

21. What is gRPC and why is it faster than REST?
22. HTTP/1.1 vs HTTP/2 difference?
23. What is Protocol Buffers?
24. How serialization works in protobuf?
25. Unary vs streaming RPC?
26. Types of streaming in gRPC?
27. How does gRPC handle deadlines/timeouts?
28. How does service discovery work in gRPC?
29. gRPC vs REST — when to use what?
30. Can browser directly call gRPC?
31. What is gRPC Gateway?
32. How authentication works in gRPC?
33. How load balancing works in gRPC?
34. How to debug gRPC calls?
35. What are limitations of gRPC?

---

## 🔐 SECTION 3: JWT & AUTH (15 Questions)

36. What is JWT structure?
37. Difference between JWT and session?
38. Stateless auth — pros/cons?
39. What is access token vs refresh token?
40. Where should JWT be stored (localStorage vs cookies)?
41. What is token expiration strategy?
42. What happens if JWT is stolen?
43. How to invalidate JWT?
44. What is signature in JWT?
45. Symmetric vs asymmetric JWT?
46. What is OAuth vs JWT?
47. How SSO works?
48. What is CSRF? How JWT affects it?
49. What is XSS risk with JWT?
50. Best practices for JWT security?

---

## 🧱 SECTION 4: Middleware (10 Questions)

51. What is middleware in Go?
52. How middleware chaining works?
53. Order of middleware — why important?
54. How to write custom middleware in Go?
55. Logging middleware design?
56. Auth middleware design?
57. Panic recovery middleware?
58. Rate limiting middleware?
59. Middleware vs interceptor (gRPC)?
60. Performance impact of middleware?

---

## 🧠 SECTION 5: API DESIGN (10 Questions)

61. How do you design a production-grade API?
62. How to design API for high traffic?
63. How do you handle backward compatibility?
64. How do you design error responses?
65. How to make APIs observable?
66. API gateway role?
67. What is request tracing?
68. How do you handle retries?
69. Circuit breaker pattern?
70. How to design idempotent APIs?

---

# 🎯 NOW — LET’S GO DEEP (CRITICAL CONCEPTS)

---

## ⚔️ REST vs gRPC (Interview Killer Question)

| Feature         | REST     | gRPC              |
| --------------- | -------- | ----------------- |
| Protocol        | HTTP/1.1 | HTTP/2            |
| Format          | JSON     | Protobuf (binary) |
| Speed           | Slower   | Faster            |
| Streaming       | No       | Yes               |
| Browser support | Native   | Needs gateway     |

👉 **Perfect Answer:**

* Use REST for:

  * Public APIs
  * Browser clients
* Use gRPC for:

  * Internal microservices
  * High-performance systems

---

## 🔐 JWT STRUCTURE (MUST KNOW)

```
HEADER.PAYLOAD.SIGNATURE
```

* Header → algorithm
* Payload → user data (claims)
* Signature → integrity

👉 Example:

```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## 🧠 MIDDLEWARE FLOW (GO)

```
Request → Logging → Auth → Rate Limit → Handler → Response
```

👉 Key insight:

* Middleware wraps handler like onion layers

---

# 💣 REAL INTERVIEW TRAPS (VERY IMPORTANT)

### ❌ Trap 1:

“JWT is stateless so logout is not possible”

✅ Correct Answer:

* Logout requires:

  * blacklist
  * short expiry + refresh tokens

---

### ❌ Trap 2:

“PUT and PATCH are same”

✅ Correct:

* PUT = full replace
* PATCH = partial update

---

### ❌ Trap 3:

“gRPC is always better”

✅ Correct:

* Not browser friendly
* Hard debugging
* Needs infra support

---

# 🔥 PRACTICAL CODING (GO MIDDLEWARE)

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Request:", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}
```




# 🚀 PHASE 3: EVENT-DRIVEN SYSTEMS (KAFKA, MQTT, BROKERS)

We’ll cover ~75 **interview-grade questions** across:

* Event-driven architecture
* Kafka
* MQTT
* Brokers & messaging patterns
* Real-world system design

---

# 🧠 SECTION 1: EVENT-DRIVEN ARCHITECTURE (15 Questions)

1. What is event-driven architecture (EDA)?
2. Event vs message — difference?
3. Producer vs consumer vs broker?
4. What is event sourcing?
5. What is CQRS?
6. Push vs pull model?
7. Synchronous vs asynchronous communication?
8. What is eventual consistency?
9. What is message ordering problem?
10. What is idempotent consumer?
11. What is at-most-once, at-least-once, exactly-once?
12. What is dead letter queue (DLQ)?
13. What is backpressure?
14. How do you handle duplicate events?
15. When NOT to use event-driven architecture?

---

# ⚡ SECTION 2: KAFKA DEEP DIVE (25 Questions)

16. What is Apache Kafka?
17. What is a topic?
18. What is partition in Kafka?
19. Why partitions are important?
20. What is offset?
21. What is consumer group?
22. How Kafka ensures scalability?
23. How Kafka ensures fault tolerance?
24. What is replication factor?
25. Leader vs follower partition?
26. ISR (in-sync replicas)?
27. What happens when broker fails?
28. What is rebalancing?
29. What is log compaction?
30. What is retention policy?
31. Kafka vs traditional queue?
32. Pull vs push in Kafka?
33. How ordering is maintained?
34. Exactly-once semantics in Kafka?
35. Kafka producer acks (0,1,all)?
36. What is Kafka lag?
37. How to monitor Kafka?
38. How to debug slow consumers?
39. Kafka vs RabbitMQ?
40. When NOT to use Kafka?

---

# 📡 SECTION 3: MQTT (20 Questions)

41. What is MQTT?
42. Why MQTT is used in IoT?
43. Broker in MQTT?
44. Topic structure in MQTT?
45. QoS levels (0,1,2)?
46. Difference between QoS levels?
47. Retained messages?
48. Last Will and Testament (LWT)?
49. Persistent sessions?
50. Clean session flag?
51. MQTT vs HTTP?
52. MQTT vs Kafka?
53. How MQTT handles unreliable networks?
54. What is keep alive?
55. What is wildcard subscription?
56. Security in MQTT?
57. How scaling works in MQTT?
58. What happens if broker crashes?
59. MQTT over WebSocket?
60. When NOT to use MQTT?

---

# 🧱 SECTION 4: BROKER & MESSAGING PATTERNS (10 Questions)

61. What is message broker?
62. Pub/Sub vs Queue model?
63. Fan-out pattern?
64. Work queue pattern?
65. Request-reply pattern?
66. Event streaming vs messaging?
67. How retries work in messaging?
68. Exponential backoff?
69. Poison messages?
70. Message durability?

---

# 🧠 SECTION 5: REAL SYSTEM DESIGN (10 Questions)

71. Design MQTT → Kafka pipeline
72. How to handle millions of IoT devices?
73. How to ensure no data loss?
74. How to scale consumers?
75. How to design real-time analytics?
76. How to handle spikes?
77. How to design fault-tolerant system?
78. How to handle schema evolution?
79. How to monitor event systems?
80. How to debug distributed system?

---

# 🔥 CORE CONCEPTS (YOU MUST MASTER)

---

## ⚔️ Kafka vs MQTT (INTERVIEW GOLD)

| Feature    | Kafka              | MQTT            |
| ---------- | ------------------ | --------------- |
| Use case   | Big data streaming | IoT messaging   |
| Protocol   | TCP                | Lightweight TCP |
| Throughput | Very high          | Moderate        |
| Latency    | Medium             | Very low        |
| Storage    | Persistent         | Optional        |

👉 **Perfect Answer:**

* Use Kafka:

  * Analytics
  * Event streaming
* Use MQTT:

  * IoT devices
  * Low bandwidth systems

---

## 📡 MQTT FLOW (YOUR PROJECT USE CASE)

![Image](https://images.openai.com/static-rsc-4/nTuhluPHHpVNWS7J6EDANIoX4hNGgdZ_UPYUh5L0V59eUuWitzyjMOxkz9aPw5ewxf-4buuQEoHci6NCEtg1T6Bt2_cq8iCzOf86jg-O9tYA-5tgj0iytgyXPBnrlgorNLE0lor5yAx24j2z2nhkPxrPHylvIg6Rj03OK3-lRvo3QJ9oKkdBQEiRkF84kPDf?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/9q9fnl8BCg13y8MvGwDkrouG7q6Ogv_oKHuu0mFBqancP8zjryLoiBopZMKlGzd2I30a3ZkY4BW0iEAEZs_EcfvsbL7oiYL8h1IgnNZ_ZcUkVBEAoloVNa2N47ip-BxbFy4_mHAuUcjbwReNFIs9ZPU77oGMcd_m9zHUYV7f6cVL_5bxmTHyxg0YYMteGEP8?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/-sdPTwMSc33To82eTA7wGWTgqy2Xu7jgCPHPXx5h_KYAa3L4reo0neAI_NBLBdX5n_hL5PZ_X5NlPm8hkIx9BmXr500w9sb5rsqnH7CkIjzExVDjRjCvTnXVz5c1XFLbOyjGOTF6DH_kHQ2CVIt9NZIak4XoXdRLzADL-5p--MvdTvYXE4e2Y5Jll89yGReW?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/Vj51EYPWAVKAtB6nFJO8RtxXc7BcLdU2cDNkmZsn38aDuKosnJuUBoGaPTkfUWXPfY5s_FhSFL6CODeJrs2CMJXAm-U2sDiNoZSIMIATcZqLkaAGI6FrIaAZuIQGsKTAlazpul6Gc5Y80ykbSY_VDzCAwiX98nslWZvibjxLhqJg6EtR7OcNDYQRDIQ0DXev?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/zTDF6n4q70odVytxksjDyTvRyEH3LfY5KLrQayIYrv7RRVtziMCFxZ1IJWzCSp6Hp5U4GOBkMP5uxkFZ-PyWnGche3SwQ9zq9IrwnK86h85VVGEBE2ltX8SzK5Q292RuHhPKBPSb_6EGO1eRpx5njHbLNHCK4xvzu9M-JK_c52tzYGn8mg1kcNIoCKijuS_L?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/70h5N3iwGkzxLrmXWF7vnG8SPaioeFYhAEgoHPEfBV442VapN_mnbH_gTAt2VZYfqpGGo4A5f3arU1KGJJNlseySlATwyp3IvW_CB4hr59l6fg6G9BYzgJ8bUhkB-rDIcZvsps4E1dZaPeF1PTgijLl_vhz3KFz06mN2kzLWmPjbugi8sOeplmVpk-BdQCmY?purpose=fullsize)

👉 Flow:

* Device → publish → broker
* Broker → distribute → subscribers

---

## ⚡ KAFKA ARCHITECTURE

![Image](https://images.openai.com/static-rsc-4/ypetG0PA9pTHNfXRiEo8117VmTtiSO9vs6KmKZM4SL-3kqmzvbERaIMRLW24tDmAL_MDkZ8EyCl_Ap_o4ucj8DDRTMPZFwlGwm-gh7cMh8BkKjodvh5MzxAFCtF38drQRzwX4i9A8urJ-QADwegJe4BdX3BOBYVRgySXWmXtseFLCYOlMIJjwleFX-up5dO0?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/h0krgBvBfHIcOOlXh8zy_2agQDYTWPI9FBHF_bZ2caui-9ev1mOmzsqLGaqw6UzWc26_m99P4DDEwR4YhQZtPHtIJ2z7njH8iIY8DQaSG0hflQWX5YNP8KUnkQQqESvEPTdiD3ddNONAOm8Nb7zYOayfiLUf67KujbcRu1Y6OvGnfjVfqv5-DmeLb_WcbkJ7?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/1eL042esENRUKO8mRmuBND1ksm78kQ2NyuVv15FBiO4lBXA4hz1cPu7ZqJFQEQYVThabD4Jn1Uo_EFCbjgYF22mVF34pfoAA_UJQLwfVVesYH5Oj2LXylQ5prg1xWwZFba96tapKmefax_wGD_j1QZVyl9fHjE4i1ZYj2hRDMVfgUT3KkXXs3F1ixqimbyO-?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/UGRza_DxXZPY1lIgESNUG4N17SX6_KzfCpsAfUjawF3MdoZsnot938u7zb0_FmKvnOpg79mWEdQ8oUhefw3Ush9wVnuAHo2SNvSKErdJcDHS7L0jlcwBfu77K8UTyY9yS7olZqNmVWfXsRxKSfALb8dcq1rdJLVPkuDk_0d6slDk5zml0lZHLG4dzozPitFk?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/gGRQFPyaO2nPO-AjPutYsuGxP3K3zdhZV2VT41tSGt1UxJP3rdbVeL3cxH3JG_7bL0eFaTDcdVoq514rZkMldq-j-KRso4Hf_XxxxXDgT8qqhCouJOgG_OKubkYnta2K2kxasvVT7NNP2T0rbMms9MyDyO9UEOEkZmdBjANnoNreCqPYisESkDn-HfD0XJ5c?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/RM8A1kMKSWe6AbIk_e3y7q6Qv3GJR0SKYQ4HreSx-yJKsQ4nmOmDDrR80yeMlI7YJgXGzCpr7OcX61V_2xjWMa1bPhSSNI_PEdV472xNzwr_lfJDnZtFRcZIeGlA0lXLmgVyKLqBLfA2ua1UgyOvI46KD7A5o9-MUsfYJM2k4bOaB9vvyvRrft6w4uxeTsdp?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/-Z0PLgzNxKcybPD2DipoVXEGEtBvCHAlB7gCZipq5TcH4aRPFFhr6dpJoAK42QGFxHOirT1-nG4hdacEOWQAV0rEfbuIPw2j2LRD9S_7rrdXQi_90S2XvrHZKgIg3K4Te5otnRX2rE3s9vtLSq8xXmEdFOuG4gvTG3one3GY9OnVzBjQ4YcET2NPcQD-kEkR?purpose=fullsize)

👉 Key Insight:

* Kafka = distributed log (not just queue)

---

# 💣 INTERVIEW TRAPS (VERY IMPORTANT)

---

### ❌ Trap 1:

“Kafka guarantees exactly once by default”

✅ Correct:

* Default = **at least once**
* Exactly-once requires:

  * idempotent producer
  * transactions

---

### ❌ Trap 2:

“MQTT QoS 2 is always best”

✅ Correct:

* Highest reliability but **slowest + costly**

---

### ❌ Trap 3:

“Event-driven = no consistency issues”

✅ Correct:

* Eventual consistency
* Requires compensation logic

---

# 🔥 REAL DESIGN (YOUR PROFILE BASED)

### 🧠 MQTT → Kafka → Processing → DB

```
Devices → MQTT Broker → Go Consumer → Kafka → Workers → InfluxDB
```

👉 Why:

* MQTT → ingestion
* Kafka → buffering + scaling
* Go → processing
* DB → storage

---

# 🧠 PRACTICAL GO PATTERN (MQTT CONSUMER → CHANNEL)

```go
messages := make(chan string)

go func() {
    for {
        msg := mqttClient.Receive()
        messages <- msg
    }
}()

for msg := range messages {
    go process(msg)
}
```

---

# ⚠️ WHY YOU FEEL NOT CONFIDENT HERE

Because:

* You use MQTT/Kafka as tools
* But interviews test:

  * **why**
  * **trade-offs**
  * **failure scenarios**

---


# 🚀 PHASE 4: INFRA & DEVOPS (KONG, NGINX, DOCKER, CI/CD)

We’ll cover ~80 **interview-grade questions** across:

* Nginx (reverse proxy, load balancing)
* Kong (API gateway)
* Docker (deep internals)
* CI/CD pipelines
* Observability (Prometheus, Grafana)

---

# 🌐 SECTION 1: NGINX (15 Questions)

1. What is NGINX?
2. Nginx vs Apache?
3. What is reverse proxy?
4. How Nginx handles concurrency?
5. Event-driven architecture in Nginx?
6. Load balancing methods (round robin, least conn, IP hash)?
7. How SSL termination works in Nginx?
8. What is upstream block?
9. Static vs dynamic content handling?
10. What is caching in Nginx?
11. Rate limiting in Nginx?
12. How to debug Nginx issues?
13. Nginx vs API Gateway?
14. How Nginx handles high traffic?
15. When NOT to use Nginx?

---

## 🔍 NGINX ARCHITECTURE

![Image](https://images.openai.com/static-rsc-4/OrVLDUfTdLSWJm1jLmD-BMKH56WLwjByvsNOq-F0AZAGnLWqmSyCFdlvEyqhGGIWAw49burnevpliDorrXaMSUyIqIMi3UBldA1ngcCvp_ObkrJMYJbC4WIh0t-26Ryg9oNW3ymOtIYZ-e42VDd-5otCUmrqnjqOMx_adNuOeiu9s8GveIxABwUsWyP26cNB?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/1yLtospTBwWmqWSVcW5EG2en8gHcryK4MextMEENVl5gflZNSTb_ZjncbTmuuLrnKMMBI2fExMBqc9PVBpHhZFy2IRtIT2nc0taG1m5jUTug36CLhwzbvMe7SgDDaCowLwviDF7-UP83kxIPctd8S-98NaDCMKNlznFyAZXIwZFwq1-salT0DRpjolZq_Kov?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/gOBc72RSZb0CdM0iQgWx45bjvSF2_8GjSuZ64GNtWopegnq6B4BHwziRH8zvSzAJ1N1yBjEWtnrW_1xXpgvcdtxUXZ69T6HxoMsgVT9R8KcMfONSrfqPfMCXPB0hCPZU0qMxPn_mlLqvFvv8wHLvezKQMWPoaBDWBAniT_ZWPU5rxoScsywlN9mRsCNcG8FL?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/OJjVUawMVDoY7MQWnYa54sOOi7N3qLVcTO8nwRDARFjXWmbK3fOJLH37qz1rvuMubNGu1_PaDBThD4UTx1xs0stY9e7n15xkTLqRc4PUWVcoLhf8fXYgugT5quYT-pw8Sz0TtEyiGGmqeV07Mx038aCApXGswb2k9z3MnQYLe411A3P-Gm33yarUSI5jDwVd?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/MXSuXy3XgkvC-jv3dvWnYStqVgdZpg9gYTre2RgF4PuCBIRV7p9ZnmhoqISuK_PrZhM1KPrIXYPPtE1vAVRAecKliD3o-oPNvxmPYawY69ClC4GMkydH5fLz6QtZS4gL7Arb63XudJfD57Oe1CXv3FwEjdhxN6bbowNsOc91kp8WUQ2pCJLcnSv9va-vbRCx?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/DLKlJRBf0kULFO8e9Z5FbzYHXSyaDY7RtVTD-hwOPeuIE0bO-ct6S2BiRaIKshS4aiCc5S60Ejo3Xo-nyAneNufKrco2vPhAIsyOofscF6lnQexV6YsdCLAZecPSDq1tHJUr0yJdMILkTBr-HYfvTVco7gdNi9Ye_SiMLHHabAPcLOc4qcMKv667oLYWkgLB?purpose=fullsize)

👉 Key insight:

* Master process + worker processes
* Event loop → handles thousands of connections efficiently

---

# 🚪 SECTION 2: KONG API GATEWAY (15 Questions)

1. What is Kong?
2. Kong vs Nginx?
3. What is API Gateway?
4. Why use API Gateway in microservices?
5. Plugins in Kong?
6. Authentication in Kong?
7. Rate limiting in Kong?
8. Logging & monitoring in Kong?
9. How Kong scales?
10. Kong architecture?
11. Kong vs AWS API Gateway?
12. When NOT to use Kong?
13. How routing works in Kong?
14. Service vs Route in Kong?
15. Kong in Kubernetes?

---

## ⚡ KONG FLOW

![Image](https://images.openai.com/static-rsc-4/-jJsRduxhCVgj6mfTokZmB0Ob0cTjgeZVE6L8iUu10_ZKXaRr0PufEDxDqlejsx7XUhnuYEZOFuWw5_3LLndjswlvKE0vZ6zSxvg0v7WKHz1UyAo-3t_IGDy3S-TAzSy8zluY6mixLA6o8kJAcbRdp13WNVHrbwwPvchDou1g3qzi88D5UBbErj_1kenJN69?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/O570jSkM2VkUBZt5qyiNdLPCIzuu8Rl7A_v4LfZ1XnkNrt363VI5K-DPvrWL5GI2XyeJlV45Q4YCgsc4p_3iTEGCKadXih14c2QAuHSyMyJJER5eB1k8dljvGBdww9Or44r5emsLLOURgqVT5jmdYLGHafrHdjIf4YhhL7XVLeQnsVtkVF50bbjABGM_T3vh?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/hUDAmxSwfLeJQ2LXhplWbYLvbaQOw06EaLbuH3tI2CtoANGsAV7SudRKdfQCINBMDZosrn2Smzo28LdniNSCKqdWd0jBRmxngULzRsz3SNRRZ9NHLH6Xqt7DW6R7U-PPswXYIQz4rVvCstinFirhH2PcbinTRp6JkQY48P7eS-HuhLKdFvDXZyrMapr3EKQr?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/kFMejozxSCy98DIRIwUu9Vbpi6VSTNhGuiG4pG25yh3E78LJoV7wUEs7DcTGnAPOMSv8La4LHep8oL6T0tX-05FI1LzJTTdgRDbmkTfYSABNbyogvXu98E1pH1RjResAufsBnQqRK-3AyY27GuRrmyW_Hd__D7TMMig3aGL7ccZYbD-ToEYnsq-HgavPojG6?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/vwQQ5YAQOxynPVh6nwDpcm4q4TPA-2q_8svmG_Gec-vtwZvCWSmt55T4ulcKGVr0j6y4cwm6KaujPREF_P8r2etn5Yktg2azRJ2euVN0Yda5aZRakxfsdR99gxp38pG2SPUArCdIkupoKgxCOHCyQydmX-ZHfiuM0G7ffU12gNGF_3mJ4Bk8TJvc1Obc5wwK?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/XVaT8YI9ShGK9tlqCytFR-pQ1yjgWUnXlvJxLgvDTMAo1T3Ma0ecSd3vCznX3cxqv7jfFrSxVhbw_-VzA-3bSVHYMl_Z5wZXzjZTwxO4f8Igp_wR2zTFbtMoOLsbJ2Sqlb1KGO4bYKZe_IwSKCZ4jnJtbhknsERiD6jxiynnxsR_HKFCnhBtwfKrS54Wk4pD?purpose=fullsize)

👉 Flow:
Client → Kong → Plugins (auth, rate limit) → Service

---

# 🐳 SECTION 3: DOCKER (25 Questions)

1. What is Docker?
2. Container vs VM?
3. How Docker works internally?
4. What is Docker image?
5. What is container?
6. What is Dockerfile?
7. Layers in Docker image?
8. What is Docker cache?
9. CMD vs ENTRYPOINT?
10. COPY vs ADD?
11. What is volume?
12. Bind mount vs volume?
13. What is Docker network?
14. Bridge vs host network?
15. What is Docker Compose?
16. How to optimize Docker image size?
17. Multi-stage builds?
18. How to debug container?
19. What happens when container crashes?
20. Restart policies?
21. Docker security concerns?
22. What is cgroups and namespaces?
23. How Docker isolates processes?
24. Logging in Docker?
25. Docker vs Kubernetes?

---

## 🐳 DOCKER ARCHITECTURE

![Image](https://images.openai.com/static-rsc-4/ioJLD1oDVTn2d0UY-krkybter1NDacdEXCqVyodTZZdDWxdVRBgT-msj7fIjjhzCrLSKmsZ3VMq6tL34itBeITwdxs-i-qUUajt6zlHFKx4Fssbws3-sXaiRP43nhECTS2h6CsrduyVPTNjIEGRUpBrERqRgUyIsIIZ1XUdFZ80JgLnm0BED70uif4hD5ZK5?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/LTJMN_jnPazGs289XU78RqCOEpQJOIT5S-wbrV7LzITlPFCnB8OezD0r44IRrtE5zWNWeFYpvN12V5zZmQ0giFNeH6bmm4eSj2ZZera4XynzTFqDy9DLeuMmgSXfJu6Z6YrNK1NefCGgFlJZO_nJAXZOOMgHaweq5fr4NKTcn7HwndvY2tS1t15HnIzoLynn?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/Upl0_pHqTUQiq6KJcp7tJJ9DUY2s9GRpVdVbbgho8qM25UF0roHJzWkHp-Olu2LoxdfTi2CudIrBP6fVb0vbd04N1dBxFmdfzUynGG0xH_rpLg3DVkRokj1Uwa6f8wxNQEt0mkiAvRLxhqA_DgSNgETnJaLUeKB8Qt40Xwj8VqhP9h7-U6cgQ5f_dIdAO3RS?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/dwhKbm1F1z62g0-pPorMe0vBEDwuZUM3IM8GQLSyBvwQ78bO0n1QyFOyvd7Rb5rraarSGZ-XWDrVELqA-vCE5XaPy48_Z-rzKUVR6aWt574aB8KWIK7rxwfTt9czsRvElClCt33vXiXFG7aanJQZcKjQQBkyvKwjwpltWwo8lK1GS044BFAJ1JLLIp9q53oa?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/ha-2afv1PIWwGetOCMuLWWHFZiFrRtnBXlMQ3QcZRWnVADtjPIH0ZlgOq3rMkfOGVtp3s3-FrZ89yfk5DhKhk9gn1qt9eMIuvfgAtgRUK5UOZnRNA8E5NH-SM4h7XQARnOolY-vq6Eh-6P4X_a_PZGxlyJhNsTz1j9MIANvzD9Rjhp4-x47Xj_1kjjZo-bYV?purpose=fullsize)

👉 Key Insight:

* Docker uses:

  * Namespaces → isolation
  * Cgroups → resource control

---

# 🔁 SECTION 4: CI/CD (15 Questions)

1. What is CI/CD?
2. Difference between CI and CD?
3. What is pipeline?
4. Stages in pipeline?
5. What is build artifact?
6. What is blue-green deployment?
7. What is canary deployment?
8. What is rollback strategy?
9. What is GitOps?
10. How to secure pipeline?
11. What is parallel execution in CI?
12. How to handle secrets?
13. How to test in CI?
14. How to deploy Docker in CI/CD?
15. Common CI/CD failures?

---

# 📊 SECTION 5: OBSERVABILITY (10 Questions)

1. What is observability?
2. Logging vs metrics vs tracing?
3. What is Prometheus?
4. What is Grafana?
5. How Prometheus collects metrics?
6. What is pull vs push model?
7. What is time-series DB?
8. What is alerting?
9. What is distributed tracing?
10. How to debug production issue?

---

## 📊 OBSERVABILITY FLOW

![Image](https://images.openai.com/static-rsc-4/3a0ZQ1j7uxrbCtD8eaCzNYunNc5zYEVDhp450-bfgcdiT_ZVDvYUeWj9BpKH-MIcX4Tl6HjZ8XuCRJuKG5ccEEuDDFfHDBq6xeVz-otXAQtW66TH3SXz1fsJlbgfGwT-MnrVRLlct9INENXn417NCNHp_yzep7CImT0CvS8c0IVQh9usfnajxHCGXEznDqvF?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/XYs36MBBDoUcxaPLiUfOcrZVK2cPVaxeQz912-ydCbJjv-7dd6YbprjJK00059J7iBNgV_8DjEAB_DmdrQ4-a8m5Aw8aMIu7JP9DXqw7kAodiu-5oH_ud_mkoeU59rDbtHmlfZXbld3l2td1fxYa6NjrRw_nGX489vDsru6bSVss9Zgo9P66IjSwPUJ_yBdg?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/mwENCzHqYk9XdbHJ2ffYMJR2ZrqNP_IwWyRSh-o3lEgGioo6SeKvADs9iIYO5U72WUyGAP2nckieAVARGB1sGz2g077KysIFajEN-B8hAUpxxHb1XjwTXy5txagBd4jiRvLJnwXwWyRf34U2jDE6TwIRHbFsNUfNiUWfft6x0zKoinhzSeaUxlrnOAorP8Mn?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/ks4VD7JN3i_4M7HMp62UDHAkptAnGIJ3VBCC8BR8fXBmD0v_NVZGdwKiJFWUef67Jt-S1USoFSWt4klI6f06H5109bB60ajtw6E0zl-9v-rMf0F-yiCzl_9LpdNxqNmNuJT6ifTcJAous6iGBDZj5jMzUqtzQwf6gHSVzfZXnlml8h_SxU02mzzVA-iIkH_S?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/MoWR_aafEKrhFdOT7CEHZ7cVzAi4nSog1NVLfEd3suyClHrWE2-ZpmI2ZtQq2RXVHh-y5jOgyxUoYB1amKOsLimkON0OdpzkDFZwFJgPFMuUmeDIMlVKq6ZajP9yKPQ-tPrOw_iJJw7fVPcVyDRAN2xoZJ3bhS0OVcF9Oui6vVemuyT1FKsrsmovqM0dX0pM?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/utVbt289isboCbQ9IeIzrJbX3guPFwVnTvNtXjP4B_oXdB5qz-kgF1E1wQzO4OT_qJeWLb9Eh23S_9VJuW9eoaUa5bITmPNxpSAML5iiPLFIQiK6CDyKWJUlXvhS1vxHhzJouUB8f8g5RxsHgqornlArbwL4f1ECC2-cZCxpCrcdfid0D7Rlk4G1jrnwS8sF?purpose=fullsize)

---

# 🔥 CORE INTERVIEW INSIGHTS

---

## ⚔️ NGINX vs KONG (VERY IMPORTANT)

| Feature  | Nginx           | Kong                      |
| -------- | --------------- | ------------------------- |
| Type     | Reverse Proxy   | API Gateway               |
| Logic    | Basic           | Advanced                  |
| Plugins  | Limited         | Rich ecosystem            |
| Use case | Routing, static | Auth, rate limit, gateway |

👉 **Perfect Answer:**

* Use Nginx → performance, proxy
* Use Kong → API management layer

---

## ⚔️ DOCKER vs VM

| Feature   | Docker        | VM      |
| --------- | ------------- | ------- |
| Boot time | Seconds       | Minutes |
| Size      | Lightweight   | Heavy   |
| Isolation | Process-level | Full OS |

---

# 💣 INTERVIEW TRAPS

---

### ❌ Trap 1:

“Docker provides full isolation like VM”

✅ Correct:

* It’s **process-level isolation**, not full OS

---

### ❌ Trap 2:

“Kong replaces Nginx”

✅ Correct:

* Kong is built **on top of Nginx**

---

### ❌ Trap 3:

“CI/CD = just automation”

✅ Correct:

* It’s about:

  * reliability
  * repeatability
  * fast feedback

---

# 🔥 REAL-WORLD DESIGN (YOUR PROFILE)

```id="yg9h0x"
Client → Nginx → Kong → Go Services → Kafka/MQTT → DB
```

👉 Why:

* Nginx → load balancing
* Kong → auth, rate limiting
* Services → logic
* Kafka/MQTT → async processing


# 🚀 PHASE 5: SPECIALIZED / EDGE SYSTEMS (BACnet, Modbus, CGO, Edge)

We’ll cover ~70 **deep, interview-level questions** across:

* Industrial protocols (BACnet, Modbus)
* Edge computing
* CGO (Go ↔ C)
* Networking + file ops
* Real-world system design

---

# 🏭 SECTION 1: BACnet (15 Questions)

1. What is BACnet?
2. Where is BACnet used?
3. BACnet vs Modbus?
4. What is a BACnet object?
5. What is object identifier?
6. What is property in BACnet?
7. What is device discovery?
8. What is Who-Is / I-Am?
9. What is ReadProperty / WriteProperty?
10. What is BACnet/IP?
11. What is segmentation in BACnet?
12. What is polling vs subscription (COV)?
13. What is priority array?
14. How scheduling works in BACnet?
15. Common challenges in BACnet systems?

---

## 🏭 BACnet ARCHITECTURE

![Image](https://images.openai.com/static-rsc-4/y9V8zxGD2PcjGViPjfl9Y0HIacJVV3uvL1Yxaq7twVC-vRIPda84slxFfkD9Wwa1sDaEPqRvTG5EdMS7cTzvbtQr5T5QhgelP5Lzs51NrNRmECHLL66GcpMzWCNQ-_nAb2J_-Vsm9dFhU-KCRFurTCJmUgqsErUdbVBlwdzXaz17lPUJPrx0VewQzBw15AhP?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/wnESdNHy_Y-FwEN0KU99SFOm7YaLoUgTAxZ7eoSvb-fzRqYsvR3cpZEY06irhJ56_TWJBMfYlDrbJz6UemmMBeYyGGGlmM_k8-zQ9pi7I9zC3enwK6bZy6WfsA2FDQCNfrA3qQJLyxZ-ly0s0HEj6pJHSJ9_sJizzIJKFgnYqgJMidJAblyiU8mTbL9I93eL?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/MjLfuSjoZCv-DybwpqBLsOSdc2oOCpN-EJR6vzpS5nj20LCyiAaxI-DicnQR252sKuMvJdUih9WSG6V_G5NUJvUyyQf_YQGVrEOEYlYaTTxFU9TWFeHJcY3ZsxFvINXJW4jLeO8t0REgISadzwd4j4D6ZSc5OnP-Zvw2Jat66Ik5lHxzqc4NSG7vGZZzv2MS?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/_OQJDR1HlSl4F0AXwl0LPCuuRARgpqvx1mJiP--G4J75SsrdT5f52K0PR2HNuLD84N8-I_i98FHkiIKs4G-fVXlig4fMht9G1ZSBLq_1AntJDRShiZsNJfAehkzKNEElq3mcBIzeMkIZmvg__5qAoIJuBtMgshFUrrCaKJ0jcX9JdgQoN6yiVzF2qyjuYjEs?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/4n5NXWOsEA8LOX7e64HnIzW0DmqvnZwLJGUQwOUoSb9vSi1nr2XtNmAK5KBXp59m78LQ2YZNeaOnpZn3AAQuDnHsCyT8swsawWYuBJQtMFd7nVdOi-rT4HAP1KyO-RJM_MfwOQ-Gne6GG7iD2U-qufkcT_eHJqQSroX2u356_64uWHGSW4LQ3qvqu4BX6NRI?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/cgDlX3kh99h-knnUm79KzFKlktyqyAhMiLpTXazHambjDdG8zgvGY2t25ZNhMucWF_qPFwwgD9FT9V5U1dwFnDMcLAU-QFcePZFgrfumGzuhSfXVyLdQCbPtQN8oLvl2xPk_WoSMqgeSWf4z4kLdm5hSqUSVTN_KunBGTxVh26R0rzSIAqg8vYhLpyQfrbfx?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/EtSkWCq02sQl2gXovolslqrqcIlHCxq-s_lGMcprb-VRNT3zrgFL8q9lTkdB_zLBQfL42_4N8ydJX-xRJUcCW1lWkYDmuUSqPvAoqoeKokx6krYrjlx164Os5wxjOKrueWOWIZK9giokNSg4xH2iq8_09lK1w0PAzns0lMoJdcrS0SGRN_TWcFenQOTJzv_h?purpose=fullsize)

👉 Key Insight:

* BACnet = object-oriented protocol (devices expose objects + properties)

---

# ⚡ SECTION 2: MODBUS (10 Questions)

1. What is Modbus?
2. Modbus RTU vs TCP?
3. Master-slave architecture?
4. What are coils and registers?
5. Function codes in Modbus?
6. Modbus vs BACnet?
7. Polling in Modbus?
8. Limitations of Modbus?
9. Error handling in Modbus?
10. When to use Modbus?

---

## 🔌 MODBUS FLOW

![Image](https://images.openai.com/static-rsc-4/n2wEdHzXJ4Uu9jxTa-VKJixHwXt52FLudX8IBjBpj5LprpPTbIENUvc7dDeg_cWxgCKNeZAiLZLiuIc26que6CbVBIRb0mg3546KeKYA6Yaq8n22-chXAhTsn81eCsFszqtkZ_nJ9D_4xuoGEx7TT-ZzmQBCWzzY8PE9gecycbl4-WbnCUaMA8YTbbWrDGI-?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/t0UIgjTBG6uo1qe8Wt0nOSLOgthk8YuKwpBJcWelJz2UaGqf9CtHUo4uDP2oeYS0mjcj2GrPHg4OHfTQtEazuxkmIyvDdZV4y2hoKlkyNfPSUjd2qdFoh360Z8DFxphEugZJYGMt0Ob-6FbwyiTk-hSNqMS5C7B8mQaU2qDB9R5f7YwTUfTlJEou1BrWBtS6?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/H42hEHHsu82ubMtZlY_NbG9X6DzEUk6rMQbJs8SqVWdtiOMnFam1nqNVWFcpHO2p4XOtvaFI63y594VEhd0gWSz4IeWoNvjYTY9CsnSyQxLEiUXGnoSywI7qVPgF79HR4zUTQ7gE7Gt0MnCO8v18jyHCw8ECTjHj0j5uhcFkuVkws_ENXGJFqEcKs7M7x5Ii?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/QuXUjZsbNFhiuu_tWyf-2znKzglR-n4OEdKffForxxGkovnCJ3hnrg_UWslNF-jCQCjO46y-3Ll2x7Axj1ikVi_Ez0Mr_xE5oMv-2pcnxt5EBg1396ctWH4X1o2j_Soq3Bl8S2w30bHW5g2ShKJva1ufxT3HVcgTvsb9T1JNsN1RJQtX4x0fInveM8fEpe2G?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/KHJDQaG4aTdPClXFlJYTUNDnbVwOBdrj45hVFeuRjSLmZHEN1f8Cb5QgR99InKEY9f7lbxgFh_jyDFbtdF7N5rOgLsPKwjfbID7O_dc7VmE2nGi6_A2tZyqOsd_IVfM9GVeYD8MroAoJKegZo8v-KMfEWv3Jo6FCri8GpRdzywKmXld5JxXbRm9GmkCmzD6V?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/eEJ0Ct4N2YsJ8RZf-0k7fqFha1KqUrNUtzZ7eXFpmtwI9EhqyiXaBKW_4zDM5KPv-vshvIYVPlUIDzdnSWaJhpKRB8tLypef1HFArjl_riNgBNNmYWez3eScE-u-cpv0c9TF1oXxw9q__CeCLA8wVP3tUT0-5mcsf2-YKnKzx5aAvRa4kMsbUoJCuVqwGlx4?purpose=fullsize)

---

# 🌐 SECTION 3: EDGE COMPUTING (15 Questions)

1. What is edge computing?
2. Edge vs cloud computing?
3. Why edge is needed in IoT?
4. Latency vs bandwidth tradeoff?
5. Offline-first systems?
6. Data aggregation at edge?
7. Edge → cloud sync challenges?
8. How to handle intermittent connectivity?
9. Security at edge?
10. Edge device constraints?
11. How to deploy apps at edge?
12. How MQTT fits into edge?
13. How to design edge pipeline?
14. What is fog computing?
15. When NOT to use edge?

---

## 🌐 EDGE COMPUTING FLOW

![Image](https://images.openai.com/static-rsc-4/h13RtJZmrtAEA8oXhVSPo0uJbObmbufLl7AEYyCeAXrhDl-aBRTlByKgfZt_zqDCn8aGHqxhV0UyJ8JEecTE0V5fTbs8p39fhEfkCCikhu8ojxlfmgmfZscwIuZv7EQu2mfnvbhFpuTfnE3_wPAlVb_gSCDBEwu3_WFU9yc3YLiu4Qwtmfpn6vvqh1TwHXQo?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/DPMzZaBirmHMf5B5X-smcxFC93LUBi2-mLTBMujxZJtkFNHq5suwOLaMJN66BboVqRHW2if8faiB3bYG_9hT9qVM43ApAhAc22kjr4k_mta8CnifkxJunkIGFYOOT5kcjPu_hVO0DUrjUwXOSBedO-Nzf-XRtpMeNNLuiyTmejr2GTKS9AS6dDyBPayfD5Ct?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/ivQbT3fwpMG4REDEjHIyCGfMFW83gyxz9zb8NWtlTJU1zWyT-MJAqyLVWlsoHRlQPMtNWDQhamV0Wwj7_kNy1GJQWR4W3s2LwTZer9_ea3QYlCLkNYZM7-LYMVYzthBqcWezhN7dxJEybqWF1V4EG4YKdVKH20McrHsUEs3ZUu0Csg_QVvmt-O688z6gbDi9?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/13t4WNn3oKoQywtaT6xVZn6D1HXIMEWLZ37GtC66Kewqnx0KZZ_FeStFvZVzrGXfmHmb4gP4rPGz2FMQMpfGy1vTfj0AZzbMcMYBEFlR5Gpz481Q7m7loUOAd_giAn0MkcxGO72g1g6XaUSQyXzZexhcC8BwFXl1rv9Aj7rNPPG3NwApGmKDMNV115J2tqUO?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/wwV1eOK4SuoupyITgzfemg5gIftBJOPt5VI71GkORnqH54tkou562y183Uav0MBirATtru12pyy08tbD9mhpwovdNV0OkCBO--m3F-Ond07GxktiqTDNvpDKYZ7JmEuj2kX9fJd6JgyvK3Bg0BbDQayaXd0l1hT43oEZJtWnUobn996cC0xxUypm7afuq-6h?purpose=fullsize)

![Image](https://images.openai.com/static-rsc-4/eUucTjw-xIqfNDxpXrKv1q0z-QjUBEkPNL3EsKpexwWK4Zj_cLzW8ef5LX1BZR4BGjSwofUZuPKiSZhuudrCWsbL7S2wdR94JJ1yGvEGy-mDtbQQWeSPYB-dnoXuq5dvExH1m9SNoEcKbgFm7RT9Oj6pedJBS0K2PKnhyEFppd1Xc2DKQS3YgDDkI2FOneGe?purpose=fullsize)

---

# 🧬 SECTION 4: CGO (15 Questions)

1. What is CGO?
2. Why use CGO in Go?
3. How Go calls C code?
4. Memory management between Go and C?
5. What is pointer passing rule?
6. What is C.CString?
7. What is unsafe package?
8. CGO performance impact?
9. When NOT to use CGO?
10. How to debug CGO?
11. Threading issues in CGO?
12. Cross-compilation challenges?
13. How to avoid memory leaks?
14. CGO vs pure Go?
15. Security concerns in CGO?

---

# 🌐 SECTION 5: NETWORKING + FILE OPS (15 Questions)

1. TCP vs UDP?
2. How sockets work?
3. What is connection pooling?
4. How to handle network retries?
5. Exponential backoff?
6. File I/O in Go?
7. Buffered vs unbuffered file reading?
8. Large file processing strategy?
9. How to handle file corruption?
10. Concurrent file access?
11. Streaming vs batch processing?
12. How to handle network failure?
13. How to debug network issues?
14. What is packet loss?
15. How to design reliable data transfer?

---

# 🔥 CORE INTERVIEW INSIGHTS (VERY IMPORTANT)

---

## ⚔️ BACnet vs Modbus

| Feature    | BACnet              | Modbus             |
| ---------- | ------------------- | ------------------ |
| Model      | Object-based        | Register-based     |
| Complexity | High                | Simple             |
| Use case   | Building automation | Industrial devices |

👉 **Perfect Answer:**

* BACnet → smart building systems
* Modbus → simple industrial communication

---

## ⚔️ EDGE vs CLOUD

| Feature      | Edge          | Cloud          |
| ------------ | ------------- | -------------- |
| Latency      | Low           | Higher         |
| Processing   | Local         | Centralized    |
| Connectivity | Works offline | Needs internet |

---

## ⚔️ CGO TRADEOFF

👉 Use CGO when:

* Need C libraries (BACnet stack etc.)

❌ Avoid when:

* Performance critical Go system
* Portability matters

---

# 💣 INTERVIEW TRAPS

---

### ❌ Trap 1:

“BACnet is just another protocol like HTTP”

✅ Correct:

* It’s **domain-specific (building automation)**
* Object-based model

---

### ❌ Trap 2:

“CGO is always fine to use”

✅ Correct:

* Adds:

  * complexity
  * memory risk
  * cross-compilation issues

---

### ❌ Trap 3:

“Edge computing replaces cloud”

✅ Correct:

* It **complements**, not replaces

---

# 🔥 REAL SYSTEM DESIGN (YOUR PROFILE — VERY IMPORTANT)

```id="h5xk2v"
Devices (BACnet/Modbus)
        ↓
Edge Gateway (Go + CGO)
        ↓
MQTT Broker
        ↓
Kafka
        ↓
Processing Services
        ↓
InfluxDB + Cloud
```

👉 Why this is powerful:

* Edge handles real-time + offline
* MQTT = lightweight ingestion
* Kafka = scalability
* Cloud = analytics

---

