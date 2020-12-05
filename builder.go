
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
}

// WithBasicAuth is a Builder option that adds a username and password for Basic API authentication
func WithBasicAuth(username, password string) {
	return func(b *Builder) {
		b.authUsername = username
		b.authPassword = password
	}
}

func (b *Builder) addFeatureData(featureName string, values []string) error {
	writeStringColumn(b.data, featureName, values)
	return nil
}

func (b *Builder) getFeatureData(featureName string) []string {
	items := make([]string, b.records)
	if _, ok := b.featureMap[featureName]; ok {
		readStringColumn(items, featureName, b.data)
		return items
	}
	// readStringColumn(items, featureName, b.data)
	return items
}

// GetFeature returns a feature in the detaset based on it's name
func (b *Builder) GetFeature(name string) *Feature {
	var feat *Feature
	if val, ok := b.featureMap[name]; ok {
		feat = val
		return feat