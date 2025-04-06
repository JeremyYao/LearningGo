package main

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}

/*
---------------------------------
0 passed, 2 failed
--- FAIL: Test (0.00s)
    main_test.go:38: ---------------------------------
        Inputs:     (contact)
        Expecting:  24 bytes
        Actual:     32 bytes
    main_test.go:38: ---------------------------------
        Inputs:     (perms)
        Expecting:  16 bytes
        Actual:     24 bytes
        Fail
FAIL
exit status 1
FAIL    4l7     0.189s

for

package main

type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
}

*/
