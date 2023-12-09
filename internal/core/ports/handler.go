package ports

type Process struct {
	ID       string
	Name     string
	LastName string
	Genre    string
	Email    string
	Age      string
}

type Handler interface {
	StartProcess(process *Process) error
	SetNext(handler Handler) Handler
}

type Response struct {
	UserID string
}
