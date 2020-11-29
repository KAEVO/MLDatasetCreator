
package dsbldr

import (
	"encoding/csv"
	"io/ioutil"
	"net/http"
)

// Builder is main type for this tool.
type Builder struct {
	BaseURL    string
	featureMap map[string]*Feature
	data       [][]string // Strings of Data to be read in to CSV
	records    int        // Number of records to be retrieved for dataset