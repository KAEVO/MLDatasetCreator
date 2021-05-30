package dsbldr

import (
	"fmt"
	"regexp"
)

// RunFunc holds the computation that processes the API responses to features
// is sent an array of JSON strings as the responses ??as well as a map of data from the features parent features??
// Basically what you do with the run function is take in a string of
// serialized 