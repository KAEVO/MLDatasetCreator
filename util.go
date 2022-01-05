package dsbldr

import (
	"fmt"
)

// BasicOAuthHeader spits out a basic OAuth Header based on access token
func BasicOAuthHeader(consumerKey, nonce, signature, signatureMethod,