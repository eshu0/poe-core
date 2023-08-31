package types

import (
	"fmt"
)

type GoMembers struct {
	Type    MemberType
	Members map[string]IMember
}

func NewGoMembers(mt MemberType) GoMembers {
	str := GoMembers{Type: mt, Members: make(map[string]IMember)}
	return str
}

func (gms GoMembers) Add(memb IMember) (GoMembers, error) {
	if memb.GetType() != gms.Type {
		return gms, fmt.Errorf("%s does not match %s - for %s", gms.Type.String(), memb.GetType().String(), memb.GetName())
	}
	members := gms.Members
	_, exists := members[memb.GetName()]
	if exists {
		return gms, fmt.Errorf("%s %s already exists", gms.Type.String(), memb.GetName())
	}
	members[memb.GetName()] = memb
	gms.Members = members
	return gms, nil
}

func (gms GoMembers) Update(memb IMember) (GoMembers, error) {
	if memb.GetType() != gms.Type {
		return gms, fmt.Errorf("%s does not match %s - for %s", gms.Type.String(), memb.GetType().String(), memb.GetName())
	}
	members := gms.Members
	_, exists := members[memb.GetName()]
	if !exists {
		return gms.Add(memb)
	}
	members[memb.GetName()] = memb
	gms.Members = members
	return gms, nil
}

func (gms GoMembers) Clear(name string) (GoMembers, error) {
	members := gms.Members
	_, exists := members[name]
	if !exists {
		return gms, nil
	}
	delete(members, name)
	gms.Members = members
	return gms, nil
}

func (gms GoMembers) GetMember(name string) (IMember, error) {
	member, exists := gms.Members[name]
	if !exists {
		return nil, fmt.Errorf("member %s missing", name)
	}
	return member, nil
}

func (gms GoMembers) GetAllPointers() map[string]IMember {

	results := make(map[string]IMember)
	for name, v := range gms.Members {
		if v.HasPointerPrefix() {
			results[name] = v
		}
	}
	return results
}

func (gms GoMembers) GetAllReferences() map[string]IMember {
	results := make(map[string]IMember)
	for name, v := range gms.Members {
		if !v.HasPointerPrefix() {
			results[name] = v
		}
	}
	return results
}
