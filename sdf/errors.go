package sdf

import "github.com/polefishu/sdk-builder/sdf/sdferr"

var (
	// ErrMissingRegion is an error that is returned if region configuration is
	// not found.
	//
	// @readonly
	ErrMissingRegion = sdferr.New("MissingRegion", "could not find region configuration", nil)

	// ErrMissingEndpoint is an error that is returned if an endpoint cannot be
	// resolved for a service.
	//
	// @readonly
	ErrMissingEndpoint = sdferr.New("MissingEndpoint", "'Endpoint' configuration is required for this service", nil)
)
