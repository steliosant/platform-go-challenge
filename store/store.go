package store

import "platform-go-challenge/models"

func ptr(s string) *string {
	return &s
}

var Favourites = []models.Favourite{
	{
		UserID:      "u1",
		AssetID:     "a1",
		Description: ptr("Key chart for weekly review"),
	},
	{
		UserID:      "u1",
		AssetID:     "a3",
		Description: ptr("Target audience for campaign"),
	},
}
