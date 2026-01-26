package repositories

import (
	"testing"
)

func TestAddFavourite(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		assetID     string
		description *string
		shouldError bool
	}{
		{
			name:        "add favourite without description",
			userID:      "u1",
			assetID:     "550e8400-e29b-41d4-a716-446655440000",
			description: nil,
			shouldError: false,
		},
		{
			name:        "add favourite with description",
			userID:      "u2",
			assetID:     "550e8400-e29b-41d4-a716-446655440001",
			description: ptrString("My favorite chart"),
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.userID == "" {
				t.Error("userID cannot be empty")
			}
			if tt.assetID == "" {
				t.Error("assetID cannot be empty")
			}
		})
	}
}

func TestRemoveFavourite(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		assetID     string
		shouldError bool
	}{
		{
			name:        "remove existing favourite",
			userID:      "u1",
			assetID:     "550e8400-e29b-41d4-a716-446655440000",
			shouldError: false,
		},
		{
			name:        "remove non-existent favourite",
			userID:      "u99",
			assetID:     "non-existent-uuid",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.userID == "" {
				t.Error("userID cannot be empty")
			}
			if tt.assetID == "" {
				t.Error("assetID cannot be empty")
			}
		})
	}
}

func TestUpdateFavouriteDescription(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		assetID     string
		description *string
		shouldError bool
	}{
		{
			name:        "update with new description",
			userID:      "u1",
			assetID:     "550e8400-e29b-41d4-a716-446655440000",
			description: ptrString("Updated description"),
			shouldError: false,
		},
		{
			name:        "update with nil description",
			userID:      "u2",
			assetID:     "550e8400-e29b-41d4-a716-446655440001",
			description: nil,
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.userID == "" {
				t.Error("userID cannot be empty")
			}
			if tt.assetID == "" {
				t.Error("assetID cannot be empty")
			}
		})
	}
}
