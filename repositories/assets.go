package repositories

import (
	"encoding/json"
	"fmt"

	"platform-go-challenge/models"
)

func unmarshalAssetData(
	assetType models.AssetType,
	raw json.RawMessage,
) (any, error) {

	switch assetType {
	case models.AssetChart:
		var c struct {
			XAxisTitle string    `json:"x_axis_title"`
			YAxisTitle string    `json:"y_axis_title"`
			DataPoints []float64 `json:"data_points"`
		}
		return c, json.Unmarshal(raw, &c)

	case models.AssetInsight:
		var i struct {
			Text string `json:"text"`
		}
		return i, json.Unmarshal(raw, &i)

	case models.AssetAudience:
		var a struct {
			Gender                string   `json:"gender"`
			BirthCountry          string   `json:"birth_country"`
			AgeGroups             []string `json:"age_groups"`
			HoursOnSocialMediaMin int      `json:"hours_on_social_media_min"`
			PurchasesLastMonthMin int      `json:"purchases_last_month_min"`
		}
		return a, json.Unmarshal(raw, &a)

	default:
		return nil, fmt.Errorf("unknown asset type: %s", assetType)
	}
}

