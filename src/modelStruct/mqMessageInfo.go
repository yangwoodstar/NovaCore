package modelStruct

type QueueInfo struct {
	Messages                  int    `json:"messages"`
	MessagesPagedOut          int    `json:"messages_paged_out"`
	MessagesPersistent        int    `json:"messages_persistent"`
	MessagesRAM               int    `json:"messages_ram"`
	MessagesReady             int    `json:"messages_ready"`
	MessagesReadyRAM          int    `json:"messages_ready_ram"`
	MessagesUnacknowledged    int    `json:"messages_unacknowledged"`
	MessagesUnacknowledgedRAM int    `json:"messages_unacknowledged_ram"`
	Name                      string `json:"name"`
}

type RoutingBinding struct {
	Source          string                 `json:"source"`
	Vhost           string                 `json:"vhost"`
	Destination     string                 `json:"destination"`
	DestinationType string                 `json:"destination_type"`
	RoutingKey      string                 `json:"routing_key"`
	Arguments       map[string]interface{} `json:"arguments"`
	PropertiesKey   string                 `json:"properties_key"`
}
