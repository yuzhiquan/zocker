package main

import (
	"os"
	"fmt"
	"os/exec"
	"syscall"
)

func main() {
	switch  os.Args[1] {
	case "run":
		run()
	case "childProc":
		childProc()
	default:
		panic("wtf?")
	}
}

func run() {
	cmd := exec.Command("/proc/self/exe", append([]string{"childProc"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWNS | syscall.CLONE_NEWPID,
	}

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func childProc() {
	fmt.Printf("running %v as pid %d", os.Args[2:], os.Getgid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Chroot("/var/lib/alpine"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	must(cmd.Run())
}
