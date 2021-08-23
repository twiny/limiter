## HTTP Rate Limiter
rate limit HTTP requests based on client IP address.

### Install
`go get github.com/twiny/limiter`

### Example

```go

import "github.com/twiny/limiter"

// App
type App struct{
    ... ...
    limiter *limiter.Limiter
    ... ...
}

// RateLimit
func (a *App) RateAllow(ip string) bool {
	limit, found := a.limiter.Get(ip)
	if !found {
		a.limiter.Set(ip)
		return true
	}

	return limit.Allow()
}
```

```go
// HTTPHandler
type HTTPHandler struct{
    ... ...
    app *App
    ... ...
}

// rate limit
func (route *HTTPHandler) ratelimit(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var remoteIP string
		remoteIP = r.RemoteAddr
		if strings.ContainsRune(r.RemoteAddr, ':') {
			remoteIP, _, _ = net.SplitHostPort(r.RemoteAddr)
		}

		if !route.app.RateAllow(remoteIP) {
			http.Error(w,"going too fast - slow down", http.StatusTooManyRequests)
			return
		}

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
```
