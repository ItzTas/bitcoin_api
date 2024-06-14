package main

import "github.com/lib/pq"

const (
	UniqueViolationErr = pq.ErrorCode("23505")
)

func isErrorCode(err error, errcode pq.ErrorCode) bool {
	if pgerr, ok := err.(*pq.Error); ok {
		return pgerr.Code == errcode
	}
	return false
}
