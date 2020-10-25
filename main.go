package main

import (
	"fmt"

	"github.com/andey-robins/go-hash-acl/acl"
)

func main() {

	// create several users and give them permissions
	userOne := "aadams12"
	userTwo := "btaylor54"
	userThree := "zdonner78"

	permissionOne := acl.CreatePermsEntry(true, false, true, false, false, true, true, true, false)
	permissionTwo := acl.CreatePermsEntry(true, true, true, true, true, true, true, true, false)

	// create the acl list
	a := acl.New()
	a.AddUser(userOne, permissionOne)
	a.AddUser(userTwo, permissionTwo)
	a.AddUser(userThree, permissionTwo)

	// shows perms for aadams12
	perms := a.GetPermissions(userOne)
	fmt.Print("aadams12 permissions: ")
	fmt.Println(perms)

	// update the permissions for aadams12
	a.UpdateUserPerms(userOne, permissionTwo)
	perms = a.GetPermissions(userOne)
	fmt.Print("updated permissions:  ")
	fmt.Println(perms)
}
