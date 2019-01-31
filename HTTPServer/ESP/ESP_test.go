package ESP

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

var esp ESP

func TestNewESP(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want *ESP
	}{
		{name: "test1", args: args{makeRequest("11:22:33:cd:62:bb")}},
		{name: "test2", args: args{makeRequest("")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewESP(tt.args.r)
			if reflect.TypeOf(got) == reflect.TypeOf(ESP{}) {
				t.Errorf("NewESP() = %s, want %v", got.MacAp(), tt.want)
			}
		})
	}
}

func makeRequest(macSta string) *http.Request {
	r1, _ := http.NewRequest("GET", "/co2.bin", strings.NewReader(""))
	r1.Header.Add("x-ESP8266-STA-MAC", macSta)
	r1.Header.Add("x-ESP8266-AP-MAC", "lnn")
	r1.Header.Add("x-ESP8266-free-space", "123")
	r1.Header.Add("x-ESP8266-sketch-size", "123")
	r1.Header.Add("x-ESP8266-sketch-md5", "")
	r1.Header.Add("x-ESP8266-chip-size", "123")
	r1.Header.Add("x-ESP8266-sdk-version", "123")
	r1.Header.Add("x-ESP8266-mode", "sss")
	r1.Header.Add("x-ESP8266-version", "123")
	return r1
}
