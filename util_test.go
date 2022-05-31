package dsbldr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBasicOAuthHeader(t *testing.T) {
	consumerKey := "consumerKey"
	nonce := "nonce"
	signature := "signature"
	signat