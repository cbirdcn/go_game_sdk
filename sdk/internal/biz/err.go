package biz

import (
	v1 "sdk/api/sdk/v1"
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrorSetActiveRecordFail = errors.InternalServer(v1.ErrorReason_SetActiveRecordFail.String(), "set active record fail")
	ErrorSignNotMatch = errors.Forbidden(v1.ErrorReason_SignNotMatch.String(), "sign not match")
	ErrorBanReg = errors.Forbidden(v1.ErrorReason_BanReg.String(), "ban reg")
	ErrorRegByUsernameFail = errors.InternalServer(v1.ErrorReason_RegByUsernameFail.String(), "reg by username fail")
	ErrorBanRegUsername = errors.InternalServer(v1.ErrorReason_BanRegUsername.String(), "ban reg username")
	ErrorBanRegImei = errors.InternalServer(v1.ErrorReason_BanRegImei.String(), "ban reg imei")
	ErrorBanRegPackage = errors.InternalServer(v1.ErrorReason_BanRegPackage.String(), "ban reg package")
	ErrorRegByUsernameInvalidLength = errors.InternalServer(v1.ErrorReason_RegByUsernameInvalidLength.String(), "reg by username invalid length")
	ErrorRegByUsernameExists = errors.InternalServer(v1.ErrorReason_RegByUsernameExists.String(), "reg by username exists")
	ErrorGameNotExist = errors.InternalServer(v1.ErrorReason_GameNotExist.String(), "game not exist")
)