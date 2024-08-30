package main

// Use 'go build main' to build the program
// ./main <cmd> <opt>...
// ./main bash -c 'echo "$(ls $(pwd)/go/bin)"| sed -ne "/sandbox/{s/sandbox/mybox/g;p}"'
// or go run main <cmd> <opt>...

// reference: https://medium.com/hackernoon/strace-in-60-lines-of-go-b4b76e3ecd64
// source: https://github.com/lizrice/strace-from-scratch/blob/master/main.go

// Table.json source: https://filippo.io/linux-syscall-table/TABELLA_64.json
// See https://filippo.io/linux-syscall-table/

import (
	"fmt"
	"os"
	"os/exec"
	"golang.org/x/sys/unix"
  "strings"
  "runtime"
  "errors"
  "syscall"
)

const (
  RedZone = 128 // x86_64 redzone for stack
)

var (
  debug = false
)

func GetSyscallRegByArgIndex(regs *unix.PtraceRegs, idx int) uint64 {
    switch idx {
    case 0:
      return regs.Rdi
    case 1:
      return regs.Rsi
    case 2:
      return regs.Rdx
    case 3:
      return regs.R10
    case 4:
      return regs.R8
    case 5:
      return regs.R9
    default:
      panic(fmt.Sprintf("invalid arg index %d (greater than 5)", idx))
    }
}

func SetSyscallRegByArgIndex(regs *unix.PtraceRegs, idx int, value uint64) {
    switch idx {
    case 0:
      regs.Rdi = value
    case 1:
      regs.Rsi = value
    case 2:
      regs.Rdx = value
    case 3:
      regs.R10 = value
    case 4:
      regs.R8 = value
    case 5:
      regs.R9 = value
    default:
      panic(fmt.Sprintf("invalid arg index %d (greater than 5)", idx))
    }
}

func GetSyscallString(pid int, addr uintptr, cnt int) (string, error) {
  var ss strings.Builder
  var buf []byte
  if cnt > 0 { // fixed-length buffer
    buf = make([]byte, cnt)
  } else if cnt == 0 {
    return "", nil
  } else { // null-terminated c-string
    buf = make([]byte, 128)
  }
  for {
    for i := 0; i < len(buf); i++ {
      buf[i] = 0
    }
    n, err := unix.PtracePeekText(pid, addr, buf)
    if err != nil { // the address may not be valid
      return "", fmt.Errorf("failed to get string in ptrace: %w", err)
    }
    if cnt > 0 && n != len(buf) {
      return "", fmt.Errorf("failed to get string in ptrace: expected %d bytes, got %d", len(buf), n)
    }
    l := len(buf)
    if cnt < 0 { // find c-string length
      for i := 0; i < n; i++ {
        if buf[i] == 0 {
          l = i
          break
        }
      }
    }
    ss.Write(buf[:l])
    if cnt > 0 || l < len(buf) {
      break
    }
  }
  return ss.String(), nil 
}

func CheckWaitStatus(ws unix.WaitStatus, exitOk bool, usePtFlag bool) (rc int, msg string) {
  if debug {
		fmt.Printf("*** Wait status: 0x%x\n", ws)
	}
  if ws.Exited() {
    heading := ""
    if !exitOk {
      heading = "ERROR: "
      rc = 1
    }
    msg = fmt.Sprintf("%sThe tracee has exited (return code: %d)!", heading, ws.ExitStatus())
    return
  } else if ws.Stopped() {
    sig := syscall.SIGTRAP
    if usePtFlag {
      sig |= 0x80
    }
    if ws.StopSignal() != sig {
      msg = fmt.Sprintf("ERROR: The tracee was stopped by a non-tracing signal (%s)!", ws.StopSignal())
      rc = 1
      return
    }
    // normal tracing stop
    return
  } else if ws.Signaled() { // signaled, core dump
    cs := ""
    if ws.CoreDump() {
      cs = " with core dump"
    }
    msg = fmt.Sprintf("ERROR: The tracee was signaled by a non-tracing signal (%s)%s!", ws.Signal(), cs)
    rc = 1
    return
  } else if ws.Continued() {
    msg = fmt.Sprintf("ERROR: The tracee was signaled to continue!")
    rc = 1
    return
  }
  msg = fmt.Sprintf("ERROR: The tracee has a strange wait satus 0x%x!", ws)
  rc = 1
  return
}

func PrintTraceInfo(si SyscallInfo, pid int, regs *unix.PtraceRegs, isExit bool) error {
  err := unix.PtraceGetRegs(pid, regs)
  if err != nil {
    return fmt.Errorf("failed to get registers from ptrace: %w!", err)
  }
  syscallID := regs.Orig_rax
  name := si.GetNameFor(syscallID)
  if isExit {
    rs := ""
    if regs.Rax < 10 {
      rs = fmt.Sprintf("%d", regs.Rax)
    } else if regs.Rax & 0x8000000000000000 != 0 {
      rs = fmt.Sprintf("%d", int64(regs.Rax))
    } else {
      rs = fmt.Sprintf("%d(0x%x)", regs.Rax, regs.Rax)
    }
    fmt.Printf("%d: << Syscall(%d) %s(%s)=%s\n", pid, syscallID, name, si.GetArgStrFor(syscallID, regs, pid), rs)
  } else {
    fmt.Printf("%d: >> Syscall(%d) %s(%s)...\n", pid, syscallID, name, si.GetArgStrFor(syscallID, regs, pid))
  }
  return nil
}

func main() {
  rc := 0
  if len(os.Args) == 1 {
    fmt.Println("ERROR: No command to execute!")
    fmt.Println("Use 'go run main.go <cmd> <arg>...' to run it.")
    os.Exit(1)
  }

  si := SyscallInfoData

	fmt.Printf("Tracing %v...\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
  // Ptrace tells the child to call ptrace(PTRACE_TRACEME).
	// Call runtime.LockOSThread before starting a process with this set,
	// and don't call UnlockOSThread until done with PtraceSyscall calls.
	cmd.SysProcAttr = &unix.SysProcAttr{
		Ptrace: true,
	}
  	
  var regs unix.PtraceRegs
  var isExit bool
  var pid int
  var ws unix.WaitStatus 
  var msg string

  runtime.LockOSThread()
	cmd.Start()
	err := cmd.Wait()
	if err != nil {
    var ee *exec.ExitError
		if errors.As(err, &ee) { 
      if ws, ok := ee.Sys().(syscall.WaitStatus); ok { // for unix
        rc, msg = CheckWaitStatus(unix.WaitStatus(ws), false, false)
        if msg != "" {
          fmt.Println(msg)
          fmt.Println(err)
          goto end
        }
      } else { // no wait status
			  fmt.Printf("ERROR: The tracee has exited (return code : %d) before tracing starts: %s!\n", ee.ExitCode(), err)
        rc = 1
        goto end
      }
		} else { // not ExitError
      fmt.Printf("ERROR: Failed to start the tracee process: %s!\n", err)
      rc = 1
      goto end
    }
    // normal tracing stop
		fmt.Printf("Tracing starts...\n")
	} else { // no error
    fmt.Printf("ERROR: The tracee was not stopped by the expected signal!\n")
    rc = 1
    goto end
  }
  pid = cmd.Process.Pid
  err = PrintTraceInfo(si, pid, &regs, true)
  if err != nil {
    fmt.Printf("ERROR: %s!", err)
    rc = 1
    goto end
  }

	isExit = false
  // add 0x80 in the ptrace signal
  if err = unix.PtraceSetOptions(pid, unix.PTRACE_O_TRACESYSGOOD); err != nil {
    fmt.Printf("ERROR: Failed to ptrace option: %s!", err)
      rc = 1
      goto end
  }

	for {
		err = unix.PtraceSyscall(pid, 0) // same as PTRACE_SINGLESTEP
		if err != nil {
			fmt.Printf("ERROR: Failed to perform singl-step ptrace: %s!\n", err)
      rc = 1
      goto end
		}

		_, err = unix.Wait4(pid, &ws, 0, nil)
		if err != nil {
			fmt.Printf("ERROR: Failed to wait for tracee: %s!\n", err)
      rc = 1
      goto end
		}
    rc, msg = CheckWaitStatus(ws, true, true)
    if msg != "" {
      fmt.Println(msg)
      goto end
    }

    err = PrintTraceInfo(si, pid, &regs, isExit)
    if err != nil {
      fmt.Printf("ERROR: %s!", err)
      rc = 1
      goto end
    }
    if !isExit {
      syscallID := regs.Orig_rax
      if syscallID == unix.SYS_OPEN || syscallID == unix.SYS_OPENAT {
        v, idx, err := si.GetArgValByName(syscallID, &regs, "filename")
        if err != nil {
          fmt.Printf("ERROR: Failed to get filename: %s!", err)
          rc = 1
          goto end
        }
        fn, err := GetSyscallString(pid, uintptr(v), -1)
        if err != nil {
          fmt.Printf("ERROR: Failed to get filename string: %s!", err)
          rc = 1
          goto end
        }
        if fn == "test1.txt" {
          fmt.Println("Found open/openat for test1.txt!")
          newName := "/home/runner/ptrace-test/test.txt"
          buf := append([]byte(newName), '\x00')
          stackAddr := regs.Rsp - RedZone - /*unix.PathMax*/ uint64(len(buf))
          n, err := unix.PtracePokeText(pid, uintptr(stackAddr), buf)
          if err != nil {
            fmt.Printf("ERROR: Failed to set new filename %s: %s!", newName, err)
            rc = 1
            goto end
          } else if n != len(buf) {
            fmt.Printf("ERROR: Failed to set new filename %s: expected %d bytes, got %d!", newName, len(buf), n)
            rc = 1
            goto end
          }
          SetSyscallRegByArgIndex(&regs, idx, uint64(stackAddr))
          err = unix.PtraceSetRegs(pid, &regs)
          if err != nil {
            fmt.Printf("ERROR: Failed to update syscall args with new filename %s: %s!", newName, err)
            rc = 1
            goto end
          }
        }
      }
    }
		isExit = !isExit
	}

end:
  runtime.UnlockOSThread()
  if rc == 0 {
	  fmt.Printf("Done tracing %v.\n", os.Args[1:])
  } else {
    os.Exit(rc)
  }
}