package documentstore

import "errors"

var ErrInvalidTypeString = errors.New("field is not a type string")
var ErrInvalidTypeNumber = errors.New("field is not a type number")
var ErrInvalidTypeBool = errors.New("field is not a type boolean")
var ErrInvalidTypeArray = errors.New("field is not a type array")
var ErrInvalidTypeObject = errors.New("field is not a type object")

var ErrDocumentPrimaryKeyIsMissing = errors.New("document primary Key is missing")
var ErrDocumentNotFound = errors.New("document not found")

var ErrCollectionAlreadyExists = errors.New("collection already exists")
var ErrCollectionNotFound = errors.New("collection not found")

var ErrIndexDoesNotExist = errors.New("index does not exist")
var ErrIndexAlreadyExists = errors.New("index already exists")
