/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.24.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/core"
	. "github.com/twilio/twilio-go/rest/trusthub/v1"
)

func ResourceCustomerProfiles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCustomerProfiles,
		ReadContext:   readCustomerProfiles,
		UpdateContext: updateCustomerProfiles,
		DeleteContext: deleteCustomerProfiles,
		Schema: map[string]*schema.Schema{
			"email":           AsString(SchemaRequired),
			"friendly_name":   AsString(SchemaRequired),
			"policy_sid":      AsString(SchemaRequired),
			"status_callback": AsString(SchemaComputedOptional),
			"sid":             AsString(SchemaComputed),
			"status":          AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseCustomerProfilesImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createCustomerProfiles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCustomerProfileParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.TrusthubV1.CreateCustomerProfile(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateCustomerProfiles(ctx, d, m)
}

func deleteCustomerProfiles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteCustomerProfile(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCustomerProfiles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchCustomerProfile(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseCustomerProfilesImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateCustomerProfiles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCustomerProfileParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.UpdateCustomerProfile(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceCustomerProfilesChannelEndpointAssignments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCustomerProfilesChannelEndpointAssignments,
		ReadContext:   readCustomerProfilesChannelEndpointAssignments,
		DeleteContext: deleteCustomerProfilesChannelEndpointAssignments,
		Schema: map[string]*schema.Schema{
			"customer_profile_sid":  AsString(SchemaForceNewRequired),
			"channel_endpoint_sid":  AsString(SchemaForceNewRequired),
			"channel_endpoint_type": AsString(SchemaForceNewRequired),
			"sid":                   AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseCustomerProfilesChannelEndpointAssignmentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createCustomerProfilesChannelEndpointAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCustomerProfileChannelEndpointAssignmentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	customerProfileSid := d.Get("customer_profile_sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.CreateCustomerProfileChannelEndpointAssignment(customerProfileSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{customerProfileSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteCustomerProfilesChannelEndpointAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	customerProfileSid := d.Get("customer_profile_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteCustomerProfileChannelEndpointAssignment(customerProfileSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCustomerProfilesChannelEndpointAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	customerProfileSid := d.Get("customer_profile_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchCustomerProfileChannelEndpointAssignment(customerProfileSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseCustomerProfilesChannelEndpointAssignmentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected customer_profile_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("customer_profile_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceCustomerProfilesEntityAssignments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCustomerProfilesEntityAssignments,
		ReadContext:   readCustomerProfilesEntityAssignments,
		DeleteContext: deleteCustomerProfilesEntityAssignments,
		Schema: map[string]*schema.Schema{
			"customer_profile_sid": AsString(SchemaForceNewRequired),
			"object_sid":           AsString(SchemaForceNewRequired),
			"sid":                  AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseCustomerProfilesEntityAssignmentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createCustomerProfilesEntityAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCustomerProfileEntityAssignmentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	customerProfileSid := d.Get("customer_profile_sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.CreateCustomerProfileEntityAssignment(customerProfileSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{customerProfileSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteCustomerProfilesEntityAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	customerProfileSid := d.Get("customer_profile_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteCustomerProfileEntityAssignment(customerProfileSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readCustomerProfilesEntityAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	customerProfileSid := d.Get("customer_profile_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchCustomerProfileEntityAssignment(customerProfileSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseCustomerProfilesEntityAssignmentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected customer_profile_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("customer_profile_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceEndUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createEndUsers,
		ReadContext:   readEndUsers,
		UpdateContext: updateEndUsers,
		DeleteContext: deleteEndUsers,
		Schema: map[string]*schema.Schema{
			"friendly_name": AsString(SchemaRequired),
			"type":          AsString(SchemaRequired),
			"attributes":    AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseEndUsersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateEndUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.TrusthubV1.CreateEndUser(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteEndUser(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchEndUser(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseEndUsersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateEndUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateEndUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.UpdateEndUser(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceSupportingDocuments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createSupportingDocuments,
		ReadContext:   readSupportingDocuments,
		UpdateContext: updateSupportingDocuments,
		DeleteContext: deleteSupportingDocuments,
		Schema: map[string]*schema.Schema{
			"friendly_name": AsString(SchemaRequired),
			"type":          AsString(SchemaRequired),
			"attributes":    AsString(SchemaComputedOptional),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseSupportingDocumentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateSupportingDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.TrusthubV1.CreateSupportingDocument(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteSupportingDocument(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchSupportingDocument(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseSupportingDocumentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateSupportingDocuments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateSupportingDocumentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.UpdateSupportingDocument(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceTrustProducts() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrustProducts,
		ReadContext:   readTrustProducts,
		UpdateContext: updateTrustProducts,
		DeleteContext: deleteTrustProducts,
		Schema: map[string]*schema.Schema{
			"email":           AsString(SchemaRequired),
			"friendly_name":   AsString(SchemaRequired),
			"policy_sid":      AsString(SchemaRequired),
			"status_callback": AsString(SchemaComputedOptional),
			"sid":             AsString(SchemaComputed),
			"status":          AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrustProductsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrustProducts(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTrustProductParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.TrusthubV1.CreateTrustProduct(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateTrustProducts(ctx, d, m)
}

func deleteTrustProducts(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteTrustProduct(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrustProducts(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchTrustProduct(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrustProductsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateTrustProducts(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateTrustProductParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.UpdateTrustProduct(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceTrustProductsChannelEndpointAssignments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrustProductsChannelEndpointAssignments,
		ReadContext:   readTrustProductsChannelEndpointAssignments,
		DeleteContext: deleteTrustProductsChannelEndpointAssignments,
		Schema: map[string]*schema.Schema{
			"trust_product_sid":     AsString(SchemaForceNewRequired),
			"channel_endpoint_sid":  AsString(SchemaForceNewRequired),
			"channel_endpoint_type": AsString(SchemaForceNewRequired),
			"sid":                   AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrustProductsChannelEndpointAssignmentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrustProductsChannelEndpointAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTrustProductChannelEndpointAssignmentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	trustProductSid := d.Get("trust_product_sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.CreateTrustProductChannelEndpointAssignment(trustProductSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{trustProductSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteTrustProductsChannelEndpointAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trustProductSid := d.Get("trust_product_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteTrustProductChannelEndpointAssignment(trustProductSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrustProductsChannelEndpointAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trustProductSid := d.Get("trust_product_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchTrustProductChannelEndpointAssignment(trustProductSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrustProductsChannelEndpointAssignmentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected trust_product_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("trust_product_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceTrustProductsEntityAssignments() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrustProductsEntityAssignments,
		ReadContext:   readTrustProductsEntityAssignments,
		DeleteContext: deleteTrustProductsEntityAssignments,
		Schema: map[string]*schema.Schema{
			"trust_product_sid": AsString(SchemaForceNewRequired),
			"object_sid":        AsString(SchemaForceNewRequired),
			"sid":               AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrustProductsEntityAssignmentsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrustProductsEntityAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTrustProductEntityAssignmentParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	trustProductSid := d.Get("trust_product_sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.CreateTrustProductEntityAssignment(trustProductSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{trustProductSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteTrustProductsEntityAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trustProductSid := d.Get("trust_product_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrusthubV1.DeleteTrustProductEntityAssignment(trustProductSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrustProductsEntityAssignments(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trustProductSid := d.Get("trust_product_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrusthubV1.FetchTrustProductEntityAssignment(trustProductSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrustProductsEntityAssignmentsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected trust_product_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("trust_product_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
