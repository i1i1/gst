package gst

/*
#cgo pkg-config: gstreamer-1.0 gstreamer-app-1.0
#include <gst/gst.h>
#include <gst/app/gstappsink.h>
*/
import "C"
import "github.com/ziutek/glib"

type (
	Sample   struct{ MiniObject }
	Buffer   struct{ MiniObject }
	BaseSink struct{ Element }
	AppSink  struct{ BaseSink }
)

func (s *Sample) g() *C.GstSample {
	return (*C.GstSample)(s.GetPtr())
}

func (b *Buffer) g() *C.GstBuffer {
	return (*C.GstBuffer)(b.GetPtr())
}

func (a *AppSink) g() *C.GstAppSink {
	return (*C.GstAppSink)(a.GetPtr())
}

func (a *AppSink) PullSample() *Sample {
	s := new(Sample)
	s.SetPtr(glib.Pointer(C.gst_app_sink_pull_sample(a.g())))
	return s
}

func NewBuffer() *Buffer {
	b := new(Buffer)
	b.SetPtr(glib.Pointer(C.gst_buffer_new()))
	return b
}

func (s *Sample) GetBuffer() *Buffer {
	b := new(Buffer)
	b.SetPtr(glib.Pointer(C.gst_sample_get_buffer(s.g())))
	return b
}

func (s *Sample) GetCaps() *Caps {
	r := new(Caps)
	r.SetPtr(glib.Pointer(C.gst_sample_get_caps(s.g())))
	return r
}

func (b *Buffer) GetSize() uint64 {
	return uint64(C.gst_buffer_get_size(b.g()))
}

func (b *Buffer) Extract(off, size uint64) ([]byte, uint64) {
	buf := make([]byte, size, size)
	ret := uint64(C.gst_buffer_extract(
		b.g(),
		C.gsize(off),
		C.gpointer(C.CBytes(buf)),
		C.gsize(size),
	))
	return buf, ret
}
