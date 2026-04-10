🔹 1. Relational Databases (SQL)  
👉 What they are  
Data stored in tables (rows + columns)
Fixed schema
Use SQL (Structured Query Language) 
👉 Examples  
MySQL
PostgreSQL
Oracle Database
👉 Features
Strong ACID properties (Atomicity, Consistency, Isolation, Durability)
Relationships via foreign keys
Best for structured data
👉 Use cases  
Banking systems
ERP systems
E-commerce transactions
👉 Example table
Users
| id | name | email |
🔹 2. NoSQL Databases

These are non-relational, flexible schema, and designed for scale.

🔸 2.1 Document Databases  
👉 Examples
MongoDB
CouchDB
👉 Structure
JSON-like documents
{
  "id": 1,
  "name": "Satya",
  "skills": ["Go", "Docker"]
}
👉 Pros
Flexible schema
Easy to evolve data
👉 Use cases
APIs
User profiles
CMS systems
🔸 2.2 Key-Value Stores
👉 Examples
Redis
Amazon DynamoDB
👉 Structure
key → value
"user:1" → "Satya"
👉 Pros
Extremely fast
Simple
👉 Use cases
Caching
Sessions
Real-time counters
🔸 2.3 Column-Family Databases
👉 Examples
Apache Cassandra
HBase
👉 Structure
Data stored by columns instead of rows
👉 Pros
High scalability
Good for large distributed systems
👉 Use cases
Time-series data
Analytics
Logs
🔸 2.4 Graph Databases
👉 Examples
Neo4j
Amazon Neptune
👉 Structure
Nodes + edges (relationships)
(User) --FRIEND--> (User)
👉 Pros
Efficient for relationship-heavy queries
👉 Use cases
Social networks
Recommendation systems
Fraud detection
🔹 3. NewSQL Databases
👉 What they are
Combine SQL + scalability of NoSQL
👉 Examples
Google Spanner
CockroachDB
👉 Pros
Horizontal scaling + ACID
Distributed systems
👉 Use cases
Large-scale financial systems
Global applications
🔹 4. Time-Series Databases
👉 Examples
InfluxDB
Prometheus
👉 Structure
Data indexed by time
timestamp | cpu_usage
👉 Use cases
Monitoring (CPU, memory)
IoT data
Metrics
🔹 5. In-Memory Databases
👉 Examples
Redis
Memcached
👉 Pros
Ultra-fast (RAM-based)
👉 Use cases
Caching layer
Real-time systems
🔥 Key Differences (Important for Interviews)
Feature	SQL DB	NoSQL DB
Schema	Fixed	Flexible
Scaling	Vertical	Horizontal
Transactions	Strong (ACID)	Eventual (mostly)
Query	SQL	Custom APIs
Use case	Structured data	Big / unstructured data
🔥 When to Use What (Real-world thinking)
✅ Use SQL when:
Data is structured
Need transactions (banking)
Relationships matter

👉 Example: Order system

✅ Use NoSQL when:
Massive scale
Flexible schema
High performance

👉 Example: Logs, chat systems

✅ Use both (Very common 🚀)

Modern systems use polyglot persistence

Example:

PostgreSQL → transactions
Redis → caching
MongoDB → flexible data
🔥 Simple Analogy
SQL → Excel sheet with strict columns
NoSQL → JSON files (flexible)
Graph DB → Network diagram
Key-value → Dictionary