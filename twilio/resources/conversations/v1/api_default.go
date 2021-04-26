/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.14.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/twilio/terraform-provider-twilio/client"
	. "github.com/twilio/terraform-provider-twilio/twilio/common"
	. "github.com/twilio/twilio-go/twilio/rest/conversations/v1"
)

func ResourceServicesConversations() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesConversations,
		ReadContext:   readServicesConversations,
		UpdateContext: updateServicesConversations,
		DeleteContext: deleteServicesConversations,
		Schema: map[string]*schema.Schema{
			"chat_service_sid":         AsString(SchemaRequired),
			"x-twilio-webhook-enabled": AsString(SchemaOptional),
			"attributes":               AsString(SchemaOptional),
			"date_created":             AsString(SchemaOptional),
			"date_updated":             AsString(SchemaOptional),
			"friendly_name":            AsString(SchemaOptional),
			"messaging_service_sid":    AsString(SchemaOptional),
			"state":                    AsString(SchemaOptional),
			"timers.closed":            AsString(SchemaOptional),
			"timers.inactive":          AsString(SchemaOptional),
			"unique_name":              AsString(SchemaOptional),
			"account_sid":              AsString(SchemaComputed),
			"links":                    AsString(SchemaComputed),
			"sid":                      AsString(SchemaComputed),
			"timers":                   AsString(SchemaComputed),
			"url":                      AsString(SchemaComputed),
		},
	}
}

func createServicesConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceConversationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateServiceConversation(chatServiceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteServiceConversation(chatServiceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchServiceConversation(chatServiceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceConversationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateServiceConversation(chatServiceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceConversationsMessages() *schema.Resource {
	return &schema.Resource{
		CreateContext: createConversationsMessages,
		ReadContext:   readConversationsMessages,
		UpdateContext: updateConversationsMessages,
		DeleteContext: deleteConversationsMessages,
		Schema: map[string]*schema.Schema{
			"conversation_sid":         AsString(SchemaRequired),
			"x-twilio-webhook-enabled": AsString(SchemaOptional),
			"attributes":               AsString(SchemaOptional),
			"author":                   AsString(SchemaOptional),
			"body":                     AsString(SchemaOptional),
			"date_created":             AsString(SchemaOptional),
			"date_updated":             AsString(SchemaOptional),
			"media_sid":                AsString(SchemaOptional),
			"account_sid":              AsString(SchemaComputed),
			"delivery":                 AsString(SchemaComputed),
			"index":                    AsString(SchemaComputed),
			"links":                    AsString(SchemaComputed),
			"media":                    AsString(SchemaComputed),
			"participant_sid":          AsString(SchemaComputed),
			"sid":                      AsString(SchemaComputed),
			"url":                      AsString(SchemaComputed),
		},
	}
}

func createConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateConversationMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	conversationSid := d.Get("conversation_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateConversationMessage(conversationSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteConversationMessage(conversationSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchConversationMessage(conversationSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateConversationMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateConversationMessage(conversationSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesConversationsParticipants() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesConversationsParticipants,
		ReadContext:   readServicesConversationsParticipants,
		UpdateContext: updateServicesConversationsParticipants,
		DeleteContext: deleteServicesConversationsParticipants,
		Schema: map[string]*schema.Schema{
			"chat_service_sid":                    AsString(SchemaRequired),
			"conversation_sid":                    AsString(SchemaRequired),
			"x-twilio-webhook-enabled":            AsString(SchemaOptional),
			"attributes":                          AsString(SchemaOptional),
			"date_created":                        AsString(SchemaOptional),
			"date_updated":                        AsString(SchemaOptional),
			"identity":                            AsString(SchemaOptional),
			"messaging_binding.address":           AsString(SchemaOptional),
			"messaging_binding.projected_address": AsString(SchemaOptional),
			"messaging_binding.proxy_address":     AsString(SchemaOptional),
			"role_sid":                            AsString(SchemaOptional),
			"account_sid":                         AsString(SchemaComputed),
			"last_read_message_index":             AsString(SchemaComputed),
			"last_read_timestamp":                 AsString(SchemaComputed),
			"messaging_binding":                   AsString(SchemaComputed),
			"sid":                                 AsString(SchemaComputed),
			"url":                                 AsString(SchemaComputed),
		},
	}
}

func createServicesConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceConversationParticipantParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateServiceConversationParticipant(chatServiceSid, conversationSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteServiceConversationParticipant(chatServiceSid, conversationSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchServiceConversationParticipant(chatServiceSid, conversationSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceConversationParticipantParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateServiceConversationParticipant(chatServiceSid, conversationSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesConversationsWebhooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesConversationsWebhooks,
		ReadContext:   readServicesConversationsWebhooks,
		UpdateContext: updateServicesConversationsWebhooks,
		DeleteContext: deleteServicesConversationsWebhooks,
		Schema: map[string]*schema.Schema{
			"chat_service_sid":           AsString(SchemaRequired),
			"conversation_sid":           AsString(SchemaRequired),
			"target":                     AsString(SchemaRequired),
			"configuration.filters":      AsString(SchemaOptional),
			"configuration.flow_sid":     AsString(SchemaOptional),
			"configuration.method":       AsString(SchemaOptional),
			"configuration.replay_after": AsString(SchemaOptional),
			"configuration.triggers":     AsString(SchemaOptional),
			"configuration.url":          AsString(SchemaOptional),
			"account_sid":                AsString(SchemaComputed),
			"configuration":              AsString(SchemaComputed),
			"date_created":               AsString(SchemaComputed),
			"date_updated":               AsString(SchemaComputed),
			"sid":                        AsString(SchemaComputed),
			"url":                        AsString(SchemaComputed),
		},
	}
}

func createServicesConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceConversationScopedWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateServiceConversationScopedWebhook(chatServiceSid, conversationSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteServiceConversationScopedWebhook(chatServiceSid, conversationSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchServiceConversationScopedWebhook(chatServiceSid, conversationSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceConversationScopedWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateServiceConversationScopedWebhook(chatServiceSid, conversationSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createRoles,
		ReadContext:   readRoles,
		UpdateContext: updateRoles,
		DeleteContext: deleteRoles,
		Schema: map[string]*schema.Schema{
			"friendly_name":    AsString(SchemaRequired),
			"permission":       AsString(SchemaRequired),
			"type":             AsString(SchemaRequired),
			"account_sid":      AsString(SchemaComputed),
			"chat_service_sid": AsString(SchemaComputed),
			"date_created":     AsString(SchemaComputed),
			"date_updated":     AsString(SchemaComputed),
			"permissions":      AsString(SchemaComputed),
			"sid":              AsString(SchemaComputed),
			"url":              AsString(SchemaComputed),
		},
	}
}

func createRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.ConversationsV1.CreateRole(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteRole(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchRole(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateRole(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createUsers,
		ReadContext:   readUsers,
		UpdateContext: updateUsers,
		DeleteContext: deleteUsers,
		Schema: map[string]*schema.Schema{
			"identity":                 AsString(SchemaRequired),
			"x-twilio-webhook-enabled": AsString(SchemaOptional),
			"attributes":               AsString(SchemaOptional),
			"friendly_name":            AsString(SchemaOptional),
			"role_sid":                 AsString(SchemaOptional),
			"account_sid":              AsString(SchemaComputed),
			"chat_service_sid":         AsString(SchemaComputed),
			"date_created":             AsString(SchemaComputed),
			"date_updated":             AsString(SchemaComputed),
			"is_notifiable":            AsString(SchemaComputed),
			"is_online":                AsString(SchemaComputed),
			"sid":                      AsString(SchemaComputed),
			"url":                      AsString(SchemaComputed),
		},
	}
}

func createUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.ConversationsV1.CreateUser(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteUser(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchUser(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateUser(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesRoles() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesRoles,
		ReadContext:   readServicesRoles,
		UpdateContext: updateServicesRoles,
		DeleteContext: deleteServicesRoles,
		Schema: map[string]*schema.Schema{
			"chat_service_sid": AsString(SchemaRequired),
			"friendly_name":    AsString(SchemaRequired),
			"permission":       AsString(SchemaRequired),
			"type":             AsString(SchemaRequired),
			"account_sid":      AsString(SchemaComputed),
			"date_created":     AsString(SchemaComputed),
			"date_updated":     AsString(SchemaComputed),
			"permissions":      AsString(SchemaComputed),
			"sid":              AsString(SchemaComputed),
			"url":              AsString(SchemaComputed),
		},
	}
}

func createServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateServiceRole(chatServiceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteServiceRole(chatServiceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchServiceRole(chatServiceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesRoles(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceRoleParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateServiceRole(chatServiceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceConversations() *schema.Resource {
	return &schema.Resource{
		CreateContext: createConversations,
		ReadContext:   readConversations,
		UpdateContext: updateConversations,
		DeleteContext: deleteConversations,
		Schema: map[string]*schema.Schema{
			"x-twilio-webhook-enabled": AsString(SchemaOptional),
			"attributes":               AsString(SchemaOptional),
			"date_created":             AsString(SchemaOptional),
			"date_updated":             AsString(SchemaOptional),
			"friendly_name":            AsString(SchemaOptional),
			"messaging_service_sid":    AsString(SchemaOptional),
			"state":                    AsString(SchemaOptional),
			"timers.closed":            AsString(SchemaOptional),
			"timers.inactive":          AsString(SchemaOptional),
			"unique_name":              AsString(SchemaOptional),
			"account_sid":              AsString(SchemaComputed),
			"chat_service_sid":         AsString(SchemaComputed),
			"links":                    AsString(SchemaComputed),
			"sid":                      AsString(SchemaComputed),
			"timers":                   AsString(SchemaComputed),
			"url":                      AsString(SchemaComputed),
		},
	}
}

func createConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateConversationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.ConversationsV1.CreateConversation(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteConversation(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchConversation(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateConversations(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateConversationParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateConversation(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesConversationsMessages() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesConversationsMessages,
		ReadContext:   readServicesConversationsMessages,
		UpdateContext: updateServicesConversationsMessages,
		DeleteContext: deleteServicesConversationsMessages,
		Schema: map[string]*schema.Schema{
			"chat_service_sid":         AsString(SchemaRequired),
			"conversation_sid":         AsString(SchemaRequired),
			"x-twilio-webhook-enabled": AsString(SchemaOptional),
			"attributes":               AsString(SchemaOptional),
			"author":                   AsString(SchemaOptional),
			"body":                     AsString(SchemaOptional),
			"date_created":             AsString(SchemaOptional),
			"date_updated":             AsString(SchemaOptional),
			"media_sid":                AsString(SchemaOptional),
			"account_sid":              AsString(SchemaComputed),
			"delivery":                 AsString(SchemaComputed),
			"index":                    AsString(SchemaComputed),
			"links":                    AsString(SchemaComputed),
			"media":                    AsString(SchemaComputed),
			"participant_sid":          AsString(SchemaComputed),
			"sid":                      AsString(SchemaComputed),
			"url":                      AsString(SchemaComputed),
		},
	}
}

func createServicesConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceConversationMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateServiceConversationMessage(chatServiceSid, conversationSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteServiceConversationMessage(chatServiceSid, conversationSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchServiceConversationMessage(chatServiceSid, conversationSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesConversationsMessages(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceConversationMessageParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateServiceConversationMessage(chatServiceSid, conversationSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceConversationsParticipants() *schema.Resource {
	return &schema.Resource{
		CreateContext: createConversationsParticipants,
		ReadContext:   readConversationsParticipants,
		UpdateContext: updateConversationsParticipants,
		DeleteContext: deleteConversationsParticipants,
		Schema: map[string]*schema.Schema{
			"conversation_sid":                    AsString(SchemaRequired),
			"x-twilio-webhook-enabled":            AsString(SchemaOptional),
			"attributes":                          AsString(SchemaOptional),
			"date_created":                        AsString(SchemaOptional),
			"date_updated":                        AsString(SchemaOptional),
			"identity":                            AsString(SchemaOptional),
			"messaging_binding.address":           AsString(SchemaOptional),
			"messaging_binding.projected_address": AsString(SchemaOptional),
			"messaging_binding.proxy_address":     AsString(SchemaOptional),
			"role_sid":                            AsString(SchemaOptional),
			"account_sid":                         AsString(SchemaComputed),
			"last_read_message_index":             AsString(SchemaComputed),
			"last_read_timestamp":                 AsString(SchemaComputed),
			"messaging_binding":                   AsString(SchemaComputed),
			"sid":                                 AsString(SchemaComputed),
			"url":                                 AsString(SchemaComputed),
		},
	}
}

func createConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateConversationParticipantParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	conversationSid := d.Get("conversation_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateConversationParticipant(conversationSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteConversationParticipant(conversationSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchConversationParticipant(conversationSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateConversationsParticipants(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateConversationParticipantParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateConversationParticipant(conversationSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceServicesUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: createServicesUsers,
		ReadContext:   readServicesUsers,
		UpdateContext: updateServicesUsers,
		DeleteContext: deleteServicesUsers,
		Schema: map[string]*schema.Schema{
			"chat_service_sid":         AsString(SchemaRequired),
			"identity":                 AsString(SchemaRequired),
			"x-twilio-webhook-enabled": AsString(SchemaOptional),
			"attributes":               AsString(SchemaOptional),
			"friendly_name":            AsString(SchemaOptional),
			"role_sid":                 AsString(SchemaOptional),
			"account_sid":              AsString(SchemaComputed),
			"date_created":             AsString(SchemaComputed),
			"date_updated":             AsString(SchemaComputed),
			"is_notifiable":            AsString(SchemaComputed),
			"is_online":                AsString(SchemaComputed),
			"sid":                      AsString(SchemaComputed),
			"url":                      AsString(SchemaComputed),
		},
	}
}

func createServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateServiceUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateServiceUser(chatServiceSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteServiceUser(chatServiceSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchServiceUser(chatServiceSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateServicesUsers(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateServiceUserParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	chatServiceSid := d.Get("chat_service_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateServiceUser(chatServiceSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceConversationsWebhooks() *schema.Resource {
	return &schema.Resource{
		CreateContext: createConversationsWebhooks,
		ReadContext:   readConversationsWebhooks,
		UpdateContext: updateConversationsWebhooks,
		DeleteContext: deleteConversationsWebhooks,
		Schema: map[string]*schema.Schema{
			"conversation_sid":           AsString(SchemaRequired),
			"target":                     AsString(SchemaRequired),
			"configuration.filters":      AsString(SchemaOptional),
			"configuration.flow_sid":     AsString(SchemaOptional),
			"configuration.method":       AsString(SchemaOptional),
			"configuration.replay_after": AsString(SchemaOptional),
			"configuration.triggers":     AsString(SchemaOptional),
			"configuration.url":          AsString(SchemaOptional),
			"account_sid":                AsString(SchemaComputed),
			"configuration":              AsString(SchemaComputed),
			"date_created":               AsString(SchemaComputed),
			"date_updated":               AsString(SchemaComputed),
			"sid":                        AsString(SchemaComputed),
			"url":                        AsString(SchemaComputed),
		},
	}
}

func createConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateConversationScopedWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	conversationSid := d.Get("conversation_sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.CreateConversationScopedWebhook(conversationSid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteConversationScopedWebhook(conversationSid, sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchConversationScopedWebhook(conversationSid, sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateConversationsWebhooks(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateConversationScopedWebhookParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	conversationSid := d.Get("conversation_sid").(string)
	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateConversationScopedWebhook(conversationSid, sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func ResourceCredentials() *schema.Resource {
	return &schema.Resource{
		CreateContext: createCredentials,
		ReadContext:   readCredentials,
		UpdateContext: updateCredentials,
		DeleteContext: deleteCredentials,
		Schema: map[string]*schema.Schema{
			"type":          AsString(SchemaRequired),
			"api_key":       AsString(SchemaOptional),
			"certificate":   AsString(SchemaOptional),
			"friendly_name": AsString(SchemaOptional),
			"private_key":   AsString(SchemaOptional),
			"sandbox":       AsString(SchemaOptional),
			"secret":        AsString(SchemaOptional),
			"account_sid":   AsString(SchemaComputed),
			"date_created":  AsString(SchemaComputed),
			"date_updated":  AsString(SchemaComputed),
			"sid":           AsString(SchemaComputed),
			"url":           AsString(SchemaComputed),
		},
	}
}

func createCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := CreateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	r, err := m.(*client.Config).Client.ConversationsV1.CreateCredential(&params)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*r.Sid)
	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func deleteCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	err := m.(*client.Config).Client.ConversationsV1.DeleteCredential(sid)
	d.SetId("")

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func readCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.FetchCredential(sid)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func updateCredentials(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	params := UpdateCredentialParams{}
	if err := UnmarshalSchema(&params, d); err != nil {
		return diag.FromErr(err)
	}

	sid := d.Get("sid").(string)

	r, err := m.(*client.Config).Client.ConversationsV1.UpdateCredential(sid, &params)
	if err != nil {
		return diag.FromErr(err)
	}

	err = MarshalSchema(d, r)

	if err != nil {
		return diag.FromErr(err)
	}
	return nil
}
