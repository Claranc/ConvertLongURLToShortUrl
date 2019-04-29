package controller

import "sync"

const jinzhi int = 62

var (
	count = 0
	num = 19960117
	code62 = []byte{'0','1','2','3','4','5','6','7','8','9','a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u',
		'v','w','x','y','z','A','B','C','D','E','F','G','H','I','J','K','L','M','N','O','P','Q','R','S','T','U','V','W','X','Y','Z'}
	mu sync.Mutex
	validhead = []string{"http://", "https://", "www."}
	Method int = 1
)