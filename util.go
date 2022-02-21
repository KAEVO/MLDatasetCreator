package dsbldr

import (
	"fmt"
)

// BasicOAuthHeader spits out a basic OAuth Header based on access token
func BasicOAuthHeader(consumerKey, nonce, signature, signatureMethod,
	timestamp, token string) string {
	return fmt.Sprintf(`OAuth oauth_consumer_key="%s",
		oauth_nonce="%s",
		oauth_signature="%s",
		oauth_signature_method="%s",
		oauth_timestamp="%s",
		oauth_token="%s`,
		consumerKey, nonce, signature, signatureMethod, timestamp, token)
}

func writeStringColumn(data [][]string, columnName string, values []string) {
	var colIndex int
	for i := range data[0] {
		// Find first empty column or column with same header to overwrite
		if data[0][i] == "" 