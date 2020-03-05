package hdm

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var mt sync.RWMutex
var hotelsDataSource map[string]Hotel
var hotelsIDMaping map[string]string
var cHotel chan QueryHotel

func init() {
	hotelsDataSource = make(map[string]Hotel, 0)
	hotelsIDMaping = make(map[string]string, 0)
	cHotel = make(chan QueryHotel)
	getData()
}

func getData() {
	go fetchV1()
	go fetchV2()
	go fetchV3()
}

func fetchV1() {
	timeC := time.After(3 * time.Second)
	hotels := make(chan []HotelV1)

	go func() {
		hotels <- fetchSource_1()
	}()

	select {
	case <-timeC:
		fmt.Println("V1 time out")
		return
	case hts := <-hotels:
		for _, _h := range hts {
			processData(_h)
		}
	}
}
func fetchV2() {
	timeC := time.After(3 * time.Second)
	hotels := make(chan []HotelV2)

	go func() {
		hotels <- fetchSource_2()
	}()

	select {
	case <-timeC:
		fmt.Println("V2 time out")
		return
	case hts := <-hotels:
		for _, _h := range hts {
			processData(_h)
		}
	}
}

func fetchV3() {
	timeC := time.After(3 * time.Second)
	hotels := make(chan []HotelV3)

	go func() {
		hotels <- fetchSource_3()
	}()

	select {
	case <-timeC:
		fmt.Println("V3 time out")
		return
	case hts := <-hotels:
		for _, _h := range hts {
			processData(_h)
		}
	}
}

var count int = 0

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
