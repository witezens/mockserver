# Mock API - Modular Strategy with Handlers and Rules

This project implements a flexible mock API system using Go + Fiber, capable of resolving HTTP requests through:

- ✅ Custom dynamic handlers for business logic
- ✅ File-based mock responses resolved via rules
- ✅ Fallbacks for default mocks
- ❌ Graceful 404s when nothing matches

---

## 🔁 Resolution Order

When a request hits `/service-name/api/v1/resource`, the system evaluates it in the following order:

1. 🔍 **Is there a registered handler?**
    - ✅ Yes → execute the Go function (`fiber.Handler`) registered for `service/resource`

2. 🔍 **Is there a matching file rule?**
    - ✅ Yes → use the matching file template (e.g. `createOrder__customerId_1234.POST.json`)

3. 🔍 **Is there a default mock file?**
    - ✅ Yes → use the fallback file (e.g. `createOrder.POST.json`)

4. ❌ **Nothing matched**
    - → return a 404 with `{ "error": "Mock no encontrado" }`

---

## 📦 Example

### POST /serviceordering/api/v1/createOrder

#### A. With a handler:
```go
registry.Register("serviceordering/createOrder", HandleCreateOrder)
```
✅ Executes the handler instead of serving a file.

#### B. With mock rule (if no handler):
```go
{ "customerId": "1234" }
```
✅ Resolves to:
```
mock-data/serviceordering/api/v1/createOrder__customerId_1234.POST.json
```

#### C. Fallback:
If no param rule matches, fallback to:
```
createOrder.POST.json
```

---

## 🧩 Rule Format

```go
MockRule{
  Param: "customerId",
  Template: "%s__customerId_%v.%s.json",
  Source: "body" // or "query"
}
```

- `%s` → resource name
- `%v` → value of the param
- `%s` → method (GET, POST...)

---

## 📁 Project Structure (simplified)

```
mock-api/
├── handler/
├── registry/
├── resolver/
├── services/
│   ├── serviceordering/
│   └── serviceinventory/
├── utils/
├── mock-data/
└── main.go
```

---

## 🛠 How to Extend

- Add a new `serviceX/handlers.go` and call `registry.Register(...)` in an `init()` function
- Add a `.json` file under `mock-data/`
- Define new `MockRule`s in `main.go` for param-based resolution

---
🛠️ Recommendations Before Building the Docker Image
Before running docker build, make sure your Go module files are up to date to prevent errors during the image build (especially in the go mod download step):


```bash
go mod tidy
go mod verify
```

This ensures:

- All used packages are listed in go.mod

- The go.sum file contains all required checksums

- The build process won't fail due to missing or outdated module info

🚀 Typical build flow
``` bash
go mod tidy && go mod verify
docker build -t mock-server .
```

Tip: You can automate these steps in a Makefile or a build.sh script for CI/CD pipelines.
---
