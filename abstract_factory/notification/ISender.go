package notification

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}
