package i18n

const (
	ErrUnexpected Message = iota + 1
	ErrBadRequest
)

var Messages = MessageMap{

	// Error messages:
	ErrUnexpected: gettext("Unexpected error, please try again"),
	ErrBadRequest: gettext("Invalid request"),
}
