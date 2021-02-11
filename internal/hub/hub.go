package hub

import (
	"sync"

	"github.com/dwarukira/findcare/internal/event"
)

var log = event.Log
var mutex = sync.Mutex{}
