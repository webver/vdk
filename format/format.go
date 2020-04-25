package format

import (
	"github.com/webver/vdk/av/avutil"
	"github.com/webver/vdk/format/aac"
	"github.com/webver/vdk/format/flv"
	"github.com/webver/vdk/format/mp4"
	"github.com/webver/vdk/format/rtmp"
	"github.com/webver/vdk/format/rtsp"
	"github.com/webver/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
