package transaction

import (
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
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
