package utils

import (
	"os"
	"testing"
	"time"
)

func TestGetIdFromJWT(t *testing.T) {
	Now = func() time.Time { return time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC) }
	os.Setenv("hmacSecret", "my_secret")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImlkIiwiZXhwX3RpbWUiOiIyMDIwLTEyLTA5VDE2OjA5OjUzKzAwOjAwIn0.JQxFqUBjU2cu4ontl2-JKrRbpxHcPFuHYFX-gZjLBzA"
	want := "id"
	id, err := GetIdFromJWT(token)

	if !(want == id) || err != nil {
		t.Fatalf(`GetIdFromJwt = %q, %v, want match for %#q, nil`, token, err, want)
	}
}

func TestCreateJWT(t *testing.T) {
	Now = func() time.Time { return time.Date(1974, time.May, 19, 1, 2, 3, 4, time.UTC) }
	os.Setenv("hmacSecret", "my_secret")

	id := "id"
	want := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBfdGltZSI6IjE5NzQtMDUtMTlUMDI6MDI6MDNaIiwiaWQiOiJpZCJ9.idps1p5uPYlGqzvI-DWwdiNehG4-M89rwODV0IQ3Tyk"
	token, err := CreateJWT(id)

	if !(want == token) || err != nil {
		t.Fatalf(`CreateJWT = %q, %v, want match for %#q, nil`, token, err, want)
	}
}
