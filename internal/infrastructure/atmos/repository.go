package repository

import (
	"context"
	"logbook-api/internal/config"

	client "logbook-api/pkg/atmos"

	"github.com/umatare5/atmos-go"
)

// Repository struct
type Repository struct {
	Config *config.Config
}

// GetUser ...
func (r *Repository) GetUser(userID string) (*atmos.GetUserResponse, error) {
	client, err := client.NewClientWithToken(r.Config.ConfigAtmosServer, r.Config.ConfigAtmosToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetUserWithResponse(
		context.Background(), userID,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUserStatistics ...
func (r *Repository) GetUserStatistics(userID string) (*atmos.GetUserStatisticsResponse, error) {
	client, err := client.NewClientWithToken(r.Config.ConfigAtmosServer, r.Config.ConfigAtmosToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetUserStatisticsWithResponse(
		context.Background(), userID,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetDivelogs ...
func (r *Repository) GetDivelogs(limit *int, cursor *string) (*atmos.GetDivelogsResponse, error) {
	client, err := client.NewClientWithToken(r.Config.ConfigAtmosServer, r.Config.ConfigAtmosToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetDivelogsWithResponse(
		context.Background(),
		&atmos.GetDivelogsParams{
			Limit:  limit,
			Cursor: cursor,
		},
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetDeletedDivelogs ...
func (r *Repository) GetDeletedDivelogs(limit *int, cursor *string) (*atmos.GetDeletedDivelogsResponse, error) {
	client, err := client.NewClientWithToken(r.Config.ConfigAtmosServer, r.Config.ConfigAtmosToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetDeletedDivelogsWithResponse(
		context.Background(),
		&atmos.GetDeletedDivelogsParams{
			Limit:  limit,
			Cursor: cursor,
		},
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetDivelog ...
func (r *Repository) GetDivelog(divelogID string) (*atmos.GetDivelogResponse, error) {
	client, err := client.NewClientWithToken(r.Config.ConfigAtmosServer, r.Config.ConfigAtmosToken)
	if err != nil {
		return nil, err
	}

	data, err := client.GetDivelogWithResponse(
		context.Background(), divelogID,
	)
	if err != nil {
		return nil, err
	}

	return data, nil
}
