package errors

import "errors"

var (
	ErrNil                        = errors.New("data nil")
	ErrUserNameOccupied           = errors.New("username is occupied")
	ErrUserNameOrPasswordTooShort = errors.New("username or password is too short")
	ErrPutUserFailed              = errors.New("put user failed")
	ErrGetUserFailed              = errors.New("get user failed")
	ErrIncorrectPassword          = errors.New("incorrect password")
	ErrNoSuchUser                 = errors.New("incorrect username")

	ErrEmptyToolParams      = errors.New("empty tool params")
	ErrProfileNotExists     = errors.New("profile not exists")
	ErrEnvironmentNotExists = errors.New("environment not exists")
	ErrEmptyColumnParams    = errors.New("empty column params")
	ErrColumnIsNotEmpty     = errors.New("remove column is not empty")

	ErrInvalidToken        = errors.New("invalid token")
	ErrParseTokenFailed    = errors.New("parse token failed")
	ErrGenerateTokenFailed = errors.New("generate token failed")

	ErrAssetsNotExists       = errors.New("assets doesn't exists")
	ErrUnknownAssetsType     = errors.New("unknown assets file type")
	ErrUnsupportedAssetsType = errors.New("unsupported assets file type")
	ErrSaveAssetsFailed      = errors.New("save assets failed")
	ErrLoadAssetsFailed      = errors.New("load assets failed")
)
