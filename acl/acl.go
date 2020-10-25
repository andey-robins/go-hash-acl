package acl

type Permissions struct {
	read    bool
	write   bool
	execute bool
}

type Entry struct {
	os      Permissions
	program Permissions
	data    Permissions
}

type AclList struct {
	list map[string]Entry
}

// helper function to generate permissions
// os; program; data; then read; write; execute
func CreatePermsEntry(a, b, c, d, e, f, g, h, i bool) Entry {
	return Entry{Permissions{read: a, write: b, execute: c}, Permissions{read: d, write: e, execute: f}, Permissions{read: g, write: h, execute: i}}
}

// ACL List data structure
func New() *AclList {
	m := make(map[string]Entry)
	return &AclList{list: m}
}

func (a *AclList) AddUser(name string, perms Entry) {
	a.list[name] = perms
}

func (a *AclList) UpdateUserPerms(name string, perms Entry) {
	a.list[name] = perms
}

func (a *AclList) GetPermissions(name string) Entry {
	return a.list[name]
}
