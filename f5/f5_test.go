// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f5

import "testing"

func TestMakeURL(t *testing.T) {
	data := [][3]string{
		[3]string{"https://some-url.com", "/resource/action", "https://some-url.com/resource/action"},
		[3]string{"https://some-url.com", "resource/action", "https://some-url.com/resource/action"},
		[3]string{"https://some-url.com/", "/resource/action", "https://some-url.com/resource/action"},
		[3]string{"https://some-url.com/", "resource/action", "https://some-url.com/resource/action"},
	}
	for _, in := range data {
		c := Client{baseURL: in[0]}
		if u := c.makeURL(in[1]); u != in[2] {
			t.Errorf("Client.makeURL('%s'): found '%s', expected '%s'", in[1], u, in[2])
		}
	}
}
