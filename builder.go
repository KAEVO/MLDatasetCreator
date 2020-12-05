
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

	// Some fields for auth credentials
	authUsername string
	authPassword string
}

// NewBuilder creates new Builder struct
func NewBuilder(featureCount, recordCount int, options ...func(*Builder)) *Builder {
	// Add extra row for header
	preallocatedData := make([][]string, recordCount+1)
	for i := range preallocatedData {
		preallocatedData[i] = make([]string, featureCount)
	}
	b := Builder{
		featureMap: make(map[string]*Feature),
		data:       preallocatedData,
		records:    recordCount,
	}

	for _, option := range options {
		option(&b)
	}

	return &b