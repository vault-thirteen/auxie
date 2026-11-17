package h

import (
	"syscall"
	"unsafe"

	"github.com/vault-thirteen/auxie/cio"
	"github.com/vault-thirteen/auxie/errors"
	"github.com/vault-thirteen/auxie/os/win32"
	b "github.com/vault-thirteen/auxie/os/win32/base"
	c "github.com/vault-thirteen/auxie/os/win32/common"
	t "github.com/vault-thirteen/auxie/os/win32/type"
)

const (
	Err_MultipleResults           = "multiple results were found"
	Err_MouseCursorMovementFailed = "mouse cursor movement failed"
	Err_WindowIsNotFound          = "window is not found"

	Errf_NoVisibleWindow = "no visible window was found for PID %d"
)

func FindProcessByWindowClassName(className string) (pid uint32, err error) {
	var cn *uint16
	cn, err = syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, err
	}

	hwnd := win32.FindWindowW(cn, nil)
	if hwnd == 0 {
		return 0, errors.ErrorFromText(Err_WindowIsNotFound)
	}

	_, err = win32.GetWindowThreadProcessId(hwnd, &pid)
	if err != nil {
		return 0, err
	}

	return pid, nil
}
func FindMainWindowByPid(targetPid uint32) (foundHwnd b.HWND, err error) {
	var foundHwnds = make([]b.HWND, 0)

	callbackFn := func(hwnd b.HWND, lparam b.LPARAM) b.BOOL {
		var pid b.DWORD
		_, err = win32.GetWindowThreadProcessId(hwnd, &pid)
		if err != nil {
			return b.TRUE // Continue enumerating
		}

		if pid != uint32(lparam) {
			return b.TRUE // Continue enumerating
		}

		if win32.IsWindowVisible(hwnd) != b.TRUE {
			return b.TRUE // Continue enumerating
		}

		foundHwnds = append(foundHwnds, hwnd)
		return b.TRUE // Continue enumerating
	}

	_, err = win32.EnumWindows(callbackFn, uintptr(targetPid))
	if err != nil {
		return 0, err
	}

	if len(foundHwnds) == 0 {
		return 0, nil
	}

	if len(foundHwnds) > 1 {
		return 0, errors.ErrorFromText(Err_MultipleResults)
	}

	return foundHwnds[0], nil
}
func MaximizeWindowByPid(targetPid uint32) (err error) {
	var foundHwnd b.HWND
	foundHwnd, err = FindMainWindowByPid(targetPid)
	if err != nil {
		return err
	}

	if foundHwnd == 0 {
		return errors.ErrorFromFormattedText(Errf_NoVisibleWindow, targetPid)
	}

	win32.ShowWindow(foundHwnd, b.SW_MAXIMIZE)

	_, err = win32.SetForegroundWindow(foundHwnd)
	if err != nil {
		return err
	}

	return nil
}

func SendMouseEvents(inputs []t.MouseInputStruct) (err error) {
	if len(inputs) == 0 {
		return nil
	}

	_, err = win32.SendInput(b.UINT(len(inputs)), unsafe.Pointer(&inputs[0]), int32(unsafe.Sizeof(inputs[0])))
	if err != nil {
		return err
	}

	return nil
}
func SendMouseMoveEvent() (err error) {
	mv := t.MouseInputStruct{
		Type: t.INPUT_MOUSE,
		Mi: t.MOUSEINPUT{
			DwFlags: b.MOUSEEVENTF_MOVE,
		},
	}

	return SendMouseEvents([]t.MouseInputStruct{mv})
}
func MakeClickLMB() (err error) {
	leftDown := t.MouseInputStruct{
		Type: t.INPUT_MOUSE,
		Mi: t.MOUSEINPUT{
			DwFlags: b.MOUSEEVENTF_LEFTDOWN,
		},
	}

	leftUp := t.MouseInputStruct{
		Type: t.INPUT_MOUSE,
		Mi: t.MOUSEINPUT{
			DwFlags: b.MOUSEEVENTF_LEFTUP,
		},
	}

	return SendMouseEvents([]t.MouseInputStruct{leftDown, leftUp})
}
func MakeClickRMB() (err error) {
	rightDown := t.MouseInputStruct{
		Type: t.INPUT_MOUSE,
		Mi: t.MOUSEINPUT{
			DwFlags: b.MOUSEEVENTF_RIGHTDOWN,
		},
	}

	rightUp := t.MouseInputStruct{
		Type: t.INPUT_MOUSE,
		Mi: t.MOUSEINPUT{
			DwFlags: b.MOUSEEVENTF_RIGHTUP,
		},
	}

	return SendMouseEvents([]t.MouseInputStruct{rightDown, rightUp})
}

func SendKeyboardEvents(inputs []t.KeyboardInputStruct) (err error) {
	if len(inputs) == 0 {
		return nil
	}

	_, err = win32.SendInput(b.UINT(len(inputs)), unsafe.Pointer(&inputs[0]), int32(unsafe.Sizeof(inputs[0])))
	if err != nil {
		return err
	}

	return nil
}
func SendKey(vk uint16, flags uint32) (err error) {
	k := t.KeyboardInputStruct{
		Type: t.INPUT_KEYBOARD,
		Ki: t.KEYBDINPUT{
			WVk:     vk,
			DwFlags: flags,
		},
	}

	return SendKeyboardEvents([]t.KeyboardInputStruct{k})
}

func GetWindowSize(hWnd b.HWND) (s *t.SIZE, err error) {
	var r t.RECT
	_, err = win32.GetWindowRect(hWnd, &r)
	if err != nil {
		return nil, err
	}

	// Top-left corner is point (0;0), =>
	// X axis goes to the right, Y axis goes downwards.
	s = &t.SIZE{
		X: r.Right - r.Left,
		Y: r.Bottom - r.Top,
	}

	return s, nil
}

func ClickKey(keyCode cio.KeyCode) (err error) {
	err = SendKey(uint16(keyCode), 0)
	if err != nil {
		return err
	}
	err = SendKey(uint16(keyCode), b.KEYEVENTF_KEYUP)
	if err != nil {
		return err
	}
	return nil
}
func ClickEnter() (err error) {
	return ClickKey(b.VK_RETURN)
}
func ClickCtrlF4() (err error) {
	err = SendKey(uint16(cio.VK_CONTROL), 0)
	if err != nil {
		return err
	}
	err = SendKey(uint16(cio.VK_F4), 0)
	if err != nil {
		return err
	}
	err = SendKey(uint16(cio.VK_F4), b.KEYEVENTF_KEYUP)
	if err != nil {
		return err
	}
	err = SendKey(uint16(cio.VK_CONTROL), b.KEYEVENTF_KEYUP)
	if err != nil {
		return err
	}
	return nil
}

func MoveMouseCursorToPos(pos *c.Pos[int32]) (err error) {
	if pos == nil {
		return nil
	}

	if win32.SetCursorPos(pos.X, pos.Y) != b.TRUE {
		return errors.ErrorFromText(Err_MouseCursorMovementFailed)
	}

	return nil
}
