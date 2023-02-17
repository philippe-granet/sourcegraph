package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/keegancsmith/sqlf"
	"github.com/sourcegraph/log"

	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/types"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

const (
	errorCodeUserWithEmailExists          = "err_user_with_such_email_exists"
	errorCodeAccessRequestWithEmailExists = "err_access_request_with_such_email_exists"
)

// errCannotCreateAccessRequest is the error that is returned when a request_access cannot be added to the DB due to a constraint.
type errCannotCreateAccessRequest struct {
	code string
}

func (err errCannotCreateAccessRequest) Error() string {
	return fmt.Sprintf("cannot create user: %v", err.code)
}

// errAccessRequestNotFound is the error that is returned when a request_access cannot be found in the DB.
type errAccessRequestNotFound struct {
	ID int32
}

func (e *errAccessRequestNotFound) Error() string {
	return fmt.Sprintf("access_request with ID %d not found", e.ID)
}

func (e *errAccessRequestNotFound) NotFound() bool {
	return true
}

// IsAccessRequestUserWithEmailExists reports whether err is an error indicating that the access request email was already taken by a signed in user.
func IsAccessRequestUserWithEmailExists(err error) bool {
	var e errCannotCreateAccessRequest
	return errors.As(err, &e) && e.code == errorCodeUserWithEmailExists
}

// IsAccessRequestWithEmailExists reports whether err is an error indicating that the access request was already created.
func IsAccessRequestWithEmailExists(err error) bool {
	var e errCannotCreateAccessRequest
	return errors.As(err, &e) && e.code == errorCodeAccessRequestWithEmailExists
}

type AccessRequestsFilterOptions struct {
	Status *types.AccessRequestStatus
}

func (o *AccessRequestsFilterOptions) sqlConditions() []*sqlf.Query {
	conds := []*sqlf.Query{sqlf.Sprintf("TRUE")}
	if o != nil && o.Status != nil {
		conds = append(conds, sqlf.Sprintf("status = %v", *o.Status))
	}
	return conds
}

type AccessRequestsListOptions struct {
	OrderBy    *string
	Descending *bool
	Limit      *int32
	Offset     *int32
}

func (o *AccessRequestsListOptions) sqlOrderBy() (*sqlf.Query, error) {
	orderDirection := "ASC"
	if o != nil && o.Descending != nil && *o.Descending {
		orderDirection = "DESC"
	}
	orderBy := sqlf.Sprintf("id " + orderDirection)
	if o != nil && o.OrderBy != nil {
		newOrderColumn, err := toAccessRequestsField(*o.OrderBy)
		orderBy = sqlf.Sprintf(newOrderColumn + " " + orderDirection)
		if err != nil {
			return nil, err
		}
	}

	return orderBy, nil
}

func (o *AccessRequestsListOptions) sqlLimit() *sqlf.Query {
	limit := int32(100)
	if o != nil && o.Limit != nil {
		limit = *o.Limit
	}

	offset := int32(0)
	if o != nil && o.Offset != nil {
		offset = *o.Offset
	}

	return sqlf.Sprintf(`%s OFFSET %s`, limit, offset)
}

type AccessRequestsFilterAndListOptions struct {
	*AccessRequestsListOptions
	*AccessRequestsFilterOptions
}

func toAccessRequestsField(orderBy string) (string, error) {
	switch orderBy {
	case "NAME":
		return "name", nil
	case "EMAIL":
		return "email", nil
	case "CREATED_AT":
		return "created_at", nil
	default:
		return "", errors.New("invalid orderBy")
	}
}

// AccessRequestStore provides access to the `access_requests` table.
//
// For a detailed overview of the schema, see schema.md.
type AccessRequestStore interface {
	basestore.ShareableStore
	Create(context.Context, *types.AccessRequest) (*types.AccessRequest, error)
	Update(context.Context, *types.AccessRequest) (*types.AccessRequest, error)
	GetByID(context.Context, int32) (*types.AccessRequest, error)
	Count(context.Context, *AccessRequestsFilterOptions) (int, error)
	List(context.Context, *AccessRequestsFilterAndListOptions) (_ []*types.AccessRequest, err error)
}

type accessRequestStore struct {
	*basestore.Store
	logger log.Logger
}

// AccessRequestsWith instantiates and returns a new accessRequestStore using the other store handle.
func AccessRequestsWith(other basestore.ShareableStore, logger log.Logger) AccessRequestStore {
	return &accessRequestStore{Store: basestore.NewWithHandle(other.Handle()), logger: logger}
}

const (
	accessRequestInsertQuery = `
		INSERT INTO
			access_requests (name, email, additional_info)
		VALUES ( %s, %s, %s )
		RETURNING id, created_at, updated_at, name, email, status, additional_info
		`
	accessRequestListQuery = `
		SELECT
			id, created_at, updated_at, name, email, status, additional_info
		FROM
			access_requests
		WHERE (%s)
		ORDER BY %s
		LIMIT %s
		`
	accessRequestUpdateQuery = `
		UPDATE access_requests
		SET status = %s
		WHERE id = %s
		RETURNING id, created_at, updated_at, name, email, status, additional_info`
)

func (s *accessRequestStore) Create(ctx context.Context, accessRequest *types.AccessRequest) (*types.AccessRequest, error) {
	// We don't allow adding a new request_access with an email address that has already been
	// verified by another user.
	exists, _, err := basestore.ScanFirstBool(s.Query(ctx, sqlf.Sprintf("SELECT TRUE WHERE EXISTS (SELECT FROM user_emails WHERE email = %s AND verified_at IS NOT NULL)", accessRequest.Email)))
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errCannotCreateAccessRequest{errorCodeUserWithEmailExists}
	}

	// We don't allow adding a new request_access with an email address that has already been used
	exists, _, err = basestore.ScanFirstBool(s.Query(ctx, sqlf.Sprintf("SELECT TRUE WHERE EXISTS (SELECT FROM access_requests WHERE email = %s)", accessRequest.Email)))
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errCannotCreateAccessRequest{errorCodeAccessRequestWithEmailExists}
	}

	// Continue with creating the new access request.
	q := sqlf.Sprintf(
		accessRequestInsertQuery,
		accessRequest.Name,
		accessRequest.Email,
		accessRequest.AdditionalInfo,
	)
	var data types.AccessRequest

	if err := s.QueryRow(ctx, q).Scan(&data.ID, &data.CreatedAt, &data.UpdatedAt, &data.Name, &data.Email, &data.Status, &data.AdditionalInfo); err != nil {
		return nil, errors.Wrap(err, "scanning access_request")
	}

	return &data, nil
}

func (s *accessRequestStore) GetByID(ctx context.Context, id int32) (*types.AccessRequest, error) {
	row := s.QueryRow(ctx, sqlf.Sprintf("SELECT id, created_at, updated_at, name, email, status, additional_info FROM access_requests WHERE id = %s", id))
	var node types.AccessRequest

	if err := row.Scan(&node.ID, &node.CreatedAt, &node.UpdatedAt, &node.Name, &node.Email, &node.Status, &node.AdditionalInfo); err != nil {
		if err == sql.ErrNoRows {
			return nil, &errAccessRequestNotFound{ID: id}
		}
		return nil, err
	}

	return &node, nil
}

func (s *accessRequestStore) Update(ctx context.Context, accessRequest *types.AccessRequest) (*types.AccessRequest, error) {
	q := sqlf.Sprintf(accessRequestUpdateQuery, accessRequest.Status, accessRequest.ID)
	var updated types.AccessRequest

	if err := s.QueryRow(ctx, q).Scan(&updated.ID, &updated.CreatedAt, &updated.UpdatedAt, &updated.Name, &updated.Email, &updated.Status, &updated.AdditionalInfo); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, &errAccessRequestNotFound{ID: accessRequest.ID}
		}
		return nil, errors.Wrap(err, "scanning access_request")
	}

	return &updated, nil
}

func (s *accessRequestStore) Count(ctx context.Context, opt *AccessRequestsFilterOptions) (int, error) {
	q := sqlf.Sprintf("SELECT COUNT(*) FROM access_requests WHERE (%s)", sqlf.Join(opt.sqlConditions(), ") AND ("))
	var count int
	if err := s.QueryRow(ctx, q).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

func (s *accessRequestStore) List(ctx context.Context, opt *AccessRequestsFilterAndListOptions) ([]*types.AccessRequest, error) {
	orderBy, err := opt.sqlOrderBy()
	if err != nil {
		return nil, err
	}

	query := sqlf.Sprintf(accessRequestListQuery, sqlf.Join(opt.sqlConditions(), ") AND ("), orderBy, opt.sqlLimit())

	rows, err := s.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	nodes := make([]*types.AccessRequest, 0)
	for rows.Next() {
		var node types.AccessRequest

		if err := rows.Scan(&node.ID, &node.CreatedAt, &node.UpdatedAt, &node.Name, &node.Email, &node.Status, &node.AdditionalInfo); err != nil {
			return nil, err
		}

		nodes = append(nodes, &node)
	}

	return nodes, nil
}
