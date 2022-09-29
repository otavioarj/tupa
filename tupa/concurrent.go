package tupa

import (
        "sync"
        "sync/atomic"
)


type WaitGroupCount struct {
    sync.WaitGroup
    count int64
    pok  int32
}

var wg WaitGroupCount

func wgAdd(delta int) {
    atomic.AddInt64(&wg.count, int64(delta))
    wg.WaitGroup.Add(delta)
}

func  wgDone() {
    atomic.AddInt64(&wg.count, -1)
    wg.WaitGroup.Done()
}

func wgGot(){
        atomic.AddInt32(&wg.pok,1)
}


func wgCount() int {
    return int(atomic.LoadInt64(&wg.count))
}


