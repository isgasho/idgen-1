package idgen

/***************************
*Genter a int64
*|reserved 1bit|timestamp 32bit|instanceid 8bit|bid 6bit|increment 17bit|
****************************/
import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

const (
	VERSION    = 1
	MAX_NUMBER = (1 << 17) - 1
	RAND_MAX   = 1 << 16
	TIME_MASK  = ((1 << 32) - 1)
	INST_MASK  = (1 << 8) - 1
	BID_MASK   = (1 << 6) - 1
	NUM_MASK   = (1 << 17) - 1
)

type IdGen struct {
	mutex sync.RWMutex
	temp  *ID
}

type ID struct {
	Time       int64
	Instanceid int64
	Bid	       int64
	Num        int64
}

func NewIdGen(i byte) *IdGen {
	temp := &ID{
		Instanceid: int64(i) & INST_MASK,
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

func (self *IdGen) Gen(bid int) (id int64) {
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
	self.temp.Bid = int64(bid) & BID_MASK
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
	if i.Num > MAX_NUMBER {
		return 0, errors.New("num is too big")
	}
	if i.Instanceid > 255 {
		return 0, errors.New("invalid instanceid")
	}
	id :=  (i.Time << 31) |
		(i.Instanceid << 23) |(i.Bid << 17) | i.Num
	return id, nil
}

func Decode(id int64) (*ID, error) {
	d := new(ID)
	d.Time       = (id >> 31) & TIME_MASK
	d.Instanceid = (id >> 23) & INST_MASK
	d.Bid        = (id >> 17) & INST_MASK
	d.Num        = id & NUM_MASK
	return d, nil
}
