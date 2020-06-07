package adapter

import "testing"

var requset = "adaptee requst"
var response = "adaptee response"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	req := target.Request()
	if req != requset {
		t.Fatalf("expect: %s, actual: %s", requset, req)
	}
	res := target.Response()
	if res != response {
		t.Fatalf("expect: %s, actual: %s", response, res)
	}
}
