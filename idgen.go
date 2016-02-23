package idgen

/***************************
*Genter a int64
*|reserved 1bit|version 4bit|timestamp 32bit|instanceid 8bit|increment 19bit|
****************************/
import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

const (
	VERSION    = 1
	MAX_NUMBER = (1 << 19) - 1
	RAND_MAX   = 1 << 8
	TIME_MASK  = ((1 << 32) - 1) << 27
	INST_MASK  = (1 << 8) - 1
	NUM_MASK   = (1 << 19) - 1
)

type IdGen struct {
	mutex sync.RWMutex
	temp  *ID
}

type ID struct {
	Version    int64
	Time       int64
	Instanceid int64
	Num        int64
}

func NewIdGen(i byte) *IdGen {
	temp := &ID{
		Version:    VERSION,
		Instanceid: int64(i),
	}
	return &IdGen{
		temp: temp,
	}
}

func (self *IdGen) waitNextSecond() {
	nextSecond := time.Unix(self.temp.Time+1, 0)
	duration := nextSecond.Sub(time.Now())
	if duration <= 0 {
		return
	}
	time.Sleep(duration)
	return
}

func (self *IdGen) Gen() (id int64) {
	self.mutex.Lock()
	defer self.mutex.Unlock()
	self.temp.Num += 1
	if self.temp.Num >= MAX_NUMBER {
		self.waitNextSecond()
	}

	t := time.Now().Unix()
	if self.temp.Time != t {
		self.temp.Time = t
		self.temp.Num = rand.Int63n(RAND_MAX)
	}
	id, _ = Encode(self.temp)
	return
}

func GetTimeFromId(id int64) (int64, error) {
	d, err := Decode(id)
	if err != nil {
		return 0, err
	}
	return d.Time, nil
}

func Encode(i *ID) (int64, error) {
	if i.Version != VERSION {
		return 0, errors.New("invalid version")
	}
	if i.Num > MAX_NUMBER {
		return 0, errors.New("num is too big")
	}
	if i.Instanceid > 255 {
		return 0, errors.New("invalid instanceid")
	}
	id := (i.Version << 59) | (i.Time << 27) | (i.Instanceid << 19) | i.Num
	return id, nil
}

func Decode(id int64) (*ID, error) {
	if id>>59 != VERSION {
		return nil, errors.New("invalid ID")
	}
	d := new(ID)
	d.Version = VERSION
	d.Time = (id & TIME_MASK) >> 27
	d.Instanceid = (id >> 19) & INST_MASK
	d.Num = id & NUM_MASK
	return d, nil
}
