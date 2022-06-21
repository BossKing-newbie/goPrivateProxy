package middleware

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestInitializationGoProxy(t *testing.T) {
	api := "api/ping"
	name, err := url.PathUnescape(api)
	fmt.Println(err)
	fmt.Println(strings.HasSuffix(name, "/"))
}
