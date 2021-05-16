package dsbldr

import (
	"fmt"
	"regexp"
)

// RunFunc holds the computation that processes the API responses to features
// is sent an array of JSON strings as the responses ??as well as a map of data from the fe