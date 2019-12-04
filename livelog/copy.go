// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE file.

package livelog

import (
	"bufio"
	"io"
)

// Copy copies from src to dst and removes until either EOF
// is reached on src or an error occurs.
func Copy(dst io.Writer, src io.ReadCloser) error {
	r := bufio.NewReader(src)
	for {
		bytes, err := r.ReadBytes('\n')
		if _, err := dst.Write(bytes); err != nil {
			return err
		}
		if flusher, ok := dst.(Flusher); ok {
			flusher.Flush()
		}
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
	}
}

// The Flusher interface is implemented by an io.Writer that
// flushes buffered data to the client.
type Flusher interface {
	// Flush sends any buffered data to the client.
	Flush() error
}