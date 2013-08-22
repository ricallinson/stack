package connect

import(
    "net/http"
)

type Request struct {
    Raw *http.Request
    Url string
    OriginalUrl string
}

func CreateRequest(raw *http.Request) (*Request) {
    req := new(Request)
    req.Raw = raw
    req.Url = raw.RequestURI
    req.OriginalUrl = raw.RequestURI
    return req
}