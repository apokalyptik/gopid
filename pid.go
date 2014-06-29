package pid

import (
	"fmt"
	"os"
	"syscall"
)

func Do(filename string, permissions ...uint32) (*os.File, error) {
	if len(permissions) == 0 {
		permissions = []uint32{0666}
	}
	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.FileMode(permissions[0]))
	if err != nil {
		return nil, err
	}
	err = syscall.Flock(int(fp.Fd()), syscall.LOCK_NB|syscall.LOCK_EX)
	if err != nil {
		return nil, err
	}
	syscall.Ftruncate(int(fp.Fd()), 0)
	syscall.Write(int(fp.Fd()), []byte(fmt.Sprintf("%d", os.Getpid())))
	return fp, nil
}
