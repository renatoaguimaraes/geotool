package ports

import "novitatus.com/geotool/internal/core/domain"

// ContaService service
type ContaService interface {
	Matrix(coordinates []*domain.Coordinate) ([]*domain.Segment, error)
}
