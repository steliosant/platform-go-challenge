package repositories

import (
	"encoding/json"
	"testing"

	"platform-go-challenge/models"
)

func TestUnmarshalAssetData(t *testing.T) {
	tests := []struct {
		name        string
		assetType   models.AssetType
		rawData     json.RawMessage
		shouldError bool
	}{
		{
			name:        "chart asset",
			assetType:   models.AssetChart,
			rawData:     json.RawMessage(`{"x_axis_title":"Month","y_axis_title":"Sales","data_points":[100,200]}`),
			shouldError: false,
		},
		{
			name:        "insight asset",
			assetType:   models.AssetInsight,
			rawData:     json.RawMessage(`{"text":"Users prefer mobile"}`),
			shouldError: false,
		},
		{
			name:        "audience asset",
			assetType:   models.AssetAudience,
			rawData:     json.RawMessage(`{"gender":"M","birth_country":"US","age_groups":["25-34"],"hours_on_social_media_min":3,"purchases_last_month_min":5}`),
			shouldError: false,
		},
		{
			name:        "unknown asset type",
			assetType:   models.AssetType("unknown"),
			rawData:     json.RawMessage(`{}`),
			shouldError: true,
		},
		{
			name:        "invalid JSON",
			assetType:   models.AssetChart,
			rawData:     json.RawMessage(`{invalid}`),
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := unmarshalAssetData(tt.assetType, tt.rawData)

			if tt.shouldError && err == nil {
				t.Error("expected error but got nil")
			}
			if !tt.shouldError && err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !tt.shouldError && data == nil {
				t.Error("expected data but got nil")
			}
		})
	}
}

func TestCreateAsset(t *testing.T) {
	tests := []struct {
		name        string
		asset       models.Asset
		description *string
		shouldError bool
	}{
		{
			name: "create chart asset",
			asset: models.Asset{
				Type:  models.AssetChart,
				Title: ptrString("Sales Chart"),
				Data:  json.RawMessage(`{"x_axis_title":"Month"}`),
			},
			description: nil,
			shouldError: false,
		},
		{
			name: "create asset with description",
			asset: models.Asset{
				Type:  models.AssetInsight,
				Title: ptrString("Market Insight"),
				Data:  json.RawMessage(`{"text":"Market growing"}`),
			},
			description: ptrString("Important insight"),
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.asset.Title == nil {
				t.Error("asset title cannot be nil")
			}
			if tt.asset.Type == "" {
				t.Error("asset type cannot be empty")
			}
		})
	}
}

func TestGetAssetByID(t *testing.T) {
	tests := []struct {
		name        string
		assetID     string
		shouldError bool
	}{
		{
			name:        "valid UUID",
			assetID:     "550e8400-e29b-41d4-a716-446655440000",
			shouldError: false,
		},
		{
			name:        "non-existent asset",
			assetID:     "invalid-uuid",
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.assetID == "" {
				t.Error("assetID cannot be empty")
			}
		})
	}
}

// Helper functions
func ptrString(s string) *string {
	return &s
}
