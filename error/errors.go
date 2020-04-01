package error

import (
	"fmt"
)

type errScope string

const (
	errUnknown errScope = ""
	errDB               = "db_error"
	errMinio            = "minio_error"
	errRedis            = "redis_error"
)

var (
	// DB error
	ErrEmptyDBUsername = makeError(errDB, "missing Username")
	ErrEmptyDBPassword = makeError(errDB, "missing Password")
	ErrEmptyDBHost     = makeError(errDB, "missing Host")
	ErrInvalidDBPort   = makeError(errDB, "invalid Port")
	ErrEmptyDBDatabase = makeError(errDB, "missing Database")

	// Minio error
	ErrEmptyMinioAccessKey = makeError(errMinio, "access_key is empty")
	ErrEmptyMinioSecretKey = makeError(errMinio, "secret_key is empty")
	ErrEmptyMinioHost      = makeError(errMinio, "host is empty")
	ErrInvalidMinioPort    = makeError(errMinio, "port is invalid")

	// Redis error
	ErrEmptyRedisHost   = makeError(errRedis, "missing Host")
	ErrInvaldiRedisPort = makeError(errRedis, "invalid Port")
)

func makeError(scope errScope, msg ...string) error {
	return fmt.Errorf("[%s]: %s", scope, msg)
}
