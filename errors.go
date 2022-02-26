package gokystr

import "errors"

var (
	// ErrNoSuchKey is thrown if a key with the given ID is not in the store
	ErrNoSuchKey = errors.New("no matching key found")
	// ErrKeyExistsAlready is thrown if a key already exists in the store
	ErrKeyExistsAlready = errors.New("key is already in the store")
	// ErrNoKeyFound is thrown if the dir contains no valid key file
	ErrNoKeyFound = errors.New("no valid keyfiles found in root")
	// ErrNotADir is thrown if the passed dir is not a directory path
	ErrNotADir = errors.New("passed path is not a dir")
	// ErrWrongExtFormat is thrown if the passed key extension does not match '.<ext>' format
	ErrWrongExtFormat = errors.New("key extension does not match '.<ext>' format")
	// ErrWrongDirFormat is thrown if the passed dir is not a valid path to open.
	// Path names must not contain an element that is “.” or “..” or the empty string,
	// except for the special case that the root directory is named “.”.
	// Paths must not start or end with a slash: “/x” and “x/” are invalid.
	ErrWrongDirFormat = errors.New("key extension does not match '.<ext>' format")
)
