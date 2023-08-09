package models

// WebhookEvent represents the event data within the incoming webhook
type WebhookEvent struct {
	Network  string         `json:"network"`
	Activity []ActivityItem `json:"activity"`
}

// ActivityItem represents each individual activity within the webhook event
type ActivityItem struct {
	BlockNum         string      `json:"blockNum"`
	Hash             string      `json:"hash"`
	FromAddress      string      `json:"fromAddress"`
	ToAddress        string      `json:"toAddress"`
	Value            float64     `json:"value"`
	ERC721TokenId    interface{} `json:"erc721TokenId"`
	ERC1155Metadata  interface{} `json:"erc1155Metadata"`
	Asset            string      `json:"asset"`
	Category         string      `json:"category"`
	RawContract      RawContract `json:"rawContract"`
	TypeTraceAddress interface{} `json:"typeTraceAddress"`
	Log              LogItem     `json:"log"`
}

// RawContract represents the raw contract data within each activity item
type RawContract struct {
	RawValue string `json:"rawValue"`
	Address  string `json:"address"`
	Decimals int    `json:"decimals"`
}

// LogItem represents the log data within each activity item
type LogItem struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
}

// WebhookData represents the entire incoming webhook payload
type WebhookData struct {
	WebhookID string       `json:"webhookId"`
	ID        string       `json:"id"`
	CreatedAt string       `json:"createdAt"`
	Type      string       `json:"type"`
	Event     WebhookEvent `json:"event"`
}
