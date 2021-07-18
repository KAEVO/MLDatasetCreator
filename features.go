package dsbldr

import (
	"fmt"
	"regexp"
)

// RunFunc holds the computation that processes the API responses to features
// is sent an array of JSON strings as the responses ??as well as a map of data from the features parent features??
// Basically what you do with the run function is take in a string of
// serialized API data (could be in JSON or XML), do parsing on your
// own or using utility functions. You do whatever computations you want and
// then spit it back as an array of strings to read to CSV or JSON
type RunFunc func(responses []string) []string // parents map[string]string

// Structs representing RetreiveType
// SingleRetrieve Features only require one request to create the JSON Dump
// that's passed to the RunFunc
// Repeated Retrieve Features require one request per value-set of
// of parent features that are concatenated into a JSON array and then passed
// to the Features RunFunc
// Almost as a given, all dependent features will be of RepeatedRetrieve per
// value sets of their parent