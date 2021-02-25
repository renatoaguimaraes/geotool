package services

import (
	"novitatus.com/geotool/internal/core/domain"
)

// MatrixService service
type MatrixService struct {
}

// New service
func New() *MatrixService {
	return &MatrixService{}
}

// Matrix distances matrix
func (srv *MatrixService) Matrix(coordinates []*domain.Coordinate) ([]*domain.Segment, error) {
	var segments []*domain.Segment
	for _, colum := range coordinates {
		for _, row := range coordinates {
			if colum != row {
				segment := &domain.Segment{From: colum, To: row}
				segment.Calculate()
				segments = append(segments, segment)
			}
		}
	}
	return segments, nil
}
