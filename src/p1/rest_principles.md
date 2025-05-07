# Basic Rest Principles

REST (Representational State Transfer) is an architectural style for designing networked applications. 
Understanding these core principles will help you design APIs that are intuitive, maintainable, and align with web standards.

## Resources as the Core Concept

In REST, everything is a resource. 
A resource is any information that can be named: a document, an image, a service, a collection of resources, or even a concept.

```go
// Resources map naturally to URLs
// /users - a collection of users
// /users/42 - a specific user
// /users/42/orders - orders belonging to user 42
```

## HTTP Methods and CRUD Operations

REST APIs use HTTP methods to represent operations on resources:

- **POST**: Create a new resource (*C*reate)
- **GET**: Retrieve a resource or collection (*R*ead)
- **PUT/PATCH**: Update an existing resource (*U*pdate)
- **DELETE**: Remove a resource (*D*elete)


```go
// Example handler mapping
func setupRoutes() {
    http.HandleFunc("GET /users", listUsers)
    http.HandleFunc("GET /users/{id}", getUser)
    http.HandleFunc("POST /users", createUser)
    http.HandleFunc("PUT /users/{id}", updateUser)
    http.HandleFunc("DELETE /users/{id}", deleteUser)
}
```

## Statelessness

Each request must contain all the information needed to understand and process it. 
The server doesn't store client state between requests.

```go
// Each request should be self-contained
func handler(w http.ResponseWriter, r *http.Request) {
    // Authentication should be in each request (e.g., via token)
    token := r.Header.Get("Authorization")
    
    // Process based only on the request content
    // ...
}
```

## Resource URLs

Design clean, hierarchical URLs:

```
/users                  // Collection of users
/users/123              // Specific user with ID 123
/users/123/orders       // Orders belonging to user 123
/users/123/orders/456   // Specific order 456 for user 123
```

When designing URLs:

- Use nouns, not verbs (e.g., /users not /getUsers)
- Use plural nouns for collections
- Use parameters for filtering: /users?status=active
- Use hierarchical relationships when they exist

## Response Formats

JSON is the most common format for REST APIs:

```go
func getUser(w http.ResponseWriter, r *http.Request) {
    // Get user data
    user := fetchUserFromDB(r.PathValue("id"))
    
    // Set content type
    w.Header().Set("Content-Type", "application/json")
    
    // Return as JSON
    json.NewEncoder(w).Encode(user)
}
```

## Idempotence

- **Idempotent operations**: Performing the same operation multiple times has the same effect as doing it once
- GET, PUT, DELETE should be idempotent
- POST is typically not idempotent

```go
// PUT is idempotent - multiple identical requests should have same result
func updateUser(w http.ResponseWriter, r *http.Request) {
    userID := r.PathValue("id")
    var userData User
    
    json.NewDecoder(r.Body).Decode(&userData)
    userData.ID = userID
    
    // Replace the entire resource
    storeUserInDB(userData)
    
    w.WriteHeader(http.StatusOK)
}
```

## Versioning

APIs evolve over time, so versioning is essential:

```
/api/v1/users
/api/v2/users
```

## Pragmatic REST API

While REST has formal principles, a pragmatic approach focuses on:

- Clear resource naming: Use intuitive, consistent URL structures
- Appropriate HTTP methods: Match HTTP verbs to operations
- Proper status codes: Communicate outcomes clearly
- Consistent responses: Use the same format throughout your API
- Stateless design: Keep requests self-contained

> REST is just a helpful approach, not a strict rulebook. 
> What matters most is building APIs that make sense to the people using them. 
> Focus on creating clear, consistent interfaces that are easy to understand and use, rather than worrying about following every academic REST rule perfectly.


