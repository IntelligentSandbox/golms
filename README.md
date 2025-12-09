# golms
### Default: Uses cache, ignores vendor/
```go
go build
```
Looks in ~/go/pkg/mod/

### Vendor mode: Uses vendor/
```go
go build -mod=vendor
```
Looks in ./vendor/