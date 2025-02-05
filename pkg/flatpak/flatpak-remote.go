// Code generated by girgen. DO NOT EDIT.

package flatpak

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <flatpak.h>
// #include <glib-object.h>
import "C"

// glib.Type values for flatpak-remote.go.
var (
	GTypeRemoteType = externglib.Type(C.flatpak_remote_type_get_type())
	GTypeRemote     = externglib.Type(C.flatpak_remote_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeRemoteType, F: marshalRemoteType},
		{T: GTypeRemote, F: marshalRemote},
	})
}

// RemoteType: different types of FlatpakRemote.
type RemoteType C.gint

const (
	// RemoteTypeStatic: statically configured remote.
	RemoteTypeStatic RemoteType = iota
	// RemoteTypeUsb: dynamically detected local pathname remote.
	RemoteTypeUsb
	// RemoteTypeLan: dynamically detected network remote.
	RemoteTypeLan
)

func marshalRemoteType(p uintptr) (interface{}, error) {
	return RemoteType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for RemoteType.
func (r RemoteType) String() string {
	switch r {
	case RemoteTypeStatic:
		return "Static"
	case RemoteTypeUsb:
		return "Usb"
	case RemoteTypeLan:
		return "Lan"
	default:
		return fmt.Sprintf("RemoteType(%d)", r)
	}
}

// RemoteOverrider contains methods that are overridable.
type RemoteOverrider interface {
}

type Remote struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Remote)(nil)
)

func classInitRemoter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRemote(obj *externglib.Object) *Remote {
	return &Remote{
		Object: obj,
	}
}

func marshalRemote(p uintptr) (interface{}, error) {
	return wrapRemote(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRemote returns a new remote object which can be used to configure a new
// remote.
//
// Note: This is a local configuration object, you must commit changes using
// flatpak_installation_modify_remote() or flatpak_installation_add_remote() for
// the changes to take effect.
//
// The function takes the following parameters:
//
//   - name: name.
//
// The function returns the following values:
//
//   - remote: new Remote.
//
func NewRemote(name string) *Remote {
	var _arg1 *C.char          // out
	var _cret *C.FlatpakRemote // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.flatpak_remote_new(_arg1)
	runtime.KeepAlive(name)

	var _remote *Remote // out

	_remote = wrapRemote(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _remote
}

// NewRemoteFromFile returns a new pre-filled remote object which can be used to
// configure a new remote. The fields in the remote are filled in according to
// the values in the passed in flatpakrepo file.
//
// Note: This is a local configuration object, you must commit changes using
// flatpak_installation_modify_remote() or flatpak_installation_add_remote() for
// the changes to take effect.
//
// The function takes the following parameters:
//
//   - name: name.
//   - data: content of a flatpakrepo file.
//
// The function returns the following values:
//
//   - remote: new Remote, or NULL on error.
//
func NewRemoteFromFile(name string, data *glib.Bytes) (*Remote, error) {
	var _arg1 *C.char          // out
	var _arg2 *C.GBytes        // out
	var _cret *C.FlatpakRemote // in
	var _cerr *C.GError        // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(data)))

	_cret = C.flatpak_remote_new_from_file(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(name)
	runtime.KeepAlive(data)

	var _remote *Remote // out
	var _goerr error    // out

	_remote = wrapRemote(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _remote, _goerr
}

// AppstreamDir returns the directory where this remote will store locally
// cached appstream information for the specified arch.
//
// The function takes the following parameters:
//
//   - arch (optional): which architecture to fetch (default: current
//     architecture).
//
// The function returns the following values:
//
//   - file: #GFile.
//
func (self *Remote) AppstreamDir(arch string) *gio.File {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out
	var _cret *C.GFile         // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if arch != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(arch)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.flatpak_remote_get_appstream_dir(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(arch)

	var _file *gio.File // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_file = &gio.File{
			Object: obj,
		}
	}

	return _file
}

// AppstreamTimestamp returns the timestamp file that will be updated whenever
// the appstream information has been updated (or tried to update) for the
// specified arch.
//
// The function takes the following parameters:
//
//   - arch (optional): which architecture to fetch (default: current
//     architecture).
//
// The function returns the following values:
//
//   - file: #GFile.
//
func (self *Remote) AppstreamTimestamp(arch string) *gio.File {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out
	var _cret *C.GFile         // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if arch != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(arch)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.flatpak_remote_get_appstream_timestamp(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(arch)

	var _file *gio.File // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_file = &gio.File{
			Object: obj,
		}
	}

	return _file
}

// CollectionID returns the repository collection ID of this remote, if set.
//
// The function returns the following values:
//
//   - utf8 (optional): collection ID, or NULL if unset.
//
func (self *Remote) CollectionID() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_collection_id(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// Comment returns the comment of the remote.
//
// The function returns the following values:
//
//   - utf8: comment.
//
func (self *Remote) Comment() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_comment(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DefaultBranch returns the default branch configured for the remote.
//
// The function returns the following values:
//
//   - utf8: default branch, or NULL.
//
func (self *Remote) DefaultBranch() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_default_branch(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Description returns the description of the remote.
//
// The function returns the following values:
//
//   - utf8: description.
//
func (self *Remote) Description() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_description(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Disabled returns whether this remote is disabled.
//
// The function returns the following values:
//
//   - ok: whether the remote is marked as disabled.
//
func (self *Remote) Disabled() bool {
	var _arg0 *C.FlatpakRemote // out
	var _cret C.gboolean       // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_disabled(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Filter returns the filter file of the remote.
//
// The function returns the following values:
//
//   - utf8: pathname to a filter file.
//
func (self *Remote) Filter() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_filter(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// GpgVerify returns whether GPG verification is enabled for the remote.
//
// The function returns the following values:
//
//   - ok: whether GPG verification is enabled.
//
func (self *Remote) GpgVerify() bool {
	var _arg0 *C.FlatpakRemote // out
	var _cret C.gboolean       // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_gpg_verify(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homepage returns the homepage url of the remote.
//
// The function returns the following values:
//
//   - utf8: homepage url.
//
func (self *Remote) Homepage() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_homepage(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Icon returns the icon url of the remote.
//
// The function returns the following values:
//
//   - utf8: icon url.
//
func (self *Remote) Icon() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_icon(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Name returns the name of the remote repository.
//
// The function returns the following values:
//
//   - utf8: name.
//
func (self *Remote) Name() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Nodeps returns whether this remote should be used to find dependencies.
//
// The function returns the following values:
//
//   - ok: whether the remote is marked as "don't use for dependencies".
//
func (self *Remote) Nodeps() bool {
	var _arg0 *C.FlatpakRemote // out
	var _cret C.gboolean       // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_nodeps(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Noenumerate returns whether this remote should be used to list applications.
//
// The function returns the following values:
//
//   - ok: whether the remote is marked as "don't enumerate".
//
func (self *Remote) Noenumerate() bool {
	var _arg0 *C.FlatpakRemote // out
	var _cret C.gboolean       // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_noenumerate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Prio returns the priority for the remote.
//
// The function returns the following values:
//
//   - gint: priority.
//
func (self *Remote) Prio() int {
	var _arg0 *C.FlatpakRemote // out
	var _cret C.int            // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_prio(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RemoteType: get the value of Remote:type.
//
// The function returns the following values:
//
//   - remoteType: type of remote this is.
//
func (self *Remote) RemoteType() RemoteType {
	var _arg0 *C.FlatpakRemote    // out
	var _cret C.FlatpakRemoteType // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_remote_type(_arg0)
	runtime.KeepAlive(self)

	var _remoteType RemoteType // out

	_remoteType = RemoteType(_cret)

	return _remoteType
}

// Title returns the title of the remote.
//
// The function returns the following values:
//
//   - utf8: title.
//
func (self *Remote) Title() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_title(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// URL returns the repository URL of this remote.
//
// The function returns the following values:
//
//   - utf8: URL.
//
func (self *Remote) URL() string {
	var _arg0 *C.FlatpakRemote // out
	var _cret *C.char          // in

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_remote_get_url(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// SetCollectionID sets the repository collection ID of this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - collectionId (optional): new collection ID, or NULL to unset.
//
func (self *Remote) SetCollectionID(collectionId string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if collectionId != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(collectionId)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.flatpak_remote_set_collection_id(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(collectionId)
}

// SetComment sets the comment of this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - comment: new comment.
//
func (self *Remote) SetComment(comment string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(comment)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_comment(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(comment)
}

// SetDefaultBranch sets the default branch configured for this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - defaultBranch: new default_branch, or NULL to unset.
//
func (self *Remote) SetDefaultBranch(defaultBranch string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(defaultBranch)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_default_branch(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(defaultBranch)
}

// SetDescription sets the description of this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - description: new description.
//
func (self *Remote) SetDescription(description string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(description)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_description(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(description)
}

// SetDisabled sets the disabled config of this remote. See
// flatpak_remote_get_disabled().
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - disabled: bool.
//
func (self *Remote) SetDisabled(disabled bool) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if disabled {
		_arg1 = C.TRUE
	}

	C.flatpak_remote_set_disabled(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(disabled)
}

// SetFilter sets a filter for this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - filterPath: pathname of the new filter file.
//
func (self *Remote) SetFilter(filterPath string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filterPath)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_filter(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(filterPath)
}

// SetGpgKey sets the trusted gpg key for this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - gpgKey with gpg binary key data.
//
func (self *Remote) SetGpgKey(gpgKey *glib.Bytes) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.GBytes        // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(gpgKey)))

	C.flatpak_remote_set_gpg_key(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(gpgKey)
}

// SetGpgVerify sets the gpg_verify config of this remote. See
// flatpak_remote_get_gpg_verify().
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - gpgVerify: bool.
//
func (self *Remote) SetGpgVerify(gpgVerify bool) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if gpgVerify {
		_arg1 = C.TRUE
	}

	C.flatpak_remote_set_gpg_verify(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(gpgVerify)
}

// SetHomepage sets the homepage of this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - homepage: new homepage.
//
func (self *Remote) SetHomepage(homepage string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(homepage)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_homepage(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(homepage)
}

// SetIcon sets the homepage of this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - icon: new homepage.
//
func (self *Remote) SetIcon(icon string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(icon)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_icon(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(icon)
}

// SetNodeps sets the nodeps config of this remote. See
// flatpak_remote_get_nodeps().
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - nodeps: bool.
//
func (self *Remote) SetNodeps(nodeps bool) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if nodeps {
		_arg1 = C.TRUE
	}

	C.flatpak_remote_set_nodeps(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(nodeps)
}

// SetNoenumerate sets the noenumeration config of this remote. See
// flatpak_remote_get_noenumerate().
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - noenumerate: bool.
//
func (self *Remote) SetNoenumerate(noenumerate bool) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if noenumerate {
		_arg1 = C.TRUE
	}

	C.flatpak_remote_set_noenumerate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(noenumerate)
}

// SetPrio sets the prio config of this remote. See flatpak_remote_get_prio().
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - prio: bool.
//
func (self *Remote) SetPrio(prio int) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 C.int            // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.int(prio)

	C.flatpak_remote_set_prio(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(prio)
}

// SetTitle sets the repository title of this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - title: new title, or NULL to unset.
//
func (self *Remote) SetTitle(title string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_title(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(title)
}

// SetURL sets the repository URL of this remote.
//
// Note: This is a local modification of this object, you must commit changes
// using flatpak_installation_modify_remote() for the changes to take effect.
//
// The function takes the following parameters:
//
//   - url: new url.
//
func (self *Remote) SetURL(url string) {
	var _arg0 *C.FlatpakRemote // out
	var _arg1 *C.char          // out

	_arg0 = (*C.FlatpakRemote)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(url)))
	defer C.free(unsafe.Pointer(_arg1))

	C.flatpak_remote_set_url(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(url)
}
