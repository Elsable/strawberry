package store

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestCreateAndGet(t *testing.T) {
	s := New()
	res := Resource{ID: "123", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
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
	res1 := Resource{ID: "1", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
	res2 := Resource{ID: "2", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
	res3 := Resource{ID: "3", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
	res4 := Resource{ID: "4", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
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
	res1 := Resource{ID: "1", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
	res2 := Resource{ID: "2", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
	res3 := Resource{ID: "3", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
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
	res1 := Resource{ID: "1", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
	res2 := Resource{ID: "2", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
	res3 := Resource{ID: "3", Type: "Type", Description: "Description", Location: Location{Lat: 0.1, Lon: 0.2}, Timestamp: time.Now()}
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
