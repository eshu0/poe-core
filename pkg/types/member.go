package types

import (
	"unicode"
)

type GoMember struct {
	IMember
	name          string
	pointerPrefix bool
	mtype         MemberType
}

func (gm GoMember) GetName() string {
	return gm.name
}

func (gm GoMember) SetName(name string) {
	gm.name = name
}

func (gm GoMember) GetType() MemberType {
	return gm.mtype
}

func (gm GoMember) SetType(mtype MemberType) {
	gm.mtype = mtype
}

func (gm GoMember) HasPointerPrefix() bool {
	return gm.pointerPrefix
}

func (gm GoMember) IsPublic() bool {
	return unicode.IsUpper(rune(gm.name[0]))
}

func (gm GoMember) SetPointerPrefix(ptr bool) {
	gm.pointerPrefix = ptr
}

func NewGoMember(name string, t MemberType) GoMember {
	gm := GoMember{name: name, mtype: t, pointerPrefix: false}
	return gm
}

func NewGoPtrMember(name string, t MemberType) GoMember {
	gm := GoMember{name: name, mtype: t, pointerPrefix: true}
	return gm
}
