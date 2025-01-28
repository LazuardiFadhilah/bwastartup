package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

type Transaction struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	UserID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (t Transaction) FormatPayment(transaction Transaction) interface{} {
	panic("unimplemented")
}

func (t Transaction) FormatTransaction(transaction Transaction) interface{} {
	panic("unimplemented")
}
