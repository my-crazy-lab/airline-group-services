Imagine: 
- I'm a founder of the airlines
- I have 5 planes (P1, ...P5), 1 airport (X) with 3 ports (X1, X2, X3)
- My planes can go to A, B, C (round trip)

### Actor
- Passenger
- Crew
- Admin: CS, Inventory management

### Airport Management Service (static)
- Manage planes, ports, airports
  - How can the system know how many planes available? Must know how many planes in system
  - Monitoring(Flight management) have multi ports, we must to know what is correct port of our plane? Must be managed ports
  - Must have multi airports in multi countries, so need to manage them

### Flight Booking Management Service
- Booking processes (8)
  - choose depart time + destination -> choose number of passenger (adult, baby, child) -> chose type of ticket (vip / normal) -> choose type of trip (normal / round trip) -> fill information -> choose seat -> choose luggage options -> pay (payment service) 

### Reservation service
- Why split it into another service?
  - This feature need realtime, and when have multi booking at once time, this is down server (common with booking).
  - Because user will stop at this step and choose, thinking very long, and realtime API always call.
- Workflow
  - View all seats in plane, display available seat empty.
  - Manage type of seat (type of ticket) like normal, vip
  
### Payment service
- Multi type of payment store at here
- For bookings
  - 

### Check-in Service
- Passenger check-in
  - Send code (8 characters) into phone number of passenger (save from online booking)

### Baggage Management Service
- Tracks the status and location of passengers' baggage. (5)
- Tracks baggage handling (just care basic baggage)
  - P send x kg baggage -> pay -> create record baggage in db -> save : owner, baggage information, plane's id (when done, the baggage management server will return full baggage of plane)
- Stored in Elasticsearch using Logstash for indexing and analysis.
- Logged and tracked using ELK Stack for centralized logging and analysis.

### Flight Management Service
- Create a flight from available planes, add crew, setup port to fly
- Setting flight: price(with each distination, with type of people, type of ticket, type of trip), number of seat, flying time
- Provides real-time information on flight statuses (3)
- Handles delays, cancellations, and diversions (4)
- Kafka for real-time monitoring and updates
- Manages flight schedules, availability and routes (choose when booking online)

### Crew Management Service
- Manages the scheduling and assignment of flight crews (1)
- Tracks crew availability (2)

### Inventory Management Service
- Manages inventory for items such as meals, entertainment, and amenities onboard flights.

## Sequence
- Note: P is passenger, AE is airport employee, C is crew, SE is system employee, S is system(automate)
- Main flow
  - SE: create available flight, prices, available seats, schedules (depart time, duration) (9)
  - P: search for available flights
  - P: (8) choose depart time + destination -> choose number of passenger (adult, baby, child) -> chose type of ticket (vip / normal) -> choose type of trip (normal / round trip) -> fill information -> choose seat -> choose luggage options -> pay (payment service) 
  - P: go to airport
  - P: show code or qr code
  - AE: verify in system (6)
  - P: ship luggage
  - AE: check luggage's weight
  - S: tracking passenger's luggage realtime (5)
  - P: go to waiting hall, see flight's status (3)
  - S: announce passenger go to port and plane
  - P: go to port and plane
  - S: check crew availability (2)
  - S: assign and announce the crew (prepare to start the flight) (1)
  - C: approve/reject assign
  - C: go to port and plane
- When a flight cancel/delay (4)
  - SE: announce why cancel/delay
  - S: update status flight at waiting hall
- System manage reservation process (7)

### Tech stack
- Micro Frontends
  - Reactjs - Nextjs
  - Webpack Module Federation
- Backend: Go, gRPC
- Database: MongoDB
- Message Brokers
  - Kafka
  - Facilitates asynchronous communication and event-driven architecture by decoupling producers and consumers of data. It provides fault tolerance, scalability, and durability for message processing in distributed systems
- API Gateway
  - Manages API traffic, handles authentication, rate limiting, and routing of requests to appropriate microservices. It provides a centralized entry point for clients and ensures security and reliability.
  - Kong
- Containerization: Docker, Kubernetes
- Service Mesh
  - Istio/Linkerd/Envoy
  - Manages communication between microservices, handling service discovery, load balancing, traffic routing, and security features such as encryption and authentication. It provides observability into service-to-service communication and ensures resilience in distributed architectures.
- Configuration Management
  - Stores and manages configuration settings for microservices, allowing dynamic configuration updates without requiring code changes or service restarts. It ensures consistency, availability, and versioning of configurations across distributed systems.
  - Consul/Zookeeper
- API Documentation: Swagger/OpenAPI
- API Versioning
  - Defines a clear and consistent versioning scheme for APIs, ensuring backward compatibility and smooth transitions between different versions. It enables clients to safely adopt new features and enhancements without breaking existing integrations.
  - Semantic Versioning (SemVer)
- Content Delivery Networks (CDN)
  - Cloudfare
  - Accelerates content delivery by caching static assets closer to end-users, reducing latency and improving website performance. It ensures high availability and scalability for global audiences accessing airline-related content.
- Distributed Caching
  - Redis
  - Stores frequently accessed data in memory to accelerate read-heavy workloads, reduce database load, and improve application performance. It provides low-latency access to cached data and supports advanced data structures and caching strategies.
- Secrets Management
  - Safely stores and manages sensitive information such as API keys, passwords, and certificates. It provides secure access control and audit logging, ensuring that secrets are never exposed in plaintext.
  - Vault
- Load Balancing
  - Distributes incoming traffic across multiple instances of microservices to ensure high availability, fault tolerance, and optimal resource utilization. It improves performance and scalability by evenly distributing the load.
  - NGINX
- Distributed Tracing
  - Jaeger
  - Captures and visualizes the flow of requests across microservices, allowing developers to identify performance issues, latency bottlenecks, and dependencies between services. It helps diagnose and troubleshoot complex distributed systems.
- Chaos Engineering
  - Chaos Toolkit, Gremlin, Chaos Mesh
  - Introduces controlled chaos into the system to proactively identify weaknesses and vulnerabilities. It simulates real-world failures and evaluates system resilience under adverse conditions, leading to more robust and reliable microservices architecture.
- (devops) Monitoring: Prometheus (Grafana)
- (devops) Infrastructure as Code (IaC)
  - Allows you to define and provision infrastructure resources (e.g., virtual machines, networks, storage) using code. It enables automated, repeatable, and version-controlled infrastructure deployment, reducing manual errors and accelerating development cycles.
  - Terraform
- (devops) Logging: ELK Stack
  - Elasticsearch: A distributed search and analytics engine that stores and indexes log data, enabling fast and scalable search capabilities. It supports full-text search, real-time indexing, and complex querying for analyzing logs generated by microservices and infrastructure components.
  - Logstash: A data processing pipeline that ingests, transforms, and enriches log data from various sources before sending it to Elasticsearch for indexing. It provides filters for parsing and formatting logs, handling timestamps, and enriching events with additional metadata.
  - Kibana: A data visualization and exploration tool that allows you to create dashboards, charts, and graphs based on data stored in Elasticsearch. It provides a user-friendly interface for monitoring, analyzing, and troubleshooting log data, helping to identify trends, anomalies, and performance issues.
