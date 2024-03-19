package http

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{id: 1, name: "test", info: "test"}`))
}

func TestMakeHttp(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	fmt.Println(server.URL)

	time.Sleep(3 * time.Second)

	want := &Response{ID: 1, Name: "test", Info: "test"}

	t.Run("should return response", func(t *testing.T) {
		got, err := MakeHttpCall(server.URL)
		// if err != nil {
		// 	t.Errorf("MakeHttpCall(%s) got error %v", server.URL, err)
		// }
		// if got.ID != want.ID {
		// 	t.Errorf("MakeHttpCall(%s) = %v; want %v", server.URL, got, want)
		// }

		if !reflect.DeepEqual(got, want) {
			t.Errorf("MakeHttpCall(%s) = %v; want %v", server.URL, got, want)
		}

		if !errors.Is(err, nil) {
			t.Errorf("is err(%s) = %v; want %v", server.URL, got, want)
		}
	})

}
