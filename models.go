package main

/*
MessageRequest represents the payload required to
simulate sending a LinkedIn message.

Fields:
- ProfileURL: URL of the target LinkedIn profile
- Message: Text content of the message to be sent
*/
type MessageRequest struct {
	ProfileURL string `json:"profile_url"`
	Message    string `json:"message"`
}

/*
ConnectionRequest represents the payload required to
simulate sending a LinkedIn connection request.

Fields:
- ProfileURL: URL of the target LinkedIn profile
- Note: Optional personalized note included with the request
*/
type ConnectionRequest struct {
	ProfileURL string `json:"profile_url"`
	Note       string `json:"note"`
}
