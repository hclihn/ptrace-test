package main

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/sys/unix"
)

var (
	// UseTab4Indent specifies is TAB should be used for the indent unit
	// used for ToInitializer
	UseTab4Indent = true
	// Spaces4Indent is the number of spaces used for indentation if UseTab4Indent is false
	// used for ToInitializer
	Spaces4Indent = 4
)

// GetIndentStr returns the indent space string for level
func GetIndentStr(level int) string {
	if UseTab4Indent {
		return strings.Repeat("\t", level)
	}
	return strings.Repeat(" ", level*Spaces4Indent)
}

// In x86_64, char=1byte, short=2bytes, int=4bytes, long&(long long)=8bytes; size_t,pointer=8bytes
// float=4bytes, double=8bytes, (long double)=16bytes
// __BITS_PER_LONG is defined per arch

// ArgData stores the syscall argument info
type ArgData struct {
	Ctype string // the c type of the arg
	Name  string // the name of the arg
	Value uint64 // the value from syscall register (untyped)
}

// ToString returns the string representation of the argument
// pid is the PID of the tracee process
func (a ArgData) ToString(pid int) string {
	if strings.Contains(a.Ctype, "char") && strings.Contains(a.Ctype, "*") && a.Name == "filename" {
		s, err := GetSyscallString(pid, uintptr(a.Value), -1)
		if err != nil { // the address may not be valid
			return fmt.Sprintf("%s %s=0x%x", a.Ctype, a.Name, a.Value)
		}
		return fmt.Sprintf("%s %s='%s'", a.Ctype, a.Name, s)
	} else if strings.Contains(a.Ctype, "*") {
		return fmt.Sprintf("%s %s=0x%x", a.Ctype, a.Name, a.Value)
	} else if strings.Contains(a.Ctype, "int") {
		if a.Value < 10 {
			return fmt.Sprintf("%s %s=%d", a.Ctype, a.Name, a.Value)
		}
		if !strings.Contains(a.Ctype, "unsigned") {
			return fmt.Sprintf("%s %s=%d[0x%x]", a.Ctype, a.Name, int64(a.Value), int64(a.Value))
		}
	} else if a.Value < 10 {
		return fmt.Sprintf("%s %s=%d", a.Ctype, a.Name, a.Value)
	}
	return fmt.Sprintf("%s %s=%d[0x%x]", a.Ctype, a.Name, a.Value, a.Value)
}

// ToInitializer returns the object's Go initializer string
// level is the indentation level for the object
// If debug is set, it will add extra comment for debug
// The first line not indented, and the last line has no comma and NL
// This allows the caller to modify it.
func (a ArgData) ToInitializer(level int, debug bool) string {
	var buf strings.Builder
	pr := GetIndentStr(level)
	pr1 := GetIndentStr(level + 1)

	buf.WriteString("ArgData{\n")
	fmt.Fprintf(&buf, "%sCtype: %q,\n", pr1, a.Ctype)
	fmt.Fprintf(&buf, "%sName: %q,\n", pr1, a.Name)
	fmt.Fprintf(&buf, "%s}", pr)
	return buf.String()
}

// ArgumentType indicates the function argument type
type ArgumentType int

// ArgumentType enums
// Make sure String() is updated when enums change!
const (
	ArgNoArgInfo ArgumentType = iota // A syscall w/o argument info (default)
	ArgNormal                        // normal args
	ArgVoid                          // the arg is void (no args)
	ArgNoArgName                     // A syscall w/o arg name(s)
)

// String returns ArgumentType enum's code representation
func (t ArgumentType) String() string {
	switch t {
	case ArgNoArgInfo:
		return "ArgNoArgInfo"
	case ArgNormal:
		return "ArgNormal"
	case ArgVoid:
		return "ArgVoid"
	case ArgNoArgName:
		return "ArgNoArgName"
	}
	panic(fmt.Sprintf("unknown ArgumentType %d", t))
}

// ArgsInfo stores syscall function arguments info
type ArgsInfo struct {
	ArgsType ArgumentType // argument(s) type
	Args     []*ArgData   // argument(s) info
}

// ToInitializer returns the object's Go initializer string
// level is the indentation level for the object
// If debug is set, it will add extra comment for debug
// The first line not indented, and the last line has no comma and NL
// This allows the caller to modify it.
func (a ArgsInfo) ToInitializer(level int, debug bool) string {
	var buf strings.Builder
	var extra string
	if debug {
		extra = " // ArgData"
	}
	pr := GetIndentStr(level)
	pr1 := GetIndentStr(level + 1)
	pr2 := GetIndentStr(level + 2)

	buf.WriteString("ArgsInfo{\n")
	fmt.Fprintf(&buf, "%sArgsType: %s,\n", pr1, a.ArgsType)
	fmt.Fprintf(&buf, "%sArgs: []*ArgData{\n", pr1)
	for _, ad := range a.Args {
		s := ad.ToInitializer(level+2, debug)
		// since we are in an array, remove the data type from the sublevel
		s = strings.Replace(s, "ArgData", "", 1)
		fmt.Fprintf(&buf, "%s%s,%s\n", pr2, s, extra)
	}
	if debug {
		fmt.Fprintf(&buf, "%s}, // Args \n", pr1) // closing Args
	} else {
		fmt.Fprintf(&buf, "%s},\n", pr1) // closing Args
	}
	fmt.Fprintf(&buf, "%s}", pr)
	return buf.String()
}

// SyscallData stores the information about the syscall
type SyscallData struct {
	Number         uint64 // syscall number
	Name           string // syscall function name (w/o the sys_ prefix)
	EntryPoint     string // syscall function name (w/ the sys_ prefix)
	CompatEntry    string // syscall compat function name (w/ the compat_sys_ prefix)
	Implementation string // syscall implementation type (common, i386, x32, 64, etc.)
	ArgsInfo
}

// ToInitializer returns the object's Go initializer string
// level is the indentation level for the object
// If debug is set, it will add extra comment for debug
// The first line not indented, and the last line has no comma and NL
// This allows the caller to modify it.
func (s SyscallData) ToInitializer(level int, debug bool) string {
	var buf strings.Builder
	var extra string
	if debug {
		extra = " // ArgsInfo"
	}
	pr := GetIndentStr(level)
	pr1 := GetIndentStr(level + 1)

	buf.WriteString("SyscallData{\n")
	fmt.Fprintf(&buf, "%sNumber: %v,\n", pr1, s.Number)
	fmt.Fprintf(&buf, "%sName: %q,\n", pr1, s.Name)
	fmt.Fprintf(&buf, "%sEntryPoint: %q,\n", pr1, s.EntryPoint)
	fmt.Fprintf(&buf, "%sCompatEntry: %q,\n", pr1, s.CompatEntry)
	fmt.Fprintf(&buf, "%sImplementation: %q,\n", pr1, s.Implementation)
	fmt.Fprintf(&buf, "%sArgsInfo: %s,%s\n", pr1, s.ArgsInfo.ToInitializer(level+1, debug), extra)
	fmt.Fprintf(&buf, "%s}", pr)
	return buf.String()
}

// SyscallInfo is the map of syscall number to the syscall info
type SyscallInfo map[uint64]*SyscallData

// ToInitializer returns the object's Go initializer string
// level is the indentation level for the object
// If debug is set, it will add extra comment for debug
// The first line not indented, and the last line has no comma and NL
// This allows the caller to modify it.
func (s SyscallInfo) ToInitializer(level int, debug bool) string {
	var buf strings.Builder
	var extra string
	if debug {
		extra = " // SyscallData"
	}
	pr := GetIndentStr(level)
	pr1 := GetIndentStr(level + 1)

	buf.WriteString("SyscallInfo{\n")

	keys := make([]uint64, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, k := range keys {
		fmt.Fprintf(&buf, "%s%d: &%s,%s\n", pr1, k, s[k].ToInitializer(level+1, debug), extra)
	}
	fmt.Fprintf(&buf, "%s}", pr)
	return buf.String()
}

// DataFor returns the syscall info for the syscall num
// Will panic if SyscallInfo map is empty
func (s SyscallInfo) DataFor(num uint64) *SyscallData {
	if len(s) == 0 {
		panic(fmt.Sprintf("failed to get name for %d: empty SyscallInfo", num))
	}
	if v, ok := s[num]; !ok {
		panic(fmt.Sprintf("failed to get name for %d: not exist", num))
	} else {
		return v
	}
}

// GetNameFor returns the name of the syscall num
func (s SyscallInfo) GetNameFor(num uint64) string {
	info := s.DataFor(num)
	return info.Name
}

// GetArgsFor returns the argument info for the syscall num
// regs are used to fill the corresponding arg's value
func (s SyscallInfo) GetArgsFor(num uint64, regs *unix.PtraceRegs) []ArgData {
	info := s.DataFor(num)
	if len(info.Args) == 0 {
		return nil
	}
	ai := make([]ArgData, len(info.Args))
	for i, arg := range info.Args {
		ai[i].Ctype = arg.Ctype
		ai[i].Name = arg.Name
		ai[i].Value = GetSyscallRegByArgIndex(regs, i)
	}
	return ai
}

// GetArgStrFor returns the string representation of the args for syscall num
// pid is the PID of the tracee process
// regs are used to fill the corresponding arg's value
func (s SyscallInfo) GetArgStrFor(num uint64, regs *unix.PtraceRegs, pid int) string {
	ai := s.GetArgsFor(num, regs)
	if len(ai) == 0 {
		return ""
	}
	var ss strings.Builder
	first := true
	for _, arg := range ai {
		if first {
			first = false
		} else {
			ss.WriteString(", ")
		}
		// TODO: add special type processing
		// https://stackoverflow.com/questions/33431994/extracting-system-call-name-and-arguments-using-ptrace
		fmt.Fprintf(&ss, "%s", arg.ToString(pid))
	}
	return ss.String()
}

// GetArgValByName returns the arg's value specified by the arg's name
// the int returned is the arg's index
func (s SyscallInfo) GetArgValByName(num uint64, regs *unix.PtraceRegs, name string) (uint64, int, error) {
	ai := s.GetArgsFor(num, regs)
	if len(ai) == 0 {
		return 0, -1, fmt.Errorf("no args defined for syscall %d", num)
	}
	for i, arg := range ai {
		if arg.Name == name {
			return arg.Value, i, nil
		}
	}
	return 0, -1, fmt.Errorf("no arg name %s found for syscall %d", name, num)
}

// GetSyscallDataByName returns a list of syscall data by the name specified
func (s SyscallInfo) GetSyscallDataByName(name string) []*SyscallData {
	sd := make([]*SyscallData, 0)
	for _, v := range s {
		if name == v.Name {
			sd = append(sd, v)
		}
	}
	if len(sd) == 0 {
		return nil
	}
	return sd
}