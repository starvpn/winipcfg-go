// Code generated by 'go generate'; DO NOT EDIT.

package winipcfg

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modiphlpapi = windows.NewLazySystemDLL("iphlpapi.dll")

	procCancelMibChangeNotify2 = modiphlpapi.NewProc("CancelMibChangeNotify2")
	procGetAdaptersAddresses   = modiphlpapi.NewProc("GetAdaptersAddresses")
)

func cancelMibChangeNotify2(NotificationHandle uintptr) (result int32) {
	r0, _, _ := syscall.Syscall(procCancelMibChangeNotify2.Addr(), 1, uintptr(NotificationHandle), 0, 0)
	result = int32(r0)
	return
}

func getAdaptersAddresses(Family uint32, Flags uint32, Reserved uintptr, AdapterAddresses *IP_ADAPTER_ADDRESSES, SizePointer *uint32) (result uint32) {
	r0, _, _ := syscall.Syscall6(procGetAdaptersAddresses.Addr(), 5, uintptr(Family), uintptr(Flags), uintptr(Reserved), uintptr(unsafe.Pointer(AdapterAddresses)), uintptr(unsafe.Pointer(SizePointer)), 0)
	result = uint32(r0)
	return
}
