package utils

import "errors"

var (
	ErrBucketAlreadyExists = errors.New("this bucket already exists")
	ErrCreatingBucket = errors.New("occured error while creating bucket")
)