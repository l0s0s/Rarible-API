package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"l0s0s/Rarible-API/handler"
	"l0s0s/Rarible-API/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type MockService struct {
	GetNFTOwnershipFunc    func(id string) (model.Ownership, error)
	GetNFTTraitsRarityFunc func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error)
}

func (m *MockService) GetNFTOwnership(id string) (model.Ownership, error) {
	if m.GetNFTOwnershipFunc != nil {
		return m.GetNFTOwnershipFunc(id)
	}
	return model.Ownership{}, nil
}

func (m *MockService) GetNFTTraitsRarity(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
	if m.GetNFTTraitsRarityFunc != nil {
		return m.GetNFTTraitsRarityFunc(collectionID, properties)
	}
	return model.TraitsRarity{}, nil
}

func setupServer(mockService *MockService) *httptest.Server {
	handler := handler.NewHandler(mockService)

	router := gin.Default()
	handler.RegisterRoutes(router)

	return httptest.NewServer(router)
}

func TestGetNFTOwnership(t *testing.T) {
	for _, tc := range []struct {
		testName            string
		getNFTOwnershipFunc func(id string) (model.Ownership, error)
		expectedStatusCode  int
		expectedOwnership   model.Ownership
		expectedError       error
	}{
		{
			testName: "successfully get nft ownership",
			getNFTOwnershipFunc: func(id string) (model.Ownership, error) {
				return model.Ownership{ID: "id1"}, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedOwnership:  model.Ownership{ID: "id1"},
			expectedError:      nil,
		},
		{
			testName: "fail to get nft ownership",
			getNFTOwnershipFunc: func(id string) (model.Ownership, error) {
				return model.Ownership{}, fmt.Errorf("failed to get ownership")
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedOwnership:  model.Ownership{},
			expectedError:      fmt.Errorf("failed to get ownership"),
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			mockService := &MockService{
				GetNFTOwnershipFunc: tc.getNFTOwnershipFunc,
			}

			server := setupServer(mockService)
			defer server.Close()

			resp, err := http.Get(fmt.Sprintf("%s/nft/ownership/%s", server.URL, "test-id"))
			require.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, tc.expectedStatusCode, resp.StatusCode)
			if tc.expectedError != nil {
				var errorResponse map[string]string
				json.NewDecoder(resp.Body).Decode(&errorResponse)
				assert.Equal(t, tc.expectedError.Error(), errorResponse["error"])
			} else {
				var ownership model.Ownership
				json.NewDecoder(resp.Body).Decode(&ownership)
				assert.Equal(t, tc.expectedOwnership, ownership)
			}
		})
	}
}

func TestGetNFTTraitsRarity(t *testing.T) {
	for _, tc := range []struct {
		testName               string
		properties             []model.TraitProperty
		getNFTTraitsRarityFunc func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error)
		expectedStatusCode     int
		expectedRarity         model.TraitsRarity
		expectedError          error
	}{
		{
			testName:   "successfully get nft traits rarity",
			properties: []model.TraitProperty{{Key: "Key1", Value: "Value1"}},
			getNFTTraitsRarityFunc: func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
				return model.TraitsRarity{
					Traits: []model.TraitRarity{
						{Key: "Key1", Value: "Value1", Rarity: "0"},
					},
				}, nil
			},
			expectedStatusCode: http.StatusOK,
			expectedRarity: model.TraitsRarity{
				Traits: []model.TraitRarity{
					{Key: "Key1", Value: "Value1", Rarity: "0"},
				},
			},
			expectedError: nil,
		},
		{
			testName:   "fail to get nft traits rarity",
			properties: []model.TraitProperty{{Key: "Key1", Value: "Value1"}},
			getNFTTraitsRarityFunc: func(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
				return model.TraitsRarity{}, fmt.Errorf("failed to get traits rarity")
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedRarity:     model.TraitsRarity{},
			expectedError:      fmt.Errorf("failed to get traits rarity"),
		},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			mockService := &MockService{
				GetNFTTraitsRarityFunc: tc.getNFTTraitsRarityFunc,
			}

			server := setupServer(mockService)
			defer server.Close()

			jsonBody, err := json.Marshal(tc.properties)
			require.NoError(t, err)

			resp, err := http.Post(fmt.Sprintf("%s/nft/traits/rarity/%s", server.URL, "test-collection-id"), "application/json", bytes.NewBuffer(jsonBody))
			require.NoError(t, err)

			defer resp.Body.Close()

			assert.Equal(t, tc.expectedStatusCode, resp.StatusCode)
			if tc.expectedError != nil {
				var errorResponse map[string]string
				json.NewDecoder(resp.Body).Decode(&errorResponse)
				assert.Equal(t, tc.expectedError.Error(), errorResponse["error"])
			} else {
				var rarity model.TraitsRarity
				json.NewDecoder(resp.Body).Decode(&rarity)
				assert.Equal(t, tc.expectedRarity, rarity)
			}
		})
	}
}
