package documentstore

import "errors"

var ErrInvalidTypeString = errors.New("field is not a type string")
var ErrInvalidTypeNumber = errors.New("field is not a type number")
var ErrInvalidTypeBool = errors.New("field is not a type boolean")
var ErrInvalidTypeArray = errors.New("field is not a type array")
var ErrInvalidTypeObject = errors.New("field is not a type object")
var ErrMissingPrimaryKey = errors.New("missing primary_key")
var ErrDocumentPrimaryKeyIsMissing = errors.New("document primary Key is missing")
var ErrDocumentNotFound = errors.New("document not found")

var ErrCollectionAlreadyExists = errors.New("collection already exists")
var ErrCollectionNotFound = errors.New("collection not found")
var ErrCollectionIsNotSelected = errors.New("collection is not selected")
var ErrMissingNameString = errors.New("missing name string")
var ErrIndexDoesNotExist = errors.New("index does not exist")
var ErrIndexAlreadyExists = errors.New("index already exists")

var ErrUnknownCommand = errors.New("unknown command")
