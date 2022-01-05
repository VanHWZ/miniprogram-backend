package errorcode

var (
	UnAuthorized = 1
	PermissionDenied = 2

	ParamError   = 10
	RegisterError = 11
	AuthError = 12
	GroupQuitError = 13
	EventRetrieveError = 14
	EventDeleteError = 15
	EventUpdateError = 16
	MessageRetrieveError = 17
	MessageDeleteError = 18
	MessageUpdateError = 19

	DatabaseError = 50

	CodeMap = map[int]string{
		UnAuthorized:          "UnAuthorized",
		PermissionDenied:      "PermissionDenied",

		ParamError:            "ParameterError",
		RegisterError:         "RegisterError",
		AuthError:             "AuthError",
		GroupQuitError:        "GroupQuitError",
		EventRetrieveError:    "EventRetrieveError",
		EventDeleteError:      "EventDeleteError",
		EventUpdateError:      "EventUpdateError",
		MessageRetrieveError:  "MessageRetrieveError",
		MessageDeleteError:    "MessageDeleteError",
		MessageUpdateError:    "MessageUpdateError",

		DatabaseError:         "DatabaseError",
	}
)


