package gortsplib

import (
	"testing"

	"github.com/bluenviron/gortsplib/v5/pkg/base"
	"github.com/stretchr/testify/require"
)

func TestAffectedGStreamerPublisher(t *testing.T) {
	for _, ca := range []struct {
		name     string
		header   base.HeaderValue
		affected bool
	}{
		{"missing", nil, false},
		{"other-client", base.HeaderValue{"Lavf/63.0"}, false},
		{"unversioned", base.HeaderValue{"GStreamer"}, false},
		{"ambiguous", base.HeaderValue{"GStreamer/1.16.2", "GStreamer/1.24.0"}, false},
		{"old-major", base.HeaderValue{"GStreamer/0.10.36"}, true},
		{"old-minor", base.HeaderValue{"GStreamer/1.14.5"}, true},
		{"last-affected", base.HeaderValue{"GStreamer/1.16.2 (rtspclientsink)"}, true},
		{"fixed-patch", base.HeaderValue{"GStreamer/1.16.3"}, false},
		{"fixed-minor", base.HeaderValue{"GStreamer/1.18.0"}, false},
		{"current", base.HeaderValue{"GStreamer/1.24.12"}, false},
		{"future-major", base.HeaderValue{"GStreamer/2.0.0"}, false},
		{"malformed", base.HeaderValue{"GStreamer/1.16.bad"}, false},
	} {
		t.Run(ca.name, func(t *testing.T) {
			require.Equal(t, ca.affected, isAffectedGStreamerPublisher(ca.header))
		})
	}
}
