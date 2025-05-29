package cloudprovider

import "yunion.io/x/pkg/errors"

const (
	ErrMissingParameter    = errors.Error("MissingParameterError")
	ErrInputParameter      = errors.Error("InputParameterError")
	ErrInvalidProvider     = errors.Error("InvalidProvider")
	ErrNoBalancePermission = errors.Error("NoBalancePermission")
	ErrForbidden           = errors.Error("ForbiddenError")
	ErrTooLarge            = errors.Error("TooLargeEntity")
	ErrUnsupportedProtocol = errors.Error("UnsupportedProtocol")
	ErrInvalidAccessKey    = errors.Error("InvalidAccessKey")
	ErrUnauthorized        = errors.Error("UnauthorizedError")
	ErrNoPermission        = errors.Error("NoPermission")
	ErrNoSuchProvder       = errors.Error("no such provider")
	ErrTooManyRequests     = errors.Error("TooManyRequests")
	ErrInvalidSku          = errors.Error("InvalidSku")

	ErrNotFound        = errors.ErrNotFound
	ErrDuplicateId     = errors.ErrDuplicateId
	ErrInvalidStatus   = errors.ErrInvalidStatus
	ErrTimeout         = errors.ErrTimeout
	ErrNotImplemented  = errors.ErrNotImplemented
	ErrNotSupported    = errors.ErrNotSupported
	ErrAccountReadOnly = errors.ErrAccountReadOnly

	ErrUnknown = errors.Error("UnknownError")
)
