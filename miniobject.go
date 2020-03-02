package gst

/*
#include <gst/gst.h>
#include <glib-object.h>

static inline
GType _object_type(GstMiniObject* o) {
	return GST_MINI_OBJECT_TYPE(o);
}

static inline
int _object_refcount(GstMiniObject* o) {
	return GST_MINI_OBJECT_REFCOUNT(o);
}
*/
import "C"
import "github.com/ziutek/glib"

type MiniObjectCaster interface {
	AsMiniObject() *MiniObject
}

type MiniObject struct {
	p C.gpointer
}

func (o *MiniObject) g() *C.GstMiniObject {
	return (*C.GstMiniObject)(o.p)
}

func (o *MiniObject) GetPtr() glib.Pointer {
	return glib.Pointer(o.p)
}

func (o *MiniObject) SetPtr(p glib.Pointer) {
	o.p = C.gpointer(p)
}

func (o *MiniObject) Type() glib.Type {
	return glib.Type(C._object_type(o.g()))
}

func (o *MiniObject) AsMiniObject() *MiniObject {
	return o
}

func (o *MiniObject) Value() *glib.Value {
	v := glib.NewValue(o.Type())
	C.g_value_set_boxed(v2g(v), (C.gconstpointer)(o.p))
	return v
}

func (o *MiniObject) Ref() *MiniObject {
	r := new(MiniObject)
	r.SetPtr(glib.Pointer(C.gst_mini_object_ref(o.g())))
	return r
}

func (o *MiniObject) Unref() {
	C.gst_mini_object_unref(o.g())
}

func (o *MiniObject) RefCount() int {
	return int(C._object_refcount(o.g()))
}
