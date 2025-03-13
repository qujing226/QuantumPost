package event

type Consumer interface {
	start() error
}
