package exception

type ErrorUnauthorized struct {
	Error string
}

func NewErrorUnauthorized(err string) ErrorUnauthorized {
	return ErrorUnauthorized{
		Error: err,
	}
}
