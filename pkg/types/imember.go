package types

type IMember interface {
	IsPublic() bool
	HasPointerPrefix() bool
	SetPointerPrefix(bool)
	GetName() string
	SetName(string)
	GetType() MemberType
	SetType(MemberType)
}
