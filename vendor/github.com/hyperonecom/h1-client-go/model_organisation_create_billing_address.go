/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type OrganisationCreateBillingAddress struct {
	Country string `json:"country,omitempty"`
	City string `json:"city,omitempty"`
	Zipcode string `json:"zipcode,omitempty"`
	Street string `json:"street,omitempty"`
}
