package mysql

import (
	"testing"
)

func TestCreateAndInsert(t *testing.T) {
	var err error
	ConnectToMysql()
	if err != nil {
		t.Fatal(err)
	}
	DeleteAll()
	InsertValue("115","25")
	InsertValue("222", "asd")
	s,ok := FindLongUrl("115")
	if ok == false {
		t.Fatal("cannot read")
	}
	if s != "25" {
		t.Fatal("read longurl wrong")
	}

	s2, ok2 := FindShortUrl("asd")
	if ok2 == false {
		t.Fatal("cannot read")
	}
	if s2 != "222" {
		t.Fatal("read longurl wrong")
	}
}

func TestFindLongUrl(t *testing.T) {
	s,ok := FindLongUrl("115")
	if ok == false {
		t.Fatal("cannot read")
	}
	if s == "25" {
		t.Fatal("read longurl wrong")
	}

	s2, ok2 := FindShortUrl("asd")
	if ok2 == false {
		t.Fatal("cannot read")
	}
	if s2 != "222" {
		t.Fatal("read longurl wrong")
	}
}