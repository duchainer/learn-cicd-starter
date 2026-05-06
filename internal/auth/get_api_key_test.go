package auth

import (
    //"reflect"
    "testing"
    "net/http"
)

func TestEmptyHeader(t *testing.T) {
    _, error := GetAPIKey(http.Header{})
    if error != nil{
        t.Fatalf("Expected to get an error, not %v", error)
    }
}
