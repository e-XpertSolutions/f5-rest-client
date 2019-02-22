package f5

/*
import (
	"sync"
	"time"
)

// This file implements static shared resource to apply some limitation on the
// authentication. Indeed, in case of credentials error, for instance, the
// the Big-IP may be flooded which causes blocking. As a result, a per-base URL
// authentication control is needed to prevent flooding.

// MaxAuth defines the maximum number of authentication per base URL in a
// time window.
var MaxAuth = 10

// TimeWindowSize defines the time window to use to throttle the authentication
// requests.
var TimeWindowSize = 5 * time.Minute

// mutex for controlling concurrent accesses on the timeWindowPerBaseURL map.
var mu sync.Mutex

var timeWindowPerBaseURL = make(map[string]timeWindow)

// none is just there to give more semantic to struct{} and to render better
// in the code than just struct{}{}.
type none struct{}

// timeWindow provides throttling for requests on a specific base URL.
type timeWindow struct {
	c    chan none
	tick *time.Ticker
	mu   sync.Mutex
}

func newTimeWindow() *timeWindow {
	tw := &timeWindow{
		c:    make(chan none, MaxAuth),
		tick: time.NewTicker(TimeWindowSize),
	}
	// This asynchronous routine will remove an item from the channel at
	// every "tick", defined as the TimeWindowSize.
	go func() {
		for range tw.tick.C {
			select {
			case <-tw.c:
			default:
			}
		}
	}()
	return tw
}

func (tw *timeWindow) can() bool {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	yes := len(tw.c) < cap(tw.c)
	if yes {
		tw.c <- none{}
	}
	return yes
}

var windowPerBaseURL = make(map[string]*timeWindow)

func canAuth(baseURL string) bool {
	mu.Lock()
	defer mu.Unlock()
	tw, ok := windowPerBaseURL[baseURL]
	if !ok {
		tw = newTimeWindow()
		windowPerBaseURL[baseURL] = tw
	}
	return tw.can()
}

*/
