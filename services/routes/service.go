package routes

import (
	"database/sql"
	"time"

	"github.com/luizalabs/go-training/util"
)

const listQuery = "select route_key, modality_id, threshold, create_date, update_date from routes limit ? offset ?"

// ServiceIface ...
type ServiceIface interface {
	List(limit int, offset int) (Routes, error)
}

// Service ...
type Service struct {
	db *sql.DB
}

// New ...
func New(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

// List ...
func (s *Service) List(limit int, offset int) (Routes, error) {
	logger := util.GetLogger().WithFields(map[string]interface{}{
		"module":        "service",
		"function_name": "routes.list",
	})

	dbMetrics := map[string]interface{}{
		"statement": listQuery,
		"params":    "50, 0",
	}

	routes := Routes{}

	start := time.Now()
	result, err := s.db.Query(listQuery, limit, offset)
	delta := time.Since(start)

	dbMetrics["latency"] = int(delta / time.Millisecond)

	if err != nil {
		dbMetrics["error"] = err.Error()
		logger.WithField("db", dbMetrics).Error()
		return routes, err
	}

	for result.Next() {
		row := Route{}
		err = result.Scan(&row.ID, &row.Modality, &row.Threshold, &row.CreatedAt, &row.UpdatedAt)
		if err != nil {
			logger.WithField("db", dbMetrics).Error()
			return routes, err
		}

		routes = append(routes, row)
	}

	dbMetrics["records"] = len(routes)
	logger.WithField("db", dbMetrics).Info()
	return routes, nil
}
