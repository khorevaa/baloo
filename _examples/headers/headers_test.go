package simple

import (
	"testing"

	"gopkg.in/h2non/baloo.v0"
)

// test stores the HTTP testing client preconfigured
var test = baloo.New("http://httpbin.org")

func TestBalooBasic(t *testing.T) {
	test.Get("/get").
		SetHeader("Foo", "Bar").
		Expect(t).
		Status(200).
		Header("Server", "apache").
		Header("Content-Type", "json").
		Type("json").
		JSON(map[string]string{"bar": "foo"}).
		Done()
}
