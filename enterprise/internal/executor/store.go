package executor

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"

	"github.com/keegancsmith/sqlf"
	"github.com/sourcegraph/log"

	"github.com/sourcegraph/sourcegraph/internal/database"
	"github.com/sourcegraph/sourcegraph/internal/database/basestore"
	"github.com/sourcegraph/sourcegraph/internal/hashutil"
	"github.com/sourcegraph/sourcegraph/internal/observation"
	"github.com/sourcegraph/sourcegraph/lib/errors"
)

type JobTokenStore interface {
	Create(ctx context.Context, jobId int, queue string) (string, error)
	Regenerate(ctx context.Context, jobId int, queue string) (string, error)
	Exists(ctx context.Context, jobId int, queue string) (bool, error)
	Get(ctx context.Context, jobId int, queue string) (JobToken, error)
	GetByToken(ctx context.Context, tokenHexEncoded string) (JobToken, error)
	Delete(ctx context.Context, jobId int, queue string) error
}

type JobToken struct {
	Id    int64
	Value []byte
	JobId int64
	Queue string
}

type jobTokenStore struct {
	*basestore.Store
	logger         log.Logger
	operations     *operations
	observationCtx *observation.Context
}

func NewJobTokenStore(observationCtx *observation.Context, db database.DB) JobTokenStore {
	return &jobTokenStore{
		Store:          basestore.NewWithHandle(db.Handle()),
		logger:         observationCtx.Logger,
		operations:     newOperations(observationCtx),
		observationCtx: observationCtx,
	}
}

func (s *jobTokenStore) Create(ctx context.Context, jobId int, queue string) (string, error) {
	var b [20]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}

	err := s.Exec(
		ctx,
		sqlf.Sprintf(
			createExecutorJobTokenFmtstr,
			hashutil.ToSHA256Bytes(b[:]), jobId, queue,
		),
	)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b[:]), nil
}

const createExecutorJobTokenFmtstr = `
INSERT INTO executor_job_tokens (value_sha256, job_id, queue)
VALUES (%s, %s, %s)
`

func (s *jobTokenStore) Regenerate(ctx context.Context, jobId int, queue string) (string, error) {
	var b [20]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}

	err := s.Exec(
		ctx,
		sqlf.Sprintf(
			updateExecutorJobTokenFmtstr,
			hashutil.ToSHA256Bytes(b[:]), jobId, queue,
		),
	)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(b[:]), nil
}

const updateExecutorJobTokenFmtstr = `
UPDATE executor_job_tokens SET value_sha256 = %s
WHERE job_id = %s AND queue = %s
`

func (s *jobTokenStore) Exists(ctx context.Context, jobId int, queue string) (bool, error) {
	exists, _, err := basestore.ScanFirstBool(s.Query(ctx, sqlf.Sprintf(existsExecutorJobTokenFmtstr, jobId, queue)))
	return exists, err
}

const existsExecutorJobTokenFmtstr = `
SELECT EXISTS(SELECT 1 FROM executor_job_tokens WHERE job_id = %s AND queue = %s)
`

func (s *jobTokenStore) Get(ctx context.Context, jobId int, queue string) (JobToken, error) {
	row := s.QueryRow(
		ctx,
		sqlf.Sprintf(
			getExecutorJobTokenFmtstr,
			jobId, queue,
		),
	)
	return scanJobToken(row)
}

const getExecutorJobTokenFmtstr = `
SELECT id, value_sha256, job_id, queue
FROM executor_job_tokens
WHERE job_id = %s AND queue = %s
`

func (s *jobTokenStore) GetByToken(ctx context.Context, tokenHexEncoded string) (JobToken, error) {
	token, err := hex.DecodeString(tokenHexEncoded)
	if err != nil {
		return JobToken{}, errors.New("invalid token")
	}
	row := s.QueryRow(
		ctx,
		sqlf.Sprintf(
			getByTokenExecutorJobTokenFmtstr,
			hashutil.ToSHA256Bytes(token),
		),
	)
	return scanJobToken(row)
}

const getByTokenExecutorJobTokenFmtstr = `
SELECT id, value_sha256, job_id, queue
FROM executor_job_tokens
WHERE value_sha256 = %s
`

func scanJobToken(row *sql.Row) (JobToken, error) {
	jobToken := JobToken{}
	err := row.Scan(
		&jobToken.Id,
		&jobToken.Value,
		&jobToken.JobId,
		&jobToken.Queue,
	)
	if err != nil {
		return jobToken, err
	}
	return jobToken, nil
}

func (s *jobTokenStore) Delete(ctx context.Context, jobId int, queue string) error {
	return s.Store.Exec(ctx, sqlf.Sprintf(deleteExecutorJobTokenFmtstr, jobId, queue))
}

const deleteExecutorJobTokenFmtstr = `
DELETE FROM executor_job_tokens WHERE job_id = %s AND queue = %s
`