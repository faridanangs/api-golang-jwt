package exception

type ErrNotFound struct {
	Error string
}

func NewNotFoundError(err string) ErrNotFound {
	return ErrNotFound{
		Error: err,
	}
}
