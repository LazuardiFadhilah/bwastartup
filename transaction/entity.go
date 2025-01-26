package transaction

import "time"

type Transaction struct {
	ID         int    `json:"id"`
	CampaignID int    `json:"campaign_id"`
	userID     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
