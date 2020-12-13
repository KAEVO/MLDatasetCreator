
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
	}
	// err := errors.New("No such Feature in builder")
	return feat
}

func (b *Builder) writeRecord(writer csv.Writer, i int) error {
	var record []string
	for index, j := range b.data[i] {
		// if data header (feature name) has noSave == false, else don't write
		if !b.featureMap[b.data[0][index]].noSave {
			record = append(record, j)
		}
	}
	err := writer.Write(record)
	if err != nil {
		return err
	}
	return nil
}

// Save commits the downloaded features to a file
func (b *Builder) Save(writer csv.Writer) error {
	for i := range b.data {
		err := b.writeRecord(writer, i)
		if err != nil {
			return err
		}
	}
	return nil
}

// SaveIf saves records only if saveCond evaluate to true
func (b *Builder) SaveIf(writer csv.Writer, saveCond func(r []string) bool) error {
	for i := range b.data {
		if saveCond(b.data[i]) {
			err := b.writeRecord(writer, i)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AddFeatures adds a Feature struct to the "Features" Field on Builder
func (b *Builder) AddFeatures(features ...*Feature) {
	for _, feature := range features {
		feature.noSave = false
		b.featureMap[feature.Name] = feature
		// TODO: Return error if feature with same name has been added
	}
	// Increase size of data matrix if feature map is larger than initially allocated
	if len(b.featureMap) > len(b.data[0]) {
		for i := range b.data {
			b.data[i] = append(b.data[i], "")
		}