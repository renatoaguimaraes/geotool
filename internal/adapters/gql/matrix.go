package gql

import (
	"github.com/graphql-go/graphql"
	"novitatus.com/geotool/internal/core/domain"
	"novitatus.com/geotool/internal/core/services"
)

// CoordinateParamType graphql type
var CoordinateParamType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "CoordinateParam",
	Description: "Geographic coordinate",
	Fields: graphql.InputObjectConfigFieldMap{
		"lgt": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "Longitude.",
		},
		"lat": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "Latitude.",
		},
	},
})

// CoordinateType graphql type
var CoordinateType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Coordinate",
	Description: "Geographic coordinate",
	Fields: graphql.Fields{
		"lgt": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "Longitude.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coordinate, ok := p.Source.(*domain.Coordinate); ok {
					return coordinate.Lgt, nil
				}
				return nil, nil
			},
		},
		"lat": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "Latitude.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if coordinate, ok := p.Source.(*domain.Coordinate); ok {
					return coordinate.Lat, nil
				}
				return nil, nil
			},
		},
	},
})

// SegmentType graphql type
var SegmentType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Segment",
	Description: "From location to other with distance between them.",
	Fields: graphql.Fields{
		"from": &graphql.Field{
			Type:        graphql.NewNonNull(CoordinateType),
			Description: "Source location.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if conta, ok := p.Source.(*domain.Segment); ok {
					return conta.From, nil
				}
				return nil, nil
			},
		},
		"to": &graphql.Field{
			Type:        graphql.NewNonNull(CoordinateType),
			Description: "Target location.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if conta, ok := p.Source.(*domain.Segment); ok {
					return conta.To, nil
				}
				return nil, nil
			},
		},
		"distance": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.Float),
			Description: "Distance in Km.",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if conta, ok := p.Source.(*domain.Segment); ok {
					return conta.Distance, nil
				}
				return nil, nil
			},
		},
	},
})

// Matrix distance matrix
func Matrix() *graphql.Field {
	var contaService = services.New()
	return &graphql.Field{
		Type:        graphql.NewList(SegmentType),
		Description: "Distance Matrix.",
		Args: graphql.FieldConfigArgument{
			"coordinates": &graphql.ArgumentConfig{
				Type: graphql.NewList(CoordinateParamType),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var coordinates []*domain.Coordinate
			paramaters := params.Args["coordinates"].([]interface{})
			for _, parameter := range paramaters {
				latLng := parameter.(map[string]interface{})
				lat := latLng["lat"].(float64)
				lng := latLng["lgt"].(float64)
				coordinates = append(coordinates, &domain.Coordinate{Lat: lat, Lgt: lng})
			}
			return contaService.Matrix(coordinates)
		},
	}
}
