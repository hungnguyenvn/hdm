package hdm

import (
	"strconv"
	"sync"
)

var mt sync.RWMutex
var hotelsDataSource map[string]Hotel
var hotelsIDMaping map[string]string

func init() {
	hotelsDataSource = make(map[string]Hotel, 0)
	hotelsIDMaping = make(map[string]string, 0)
	getData()
}

func getData() {
	go func() {
		hotels := fetchSource_1()
		mt.Lock()
		defer mt.Unlock()
		for _, h := range hotels {
			processData(h)
		}
	}()
	go func() {
		hotels := fetchSource_2()
		mt.Lock()
		defer mt.Unlock()
		for _, h := range hotels {
			processData(h)
		}
	}()
	go func() {
		hotels := fetchSource_3()
		mt.Lock()
		defer mt.Unlock()
		for _, h := range hotels {
			processData(h)
		}
	}()
}

func processData(h QueryHotel) {
	_h, ok := hotelsDataSource[h.getHotelID()]
	if ok {
		_h.merge(h)
		hotelsDataSource[h.getHotelID()] = _h
	} else {
		hotelsDataSource[h.getHotelID()] = makeHotelFromVarian(h)
		hotelsIDMaping[strconv.Itoa(int(h.getDestinationID()))] = h.getHotelID()
	}
}
