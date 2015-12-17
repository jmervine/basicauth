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

### example w/ curl

```
curl -i -H "$(basicauth -u 'user' -p 'pass' -header)" http://localhost:5000/endpoint/foo
```

> I know `curl` supports `-u` for this, I including this as a simple example.
> Typically I use it for things like `httperf` and/or `siege`.

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
