// +build windows

package lfs

import (
	"os"
	"syscall"
	"time"

	"github.com/yuin/gopher-lua"
)

func attributesFill(tbl *lua.LTable, stat os.FileInfo) error {
	sys := stat.Sys().(*syscall.Win32FileAttributeData)
	tbl.RawSetH(lua.LString("dev"), lua.LNumber(0))
	tbl.RawSetH(lua.LString("ino"), lua.LNumber(0))

	if stat.IsDir() {
		tbl.RawSetH(lua.LString("mode"), lua.LString("directory"))
	} else {
		tbl.RawSetH(lua.LString("mode"), lua.LString("file"))
	}

	tbl.RawSetH(lua.LString("nlink"), lua.LNumber(0))
	tbl.RawSetH(lua.LString("uid"), lua.LNumber(0))
	tbl.RawSetH(lua.LString("gid"), lua.LNumber(0))
	tbl.RawSetH(lua.LString("rdev"), lua.LNumber(0))

	tbl.RawSetH(lua.LString("access"), lua.LNumber(time.Unix(0, sys.LastAccessTime.Nanoseconds()/1e9).Second()))
	tbl.RawSetH(lua.LString("modification"), lua.LNumber(time.Unix(0, sys.CreationTime.Nanoseconds()/1e9).Second()))
	tbl.RawSetH(lua.LString("change"), lua.LNumber(time.Unix(0, sys.LastWriteTime.Nanoseconds()/1e9).Second()))
	tbl.RawSetH(lua.LString("size"), lua.LNumber(stat.Size()))

	tbl.RawSetH(lua.LString("blocks"), lua.LNumber(0))
	tbl.RawSetH(lua.LString("blksize"), lua.LNumber(0))
	return nil
}
