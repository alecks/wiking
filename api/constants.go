package main

const apiPath = "/api/v1/"

// This is just here in case things change (possibly to ints)
var superuserAllow = map[string]byte{
	"ownerUser":  3,
	"adminUser":  2,
	"editorUser": 1,
	"normalUser": 0,
}
