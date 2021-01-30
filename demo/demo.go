package main

import (
	d "dsbldr"
	"encoding/json"
	"fmt"
)

builder d.Builder := d.Builder{
	BaseURL: "localhost:8080",
	RequestHeaders: map[string]string{
		"Authoriza