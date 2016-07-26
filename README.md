# go-aha

Go client library for accessing the [Aha! API](http://www.aha.io/api)

## Usage

```go
import "github.com/Tonkpils/go-aha/aha"
```

The main entrypoint for the package is the client which can be constructed with `NewClient` or `NewBasicAuthClient`.
Both clients require the Aha! Account Name.

```go
client, err := aha.NewClient(nil, "my account")
```

### Authentication

#### OAuth2

To use oauth2, provide an `http.Client` as the first argument to `NewClient`. A library like [oauth2](https://github.com/golang/oauth2) provides this functionality.

#### Basic Auth

The Basic Auth Client will use `http.DefaultClient` and send along the username and password.

```go
client, err := aha.NewBasicAuthClient("username", "password", "myaccount")
```

## TODO

- [ ] Pagination
- [ ] Response struct and errors
- [ ] Implement all of the APIs
- [ ] Tests

## LICENSE

[LICENSE](LICENSE)
