# Request/Response Patterns

## Parsing JSON Requests

Most modern REST APIs work with JSON. Here's a simple pattern for parsing JSON request bodies:

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func handleUserCreate(w http.ResponseWriter, r *http.Request) {
    // Check method
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

        return
    }
    
    // Parse the request body
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)

        return
    }
    
    // Process the user data...
    
    // Send response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
```

## Form Data

```go
func handleFormSubmission(w http.ResponseWriter, r *http.Request) {
    // Parse form data
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Error parsing form data", http.StatusBadRequest)

        return
    }
    
    // Access form values
    name := r.Form.Get("name")
    email := r.Form.Get("email")
    
    // Process the data...
    
    // Send response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
```

## Url parameters and Query strings

```go
func getUserHandler(w http.ResponseWriter, r *http.Request) {
    // Get URL parameter
    userID := r.PathValue("id")
    
    // Get query parameter
    format := r.URL.Query().Get("format")
    
    // Use the parameters...
    
    // Send response
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string{
        "userId": userID,
        "format": format,
    }
    json.NewEncoder(w).Encode(response)
}
```

## Structured Response Patterns

> Using consistent response structures helps API consumers and maintainers:
> Always use a structured response

```go
type SuccessResponse struct {
    Status  string      `json:"status"`
    Data    interface{} `json:"data"`
}

func sendSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
    response := SuccessResponse{
        Status: "success",
        Data:   data,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}
```

## Error Responses

```go
type ErrorResponse struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
    response := ErrorResponse{
        Status:  "error",
        Message: message,
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(response)
}
```


## Status Code Selection

A simplified guide to common status codes:

- `200 OK`: Request succeeded
- `201 Created`: Resource successfully created
- `400 Bad Request`: Invalid request format or parameters
- `401 Unauthorized`: Authentication required
- `403 Forbidden`: Authenticated but not authorized
- `404 Not Found`: Resource not found
- `405 Method Not Allowed`: HTTP method not supported
- `422 Unprocessable Entity`: Well-formed request but semantically invalid (validation errors)
- `500 Internal Server Error`: Unexpected server error

```go
func createResourceHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request
    var req CreateRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        // Syntax error in the request
        sendErrorResponse(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }
    
    // Validate request
    validationErrors := validateRequest(req)
    if len(validationErrors) > 0 {
        // Request syntax was valid but the content had validation errors
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":  "error",
            "message": "Validation failed",
            "errors":  validationErrors,
        })
        return
    }
    
    // Create resource
    resource, err := createResource(req)
    if err != nil {
        sendErrorResponse(w, "Failed to create resource", http.StatusInternalServerError)
        return
    }
    
    // Success
    sendSuccessResponse(w, resource, http.StatusCreated)
}
```

## Simple request validation

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func (r *CreateUserRequest) Validate() map[string]string {
    errors := make(map[string]string)
    
    if r.Name == "" {
        errors["name"] = "name is required"
    }
    if r.Email == "" {
        errors["email"] = "email is required"
    } else if !strings.Contains(r.Email, "@") {
        errors["email"] = "invalid email format"
    }
    
    return errors
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
        return
    }
    
    // Validate
    validationErrors := req.Validate()
    if len(validationErrors) > 0 {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusUnprocessableEntity)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status":  "error",
            "message": "Validation failed",
            "errors":  validationErrors,
        })
        return
    }
    
    // Process valid request...
}
```

## Best Practices

1. **Be consistent**: Use the same response format across your API
3. **Use appropriate status codes**: Choose status codes that clearly communicate what happened
4. **Validate early**: Check request data before performing business logic
5. **Set proper headers**: Always set Content-Type header
6. **Distinguish between different error types**: Use 400 for syntax errors and 422 for validation errors
7. **Handle errors gracefully**: Provide meaningful error messages

By following these simple patterns, you'll create a consistent, intuitive API that's easy for clients to use and for your team to maintain.


