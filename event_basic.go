package fomo

type EventBasic struct {
	EventTypeId  int               `json:"event_type_id,omitempty"`
	EventTypeTag int               `json:"event_type_tag,omitempty"`
	Message      string            `json:"message,omitempty"`
	Url          string            `json:"url,omitempty"`
	FirstName    string            `json:"first_name,omitempty"`
	EmailAddress string            `json:"emmail_address,omitempty"`
	IpAddress    string            `json:"ip_address,omitempty"`
	City         string            `json:"city,omitempty"`
	Province     string            `json:"province,omitempty"`
	Country      string            `json:"country,omitempty"`
	Title        string            `json:"title,omitempty"`
	ImageUrl     string            `json:"image_url,omitempty"`
	ExternalID   string            `json:"external_id,omitempty"`
	Attributes   map[string]string `json:"custom_event_fields_attributes"`
}
