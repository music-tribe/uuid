

# uuid - forked from github.com/google/uuid

## Methods added to original package

 - `GetBSON` & `SetBSON` have been added such that the UUID type integrates seamlessly (hopefully) with the MongoDB driver, github.com/globalsign/mgo. Note that this specific package _must_ be used for the integration to work correctly.

 - `UnmarshalParam` has been added specifically for integration with the Echo packages `Bind` method.

## Original Readme


The uuid package generates and inspects UUIDs based on
[RFC 4122](http://tools.ietf.org/html/rfc4122)
and DCE 1.1: Authentication and Security Services. 

This package is based on the github.com/pborman/uuid package (previously named
code.google.com/p/go-uuid).  It differs from these earlier packages in that
a UUID is a 16 byte array rather than a byte slice.  One loss due to this
change is the ability to represent an invalid UUID (vs a NIL UUID).

###### Install
`go get github.com/music-tribe/uuid`

###### Documentation 
[![GoDoc](https://godoc.org/github.com/google/uuid?status.svg)](http://godoc.org/github.com/google/uuid)

Full `go doc` style documentation for the package can be viewed online without
installing this package by using the GoDoc site here: 
http://godoc.org/github.com/google/uuid
