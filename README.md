## basicauth

Simple CLI to convert username and password to basic auth string or header.

### install

```
go get github.com/jmervine/basicauth
```

### usage

```
basicauth -h
```


### in code

```go
import (
    // ...

    "github.com/jmervine/basicauth/basicauth"
)


func main() {
    // get user, pass however you want

    ba := basicauth.NewBasicAuth(user, pass)

    print(ba.String())
    print(ba.Header())
}
```
