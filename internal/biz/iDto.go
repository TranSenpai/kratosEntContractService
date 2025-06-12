package biz

type IDto interface {
	Convert() error
	Getter() any
}
