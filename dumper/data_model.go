package dumper

type User struct {
	UserId int
	Name   string
}

type Sale struct {
	OrderId     int
	UserId      int
	OrderAmount int
}
