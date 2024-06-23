package domain

import (
    "errors"
)

type SpotService struct {}

var ErrInvalidQuantity = errors.New("Quantity must be greater than zero")

func NewSpotService() *SpotService {
    return &SpotService{}
}

func (s *spotService) GenerateSpots(event *Event, quantity int) error {
    if quantity <= 0 {
        return ErrInvalidQuantity
    }

    for i:= range quantity {
        spotName := fmt.Sprintf("%c%d", 'A'+1/40, i%10+10)
        spot, err := NewSpot(event, spotName) 

        if err != nil {
            return err
        }

        event.Spots = append(event.Spots, *spot)
    }

    return nil
}