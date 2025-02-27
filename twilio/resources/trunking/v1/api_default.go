/*
 * Twilio - Trunking
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
	. "github.com/twilio/twilio-go/rest/trunking/v1"
)

func ResourceTrunksCredentialLists() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrunksCredentialLists,
		ReadContext:   readTrunksCredentialLists,
		DeleteContext: deleteTrunksCredentialLists,
		Schema: map[string]*schema.Schema{
			"trunk_sid":           AsString(SchemaForceNewRequired),
			"credential_list_sid": AsString(SchemaForceNewRequired),
			"sid":                 AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrunksCredentialListsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrunksCredentialLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialListParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	trunkSid := d.Get("trunk_sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.CreateCredentialList(trunkSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{trunkSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteTrunksCredentialLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrunkingV1.DeleteCredentialList(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrunksCredentialLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.FetchCredentialList(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrunksCredentialListsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected trunk_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("trunk_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceTrunksIpAccessControlLists() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrunksIpAccessControlLists,
		ReadContext:   readTrunksIpAccessControlLists,
		DeleteContext: deleteTrunksIpAccessControlLists,
		Schema: map[string]*schema.Schema{
			"trunk_sid":                  AsString(SchemaForceNewRequired),
			"ip_access_control_list_sid": AsString(SchemaForceNewRequired),
			"sid":                        AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrunksIpAccessControlListsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrunksIpAccessControlLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateIpAccessControlListParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	trunkSid := d.Get("trunk_sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.CreateIpAccessControlList(trunkSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{trunkSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteTrunksIpAccessControlLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrunkingV1.DeleteIpAccessControlList(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrunksIpAccessControlLists(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.FetchIpAccessControlList(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrunksIpAccessControlListsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected trunk_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("trunk_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceTrunksOriginationUrls() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrunksOriginationUrls,
		ReadContext:   readTrunksOriginationUrls,
		UpdateContext: updateTrunksOriginationUrls,
		DeleteContext: deleteTrunksOriginationUrls,
		Schema: map[string]*schema.Schema{
			"trunk_sid":     AsString(SchemaRequired),
			"enabled":       AsBool(SchemaRequired),
			"friendly_name": AsString(SchemaRequired),
			"priority":      AsInt(SchemaRequired),
			"sip_url":       AsString(SchemaRequired),
			"weight":        AsInt(SchemaRequired),
			"sid":           AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrunksOriginationUrlsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrunksOriginationUrls(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateOriginationUrlParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	trunkSid := d.Get("trunk_sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.CreateOriginationUrl(trunkSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{trunkSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteTrunksOriginationUrls(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrunkingV1.DeleteOriginationUrl(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrunksOriginationUrls(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.FetchOriginationUrl(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrunksOriginationUrlsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected trunk_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("trunk_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateTrunksOriginationUrls(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateOriginationUrlParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.UpdateOriginationUrl(trunkSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func ResourceTrunksPhoneNumbers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrunksPhoneNumbers,
		ReadContext:   readTrunksPhoneNumbers,
		DeleteContext: deleteTrunksPhoneNumbers,
		Schema: map[string]*schema.Schema{
			"trunk_sid":        AsString(SchemaForceNewRequired),
			"phone_number_sid": AsString(SchemaForceNewRequired),
			"sid":              AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrunksPhoneNumbersImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrunksPhoneNumbers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreatePhoneNumberParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	trunkSid := d.Get("trunk_sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.CreatePhoneNumber(trunkSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{trunkSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteTrunksPhoneNumbers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrunkingV1.DeletePhoneNumber(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrunksPhoneNumbers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	trunkSid := d.Get("trunk_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.FetchPhoneNumber(trunkSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrunksPhoneNumbersImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected trunk_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("trunk_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceTrunks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createTrunks,
		ReadContext:   readTrunks,
		UpdateContext: updateTrunks,
		DeleteContext: deleteTrunks,
		Schema: map[string]*schema.Schema{
			"cnam_lookup_enabled":      AsBool(SchemaComputedOptional),
			"disaster_recovery_method": AsString(SchemaComputedOptional),
			"disaster_recovery_url":    AsString(SchemaComputedOptional),
			"domain_name":              AsString(SchemaComputedOptional),
			"friendly_name":            AsString(SchemaComputedOptional),
			"secure":                   AsBool(SchemaComputedOptional),
			"transfer_caller_id":       AsString(SchemaComputedOptional),
			"transfer_mode":            AsString(SchemaComputedOptional),
			"sid":                      AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseTrunksImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateTrunkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.TrunkingV1.CreateTrunk(&params)
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

func deleteTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.TrunkingV1.DeleteTrunk(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.FetchTrunk(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseTrunksImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected sid"

	if len(importParts) != 1 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("sid", importParts[0])

	return nil
}
func updateTrunks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateTrunkParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.TrunkingV1.UpdateTrunk(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
