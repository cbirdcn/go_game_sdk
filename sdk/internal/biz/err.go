package biz

import (
	v1 "sdk/api/sdk/v1"
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrorSetActiveRecordFail = errors.InternalServer(v1.ErrorReason_SetActiveRecordFail.String(), "set active record fail")
	ErrorSignNotMatch = errors.Forbidden(v1.ErrorReason_SignNotMatch.String(), "sign not match")
)