package twilio

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	client "github.com/twilio/twilio-go/client"
)

// Provider initializes terraform-provider-twilio
func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_sid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "You Account SID can be found on the Twilio dashboard.",
			},
			"auth_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "You Auth Token can be found on the Twilio dashboard.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"twilio_chat_service": resourceChatService(),
		},
	}

	p.ConfigureFunc = providerClient(p)

	return p
}

type Config struct {
	AccountSID string
	AuthToken  string
	Client     *client.Twilio
}

func providerClient(p *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		config := &Config{
			AccountSID: d.Get("account_sid").(string),
			AuthToken:  d.Get("auth_token").(string),
			Client:     client.NewClient(d.Get("account_sid").(string), d.Get("auth_token").(string)),
		}

		return config, nil
	}
}
