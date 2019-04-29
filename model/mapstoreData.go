package model

import "sync"

//并发安全字典，key为长网址，value为短网址
var LongToShort sync.Map

//并发安全字典，key为短网址，value为长网址
var ShortToLong sync.Map
