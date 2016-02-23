package idgen

import (
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	genter := NewIdGen('t')
	id := genter.Gen()
	if id <= 0 {
		t.Fatalf("Genter nagetive number %d\n", id)
	}
}

func TestDecodeAndEncode(t *testing.T) {
	genter := NewIdGen('i')
	id := genter.Gen()
	d, e := Decode(id)
	if e != nil {
		t.Fatalf("Decode error %s\n", e.Error())
	}
	id2, e := Encode(d)
	if e != nil {
		t.Fatalf("Encode error %s\n", e.Error())
	}
	if id != id2 {
		t.Fatalf("Decode or encode fail")
	}
}

func TestGetTimeFromId(t *testing.T) {
	genter := NewIdGen('t')
	id := genter.Gen()
	timestamp, err := GetTimeFromId(id)
	if timestamp > time.Now().Unix() || timestamp < time.Now().Unix()-1 {
		t.Fatal("Genter id format  error", err, timestamp, id)
	}
}

func BenchmarkGenDecodeEncode(b *testing.B) {
	b.StopTimer()
	genter := NewIdGen('i')
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		id := genter.Gen()
		d, _ := Decode(id)
		id2, _ := Encode(d)
		if id != id2 {
			return
		}
	}
}

func BenchmarkGen(b *testing.B) {
	b.StopTimer()
	genter := NewIdGen('b')
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		genter.Gen()
	}
}
