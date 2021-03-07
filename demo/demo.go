package main

import (
	d "dsbldr"
	"encoding/json"
	"fmt"
)

builder d.Builder := d.Builder{
	BaseURL: "localhost:8080",
	RequestHeaders: map[string]string{
		"Authorization": BasicOAuthHeader(
			"OAUTH_CONSUMER_KEY",
			"OAUTH_NONCE",
			"OAUTH_SIGNATURE",
			"OAUTH_SIGNATURE_METHOD", "OAUTH_TIMESTAMP",
			"OAUTH_TOKEN",
		),
	},
}

builder.AddFeatures(
	&d.Feature{
		Name: "item_ids",
		Endpoint: "/items/",
		RunFunc: fun