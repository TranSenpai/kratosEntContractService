package data

import (
	"database/sql"
	kerror "github.com/go-kratos/kratos/v2/errors"
	"github.com/lib/pq"
	"kratosEntContractService/internal/ent"
	"net/http"
)

func GetError(err error) error {
	if err == nil {
		return nil
	}

	if ent.IsNotFound(err) || err == sql.ErrNoRows {
		return kerror.New(http.StatusNotFound, "Repo error", "resource not found")
	}

	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code.Name() {
		case "unique_violation":
			return kerror.New(http.StatusBadRequest, "Repo error", "duplicate value")
		case "foreign_key_violation":
			return kerror.New(http.StatusBadRequest, "Repo error", "invalid reference")
		case "check_violation":
			return kerror.New(http.StatusBadRequest, "Repo error", "constraint failed")
		case "not_null_violation":
			return kerror.New(http.StatusBadRequest, "Repo error", "missing required field")
		}
	}

	return kerror.New(http.StatusInternalServerError, "Repo error", err.Error())
}
