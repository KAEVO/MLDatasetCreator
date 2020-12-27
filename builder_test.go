
package dsbldr

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewBuilder(t *testing.T) {
	b := NewBuilder(4, 100)
	if got, want := len(b.data), 101; got != want {
		t.Fatalf("got: %v\n want: %v\n ", got, want)
	}
	if got, want := len(b.data[0]), 4; got != want {
		t.Fatalf("got: %v\n want: %v\n ", got, want)
	}
	if got, want := b.records, 100; got != want {
		t.Fatalf("got: %v\n want: %v\n ", got, want)
	}
}

func TestAddFeatureData(t *testing.T) {
	b := NewBuilder(4, 3)
	b.addFeatureData("newFeature", []string{"one", "two", "three"})
	expectedData := [][]string{
		[]string{"newFeature", "", "", ""},
		[]string{"one", "", "", ""},
		[]string{"two", "", "", ""},