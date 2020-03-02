package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
#include <glib-object.h>

int capsRefCount(GstCaps *c) {
	return GST_CAPS_REFCOUNT(c);
}
*/
import "C"

import (
	"github.com/ziutek/glib"
	"unsafe"
)

type (
	Caps      struct{ MiniObject }
	Structure C.GstStructure
)

func (s *Structure) g() *C.GstStructure {
	return (*C.GstStructure)(s)
}

func (c *Caps) g() *C.GstCaps {
	return (*C.GstCaps)(c.GetPtr())
}

func (c *Caps) AppendStructure(media_type string, fields glib.Params) {
	C.gst_caps_append_structure(c.g(), makeGstStructure(media_type, fields))
}

func (c *Caps) GetSize() int {
	return int(C.gst_caps_get_size(c.g()))
}

func (c *Caps) GetStructure(idx uint) *Structure {
	return (*Structure)(unsafe.Pointer(
		C.gst_caps_get_structure(c.g(), C.guint(idx)),
	))
}

func (s *Structure) GetValue(fname string) *glib.Value {
	cs := (*C.gchar)(C.CString(fname))
	defer C.free(unsafe.Pointer(cs))
	return (*glib.Value)(unsafe.Pointer(C.gst_structure_get_value(s.g(), cs)))
}

func (c *Caps) String() string {
	s := (*C.char)(C.gst_caps_to_string(c.g()))
	defer C.free(unsafe.Pointer(s))
	return C.GoString(s)
}

func NewCapsAny() *Caps {
	c := new(Caps)
	c.SetPtr(glib.Pointer(C.gst_caps_new_any()))
	return c
}

func NewCapsEmpty() *Caps {
	c := new(Caps)
	c.SetPtr(glib.Pointer(C.gst_caps_new_empty()))
	return c
}

func NewCapsSimple(media_type string, fields glib.Params) *Caps {
	c := NewCapsEmpty()
	c.AppendStructure(media_type, fields)
	return c
}

func CapsFromString(s string) *Caps {
	cs := (*C.gchar)(C.CString(s))
	defer C.free(unsafe.Pointer(cs))
	c := new(Caps)
	c.SetPtr(glib.Pointer(C.gst_caps_from_string(cs)))
	return c
}
