package http_context

import (
    "net/http"
    "sync"
)

var (
    contextMutex, requestMutex sync.RWMutex
    contexts                   = make(map[*http.Request]Context)
    requests                   = make(map[uint64]*http.Request)
)

type IContext interface {
    Get(interface{}) interface{}
    Set(interface{}, interface{})
}

type Context map[interface{}]interface{}

func (c Context) Get(key interface{}) interface{} {
    return c[key]
}

func (c Context) Set(key interface{}, value interface{}) {
    c[key] = value
}

func GetContext(r *http.Request) (context Context) {
    contextMutex.RLock()
    context = contexts[r]
    contextMutex.RUnlock()
    return
}

func GetRequest() (request *http.Request) {
    goid := CurGoroutineID()
    requestMutex.RLock()
    request = requests[goid]
    requestMutex.RUnlock()
    return
}

func Setup(r *http.Request) {
    goid := CurGoroutineID()
    context := Context(make(map[interface{}]interface{}))

    contextMutex.Lock()
    contexts[r] = context
    contextMutex.Unlock()

    requestMutex.Lock()
    requests[goid] = r
    requestMutex.Unlock()
}

func Teardown(r *http.Request) {
    goid := CurGoroutineID()

    contextMutex.Lock()
    delete(contexts, r)
    contextMutex.Unlock()

    requestMutex.Lock()
    delete(requests, goid)
    requestMutex.Unlock()
}
