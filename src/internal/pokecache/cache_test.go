package pokecache

import (
	"testing"
	"time"
	"fmt"
	// "reflect"
)


func TestCache(t *testing.T) {
	cases := []struct{
		reap_dur time.Duration
		wait_dur bool
		key []string 
		cache_data [][]byte
		expect_data [][]byte
	}{
		{
			5*time.Second,
			false,
			[]string{
				"item1",
				"item2",
				"item3",
			},
			[][]byte{
				{0,1,2},
				{3,4,5},
				{6,7,8},
			},
			[][]byte{
				{0,1,2},
				{3,4,5},
				{6,7,8},
			},
		},
		{
			1*time.Second,
			true,
			[]string{
				"item1",
				"item2",
				"item3",
			},
			[][]byte{
				{0,1,2},
				{3,4,5},
				{6,7,8},
			},
			nil,
		},
	}
	for fural, c := range cases {
		c_ := NewCache(c.reap_dur)
		for idx, _  := range c.key {
			c_.Add(c.key[idx], c.cache_data[idx])
		}
		if c.wait_dur {
			time.Sleep(c.reap_dur)
			time.Sleep(c.reap_dur)
			if len(c_.cmap) > 0 {
				c_.m.Lock()
				for key, item := range c_.cmap {
					fmt.Println(key,item)
				}
				c_.m.Unlock()
				t.Fatalf("Reap Loop is not deleting cache entries in the expected amount of time, Size: %v", len(c_.cmap))
			} else {
				fmt.Printf("[X] Item %v Passed!! \n\n", fural)
			}
		} else {
			fmt.Printf("[X] Item %v Passed!! \n\n", fural)
			// for idx, item := range c.expect_data {
			//
			// }
		}

	}
}
