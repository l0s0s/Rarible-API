package service

import (
	"fmt"
	"l0s0s/Rarible-API/model"
)

type NFTClient interface {
	GetNFTOwnership(id string) (model.Ownership, error)
	GetNFTTraitsRarity(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error)
}

func NewService(client NFTClient) *Service {
	return &Service{
		client: client,
	}
}

type Service struct {
	client NFTClient
}

func (s *Service) GetNFTOwnership(id string) (model.Ownership, error) {
	ownership, err := s.client.GetNFTOwnership(id)
	if err != nil {
		return model.Ownership{}, fmt.Errorf("failed to get NFT ownership: %w", err)
	}

	return ownership, nil
}

func (s *Service) GetNFTTraitsRarity(collectionID string, properties []model.TraitProperty) (model.TraitsRarity, error) {
	traitsRarity, err := s.client.GetNFTTraitsRarity(collectionID, properties)
	if err != nil {
		return model.TraitsRarity{}, fmt.Errorf("failed to get NFT traits rarity: %w", err)
	}

	return traitsRarity, nil
}
