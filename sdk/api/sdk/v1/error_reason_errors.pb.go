// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsOK(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_OK.String() && e.Code == 200
}

func ErrorOK(format string, args ...interface{}) *errors.Error {
	return errors.New(200, ErrorReason_OK.String(), fmt.Sprintf(format, args...))
}

func IsSetActiveRecordFail(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SetActiveRecordFail.String() && e.Code == 500
}

func ErrorSetActiveRecordFail(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_SetActiveRecordFail.String(), fmt.Sprintf(format, args...))
}
