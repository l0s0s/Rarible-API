package rarible_test

import (
	"fmt"
	"l0s0s/Rarible-API/model"
	"l0s0s/Rarible-API/rarible"
	"testing"

	"github.com/stretchr/testify/assert"
)

const TestApiKey = "11111111-1111-1111-1111-111111111111"
const TestReferer = "https://docs.rarible.org"

func TestGetNFTOwnership(t *testing.T) {
	client := rarible.NewClient(TestApiKey, TestReferer)

	for _, tc := range []struct {
		testName          string
		id                string
		expectedOwnership model.Ownership
		expectedError     error
	}{
		{
			testName: "successfully get nft ownership",
			id:       "ETHEREUM%3A0xb66a603f4cfe17e3d27b87a8bfcad319856518b8%3A32292934596187112148346015918544186536963932779440027682601542850818403729410%3A0x4765273c477c2dc484da4f1984639e943adccfeb",
			expectedOwnership: model.Ownership{
				ID:            "ETHEREUM:0xb66a603f4cfe17e3d27b87a8bfcad319856518b8:32292934596187112148346015918544186536963932779440027682601542850818403729410:0x4765273c477c2dc484da4f1984639e943adccfeb",
				Blockchain:    "ETHEREUM",
				ItemID:        "ETHEREUM:0xb66a603f4cfe17e3d27b87a8bfcad319856518b8:32292934596187112148346015918544186536963932779440027682601542850818403729410",
				Contract:      "ETHEREUM:0xb66a603f4cfe17e3d27b87a8bfcad319856518b8",
				Collection:    "ETHEREUM:0xb66a603f4cfe17e3d27b87a8bfcad319856518b8",
				TokenID:       "32292934596187112148346015918544186536963932779440027682601542850818403729410",
				Owner:         "ETHEREUM:0x4765273c477c2dc484da4f1984639e943adccfeb",
				Value:         "21",
				Source:        "",
				CreatedAt:     "2022-04-15T10:59:03Z",
				LastUpdatedAt: "2024-02-19T11:47:36.262Z",
				Creators:      []model.Creator{{Account: "ETHEREUM:0x4765273c477c2dc484da4f1984639e943adccfeb", Value: 10000}},
				LazyValue:     "0",
				Version:       0,
			},
		},
		{
			testName:      "fail to get nft ownership: invalid ID format",
			id:            "invalid-id",
			expectedError: fmt.Errorf("Illegal format for ID: 'invalid-id', blockchain prefix not found"),
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			resultOwnership, err := client.GetNFTOwnership(tc.id)
			assert.Equal(t, tc.expectedOwnership, resultOwnership)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetNFTTraitsRarity(t *testing.T) {
	client := rarible.NewClient(TestApiKey, TestReferer)

	for _, tc := range []struct {
		testName       string
		collectionID   string
		properties     []model.TraitProperty
		expectedRarity model.TraitsRarity
		expectedError  error
	}{
		{
			testName:     "successfully get nft traits rarity",
			collectionID: "ETHEREUM:0x60e4d786628fea6478f785a6d7e704777c86a7c6",
			properties: []model.TraitProperty{
				{Key: "Key1", Value: "Value1"},
			},
			expectedRarity: model.TraitsRarity{
				Traits: []model.TraitRarity{
					{Key: "Key1", Value: "Value1", Rarity: "0"},
				},
			},
		},
		{
			testName:      "fail to get nft traits rarity: invalid collection ID format",
			collectionID:  "invalid-collection-id",
			properties:    []model.TraitProperty{{Key: "Key1", Value: "Value1"}},
			expectedError: fmt.Errorf("Unexpected server error"),
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			resultRarity, err := client.GetNFTTraitsRarity(tc.collectionID, tc.properties)
			assert.Equal(t, tc.expectedRarity, resultRarity)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
