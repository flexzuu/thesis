package client

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/lucas-clemente/quic-go/h2quic"
)

var N = 1
var simulatedRTT = 5 /*30*/ * time.Millisecond

func messure(t *testing.T, toMessure func()) {
	var times []time.Duration
	for n := 0; n < N; n++ {
		t0 := time.Now()
		toMessure()
		t1 := time.Now()
		times = append(times, t1.Sub(t0))
	}
	var sum time.Duration
	var max time.Duration
	var min time.Duration

	for _, t := range times {
		sum += t
		if t > max {
			max = t
		}
		if t < min || min == 0 {
			min = t
		}
	}
	n := time.Duration(len(times))
	avg := sum / n
	fmt.Printf("%s\nAvg: %v, Max: %v, Min: %v\n", t.Name(), avg, max, min)
}
func TestFetchNewsFeed(t *testing.T) {
	messure(t, func() {
		FetchNewsFeed("https://thesis:7771", simulatedRTT, http.Client{})
	})
}

func TestFetchNewsFeedWithAuthor(t *testing.T) {
	messure(t, func() {
		FetchNewsFeedWithAuthor("https://thesis:7771", simulatedRTT, http.Client{})
	})
}

func TestFetchNewsFeedWithAuthorAndComment(t *testing.T) {
	messure(t, func() {
		FetchNewsFeedWithAuthorAndComments("https://thesis:7771", simulatedRTT, http.Client{})
	})
}

func TestFetchNewsFeedQUIC(t *testing.T) {
	messure(t, func() {
		FetchNewsFeed("https://thesis:7772", simulatedRTT, http.Client{Transport: &h2quic.RoundTripper{}})
	})
}

func TestFetchNewsFeedWithAuthorQUIC(t *testing.T) {
	messure(t, func() {
		FetchNewsFeedWithAuthor("https://thesis:7772", simulatedRTT, http.Client{Transport: &h2quic.RoundTripper{}})
	})
}

func TestFetchNewsFeedWithAuthorAndCommentQUIC(t *testing.T) {
	messure(t, func() {
		FetchNewsFeedWithAuthorAndComments("https://thesis:7772", simulatedRTT, http.Client{Transport: &h2quic.RoundTripper{}})
	})
}
