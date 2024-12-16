package console

import (
	"bufio"
	"os"
	"os/signal"
	"syscall"
	"unsafe"
)

func ReadInputWithConfirm() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func ReadInput() string {
	fd := int(os.Stdin.Fd())
	oldState, err := setRawMode(fd)
	if err != nil {
		panic(err)
	}
	defer restoreState(fd, oldState)

	handleSignal(fd, oldState)

	char, err := readChar()
	if err != nil {
		panic(err)
	}
	return string(char)
}

func setRawMode(fd int) (*syscall.Termios, error) {
	var oldState syscall.Termios
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&oldState)))
	if errno != 0 {
		return nil, errno
	}
	newState := oldState
	newState.Lflag &^= syscall.ICANON | syscall.ECHO
	_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&newState)))
	if errno != 0 {
		return nil, errno
	}
	return &oldState, nil
}

func restoreState(fd int, state *syscall.Termios) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(state)))
	if errno != 0 {
		return errno
	}
	return nil
}

func readChar() (byte, error) {
	var buf [1]byte
	_, err := os.Stdin.Read(buf[:])
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func handleSignal(fd int, state *syscall.Termios) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for range sigChan {
			restoreState(fd, state)
			os.Exit(0)
		}
	}()
}
