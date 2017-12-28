package idgen

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGen(t *testing.T) {
	genter := NewIdGen('t')
	id := genter.Gen(3)
	if id <= 0 {
		t.Fatalf("Genter nagetive number %d\n", id)
	}
}

func TestDecodeAndEncode(t *testing.T) {
	genter := NewIdGen('i')
	id := genter.Gen(5)
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

func BencmarkGenpb(b *testing.B) {
	b.StopTimer()
	genter := NewIdGen('i')
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		if pb.Next() {
			genter.Gen(2)
		}
	})
}

type d struct {
	m *sync.Mutex
	s map[int64]int
}

func (d *d) Set(i int64) error {
	d.m.Lock()
	defer d.m.Unlock()
	if _, ok := d.s[i]; ok {
		return fmt.Errorf("i:%d exists!", i)
	}
	d.s[i] = 1
	return nil
}

func TestParalle(t *testing.T) {
	N := 5000000
	a := &d{
		m: new(sync.Mutex),
		s: make(map[int64]int, N),
	}

	genter := NewIdGen('i')

	for i := 0; i < N; i++ {
		go func() {
			id := genter.Gen(6)
			if err := a.Set(id); err != nil {
				t.Fatalf("Paralle error:%v", err)
			}
		}()
	}
}

func TestGetTimeFromId(t *testing.T) {
	genter := NewIdGen('t')
	id := genter.Gen(7)
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
		id := genter.Gen(8)
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
		genter.Gen(9)
	}
}
