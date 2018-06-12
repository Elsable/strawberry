package store

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"reflect"
)

func TestCreateAndGet(t *testing.T) {
	s := New()
	res := Resource{ID:"123", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	resourceId, err := s.Create(res)
	assert.Nil(t, err)
	assert.Equal(t, "123", resourceId)
	sameRes, err := s.Get("123")
	assert.Nil(t, err)
	assert.EqualValues(t, res, sameRes)
	_, err = s.Get("123456")
	assert.Equal(t, reflect.TypeOf(err).String(), "store.ErrKeyNotFound")
}

func TestList(t *testing.T) {
	s := New()
	res1 := Resource{ID:"1", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	res2 := Resource{ID:"2", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	res3 := Resource{ID:"3", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	res4 := Resource{ID:"4", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	s.Create(res1)
	s.Create(res2)
	s.Create(res3)
	l, err := s.List(0)
	list := *l
	assert.Nil(t, err)
	assert.Equal(t, 3, len(list))
	assert.Contains(t, list, res1)
	assert.Contains(t, list, res2)
	assert.Contains(t, list, res3)
	assert.NotContains(t, list, res4)
}

func TestListWithLimit(t *testing.T) {
	s := New()
	res1 := Resource{ID:"1", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	res2 := Resource{ID:"2", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	res3 := Resource{ID:"3", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	s.Create(res1)
	s.Create(res2)
	s.Create(res3)
	list, err := s.List(5)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(*list))

	list, err = s.List(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(*list))
}

func TestDelete(t *testing.T) {
	s := New()
	res1 := Resource{ID:"1", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	res2 := Resource{ID:"2", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	res3 := Resource{ID:"3", Type:"Type", Description:"Description", Location:Location{Lat:0.1, Lon:0.2}, Timestamp:time.Now()}
	s.Create(res1)
	s.Create(res2)
	s.Create(res3)

	list, err := s.List(0)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(*list))

	err = s.Delete(res1.ID)
	assert.Nil(t, err)

	list, _ = s.List(0)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(*list))

	err = s.Delete(res1.ID)
	list, _ = s.List(0)
	assert.Equal(t, 2, len(*list))
}
/*
func TestIncErr(t *testing.T) {
	s := NewInMemory(time.Second)
	msg := Message{Key: "key123456", Exp: time.Now(), Data: "data string", PinHash: "123456"}
	assert.Nil(t, s.Save(&msg))

	cnt, err := s.IncErr("key123456")
	assert.Nil(t, err)
	assert.Equal(t, 1, cnt)

	cnt, err = s.IncErr("key123456")
	assert.Nil(t, err)
	assert.Equal(t, 2, cnt)

	_, err = s.IncErr("aaakey123456")
	assert.Equal(t, ErrKeyNotFound, err)
}

func TestCleaner(t *testing.T) {
	s := NewInMemory(time.Millisecond * 50)
	msg := Message{Key: "key123456", Exp: time.Now(), Data: "data string", PinHash: "123456"}
	assert.Nil(t, s.Save(&msg), "saved fine")

	_, err := s.Load("key123456")
	assert.Nil(t, err, "key still in store")
	time.Sleep(time.Millisecond * 101)

	_, err = s.Load("key123456")
	assert.Equal(t, ErrKeyNotFound, err, "msg gone")
}*/