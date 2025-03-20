package model

type LikeEvent struct {
	UserID int `json:"user_id"`
	ItemID int `json:"item_id"`
}
