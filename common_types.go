package fomo

import "fmt"

// APIError is what the what the api returns on error
type APIError struct {
	Errors  string `json:"error,omitempty"`
	Success bool   `json:"success,omitempty"`
}

func (err APIError) String() string {
	return fmt.Sprintf("%s", err.Errors)
}
func (err APIError) Error() string {
	return err.String()
}

// HasError checks if this call had an error
func (err APIError) HasError() bool {
	return err.Success
}

// QueryParams defines the different params
type QueryParams interface {
	Params() map[string]string
}
