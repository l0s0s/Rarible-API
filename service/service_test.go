package service_test

import (
	"fmt"
	"l0s0s/Rarible-API/model"
	"l0s0s/Rarible-API/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockNFTClient struct {
	GetNFTOwnershipFunc    func(id string) (model.Ownership, error)
	GetNFTTraitsRarityFunc func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error)
}

func (m *MockNFTClient) GetNFTOwnership(id string) (model.Ownership, error) {
	if m.GetNFTOwnershipFunc != nil {
		return m.GetNFTOwnershipFunc(id)
	}
	return model.Ownership{}, nil
}

func (m *MockNFTClient) GetNFTTraitsRarity(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
	if m.GetNFTTraitsRarityFunc != nil {
		return m.GetNFTTraitsRarityFunc(collectionID, properties)
	}
	return model.TraitsRarity{}, nil
}

func TestGetNFTOwnership(t *testing.T) {
	for _, tc := range []struct {
		testName            string
		getNFTOwnershipFunc func(id string) (model.Ownership, error)
		expectedOwnership   model.Ownership
		expectedError       error
	}{
		{
			testName: "successfully get nft ownership",
			getNFTOwnershipFunc: func(id string) (model.Ownership, error) {
				return model.Ownership{
					ID: "id1",
				}, nil
			},
			expectedOwnership: model.Ownership{
				ID: "id1",
			},
			expectedError: nil,
		},
		{
			testName: "fail to get nft ownership",
			getNFTOwnershipFunc: func(id string) (model.Ownership, error) {
				return model.Ownership{}, fmt.Errorf("failed to get ownership")
			},
			expectedOwnership: model.Ownership{},
			expectedError:     fmt.Errorf("failed to get NFT ownership: %w", fmt.Errorf("failed to get ownership")),
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			mockClient := &MockNFTClient{
				GetNFTOwnershipFunc: tc.getNFTOwnershipFunc,
			}

			service := service.NewService(mockClient)
			ownership, err := service.GetNFTOwnership("test-id")
			assert.Equal(t, tc.expectedOwnership, ownership)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGetNFTTraitsRarity(t *testing.T) {
	for _, tc := range []struct {
		testName               string
		getNFTTraitsRarityFunc func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error)
		expectedRarity         model.TraitsRarity
		expectedError          error
	}{
		{
			testName: "successfully get nft traits rarity",
			getNFTTraitsRarityFunc: func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
				return model.TraitsRarity{
					Traits: []model.TraitRarity{
						{Key: "Key1", Value: "Value1", Rarity: "0"},
					},
				}, nil
			},
			expectedRarity: model.TraitsRarity{
				Traits: []model.TraitRarity{
					{Key: "Key1", Value: "Value1", Rarity: "0"},
				},
			},
			expectedError: nil,
		},
		{
			testName: "fail to get nft traits rarity",
			getNFTTraitsRarityFunc: func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
				return model.TraitsRarity{}, fmt.Errorf("failed to get traits rarity")
			},
			expectedRarity: model.TraitsRarity{},
			expectedError:  fmt.Errorf("failed to get NFT traits rarity: %w", fmt.Errorf("failed to get traits rarity")),
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			mockClient := &MockNFTClient{
				GetNFTTraitsRarityFunc: tc.getNFTTraitsRarityFunc,
			}

			service := service.NewService(mockClient)
			traitsRarity, err := service.GetNFTTraitsRarity("test-collection-id", []model.TraitProperty{{Key: "Key1", Value: "Value1"}})
			assert.Equal(t, tc.expectedRarity, traitsRarity)
			assert.Equal(t, tc.expectedError, err)
		})
	}
}
