package Todo

type Item struct {
	Timestamp string `json:"timestamp" dynamodbav:"timestamp"`
	Content   string `json:"content"   dynamodbav:"content"`
}
