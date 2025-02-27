/*
 * Twilio - Studio
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
	. "github.com/twilio/twilio-go/rest/studio/v1"
)

func ResourceFlowsEngagements() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlowsEngagements,
		ReadContext:   readFlowsEngagements,
		DeleteContext: deleteFlowsEngagements,
		Schema: map[string]*schema.Schema{
			"flow_sid":   AsString(SchemaForceNewRequired),
			"from":       AsString(SchemaForceNewRequired),
			"to":         AsString(SchemaForceNewRequired),
			"parameters": AsString(SchemaForceNewOptional),
			"sid":        AsString(SchemaComputed),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseFlowsEngagementsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createFlowsEngagements(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateEngagementParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	flowSid := d.Get("flow_sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.CreateEngagement(flowSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{flowSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func deleteFlowsEngagements(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.StudioV1.DeleteEngagement(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlowsEngagements(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.FetchEngagement(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseFlowsEngagementsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected flow_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("flow_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func ResourceFlowsExecutions() *schema.Resource {
	return &schema.Resource{
		CreateContext: createFlowsExecutions,
		ReadContext:   readFlowsExecutions,
		UpdateContext: updateFlowsExecutions,
		DeleteContext: deleteFlowsExecutions,
		Schema: map[string]*schema.Schema{
			"flow_sid":   AsString(SchemaRequired),
			"from":       AsString(SchemaRequired),
			"to":         AsString(SchemaRequired),
			"parameters": AsString(SchemaComputedOptional),
			"sid":        AsString(SchemaComputed),
			"status":     AsString(SchemaComputedOptional),
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := parseFlowsExecutionsImportId(d.Id(), d)
				if err != nil {
					return nil, err
				}

				return []*schema.ResourceData{d}, nil
			},
		},
	}
}

func createFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateExecutionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	flowSid := d.Get("flow_sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.CreateExecution(flowSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	idParts := []string{flowSid}
	idParts = append(idParts, (*r.Sid))
	d.SetId(strings.Join(idParts, "/"))
	d.Set("sid", *r.Sid)

	return updateFlowsExecutions(ctx, d, m)
}

func deleteFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.StudioV1.DeleteExecution(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")

	return nil
}

func readFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.FetchExecution(flowSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func parseFlowsExecutionsImportId(importId string, d *schema.ResourceData) error {
	importParts := strings.Split(importId, "/")
	errStr := "invalid import ID (%q), expected flow_sid/sid"

	if len(importParts) != 2 {
		return fmt.Errorf(errStr, importId)
	}

	d.Set("flow_sid", importParts[0])
	d.Set("sid", importParts[1])

	return nil
}
func updateFlowsExecutions(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateExecutionParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	flowSid := d.Get("flow_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.StudioV1.UpdateExecution(flowSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)
	if err != nil {
		return diag.FromErr(err)
	}

	return nil
}
