package main

// SyscallInfoData is the SyscallInfo object
// Total 382 syscalls.
var SyscallInfoData = SyscallInfo{
	0: &SyscallData{
		Number: 0,
		Name: "read",
		EntryPoint: "sys_read",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "count",
				},
			},
		},
	},
	1: &SyscallData{
		Number: 1,
		Name: "write",
		EntryPoint: "sys_write",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "count",
				},
			},
		},
	},
	2: &SyscallData{
		Number: 2,
		Name: "open",
		EntryPoint: "sys_open",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	3: &SyscallData{
		Number: 3,
		Name: "close",
		EntryPoint: "sys_close",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
			},
		},
	},
	4: &SyscallData{
		Number: 4,
		Name: "stat",
		EntryPoint: "sys_newstat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "struct stat *",
					Name: "statbuf",
				},
			},
		},
	},
	5: &SyscallData{
		Number: 5,
		Name: "fstat",
		EntryPoint: "sys_newfstat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "struct stat *",
					Name: "statbuf",
				},
			},
		},
	},
	6: &SyscallData{
		Number: 6,
		Name: "lstat",
		EntryPoint: "sys_newlstat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "struct stat *",
					Name: "statbuf",
				},
			},
		},
	},
	7: &SyscallData{
		Number: 7,
		Name: "poll",
		EntryPoint: "sys_poll",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct pollfd *",
					Name: "ufds",
				},
				{
					Ctype: "unsigned int",
					Name: "nfds",
				},
				{
					Ctype: "int",
					Name: "timeout",
				},
			},
		},
	},
	8: &SyscallData{
		Number: 8,
		Name: "lseek",
		EntryPoint: "sys_lseek",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "off_t",
					Name: "offset",
				},
				{
					Ctype: "unsigned int",
					Name: "whence",
				},
			},
		},
	},
	9: &SyscallData{
		Number: 9,
		Name: "mmap",
		EntryPoint: "sys_mmap",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "addr",
				},
				{
					Ctype: "unsigned long",
					Name: "len",
				},
				{
					Ctype: "unsigned long",
					Name: "prot",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "off_t",
					Name: "pgoff",
				},
			},
		},
	},
	10: &SyscallData{
		Number: 10,
		Name: "mprotect",
		EntryPoint: "sys_mprotect",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned long",
					Name: "prot",
				},
			},
		},
	},
	11: &SyscallData{
		Number: 11,
		Name: "munmap",
		EntryPoint: "sys_munmap",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "addr",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
			},
		},
	},
	12: &SyscallData{
		Number: 12,
		Name: "brk",
		EntryPoint: "sys_brk",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "brk",
				},
			},
		},
	},
	13: &SyscallData{
		Number: 13,
		Name: "rt_sigaction",
		EntryPoint: "sys_rt_sigaction",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "sig",
				},
				{
					Ctype: "const struct sigaction *",
					Name: "act",
				},
				{
					Ctype: "struct sigaction *",
					Name: "oact",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	14: &SyscallData{
		Number: 14,
		Name: "rt_sigprocmask",
		EntryPoint: "sys_rt_sigprocmask",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "how",
				},
				{
					Ctype: "sigset_t *",
					Name: "set",
				},
				{
					Ctype: "sigset_t *",
					Name: "oset",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	15: &SyscallData{
		Number: 15,
		Name: "rt_sigreturn",
		EntryPoint: "sys_rt_sigreturn",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct pt_regs *",
					Name: "regs",
				},
			},
		},
	},
	16: &SyscallData{
		Number: 16,
		Name: "ioctl",
		EntryPoint: "sys_ioctl",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
				{
					Ctype: "unsigned long",
					Name: "arg",
				},
			},
		},
	},
	17: &SyscallData{
		Number: 17,
		Name: "pread64",
		EntryPoint: "sys_pread64",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "count",
				},
				{
					Ctype: "loff_t",
					Name: "pos",
				},
			},
		},
	},
	18: &SyscallData{
		Number: 18,
		Name: "pwrite64",
		EntryPoint: "sys_pwrite64",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "count",
				},
				{
					Ctype: "loff_t",
					Name: "pos",
				},
			},
		},
	},
	19: &SyscallData{
		Number: 19,
		Name: "readv",
		EntryPoint: "sys_readv",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
			},
		},
	},
	20: &SyscallData{
		Number: 20,
		Name: "writev",
		EntryPoint: "sys_writev",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
			},
		},
	},
	21: &SyscallData{
		Number: 21,
		Name: "access",
		EntryPoint: "sys_access",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "int",
					Name: "mode",
				},
			},
		},
	},
	22: &SyscallData{
		Number: 22,
		Name: "pipe",
		EntryPoint: "sys_pipe",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int *",
					Name: "fildes",
				},
			},
		},
	},
	23: &SyscallData{
		Number: 23,
		Name: "select",
		EntryPoint: "sys_select",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "n",
				},
				{
					Ctype: "fd_set *",
					Name: "inp",
				},
				{
					Ctype: "fd_set *",
					Name: "outp",
				},
				{
					Ctype: "fd_set *",
					Name: "exp",
				},
				{
					Ctype: "struct __kernel_old_timeval *",
					Name: "tvp",
				},
			},
		},
	},
	24: &SyscallData{
		Number: 24,
		Name: "sched_yield",
		EntryPoint: "sys_sched_yield",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	25: &SyscallData{
		Number: 25,
		Name: "mremap",
		EntryPoint: "sys_mremap",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "addr",
				},
				{
					Ctype: "unsigned long",
					Name: "old_len",
				},
				{
					Ctype: "unsigned long",
					Name: "new_len",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
				{
					Ctype: "unsigned long",
					Name: "new_addr",
				},
			},
		},
	},
	26: &SyscallData{
		Number: 26,
		Name: "msync",
		EntryPoint: "sys_msync",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	27: &SyscallData{
		Number: 27,
		Name: "mincore",
		EntryPoint: "sys_mincore",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned char *",
					Name: "vec",
				},
			},
		},
	},
	28: &SyscallData{
		Number: 28,
		Name: "madvise",
		EntryPoint: "sys_madvise",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "int",
					Name: "behavior",
				},
			},
		},
	},
	29: &SyscallData{
		Number: 29,
		Name: "shmget",
		EntryPoint: "sys_shmget",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "key_t",
					Name: "key",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
				{
					Ctype: "int",
					Name: "flag",
				},
			},
		},
	},
	30: &SyscallData{
		Number: 30,
		Name: "shmat",
		EntryPoint: "sys_shmat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "shmid",
				},
				{
					Ctype: "char *",
					Name: "shmaddr",
				},
				{
					Ctype: "int",
					Name: "shmflg",
				},
			},
		},
	},
	31: &SyscallData{
		Number: 31,
		Name: "shmctl",
		EntryPoint: "sys_shmctl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "shmid",
				},
				{
					Ctype: "int",
					Name: "cmd",
				},
				{
					Ctype: "struct shmid_ds *",
					Name: "buf",
				},
			},
		},
	},
	32: &SyscallData{
		Number: 32,
		Name: "dup",
		EntryPoint: "sys_dup",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fildes",
				},
			},
		},
	},
	33: &SyscallData{
		Number: 33,
		Name: "dup2",
		EntryPoint: "sys_dup2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "oldfd",
				},
				{
					Ctype: "unsigned int",
					Name: "newfd",
				},
			},
		},
	},
	34: &SyscallData{
		Number: 34,
		Name: "pause",
		EntryPoint: "sys_pause",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	35: &SyscallData{
		Number: 35,
		Name: "nanosleep",
		EntryPoint: "sys_nanosleep",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct __kernel_timespec *",
					Name: "rqtp",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "rmtp",
				},
			},
		},
	},
	36: &SyscallData{
		Number: 36,
		Name: "getitimer",
		EntryPoint: "sys_getitimer",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "struct __kernel_old_itimerval *",
					Name: "value",
				},
			},
		},
	},
	37: &SyscallData{
		Number: 37,
		Name: "alarm",
		EntryPoint: "sys_alarm",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "seconds",
				},
			},
		},
	},
	38: &SyscallData{
		Number: 38,
		Name: "setitimer",
		EntryPoint: "sys_setitimer",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "struct __kernel_old_itimerval *",
					Name: "value",
				},
				{
					Ctype: "struct __kernel_old_itimerval *",
					Name: "ovalue",
				},
			},
		},
	},
	39: &SyscallData{
		Number: 39,
		Name: "getpid",
		EntryPoint: "sys_getpid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	40: &SyscallData{
		Number: 40,
		Name: "sendfile",
		EntryPoint: "sys_sendfile64",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "out_fd",
				},
				{
					Ctype: "int",
					Name: "in_fd",
				},
				{
					Ctype: "loff_t *",
					Name: "offset",
				},
				{
					Ctype: "size_t",
					Name: "count",
				},
			},
		},
	},
	41: &SyscallData{
		Number: 41,
		Name: "socket",
		EntryPoint: "sys_socket",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "family",
				},
				{
					Ctype: "int",
					Name: "type",
				},
				{
					Ctype: "int",
					Name: "protocol",
				},
			},
		},
	},
	42: &SyscallData{
		Number: 42,
		Name: "connect",
		EntryPoint: "sys_connect",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int",
					Name: "addrlen",
				},
			},
		},
	},
	43: &SyscallData{
		Number: 43,
		Name: "accept",
		EntryPoint: "sys_accept",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int *",
					Name: "addrlen",
				},
			},
		},
	},
	44: &SyscallData{
		Number: 44,
		Name: "sendto",
		EntryPoint: "sys_sendto",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "void *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int",
					Name: "addrlen",
				},
			},
		},
	},
	45: &SyscallData{
		Number: 45,
		Name: "recvfrom",
		EntryPoint: "sys_recvfrom",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "void *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int *",
					Name: "addrlen",
				},
			},
		},
	},
	46: &SyscallData{
		Number: 46,
		Name: "sendmsg",
		EntryPoint: "sys_sendmsg",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct user_msghdr *",
					Name: "msg",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
			},
		},
	},
	47: &SyscallData{
		Number: 47,
		Name: "recvmsg",
		EntryPoint: "sys_recvmsg",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct user_msghdr *",
					Name: "msg",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
			},
		},
	},
	48: &SyscallData{
		Number: 48,
		Name: "shutdown",
		EntryPoint: "sys_shutdown",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "sockfd",
				},
				{
					Ctype: "int",
					Name: "how",
				},
			},
		},
	},
	49: &SyscallData{
		Number: 49,
		Name: "bind",
		EntryPoint: "sys_bind",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int",
					Name: "addrlen",
				},
			},
		},
	},
	50: &SyscallData{
		Number: 50,
		Name: "listen",
		EntryPoint: "sys_listen",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "int",
					Name: "backlog",
				},
			},
		},
	},
	51: &SyscallData{
		Number: 51,
		Name: "getsockname",
		EntryPoint: "sys_getsockname",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int *",
					Name: "addrlen",
				},
			},
		},
	},
	52: &SyscallData{
		Number: 52,
		Name: "getpeername",
		EntryPoint: "sys_getpeername",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int *",
					Name: "addrlen",
				},
			},
		},
	},
	53: &SyscallData{
		Number: 53,
		Name: "socketpair",
		EntryPoint: "sys_socketpair",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "family",
				},
				{
					Ctype: "int",
					Name: "type",
				},
				{
					Ctype: "int",
					Name: "protocol",
				},
				{
					Ctype: "int *",
					Name: "usockvec",
				},
			},
		},
	},
	54: &SyscallData{
		Number: 54,
		Name: "setsockopt",
		EntryPoint: "sys_setsockopt",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "int",
					Name: "level",
				},
				{
					Ctype: "int",
					Name: "optname",
				},
				{
					Ctype: "char *",
					Name: "optval",
				},
				{
					Ctype: "int",
					Name: "optlen",
				},
			},
		},
	},
	55: &SyscallData{
		Number: 55,
		Name: "getsockopt",
		EntryPoint: "sys_getsockopt",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "int",
					Name: "level",
				},
				{
					Ctype: "int",
					Name: "optname",
				},
				{
					Ctype: "char *",
					Name: "optval",
				},
				{
					Ctype: "int *",
					Name: "optlen",
				},
			},
		},
	},
	56: &SyscallData{
		Number: 56,
		Name: "clone",
		EntryPoint: "sys_clone",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
				{
					Ctype: "void *",
					Name: "stack",
				},
				{
					Ctype: "int *",
					Name: "parent_tid",
				},
				{
					Ctype: "int *",
					Name: "child_tid",
				},
				{
					Ctype: "unsigned long",
					Name: "tls",
				},
			},
		},
	},
	57: &SyscallData{
		Number: 57,
		Name: "fork",
		EntryPoint: "sys_fork",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	58: &SyscallData{
		Number: 58,
		Name: "vfork",
		EntryPoint: "sys_vfork",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	59: &SyscallData{
		Number: 59,
		Name: "execve",
		EntryPoint: "sys_execve",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "const char *const *",
					Name: "argv",
				},
				{
					Ctype: "const char *const *",
					Name: "envp",
				},
			},
		},
	},
	60: &SyscallData{
		Number: 60,
		Name: "exit",
		EntryPoint: "sys_exit",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "error_code",
				},
			},
		},
	},
	61: &SyscallData{
		Number: 61,
		Name: "wait4",
		EntryPoint: "sys_wait4",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int *",
					Name: "stat_addr",
				},
				{
					Ctype: "int",
					Name: "options",
				},
				{
					Ctype: "struct rusage *",
					Name: "ru",
				},
			},
		},
	},
	62: &SyscallData{
		Number: 62,
		Name: "kill",
		EntryPoint: "sys_kill",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
			},
		},
	},
	63: &SyscallData{
		Number: 63,
		Name: "uname",
		EntryPoint: "sys_newuname",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct new_utsname *",
					Name: "name",
				},
			},
		},
	},
	64: &SyscallData{
		Number: 64,
		Name: "semget",
		EntryPoint: "sys_semget",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "key_t",
					Name: "key",
				},
				{
					Ctype: "int",
					Name: "nsems",
				},
				{
					Ctype: "int",
					Name: "semflg",
				},
			},
		},
	},
	65: &SyscallData{
		Number: 65,
		Name: "semop",
		EntryPoint: "sys_semop",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "semid",
				},
				{
					Ctype: "struct sembuf *",
					Name: "sops",
				},
				{
					Ctype: "unsigned",
					Name: "nsops",
				},
			},
		},
	},
	66: &SyscallData{
		Number: 66,
		Name: "semctl",
		EntryPoint: "sys_semctl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "semid",
				},
				{
					Ctype: "int",
					Name: "semnum",
				},
				{
					Ctype: "int",
					Name: "cmd",
				},
				{
					Ctype: "unsigned long",
					Name: "arg",
				},
			},
		},
	},
	67: &SyscallData{
		Number: 67,
		Name: "shmdt",
		EntryPoint: "sys_shmdt",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "shmaddr",
				},
			},
		},
	},
	68: &SyscallData{
		Number: 68,
		Name: "msgget",
		EntryPoint: "sys_msgget",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "key_t",
					Name: "key",
				},
				{
					Ctype: "int",
					Name: "msgflg",
				},
			},
		},
	},
	69: &SyscallData{
		Number: 69,
		Name: "msgsnd",
		EntryPoint: "sys_msgsnd",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "msqid",
				},
				{
					Ctype: "struct msgbuf *",
					Name: "msgp",
				},
				{
					Ctype: "size_t",
					Name: "msgsz",
				},
				{
					Ctype: "int",
					Name: "msgflg",
				},
			},
		},
	},
	70: &SyscallData{
		Number: 70,
		Name: "msgrcv",
		EntryPoint: "sys_msgrcv",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "msqid",
				},
				{
					Ctype: "struct msgbuf *",
					Name: "msgp",
				},
				{
					Ctype: "size_t",
					Name: "msgsz",
				},
				{
					Ctype: "long",
					Name: "msgtyp",
				},
				{
					Ctype: "int",
					Name: "msgflg",
				},
			},
		},
	},
	71: &SyscallData{
		Number: 71,
		Name: "msgctl",
		EntryPoint: "sys_msgctl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "msqid",
				},
				{
					Ctype: "int",
					Name: "cmd",
				},
				{
					Ctype: "struct msqid_ds *",
					Name: "buf",
				},
			},
		},
	},
	72: &SyscallData{
		Number: 72,
		Name: "fcntl",
		EntryPoint: "sys_fcntl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
				{
					Ctype: "unsigned long",
					Name: "arg",
				},
			},
		},
	},
	73: &SyscallData{
		Number: 73,
		Name: "flock",
		EntryPoint: "sys_flock",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
			},
		},
	},
	74: &SyscallData{
		Number: 74,
		Name: "fsync",
		EntryPoint: "sys_fsync",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
			},
		},
	},
	75: &SyscallData{
		Number: 75,
		Name: "fdatasync",
		EntryPoint: "sys_fdatasync",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
			},
		},
	},
	76: &SyscallData{
		Number: 76,
		Name: "truncate",
		EntryPoint: "sys_truncate",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "long",
					Name: "length",
				},
			},
		},
	},
	77: &SyscallData{
		Number: 77,
		Name: "ftruncate",
		EntryPoint: "sys_ftruncate",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned long",
					Name: "length",
				},
			},
		},
	},
	78: &SyscallData{
		Number: 78,
		Name: "getdents",
		EntryPoint: "sys_getdents",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "struct linux_dirent *",
					Name: "dirent",
				},
				{
					Ctype: "unsigned int",
					Name: "count",
				},
			},
		},
	},
	79: &SyscallData{
		Number: 79,
		Name: "getcwd",
		EntryPoint: "sys_getcwd",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "unsigned long",
					Name: "size",
				},
			},
		},
	},
	80: &SyscallData{
		Number: 80,
		Name: "chdir",
		EntryPoint: "sys_chdir",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
			},
		},
	},
	81: &SyscallData{
		Number: 81,
		Name: "fchdir",
		EntryPoint: "sys_fchdir",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
			},
		},
	},
	82: &SyscallData{
		Number: 82,
		Name: "rename",
		EntryPoint: "sys_rename",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "oldname",
				},
				{
					Ctype: "const char *",
					Name: "newname",
				},
			},
		},
	},
	83: &SyscallData{
		Number: 83,
		Name: "mkdir",
		EntryPoint: "sys_mkdir",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "pathname",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	84: &SyscallData{
		Number: 84,
		Name: "rmdir",
		EntryPoint: "sys_rmdir",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "pathname",
				},
			},
		},
	},
	85: &SyscallData{
		Number: 85,
		Name: "creat",
		EntryPoint: "sys_creat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "pathname",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	86: &SyscallData{
		Number: 86,
		Name: "link",
		EntryPoint: "sys_link",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "oldname",
				},
				{
					Ctype: "const char *",
					Name: "newname",
				},
			},
		},
	},
	87: &SyscallData{
		Number: 87,
		Name: "unlink",
		EntryPoint: "sys_unlink",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "pathname",
				},
			},
		},
	},
	88: &SyscallData{
		Number: 88,
		Name: "symlink",
		EntryPoint: "sys_symlink",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "old",
				},
				{
					Ctype: "const char *",
					Name: "new",
				},
			},
		},
	},
	89: &SyscallData{
		Number: 89,
		Name: "readlink",
		EntryPoint: "sys_readlink",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "int",
					Name: "bufsiz",
				},
			},
		},
	},
	90: &SyscallData{
		Number: 90,
		Name: "chmod",
		EntryPoint: "sys_chmod",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	91: &SyscallData{
		Number: 91,
		Name: "fchmod",
		EntryPoint: "sys_fchmod",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	92: &SyscallData{
		Number: 92,
		Name: "chown",
		EntryPoint: "sys_chown",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "uid_t",
					Name: "user",
				},
				{
					Ctype: "gid_t",
					Name: "group",
				},
			},
		},
	},
	93: &SyscallData{
		Number: 93,
		Name: "fchown",
		EntryPoint: "sys_fchown",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "uid_t",
					Name: "user",
				},
				{
					Ctype: "gid_t",
					Name: "group",
				},
			},
		},
	},
	94: &SyscallData{
		Number: 94,
		Name: "lchown",
		EntryPoint: "sys_lchown",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "uid_t",
					Name: "user",
				},
				{
					Ctype: "gid_t",
					Name: "group",
				},
			},
		},
	},
	95: &SyscallData{
		Number: 95,
		Name: "umask",
		EntryPoint: "sys_umask",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "mask",
				},
			},
		},
	},
	96: &SyscallData{
		Number: 96,
		Name: "gettimeofday",
		EntryPoint: "sys_gettimeofday",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct __kernel_old_timeval *",
					Name: "tv",
				},
				{
					Ctype: "struct timezone *",
					Name: "tz",
				},
			},
		},
	},
	97: &SyscallData{
		Number: 97,
		Name: "getrlimit",
		EntryPoint: "sys_getrlimit",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "resource",
				},
				{
					Ctype: "struct rlimit *",
					Name: "rlim",
				},
			},
		},
	},
	98: &SyscallData{
		Number: 98,
		Name: "getrusage",
		EntryPoint: "sys_getrusage",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "who",
				},
				{
					Ctype: "struct rusage *",
					Name: "ru",
				},
			},
		},
	},
	99: &SyscallData{
		Number: 99,
		Name: "sysinfo",
		EntryPoint: "sys_sysinfo",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct sysinfo *",
					Name: "info",
				},
			},
		},
	},
	100: &SyscallData{
		Number: 100,
		Name: "times",
		EntryPoint: "sys_times",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct tms *",
					Name: "tbuf",
				},
			},
		},
	},
	101: &SyscallData{
		Number: 101,
		Name: "ptrace",
		EntryPoint: "sys_ptrace",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "long",
					Name: "request",
				},
				{
					Ctype: "long",
					Name: "pid",
				},
				{
					Ctype: "unsigned long",
					Name: "addr",
				},
				{
					Ctype: "unsigned long",
					Name: "data",
				},
			},
		},
	},
	102: &SyscallData{
		Number: 102,
		Name: "getuid",
		EntryPoint: "sys_getuid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	103: &SyscallData{
		Number: 103,
		Name: "syslog",
		EntryPoint: "sys_syslog",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "type",
				},
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "int",
					Name: "len",
				},
			},
		},
	},
	104: &SyscallData{
		Number: 104,
		Name: "getgid",
		EntryPoint: "sys_getgid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	105: &SyscallData{
		Number: 105,
		Name: "setuid",
		EntryPoint: "sys_setuid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "uid_t",
					Name: "uid",
				},
			},
		},
	},
	106: &SyscallData{
		Number: 106,
		Name: "setgid",
		EntryPoint: "sys_setgid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "gid_t",
					Name: "gid",
				},
			},
		},
	},
	107: &SyscallData{
		Number: 107,
		Name: "geteuid",
		EntryPoint: "sys_geteuid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	108: &SyscallData{
		Number: 108,
		Name: "getegid",
		EntryPoint: "sys_getegid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	109: &SyscallData{
		Number: 109,
		Name: "setpgid",
		EntryPoint: "sys_setpgid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "pid_t",
					Name: "pgid",
				},
			},
		},
	},
	110: &SyscallData{
		Number: 110,
		Name: "getppid",
		EntryPoint: "sys_getppid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	111: &SyscallData{
		Number: 111,
		Name: "getpgrp",
		EntryPoint: "sys_getpgrp",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	112: &SyscallData{
		Number: 112,
		Name: "setsid",
		EntryPoint: "sys_setsid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	113: &SyscallData{
		Number: 113,
		Name: "setreuid",
		EntryPoint: "sys_setreuid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "uid_t",
					Name: "ruid",
				},
				{
					Ctype: "uid_t",
					Name: "euid",
				},
			},
		},
	},
	114: &SyscallData{
		Number: 114,
		Name: "setregid",
		EntryPoint: "sys_setregid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "gid_t",
					Name: "rgid",
				},
				{
					Ctype: "gid_t",
					Name: "egid",
				},
			},
		},
	},
	115: &SyscallData{
		Number: 115,
		Name: "getgroups",
		EntryPoint: "sys_getgroups",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "gidsetsize",
				},
				{
					Ctype: "gid_t *",
					Name: "grouplist",
				},
			},
		},
	},
	116: &SyscallData{
		Number: 116,
		Name: "setgroups",
		EntryPoint: "sys_setgroups",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "gidsetsize",
				},
				{
					Ctype: "gid_t *",
					Name: "grouplist",
				},
			},
		},
	},
	117: &SyscallData{
		Number: 117,
		Name: "setresuid",
		EntryPoint: "sys_setresuid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "uid_t",
					Name: "ruid",
				},
				{
					Ctype: "uid_t",
					Name: "euid",
				},
				{
					Ctype: "uid_t",
					Name: "suid",
				},
			},
		},
	},
	118: &SyscallData{
		Number: 118,
		Name: "getresuid",
		EntryPoint: "sys_getresuid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "uid_t *",
					Name: "ruid",
				},
				{
					Ctype: "uid_t *",
					Name: "euid",
				},
				{
					Ctype: "uid_t *",
					Name: "suid",
				},
			},
		},
	},
	119: &SyscallData{
		Number: 119,
		Name: "setresgid",
		EntryPoint: "sys_setresgid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "gid_t",
					Name: "rgid",
				},
				{
					Ctype: "gid_t",
					Name: "egid",
				},
				{
					Ctype: "gid_t",
					Name: "sgid",
				},
			},
		},
	},
	120: &SyscallData{
		Number: 120,
		Name: "getresgid",
		EntryPoint: "sys_getresgid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "gid_t *",
					Name: "rgid",
				},
				{
					Ctype: "gid_t *",
					Name: "egid",
				},
				{
					Ctype: "gid_t *",
					Name: "sgid",
				},
			},
		},
	},
	121: &SyscallData{
		Number: 121,
		Name: "getpgid",
		EntryPoint: "sys_getpgid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
			},
		},
	},
	122: &SyscallData{
		Number: 122,
		Name: "setfsuid",
		EntryPoint: "sys_setfsuid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "uid_t",
					Name: "uid",
				},
			},
		},
	},
	123: &SyscallData{
		Number: 123,
		Name: "setfsgid",
		EntryPoint: "sys_setfsgid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "gid_t",
					Name: "gid",
				},
			},
		},
	},
	124: &SyscallData{
		Number: 124,
		Name: "getsid",
		EntryPoint: "sys_getsid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
			},
		},
	},
	125: &SyscallData{
		Number: 125,
		Name: "capget",
		EntryPoint: "sys_capget",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "cap_user_header_t",
					Name: "header",
				},
				{
					Ctype: "cap_user_data_t",
					Name: "dataptr",
				},
			},
		},
	},
	126: &SyscallData{
		Number: 126,
		Name: "capset",
		EntryPoint: "sys_capset",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "cap_user_header_t",
					Name: "header",
				},
				{
					Ctype: "const cap_user_data_t",
					Name: "data",
				},
			},
		},
	},
	127: &SyscallData{
		Number: 127,
		Name: "rt_sigpending",
		EntryPoint: "sys_rt_sigpending",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "sigset_t *",
					Name: "set",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	128: &SyscallData{
		Number: 128,
		Name: "rt_sigtimedwait",
		EntryPoint: "sys_rt_sigtimedwait",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const sigset_t *",
					Name: "uthese",
				},
				{
					Ctype: "siginfo_t *",
					Name: "uinfo",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "uts",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	129: &SyscallData{
		Number: 129,
		Name: "rt_sigqueueinfo",
		EntryPoint: "sys_rt_sigqueueinfo",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
				{
					Ctype: "siginfo_t *",
					Name: "uinfo",
				},
			},
		},
	},
	130: &SyscallData{
		Number: 130,
		Name: "rt_sigsuspend",
		EntryPoint: "sys_rt_sigsuspend",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "sigset_t *",
					Name: "unewset",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	131: &SyscallData{
		Number: 131,
		Name: "sigaltstack",
		EntryPoint: "sys_sigaltstack",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const struct sigaltstack *",
					Name: "uss",
				},
				{
					Ctype: "struct sigaltstack *",
					Name: "uoss",
				},
			},
		},
	},
	132: &SyscallData{
		Number: 132,
		Name: "utime",
		EntryPoint: "sys_utime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "filename",
				},
				{
					Ctype: "struct utimbuf *",
					Name: "times",
				},
			},
		},
	},
	133: &SyscallData{
		Number: 133,
		Name: "mknod",
		EntryPoint: "sys_mknod",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
				{
					Ctype: "unsigned",
					Name: "dev",
				},
			},
		},
	},
	135: &SyscallData{
		Number: 135,
		Name: "personality",
		EntryPoint: "sys_personality",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "personality",
				},
			},
		},
	},
	136: &SyscallData{
		Number: 136,
		Name: "ustat",
		EntryPoint: "sys_ustat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned",
					Name: "dev",
				},
				{
					Ctype: "struct ustat *",
					Name: "ubuf",
				},
			},
		},
	},
	137: &SyscallData{
		Number: 137,
		Name: "statfs",
		EntryPoint: "sys_statfs",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "struct statfs *",
					Name: "buf",
				},
			},
		},
	},
	138: &SyscallData{
		Number: 138,
		Name: "fstatfs",
		EntryPoint: "sys_fstatfs",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "struct statfs *",
					Name: "buf",
				},
			},
		},
	},
	139: &SyscallData{
		Number: 139,
		Name: "sysfs",
		EntryPoint: "sys_sysfs",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "option",
				},
				{
					Ctype: "unsigned long",
					Name: "arg1",
				},
				{
					Ctype: "unsigned long",
					Name: "arg2",
				},
			},
		},
	},
	140: &SyscallData{
		Number: 140,
		Name: "getpriority",
		EntryPoint: "sys_getpriority",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "int",
					Name: "who",
				},
			},
		},
	},
	141: &SyscallData{
		Number: 141,
		Name: "setpriority",
		EntryPoint: "sys_setpriority",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "int",
					Name: "who",
				},
				{
					Ctype: "int",
					Name: "niceval",
				},
			},
		},
	},
	142: &SyscallData{
		Number: 142,
		Name: "sched_setparam",
		EntryPoint: "sys_sched_setparam",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "struct sched_param *",
					Name: "param",
				},
			},
		},
	},
	143: &SyscallData{
		Number: 143,
		Name: "sched_getparam",
		EntryPoint: "sys_sched_getparam",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "struct sched_param *",
					Name: "param",
				},
			},
		},
	},
	144: &SyscallData{
		Number: 144,
		Name: "sched_setscheduler",
		EntryPoint: "sys_sched_setscheduler",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "policy",
				},
				{
					Ctype: "struct sched_param *",
					Name: "param",
				},
			},
		},
	},
	145: &SyscallData{
		Number: 145,
		Name: "sched_getscheduler",
		EntryPoint: "sys_sched_getscheduler",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
			},
		},
	},
	146: &SyscallData{
		Number: 146,
		Name: "sched_get_priority_max",
		EntryPoint: "sys_sched_get_priority_max",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "policy",
				},
			},
		},
	},
	147: &SyscallData{
		Number: 147,
		Name: "sched_get_priority_min",
		EntryPoint: "sys_sched_get_priority_min",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "policy",
				},
			},
		},
	},
	148: &SyscallData{
		Number: 148,
		Name: "sched_rr_get_interval",
		EntryPoint: "sys_sched_rr_get_interval",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "interval",
				},
			},
		},
	},
	149: &SyscallData{
		Number: 149,
		Name: "mlock",
		EntryPoint: "sys_mlock",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
			},
		},
	},
	150: &SyscallData{
		Number: 150,
		Name: "munlock",
		EntryPoint: "sys_munlock",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
			},
		},
	},
	151: &SyscallData{
		Number: 151,
		Name: "mlockall",
		EntryPoint: "sys_mlockall",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	152: &SyscallData{
		Number: 152,
		Name: "munlockall",
		EntryPoint: "sys_munlockall",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	153: &SyscallData{
		Number: 153,
		Name: "vhangup",
		EntryPoint: "sys_vhangup",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	154: &SyscallData{
		Number: 154,
		Name: "modify_ldt",
		EntryPoint: "sys_modify_ldt",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "func",
				},
				{
					Ctype: "void *",
					Name: "ptr",
				},
				{
					Ctype: "unsigned long",
					Name: "bytecount",
				},
			},
		},
	},
	155: &SyscallData{
		Number: 155,
		Name: "pivot_root",
		EntryPoint: "sys_pivot_root",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "new_root",
				},
				{
					Ctype: "const char *",
					Name: "put_old",
				},
			},
		},
	},
	156: &SyscallData{
		Number: 156,
		Name: "_sysctl",
		EntryPoint: "sys_ni_syscall",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	157: &SyscallData{
		Number: 157,
		Name: "prctl",
		EntryPoint: "sys_prctl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "option",
				},
				{
					Ctype: "unsigned long",
					Name: "arg2",
				},
				{
					Ctype: "unsigned long",
					Name: "arg3",
				},
				{
					Ctype: "unsigned long",
					Name: "arg4",
				},
				{
					Ctype: "unsigned long",
					Name: "arg5",
				},
			},
		},
	},
	158: &SyscallData{
		Number: 158,
		Name: "arch_prctl",
		EntryPoint: "sys_arch_prctl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "code",
				},
				{
					Ctype: "unsigned long",
					Name: "addr",
				},
			},
		},
	},
	159: &SyscallData{
		Number: 159,
		Name: "adjtimex",
		EntryPoint: "sys_adjtimex",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct __kernel_timex *",
					Name: "txc_p",
				},
			},
		},
	},
	160: &SyscallData{
		Number: 160,
		Name: "setrlimit",
		EntryPoint: "sys_setrlimit",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "resource",
				},
				{
					Ctype: "struct rlimit *",
					Name: "rlim",
				},
			},
		},
	},
	161: &SyscallData{
		Number: 161,
		Name: "chroot",
		EntryPoint: "sys_chroot",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
			},
		},
	},
	162: &SyscallData{
		Number: 162,
		Name: "sync",
		EntryPoint: "sys_sync",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	163: &SyscallData{
		Number: 163,
		Name: "acct",
		EntryPoint: "sys_acct",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "name",
				},
			},
		},
	},
	164: &SyscallData{
		Number: 164,
		Name: "settimeofday",
		EntryPoint: "sys_settimeofday",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct __kernel_old_timeval *",
					Name: "tv",
				},
				{
					Ctype: "struct timezone *",
					Name: "tz",
				},
			},
		},
	},
	165: &SyscallData{
		Number: 165,
		Name: "mount",
		EntryPoint: "sys_mount",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "dev_name",
				},
				{
					Ctype: "char *",
					Name: "dir_name",
				},
				{
					Ctype: "char *",
					Name: "type",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
				{
					Ctype: "void *",
					Name: "data",
				},
			},
		},
	},
	166: &SyscallData{
		Number: 166,
		Name: "umount2",
		EntryPoint: "sys_umount",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "name",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	167: &SyscallData{
		Number: 167,
		Name: "swapon",
		EntryPoint: "sys_swapon",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "specialfile",
				},
				{
					Ctype: "int",
					Name: "swap_flags",
				},
			},
		},
	},
	168: &SyscallData{
		Number: 168,
		Name: "swapoff",
		EntryPoint: "sys_swapoff",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "specialfile",
				},
			},
		},
	},
	169: &SyscallData{
		Number: 169,
		Name: "reboot",
		EntryPoint: "sys_reboot",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "magic1",
				},
				{
					Ctype: "int",
					Name: "magic2",
				},
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
				{
					Ctype: "void *",
					Name: "arg",
				},
			},
		},
	},
	170: &SyscallData{
		Number: 170,
		Name: "sethostname",
		EntryPoint: "sys_sethostname",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "name",
				},
				{
					Ctype: "int",
					Name: "len",
				},
			},
		},
	},
	171: &SyscallData{
		Number: 171,
		Name: "setdomainname",
		EntryPoint: "sys_setdomainname",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "name",
				},
				{
					Ctype: "int",
					Name: "len",
				},
			},
		},
	},
	172: &SyscallData{
		Number: 172,
		Name: "iopl",
		EntryPoint: "sys_iopl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "level",
				},
			},
		},
	},
	173: &SyscallData{
		Number: 173,
		Name: "ioperm",
		EntryPoint: "sys_ioperm",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "from",
				},
				{
					Ctype: "unsigned long",
					Name: "num",
				},
				{
					Ctype: "int",
					Name: "on",
				},
			},
		},
	},
	175: &SyscallData{
		Number: 175,
		Name: "init_module",
		EntryPoint: "sys_init_module",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "void *",
					Name: "umod",
				},
				{
					Ctype: "unsigned long",
					Name: "len",
				},
				{
					Ctype: "const char *",
					Name: "uargs",
				},
			},
		},
	},
	176: &SyscallData{
		Number: 176,
		Name: "delete_module",
		EntryPoint: "sys_delete_module",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "name_user",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	179: &SyscallData{
		Number: 179,
		Name: "quotactl",
		EntryPoint: "sys_quotactl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
				{
					Ctype: "const char *",
					Name: "special",
				},
				{
					Ctype: "qid_t",
					Name: "id",
				},
				{
					Ctype: "void *",
					Name: "addr",
				},
			},
		},
	},
	186: &SyscallData{
		Number: 186,
		Name: "gettid",
		EntryPoint: "sys_gettid",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	187: &SyscallData{
		Number: 187,
		Name: "readahead",
		EntryPoint: "sys_readahead",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "loff_t",
					Name: "offset",
				},
				{
					Ctype: "size_t",
					Name: "count",
				},
			},
		},
	},
	188: &SyscallData{
		Number: 188,
		Name: "setxattr",
		EntryPoint: "sys_setxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "const void *",
					Name: "value",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	189: &SyscallData{
		Number: 189,
		Name: "lsetxattr",
		EntryPoint: "sys_lsetxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "const void *",
					Name: "value",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	190: &SyscallData{
		Number: 190,
		Name: "fsetxattr",
		EntryPoint: "sys_fsetxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "const void *",
					Name: "value",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	191: &SyscallData{
		Number: 191,
		Name: "getxattr",
		EntryPoint: "sys_getxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "void *",
					Name: "value",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	192: &SyscallData{
		Number: 192,
		Name: "lgetxattr",
		EntryPoint: "sys_lgetxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "void *",
					Name: "value",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	193: &SyscallData{
		Number: 193,
		Name: "fgetxattr",
		EntryPoint: "sys_fgetxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "void *",
					Name: "value",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	194: &SyscallData{
		Number: 194,
		Name: "listxattr",
		EntryPoint: "sys_listxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "char *",
					Name: "list",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	195: &SyscallData{
		Number: 195,
		Name: "llistxattr",
		EntryPoint: "sys_llistxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "char *",
					Name: "list",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	196: &SyscallData{
		Number: 196,
		Name: "flistxattr",
		EntryPoint: "sys_flistxattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "char *",
					Name: "list",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	197: &SyscallData{
		Number: 197,
		Name: "removexattr",
		EntryPoint: "sys_removexattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
			},
		},
	},
	198: &SyscallData{
		Number: 198,
		Name: "lremovexattr",
		EntryPoint: "sys_lremovexattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
			},
		},
	},
	199: &SyscallData{
		Number: 199,
		Name: "fremovexattr",
		EntryPoint: "sys_fremovexattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
			},
		},
	},
	200: &SyscallData{
		Number: 200,
		Name: "tkill",
		EntryPoint: "sys_tkill",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
			},
		},
	},
	201: &SyscallData{
		Number: 201,
		Name: "time",
		EntryPoint: "sys_time",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "__kernel_old_time_t *",
					Name: "tloc",
				},
			},
		},
	},
	202: &SyscallData{
		Number: 202,
		Name: "futex",
		EntryPoint: "sys_futex",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "u32 *",
					Name: "uaddr",
				},
				{
					Ctype: "int",
					Name: "op",
				},
				{
					Ctype: "u32",
					Name: "val",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "utime",
				},
				{
					Ctype: "u32 *",
					Name: "uaddr2",
				},
				{
					Ctype: "u32",
					Name: "val3",
				},
			},
		},
	},
	203: &SyscallData{
		Number: 203,
		Name: "sched_setaffinity",
		EntryPoint: "sys_sched_setaffinity",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "unsigned int",
					Name: "len",
				},
				{
					Ctype: "unsigned long *",
					Name: "user_mask_ptr",
				},
			},
		},
	},
	204: &SyscallData{
		Number: 204,
		Name: "sched_getaffinity",
		EntryPoint: "sys_sched_getaffinity",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "unsigned int",
					Name: "len",
				},
				{
					Ctype: "unsigned long *",
					Name: "user_mask_ptr",
				},
			},
		},
	},
	206: &SyscallData{
		Number: 206,
		Name: "io_setup",
		EntryPoint: "sys_io_setup",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned",
					Name: "nr_reqs",
				},
				{
					Ctype: "aio_context_t *",
					Name: "ctx",
				},
			},
		},
	},
	207: &SyscallData{
		Number: 207,
		Name: "io_destroy",
		EntryPoint: "sys_io_destroy",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "aio_context_t",
					Name: "ctx",
				},
			},
		},
	},
	208: &SyscallData{
		Number: 208,
		Name: "io_getevents",
		EntryPoint: "sys_io_getevents",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "aio_context_t",
					Name: "ctx_id",
				},
				{
					Ctype: "long",
					Name: "min_nr",
				},
				{
					Ctype: "long",
					Name: "nr",
				},
				{
					Ctype: "struct io_event *",
					Name: "events",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "timeout",
				},
			},
		},
	},
	209: &SyscallData{
		Number: 209,
		Name: "io_submit",
		EntryPoint: "sys_io_submit",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "aio_context_t",
					Name: "ctx_id",
				},
				{
					Ctype: "long",
					Name: "nr",
				},
				{
					Ctype: "struct iocb **",
					Name: "iocbpp",
				},
			},
		},
	},
	210: &SyscallData{
		Number: 210,
		Name: "io_cancel",
		EntryPoint: "sys_io_cancel",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "aio_context_t",
					Name: "ctx_id",
				},
				{
					Ctype: "struct iocb *",
					Name: "iocb",
				},
				{
					Ctype: "struct io_event *",
					Name: "result",
				},
			},
		},
	},
	212: &SyscallData{
		Number: 212,
		Name: "lookup_dcookie",
		EntryPoint: "sys_lookup_dcookie",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "u64",
					Name: "cookie64",
				},
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
			},
		},
	},
	213: &SyscallData{
		Number: 213,
		Name: "epoll_create",
		EntryPoint: "sys_epoll_create",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "size",
				},
			},
		},
	},
	216: &SyscallData{
		Number: 216,
		Name: "remap_file_pages",
		EntryPoint: "sys_remap_file_pages",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "unsigned long",
					Name: "size",
				},
				{
					Ctype: "unsigned long",
					Name: "prot",
				},
				{
					Ctype: "unsigned long",
					Name: "pgoff",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	217: &SyscallData{
		Number: 217,
		Name: "getdents64",
		EntryPoint: "sys_getdents64",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "struct linux_dirent64 *",
					Name: "dirent",
				},
				{
					Ctype: "unsigned int",
					Name: "count",
				},
			},
		},
	},
	218: &SyscallData{
		Number: 218,
		Name: "set_tid_address",
		EntryPoint: "sys_set_tid_address",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int *",
					Name: "tidptr",
				},
			},
		},
	},
	219: &SyscallData{
		Number: 219,
		Name: "restart_syscall",
		EntryPoint: "sys_restart_syscall",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	220: &SyscallData{
		Number: 220,
		Name: "semtimedop",
		EntryPoint: "sys_semtimedop",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "semid",
				},
				{
					Ctype: "struct sembuf *",
					Name: "sops",
				},
				{
					Ctype: "unsigned",
					Name: "nsops",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "timeout",
				},
			},
		},
	},
	221: &SyscallData{
		Number: 221,
		Name: "fadvise64",
		EntryPoint: "sys_fadvise64",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "loff_t",
					Name: "offset",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "int",
					Name: "advice",
				},
			},
		},
	},
	222: &SyscallData{
		Number: 222,
		Name: "timer_create",
		EntryPoint: "sys_timer_create",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "clockid_t",
					Name: "which_clock",
				},
				{
					Ctype: "struct sigevent *",
					Name: "timer_event_spec",
				},
				{
					Ctype: "timer_t *",
					Name: "created_timer_id",
				},
			},
		},
	},
	223: &SyscallData{
		Number: 223,
		Name: "timer_settime",
		EntryPoint: "sys_timer_settime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "timer_t",
					Name: "timer_id",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
				{
					Ctype: "const struct __kernel_itimerspec *",
					Name: "new_setting",
				},
				{
					Ctype: "struct __kernel_itimerspec *",
					Name: "old_setting",
				},
			},
		},
	},
	224: &SyscallData{
		Number: 224,
		Name: "timer_gettime",
		EntryPoint: "sys_timer_gettime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "timer_t",
					Name: "timer_id",
				},
				{
					Ctype: "struct __kernel_itimerspec *",
					Name: "setting",
				},
			},
		},
	},
	225: &SyscallData{
		Number: 225,
		Name: "timer_getoverrun",
		EntryPoint: "sys_timer_getoverrun",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "timer_t",
					Name: "timer_id",
				},
			},
		},
	},
	226: &SyscallData{
		Number: 226,
		Name: "timer_delete",
		EntryPoint: "sys_timer_delete",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "timer_t",
					Name: "timer_id",
				},
			},
		},
	},
	227: &SyscallData{
		Number: 227,
		Name: "clock_settime",
		EntryPoint: "sys_clock_settime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "clockid_t",
					Name: "which_clock",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "tp",
				},
			},
		},
	},
	228: &SyscallData{
		Number: 228,
		Name: "clock_gettime",
		EntryPoint: "sys_clock_gettime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "clockid_t",
					Name: "which_clock",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "tp",
				},
			},
		},
	},
	229: &SyscallData{
		Number: 229,
		Name: "clock_getres",
		EntryPoint: "sys_clock_getres",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "clockid_t",
					Name: "which_clock",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "tp",
				},
			},
		},
	},
	230: &SyscallData{
		Number: 230,
		Name: "clock_nanosleep",
		EntryPoint: "sys_clock_nanosleep",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "clockid_t",
					Name: "which_clock",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "rqtp",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "rmtp",
				},
			},
		},
	},
	231: &SyscallData{
		Number: 231,
		Name: "exit_group",
		EntryPoint: "sys_exit_group",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "error_code",
				},
			},
		},
	},
	232: &SyscallData{
		Number: 232,
		Name: "epoll_wait",
		EntryPoint: "sys_epoll_wait",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "epfd",
				},
				{
					Ctype: "struct epoll_event *",
					Name: "events",
				},
				{
					Ctype: "int",
					Name: "maxevents",
				},
				{
					Ctype: "int",
					Name: "timeout",
				},
			},
		},
	},
	233: &SyscallData{
		Number: 233,
		Name: "epoll_ctl",
		EntryPoint: "sys_epoll_ctl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "epfd",
				},
				{
					Ctype: "int",
					Name: "op",
				},
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct epoll_event *",
					Name: "event",
				},
			},
		},
	},
	234: &SyscallData{
		Number: 234,
		Name: "tgkill",
		EntryPoint: "sys_tgkill",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "tgid",
				},
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
			},
		},
	},
	235: &SyscallData{
		Number: 235,
		Name: "utimes",
		EntryPoint: "sys_utimes",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "filename",
				},
				{
					Ctype: "struct __kernel_old_timeval *",
					Name: "utimes",
				},
			},
		},
	},
	237: &SyscallData{
		Number: 237,
		Name: "mbind",
		EntryPoint: "sys_mbind",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "unsigned long",
					Name: "len",
				},
				{
					Ctype: "unsigned long",
					Name: "mode",
				},
				{
					Ctype: "const unsigned long *",
					Name: "nmask",
				},
				{
					Ctype: "unsigned long",
					Name: "maxnode",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
			},
		},
	},
	238: &SyscallData{
		Number: 238,
		Name: "set_mempolicy",
		EntryPoint: "sys_set_mempolicy",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "mode",
				},
				{
					Ctype: "const unsigned long *",
					Name: "nmask",
				},
				{
					Ctype: "unsigned long",
					Name: "maxnode",
				},
			},
		},
	},
	239: &SyscallData{
		Number: 239,
		Name: "get_mempolicy",
		EntryPoint: "sys_get_mempolicy",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int *",
					Name: "policy",
				},
				{
					Ctype: "unsigned long *",
					Name: "nmask",
				},
				{
					Ctype: "unsigned long",
					Name: "maxnode",
				},
				{
					Ctype: "unsigned long",
					Name: "addr",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	240: &SyscallData{
		Number: 240,
		Name: "mq_open",
		EntryPoint: "sys_mq_open",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "int",
					Name: "oflag",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
				{
					Ctype: "struct mq_attr *",
					Name: "attr",
				},
			},
		},
	},
	241: &SyscallData{
		Number: 241,
		Name: "mq_unlink",
		EntryPoint: "sys_mq_unlink",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "name",
				},
			},
		},
	},
	242: &SyscallData{
		Number: 242,
		Name: "mq_timedsend",
		EntryPoint: "sys_mq_timedsend",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "mqd_t",
					Name: "mqdes",
				},
				{
					Ctype: "const char *",
					Name: "msg_ptr",
				},
				{
					Ctype: "size_t",
					Name: "msg_len",
				},
				{
					Ctype: "unsigned int",
					Name: "msg_prio",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "abs_timeout",
				},
			},
		},
	},
	243: &SyscallData{
		Number: 243,
		Name: "mq_timedreceive",
		EntryPoint: "sys_mq_timedreceive",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "mqd_t",
					Name: "mqdes",
				},
				{
					Ctype: "char *",
					Name: "msg_ptr",
				},
				{
					Ctype: "size_t",
					Name: "msg_len",
				},
				{
					Ctype: "unsigned int *",
					Name: "msg_prio",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "abs_timeout",
				},
			},
		},
	},
	244: &SyscallData{
		Number: 244,
		Name: "mq_notify",
		EntryPoint: "sys_mq_notify",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "mqd_t",
					Name: "mqdes",
				},
				{
					Ctype: "const struct sigevent *",
					Name: "notification",
				},
			},
		},
	},
	245: &SyscallData{
		Number: 245,
		Name: "mq_getsetattr",
		EntryPoint: "sys_mq_getsetattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "mqd_t",
					Name: "mqdes",
				},
				{
					Ctype: "const struct mq_attr *",
					Name: "mqstat",
				},
				{
					Ctype: "struct mq_attr *",
					Name: "omqstat",
				},
			},
		},
	},
	246: &SyscallData{
		Number: 246,
		Name: "kexec_load",
		EntryPoint: "sys_kexec_load",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "entry",
				},
				{
					Ctype: "unsigned long",
					Name: "nr_segments",
				},
				{
					Ctype: "struct kexec_segment *",
					Name: "segments",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	247: &SyscallData{
		Number: 247,
		Name: "waitid",
		EntryPoint: "sys_waitid",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "struct siginfo *",
					Name: "infop",
				},
				{
					Ctype: "int",
					Name: "options",
				},
				{
					Ctype: "struct rusage *",
					Name: "ru",
				},
			},
		},
	},
	248: &SyscallData{
		Number: 248,
		Name: "add_key",
		EntryPoint: "sys_add_key",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "_type",
				},
				{
					Ctype: "const char *",
					Name: "_description",
				},
				{
					Ctype: "const void *",
					Name: "_payload",
				},
				{
					Ctype: "size_t",
					Name: "plen",
				},
				{
					Ctype: "key_serial_t",
					Name: "destringid",
				},
			},
		},
	},
	249: &SyscallData{
		Number: 249,
		Name: "request_key",
		EntryPoint: "sys_request_key",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "_type",
				},
				{
					Ctype: "const char *",
					Name: "_description",
				},
				{
					Ctype: "const char *",
					Name: "_callout_info",
				},
				{
					Ctype: "key_serial_t",
					Name: "destringid",
				},
			},
		},
	},
	250: &SyscallData{
		Number: 250,
		Name: "keyctl",
		EntryPoint: "sys_keyctl",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "cmd",
				},
				{
					Ctype: "unsigned long",
					Name: "arg2",
				},
				{
					Ctype: "unsigned long",
					Name: "arg3",
				},
				{
					Ctype: "unsigned long",
					Name: "arg4",
				},
				{
					Ctype: "unsigned long",
					Name: "arg5",
				},
			},
		},
	},
	251: &SyscallData{
		Number: 251,
		Name: "ioprio_set",
		EntryPoint: "sys_ioprio_set",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "int",
					Name: "who",
				},
				{
					Ctype: "int",
					Name: "ioprio",
				},
			},
		},
	},
	252: &SyscallData{
		Number: 252,
		Name: "ioprio_get",
		EntryPoint: "sys_ioprio_get",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "int",
					Name: "who",
				},
			},
		},
	},
	253: &SyscallData{
		Number: 253,
		Name: "inotify_init",
		EntryPoint: "sys_inotify_init",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgVoid,
			Args: []*ArgData{
			},
		},
	},
	254: &SyscallData{
		Number: 254,
		Name: "inotify_add_watch",
		EntryPoint: "sys_inotify_add_watch",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "u32",
					Name: "mask",
				},
			},
		},
	},
	255: &SyscallData{
		Number: 255,
		Name: "inotify_rm_watch",
		EntryPoint: "sys_inotify_rm_watch",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "__s32",
					Name: "wd",
				},
			},
		},
	},
	256: &SyscallData{
		Number: 256,
		Name: "migrate_pages",
		EntryPoint: "sys_migrate_pages",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "unsigned long",
					Name: "maxnode",
				},
				{
					Ctype: "const unsigned long *",
					Name: "from",
				},
				{
					Ctype: "const unsigned long *",
					Name: "to",
				},
			},
		},
	},
	257: &SyscallData{
		Number: 257,
		Name: "openat",
		EntryPoint: "sys_openat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	258: &SyscallData{
		Number: 258,
		Name: "mkdirat",
		EntryPoint: "sys_mkdirat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "pathname",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	259: &SyscallData{
		Number: 259,
		Name: "mknodat",
		EntryPoint: "sys_mknodat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
				{
					Ctype: "unsigned",
					Name: "dev",
				},
			},
		},
	},
	260: &SyscallData{
		Number: 260,
		Name: "fchownat",
		EntryPoint: "sys_fchownat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "uid_t",
					Name: "user",
				},
				{
					Ctype: "gid_t",
					Name: "group",
				},
				{
					Ctype: "int",
					Name: "flag",
				},
			},
		},
	},
	261: &SyscallData{
		Number: 261,
		Name: "futimesat",
		EntryPoint: "sys_futimesat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "struct __kernel_old_timeval *",
					Name: "utimes",
				},
			},
		},
	},
	262: &SyscallData{
		Number: 262,
		Name: "newfstatat",
		EntryPoint: "sys_newfstatat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "struct stat *",
					Name: "statbuf",
				},
				{
					Ctype: "int",
					Name: "flag",
				},
			},
		},
	},
	263: &SyscallData{
		Number: 263,
		Name: "unlinkat",
		EntryPoint: "sys_unlinkat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "pathname",
				},
				{
					Ctype: "int",
					Name: "flag",
				},
			},
		},
	},
	264: &SyscallData{
		Number: 264,
		Name: "renameat",
		EntryPoint: "sys_renameat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "olddfd",
				},
				{
					Ctype: "const char *",
					Name: "oldname",
				},
				{
					Ctype: "int",
					Name: "newdfd",
				},
				{
					Ctype: "const char *",
					Name: "newname",
				},
			},
		},
	},
	265: &SyscallData{
		Number: 265,
		Name: "linkat",
		EntryPoint: "sys_linkat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "olddfd",
				},
				{
					Ctype: "const char *",
					Name: "oldname",
				},
				{
					Ctype: "int",
					Name: "newdfd",
				},
				{
					Ctype: "const char *",
					Name: "newname",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	266: &SyscallData{
		Number: 266,
		Name: "symlinkat",
		EntryPoint: "sys_symlinkat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "oldname",
				},
				{
					Ctype: "int",
					Name: "newdfd",
				},
				{
					Ctype: "const char *",
					Name: "newname",
				},
			},
		},
	},
	267: &SyscallData{
		Number: 267,
		Name: "readlinkat",
		EntryPoint: "sys_readlinkat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "int",
					Name: "bufsiz",
				},
			},
		},
	},
	268: &SyscallData{
		Number: 268,
		Name: "fchmodat",
		EntryPoint: "sys_fchmodat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "umode_t",
					Name: "mode",
				},
			},
		},
	},
	269: &SyscallData{
		Number: 269,
		Name: "faccessat",
		EntryPoint: "sys_faccessat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "int",
					Name: "mode",
				},
			},
		},
	},
	270: &SyscallData{
		Number: 270,
		Name: "pselect6",
		EntryPoint: "sys_pselect6",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "n",
				},
				{
					Ctype: "fd_set *",
					Name: "inp",
				},
				{
					Ctype: "fd_set *",
					Name: "outp",
				},
				{
					Ctype: "fd_set *",
					Name: "exp",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "tsp",
				},
				{
					Ctype: "void *",
					Name: "sig",
				},
			},
		},
	},
	271: &SyscallData{
		Number: 271,
		Name: "ppoll",
		EntryPoint: "sys_ppoll",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct pollfd *",
					Name: "ufds",
				},
				{
					Ctype: "unsigned int",
					Name: "nfds",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "tsp",
				},
				{
					Ctype: "const sigset_t *",
					Name: "sigmask",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	272: &SyscallData{
		Number: 272,
		Name: "unshare",
		EntryPoint: "sys_unshare",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "unshare_flags",
				},
			},
		},
	},
	273: &SyscallData{
		Number: 273,
		Name: "set_robust_list",
		EntryPoint: "sys_set_robust_list",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct robust_list_head *",
					Name: "head",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
			},
		},
	},
	274: &SyscallData{
		Number: 274,
		Name: "get_robust_list",
		EntryPoint: "sys_get_robust_list",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "pid",
				},
				{
					Ctype: "struct robust_list_head **",
					Name: "head_ptr",
				},
				{
					Ctype: "size_t *",
					Name: "len_ptr",
				},
			},
		},
	},
	275: &SyscallData{
		Number: 275,
		Name: "splice",
		EntryPoint: "sys_splice",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd_in",
				},
				{
					Ctype: "loff_t *",
					Name: "off_in",
				},
				{
					Ctype: "int",
					Name: "fd_out",
				},
				{
					Ctype: "loff_t *",
					Name: "off_out",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	276: &SyscallData{
		Number: 276,
		Name: "tee",
		EntryPoint: "sys_tee",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fdin",
				},
				{
					Ctype: "int",
					Name: "fdout",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	277: &SyscallData{
		Number: 277,
		Name: "sync_file_range",
		EntryPoint: "sys_sync_file_range",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "loff_t",
					Name: "offset",
				},
				{
					Ctype: "loff_t",
					Name: "nbytes",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	278: &SyscallData{
		Number: 278,
		Name: "vmsplice",
		EntryPoint: "sys_vmsplice",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "iov",
				},
				{
					Ctype: "unsigned long",
					Name: "nr_segs",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	279: &SyscallData{
		Number: 279,
		Name: "move_pages",
		EntryPoint: "sys_move_pages",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "unsigned long",
					Name: "nr_pages",
				},
				{
					Ctype: "const void * *",
					Name: "pages",
				},
				{
					Ctype: "const int *",
					Name: "nodes",
				},
				{
					Ctype: "int *",
					Name: "status",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	280: &SyscallData{
		Number: 280,
		Name: "utimensat",
		EntryPoint: "sys_utimensat",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "utimes",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	281: &SyscallData{
		Number: 281,
		Name: "epoll_pwait",
		EntryPoint: "sys_epoll_pwait",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "epfd",
				},
				{
					Ctype: "struct epoll_event *",
					Name: "events",
				},
				{
					Ctype: "int",
					Name: "maxevents",
				},
				{
					Ctype: "int",
					Name: "timeout",
				},
				{
					Ctype: "const sigset_t *",
					Name: "sigmask",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	282: &SyscallData{
		Number: 282,
		Name: "signalfd",
		EntryPoint: "sys_signalfd",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "ufd",
				},
				{
					Ctype: "sigset_t *",
					Name: "user_mask",
				},
				{
					Ctype: "size_t",
					Name: "sizemask",
				},
			},
		},
	},
	283: &SyscallData{
		Number: 283,
		Name: "timerfd_create",
		EntryPoint: "sys_timerfd_create",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "clockid",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	284: &SyscallData{
		Number: 284,
		Name: "eventfd",
		EntryPoint: "sys_eventfd",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "count",
				},
			},
		},
	},
	285: &SyscallData{
		Number: 285,
		Name: "fallocate",
		EntryPoint: "sys_fallocate",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "int",
					Name: "mode",
				},
				{
					Ctype: "loff_t",
					Name: "offset",
				},
				{
					Ctype: "loff_t",
					Name: "len",
				},
			},
		},
	},
	286: &SyscallData{
		Number: 286,
		Name: "timerfd_settime",
		EntryPoint: "sys_timerfd_settime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "ufd",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
				{
					Ctype: "const struct __kernel_itimerspec *",
					Name: "utmr",
				},
				{
					Ctype: "struct __kernel_itimerspec *",
					Name: "otmr",
				},
			},
		},
	},
	287: &SyscallData{
		Number: 287,
		Name: "timerfd_gettime",
		EntryPoint: "sys_timerfd_gettime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "ufd",
				},
				{
					Ctype: "struct __kernel_itimerspec *",
					Name: "otmr",
				},
			},
		},
	},
	288: &SyscallData{
		Number: 288,
		Name: "accept4",
		EntryPoint: "sys_accept4",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int *",
					Name: "addrlen",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	289: &SyscallData{
		Number: 289,
		Name: "signalfd4",
		EntryPoint: "sys_signalfd4",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "ufd",
				},
				{
					Ctype: "sigset_t *",
					Name: "user_mask",
				},
				{
					Ctype: "size_t",
					Name: "sizemask",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	290: &SyscallData{
		Number: 290,
		Name: "eventfd2",
		EntryPoint: "sys_eventfd2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "count",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	291: &SyscallData{
		Number: 291,
		Name: "epoll_create1",
		EntryPoint: "sys_epoll_create1",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	292: &SyscallData{
		Number: 292,
		Name: "dup3",
		EntryPoint: "sys_dup3",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "oldfd",
				},
				{
					Ctype: "unsigned int",
					Name: "newfd",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	293: &SyscallData{
		Number: 293,
		Name: "pipe2",
		EntryPoint: "sys_pipe2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int *",
					Name: "fildes",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	294: &SyscallData{
		Number: 294,
		Name: "inotify_init1",
		EntryPoint: "sys_inotify_init1",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	295: &SyscallData{
		Number: 295,
		Name: "preadv",
		EntryPoint: "sys_preadv",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_l",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_h",
				},
			},
		},
	},
	296: &SyscallData{
		Number: 296,
		Name: "pwritev",
		EntryPoint: "sys_pwritev",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_l",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_h",
				},
			},
		},
	},
	297: &SyscallData{
		Number: 297,
		Name: "rt_tgsigqueueinfo",
		EntryPoint: "sys_rt_tgsigqueueinfo",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "tgid",
				},
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
				{
					Ctype: "siginfo_t *",
					Name: "uinfo",
				},
			},
		},
	},
	298: &SyscallData{
		Number: 298,
		Name: "perf_event_open",
		EntryPoint: "sys_perf_event_open",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct perf_event_attr *",
					Name: "attr_uptr",
				},
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "cpu",
				},
				{
					Ctype: "int",
					Name: "group_fd",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	299: &SyscallData{
		Number: 299,
		Name: "recvmmsg",
		EntryPoint: "sys_recvmmsg",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct mmsghdr *",
					Name: "msg",
				},
				{
					Ctype: "unsigned int",
					Name: "vlen",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "timeout",
				},
			},
		},
	},
	300: &SyscallData{
		Number: 300,
		Name: "fanotify_init",
		EntryPoint: "sys_fanotify_init",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "unsigned int",
					Name: "event_f_flags",
				},
			},
		},
	},
	301: &SyscallData{
		Number: 301,
		Name: "fanotify_mark",
		EntryPoint: "sys_fanotify_mark",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fanotify_fd",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "u64",
					Name: "mask",
				},
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "pathname",
				},
			},
		},
	},
	302: &SyscallData{
		Number: 302,
		Name: "prlimit64",
		EntryPoint: "sys_prlimit64",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "unsigned int",
					Name: "resource",
				},
				{
					Ctype: "const struct rlimit64 *",
					Name: "new_rlim",
				},
				{
					Ctype: "struct rlimit64 *",
					Name: "old_rlim",
				},
			},
		},
	},
	303: &SyscallData{
		Number: 303,
		Name: "name_to_handle_at",
		EntryPoint: "sys_name_to_handle_at",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "name",
				},
				{
					Ctype: "struct file_handle *",
					Name: "handle",
				},
				{
					Ctype: "int *",
					Name: "mnt_id",
				},
				{
					Ctype: "int",
					Name: "flag",
				},
			},
		},
	},
	304: &SyscallData{
		Number: 304,
		Name: "open_by_handle_at",
		EntryPoint: "sys_open_by_handle_at",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "mountdirfd",
				},
				{
					Ctype: "struct file_handle *",
					Name: "handle",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	305: &SyscallData{
		Number: 305,
		Name: "clock_adjtime",
		EntryPoint: "sys_clock_adjtime",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "clockid_t",
					Name: "which_clock",
				},
				{
					Ctype: "struct __kernel_timex *",
					Name: "tx",
				},
			},
		},
	},
	306: &SyscallData{
		Number: 306,
		Name: "syncfs",
		EntryPoint: "sys_syncfs",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
			},
		},
	},
	307: &SyscallData{
		Number: 307,
		Name: "sendmmsg",
		EntryPoint: "sys_sendmmsg",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct mmsghdr *",
					Name: "msg",
				},
				{
					Ctype: "unsigned int",
					Name: "vlen",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
			},
		},
	},
	308: &SyscallData{
		Number: 308,
		Name: "setns",
		EntryPoint: "sys_setns",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "int",
					Name: "nstype",
				},
			},
		},
	},
	309: &SyscallData{
		Number: 309,
		Name: "getcpu",
		EntryPoint: "sys_getcpu",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned *",
					Name: "cpu",
				},
				{
					Ctype: "unsigned *",
					Name: "node",
				},
				{
					Ctype: "struct getcpu_cache *",
					Name: "cache",
				},
			},
		},
	},
	310: &SyscallData{
		Number: 310,
		Name: "process_vm_readv",
		EntryPoint: "sys_process_vm_readv",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "const struct iovec *",
					Name: "lvec",
				},
				{
					Ctype: "unsigned long",
					Name: "liovcnt",
				},
				{
					Ctype: "const struct iovec *",
					Name: "rvec",
				},
				{
					Ctype: "unsigned long",
					Name: "riovcnt",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	311: &SyscallData{
		Number: 311,
		Name: "process_vm_writev",
		EntryPoint: "sys_process_vm_writev",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "const struct iovec *",
					Name: "lvec",
				},
				{
					Ctype: "unsigned long",
					Name: "liovcnt",
				},
				{
					Ctype: "const struct iovec *",
					Name: "rvec",
				},
				{
					Ctype: "unsigned long",
					Name: "riovcnt",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	312: &SyscallData{
		Number: 312,
		Name: "kcmp",
		EntryPoint: "sys_kcmp",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid1",
				},
				{
					Ctype: "pid_t",
					Name: "pid2",
				},
				{
					Ctype: "int",
					Name: "type",
				},
				{
					Ctype: "unsigned long",
					Name: "idx1",
				},
				{
					Ctype: "unsigned long",
					Name: "idx2",
				},
			},
		},
	},
	313: &SyscallData{
		Number: 313,
		Name: "finit_module",
		EntryPoint: "sys_finit_module",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const char *",
					Name: "uargs",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	314: &SyscallData{
		Number: 314,
		Name: "sched_setattr",
		EntryPoint: "sys_sched_setattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "struct sched_attr *",
					Name: "attr",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	315: &SyscallData{
		Number: 315,
		Name: "sched_getattr",
		EntryPoint: "sys_sched_getattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "struct sched_attr *",
					Name: "attr",
				},
				{
					Ctype: "unsigned int",
					Name: "size",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	316: &SyscallData{
		Number: 316,
		Name: "renameat2",
		EntryPoint: "sys_renameat2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "olddfd",
				},
				{
					Ctype: "const char *",
					Name: "oldname",
				},
				{
					Ctype: "int",
					Name: "newdfd",
				},
				{
					Ctype: "const char *",
					Name: "newname",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	317: &SyscallData{
		Number: 317,
		Name: "seccomp",
		EntryPoint: "sys_seccomp",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "op",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "void *",
					Name: "uargs",
				},
			},
		},
	},
	318: &SyscallData{
		Number: 318,
		Name: "getrandom",
		EntryPoint: "sys_getrandom",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "char *",
					Name: "buf",
				},
				{
					Ctype: "size_t",
					Name: "count",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	319: &SyscallData{
		Number: 319,
		Name: "memfd_create",
		EntryPoint: "sys_memfd_create",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "uname_ptr",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	320: &SyscallData{
		Number: 320,
		Name: "kexec_file_load",
		EntryPoint: "sys_kexec_file_load",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "kernel_fd",
				},
				{
					Ctype: "int",
					Name: "initrd_fd",
				},
				{
					Ctype: "unsigned long",
					Name: "cmdline_len",
				},
				{
					Ctype: "const char *",
					Name: "cmdline_ptr",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	321: &SyscallData{
		Number: 321,
		Name: "bpf",
		EntryPoint: "sys_bpf",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "cmd",
				},
				{
					Ctype: "union bpf_attr *",
					Name: "attr",
				},
				{
					Ctype: "unsigned int",
					Name: "size",
				},
			},
		},
	},
	322: &SyscallData{
		Number: 322,
		Name: "execveat",
		EntryPoint: "sys_execveat",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "const char *const *",
					Name: "argv",
				},
				{
					Ctype: "const char *const *",
					Name: "envp",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	323: &SyscallData{
		Number: 323,
		Name: "userfaultfd",
		EntryPoint: "sys_userfaultfd",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	324: &SyscallData{
		Number: 324,
		Name: "membarrier",
		EntryPoint: "sys_membarrier",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "cmd",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "int",
					Name: "cpu_id",
				},
			},
		},
	},
	325: &SyscallData{
		Number: 325,
		Name: "mlock2",
		EntryPoint: "sys_mlock2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	326: &SyscallData{
		Number: 326,
		Name: "copy_file_range",
		EntryPoint: "sys_copy_file_range",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd_in",
				},
				{
					Ctype: "loff_t *",
					Name: "off_in",
				},
				{
					Ctype: "int",
					Name: "fd_out",
				},
				{
					Ctype: "loff_t *",
					Name: "off_out",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	327: &SyscallData{
		Number: 327,
		Name: "preadv2",
		EntryPoint: "sys_preadv2",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_l",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_h",
				},
				{
					Ctype: "rwf_t",
					Name: "flags",
				},
			},
		},
	},
	328: &SyscallData{
		Number: 328,
		Name: "pwritev2",
		EntryPoint: "sys_pwritev2",
		CompatEntry: "",
		Implementation: "64",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_l",
				},
				{
					Ctype: "unsigned long",
					Name: "pos_h",
				},
				{
					Ctype: "rwf_t",
					Name: "flags",
				},
			},
		},
	},
	329: &SyscallData{
		Number: 329,
		Name: "pkey_mprotect",
		EntryPoint: "sys_pkey_mprotect",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "start",
				},
				{
					Ctype: "size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned long",
					Name: "prot",
				},
				{
					Ctype: "int",
					Name: "pkey",
				},
			},
		},
	},
	330: &SyscallData{
		Number: 330,
		Name: "pkey_alloc",
		EntryPoint: "sys_pkey_alloc",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
				{
					Ctype: "unsigned long",
					Name: "init_val",
				},
			},
		},
	},
	331: &SyscallData{
		Number: 331,
		Name: "pkey_free",
		EntryPoint: "sys_pkey_free",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "pkey",
				},
			},
		},
	},
	332: &SyscallData{
		Number: 332,
		Name: "statx",
		EntryPoint: "sys_statx",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
				{
					Ctype: "unsigned",
					Name: "mask",
				},
				{
					Ctype: "struct statx *",
					Name: "buffer",
				},
			},
		},
	},
	333: &SyscallData{
		Number: 333,
		Name: "io_pgetevents",
		EntryPoint: "sys_io_pgetevents",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "aio_context_t",
					Name: "ctx_id",
				},
				{
					Ctype: "long",
					Name: "min_nr",
				},
				{
					Ctype: "long",
					Name: "nr",
				},
				{
					Ctype: "struct io_event *",
					Name: "events",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "timeout",
				},
				{
					Ctype: "const struct __aio_sigset *",
					Name: "sig",
				},
			},
		},
	},
	334: &SyscallData{
		Number: 334,
		Name: "rseq",
		EntryPoint: "sys_rseq",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct rseq *",
					Name: "rseq",
				},
				{
					Ctype: "uint32_t",
					Name: "rseq_len",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
				{
					Ctype: "uint32_t",
					Name: "sig",
				},
			},
		},
	},
	424: &SyscallData{
		Number: 424,
		Name: "pidfd_send_signal",
		EntryPoint: "sys_pidfd_send_signal",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "pidfd",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
				{
					Ctype: "siginfo_t *",
					Name: "info",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	425: &SyscallData{
		Number: 425,
		Name: "io_uring_setup",
		EntryPoint: "sys_io_uring_setup",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "u32",
					Name: "entries",
				},
				{
					Ctype: "struct io_uring_params *",
					Name: "p",
				},
			},
		},
	},
	426: &SyscallData{
		Number: 426,
		Name: "io_uring_enter",
		EntryPoint: "sys_io_uring_enter",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "u32",
					Name: "to_submit",
				},
				{
					Ctype: "u32",
					Name: "min_complete",
				},
				{
					Ctype: "u32",
					Name: "flags",
				},
				{
					Ctype: "const void *",
					Name: "argp",
				},
				{
					Ctype: "size_t",
					Name: "argsz",
				},
			},
		},
	},
	427: &SyscallData{
		Number: 427,
		Name: "io_uring_register",
		EntryPoint: "sys_io_uring_register",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "op",
				},
				{
					Ctype: "void *",
					Name: "arg",
				},
				{
					Ctype: "unsigned int",
					Name: "nr_args",
				},
			},
		},
	},
	428: &SyscallData{
		Number: 428,
		Name: "open_tree",
		EntryPoint: "sys_open_tree",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
			},
		},
	},
	429: &SyscallData{
		Number: 429,
		Name: "move_mount",
		EntryPoint: "sys_move_mount",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "from_dfd",
				},
				{
					Ctype: "const char *",
					Name: "from_path",
				},
				{
					Ctype: "int",
					Name: "to_dfd",
				},
				{
					Ctype: "const char *",
					Name: "to_path",
				},
				{
					Ctype: "unsigned int",
					Name: "ms_flags",
				},
			},
		},
	},
	430: &SyscallData{
		Number: 430,
		Name: "fsopen",
		EntryPoint: "sys_fsopen",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "fs_name",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	431: &SyscallData{
		Number: 431,
		Name: "fsconfig",
		EntryPoint: "sys_fsconfig",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fs_fd",
				},
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
				{
					Ctype: "const char *",
					Name: "key",
				},
				{
					Ctype: "const void *",
					Name: "value",
				},
				{
					Ctype: "int",
					Name: "aux",
				},
			},
		},
	},
	432: &SyscallData{
		Number: 432,
		Name: "fsmount",
		EntryPoint: "sys_fsmount",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fs_fd",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "unsigned int",
					Name: "ms_flags",
				},
			},
		},
	},
	433: &SyscallData{
		Number: 433,
		Name: "fspick",
		EntryPoint: "sys_fspick",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	434: &SyscallData{
		Number: 434,
		Name: "pidfd_open",
		EntryPoint: "sys_pidfd_open",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	435: &SyscallData{
		Number: 435,
		Name: "clone3",
		EntryPoint: "sys_clone3",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct clone_args *",
					Name: "uargs",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	436: &SyscallData{
		Number: 436,
		Name: "close_range",
		EntryPoint: "sys_close_range",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "max_fd",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	437: &SyscallData{
		Number: 437,
		Name: "openat2",
		EntryPoint: "sys_openat2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "struct open_how *",
					Name: "how",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
			},
		},
	},
	438: &SyscallData{
		Number: 438,
		Name: "pidfd_getfd",
		EntryPoint: "sys_pidfd_getfd",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "pidfd",
				},
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	439: &SyscallData{
		Number: 439,
		Name: "faccessat2",
		EntryPoint: "sys_faccessat2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "int",
					Name: "mode",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	440: &SyscallData{
		Number: 440,
		Name: "process_madvise",
		EntryPoint: "sys_process_madvise",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "pidfd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "size_t",
					Name: "vlen",
				},
				{
					Ctype: "int",
					Name: "behavior",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	441: &SyscallData{
		Number: 441,
		Name: "epoll_pwait2",
		EntryPoint: "sys_epoll_pwait2",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "epfd",
				},
				{
					Ctype: "struct epoll_event *",
					Name: "events",
				},
				{
					Ctype: "int",
					Name: "maxevents",
				},
				{
					Ctype: "const struct __kernel_timespec *",
					Name: "timeout",
				},
				{
					Ctype: "const sigset_t *",
					Name: "sigmask",
				},
				{
					Ctype: "size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	442: &SyscallData{
		Number: 442,
		Name: "mount_setattr",
		EntryPoint: "sys_mount_setattr",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "path",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "struct mount_attr *",
					Name: "uattr",
				},
				{
					Ctype: "size_t",
					Name: "usize",
				},
			},
		},
	},
	443: &SyscallData{
		Number: 443,
		Name: "quotactl_fd",
		EntryPoint: "sys_quotactl_fd",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
				{
					Ctype: "qid_t",
					Name: "id",
				},
				{
					Ctype: "void *",
					Name: "addr",
				},
			},
		},
	},
	444: &SyscallData{
		Number: 444,
		Name: "landlock_create_ruleset",
		EntryPoint: "sys_landlock_create_ruleset",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const struct landlock_ruleset_attr *",
					Name: "attr",
				},
				{
					Ctype: "size_t",
					Name: "size",
				},
				{
					Ctype: "__u32",
					Name: "flags",
				},
			},
		},
	},
	445: &SyscallData{
		Number: 445,
		Name: "landlock_add_rule",
		EntryPoint: "sys_landlock_add_rule",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "ruleset_fd",
				},
				{
					Ctype: "enum landlock_rule_type",
					Name: "rule_type",
				},
				{
					Ctype: "const void *",
					Name: "rule_attr",
				},
				{
					Ctype: "__u32",
					Name: "flags",
				},
			},
		},
	},
	446: &SyscallData{
		Number: 446,
		Name: "landlock_restrict_self",
		EntryPoint: "sys_landlock_restrict_self",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "ruleset_fd",
				},
				{
					Ctype: "__u32",
					Name: "flags",
				},
			},
		},
	},
	447: &SyscallData{
		Number: 447,
		Name: "memfd_secret",
		EntryPoint: "sys_memfd_secret",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	448: &SyscallData{
		Number: 448,
		Name: "process_mrelease",
		EntryPoint: "sys_process_mrelease",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "pidfd",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	449: &SyscallData{
		Number: 449,
		Name: "futex_waitv",
		EntryPoint: "sys_futex_waitv",
		CompatEntry: "",
		Implementation: "common",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct futex_waitv *",
					Name: "waiters",
				},
				{
					Ctype: "unsigned int",
					Name: "nr_futexes",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "timeout",
				},
				{
					Ctype: "clockid_t",
					Name: "clockid",
				},
			},
		},
	},
	512: &SyscallData{
		Number: 512,
		Name: "rt_sigaction",
		EntryPoint: "compat_sys_rt_sigaction",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "sig",
				},
				{
					Ctype: "const struct compat_sigaction *",
					Name: "act",
				},
				{
					Ctype: "struct compat_sigaction *",
					Name: "oact",
				},
				{
					Ctype: "compat_size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	513: &SyscallData{
		Number: 513,
		Name: "rt_sigreturn",
		EntryPoint: "compat_sys_x32_rt_sigreturn",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "__unused",
				},
			},
		},
	},
	514: &SyscallData{
		Number: 514,
		Name: "ioctl",
		EntryPoint: "compat_sys_ioctl",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned int",
					Name: "fd",
				},
				{
					Ctype: "unsigned int",
					Name: "cmd",
				},
				{
					Ctype: "compat_ulong_t",
					Name: "arg",
				},
			},
		},
	},
	515: &SyscallData{
		Number: 515,
		Name: "readv",
		EntryPoint: "sys_readv",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
			},
		},
	},
	516: &SyscallData{
		Number: 516,
		Name: "writev",
		EntryPoint: "sys_writev",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
			},
		},
	},
	517: &SyscallData{
		Number: 517,
		Name: "recvfrom",
		EntryPoint: "compat_sys_recvfrom",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "void *",
					Name: "buf",
				},
				{
					Ctype: "compat_size_t",
					Name: "len",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
				{
					Ctype: "struct sockaddr *",
					Name: "addr",
				},
				{
					Ctype: "int *",
					Name: "addrlen",
				},
			},
		},
	},
	518: &SyscallData{
		Number: 518,
		Name: "sendmsg",
		EntryPoint: "compat_sys_sendmsg",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct compat_msghdr *",
					Name: "msg",
				},
				{
					Ctype: "unsigned",
					Name: "flags",
				},
			},
		},
	},
	519: &SyscallData{
		Number: 519,
		Name: "recvmsg",
		EntryPoint: "compat_sys_recvmsg",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct compat_msghdr *",
					Name: "msg",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	520: &SyscallData{
		Number: 520,
		Name: "execve",
		EntryPoint: "compat_sys_execve",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "const compat_uptr_t *",
					Name: "argv",
				},
				{
					Ctype: "const compat_uptr_t *",
					Name: "envp",
				},
			},
		},
	},
	521: &SyscallData{
		Number: 521,
		Name: "ptrace",
		EntryPoint: "compat_sys_ptrace",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "compat_long_t",
					Name: "request",
				},
				{
					Ctype: "compat_long_t",
					Name: "pid",
				},
				{
					Ctype: "compat_long_t",
					Name: "addr",
				},
				{
					Ctype: "compat_long_t",
					Name: "data",
				},
			},
		},
	},
	522: &SyscallData{
		Number: 522,
		Name: "rt_sigpending",
		EntryPoint: "compat_sys_rt_sigpending",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "compat_sigset_t *",
					Name: "uset",
				},
				{
					Ctype: "compat_size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	523: &SyscallData{
		Number: 523,
		Name: "rt_sigtimedwait",
		EntryPoint: "compat_sys_rt_sigtimedwait_time64",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "compat_sigset_t *",
					Name: "uthese",
				},
				{
					Ctype: "struct compat_siginfo *",
					Name: "uinfo",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "uts",
				},
				{
					Ctype: "compat_size_t",
					Name: "sigsetsize",
				},
			},
		},
	},
	524: &SyscallData{
		Number: 524,
		Name: "rt_sigqueueinfo",
		EntryPoint: "compat_sys_rt_sigqueueinfo",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "compat_pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
				{
					Ctype: "struct compat_siginfo *",
					Name: "uinfo",
				},
			},
		},
	},
	525: &SyscallData{
		Number: 525,
		Name: "sigaltstack",
		EntryPoint: "compat_sys_sigaltstack",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "const compat_stack_t *",
					Name: "uss_ptr",
				},
				{
					Ctype: "compat_stack_t *",
					Name: "uoss_ptr",
				},
			},
		},
	},
	526: &SyscallData{
		Number: 526,
		Name: "timer_create",
		EntryPoint: "compat_sys_timer_create",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "clockid_t",
					Name: "which_clock",
				},
				{
					Ctype: "struct compat_sigevent *",
					Name: "timer_event_spec",
				},
				{
					Ctype: "timer_t *",
					Name: "created_timer_id",
				},
			},
		},
	},
	527: &SyscallData{
		Number: 527,
		Name: "mq_notify",
		EntryPoint: "compat_sys_mq_notify",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "mqd_t",
					Name: "mqdes",
				},
				{
					Ctype: "const struct compat_sigevent *",
					Name: "u_notification",
				},
			},
		},
	},
	528: &SyscallData{
		Number: 528,
		Name: "kexec_load",
		EntryPoint: "compat_sys_kexec_load",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "compat_ulong_t",
					Name: "entry",
				},
				{
					Ctype: "compat_ulong_t",
					Name: "nr_segments",
				},
				{
					Ctype: "struct compat_kexec_segment *",
					Name: "segments",
				},
				{
					Ctype: "compat_ulong_t",
					Name: "flags",
				},
			},
		},
	},
	529: &SyscallData{
		Number: 529,
		Name: "waitid",
		EntryPoint: "compat_sys_waitid",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "which",
				},
				{
					Ctype: "compat_pid_t",
					Name: "pid",
				},
				{
					Ctype: "struct compat_siginfo *",
					Name: "info",
				},
				{
					Ctype: "int",
					Name: "options",
				},
				{
					Ctype: "struct compat_rusage *",
					Name: "ru",
				},
			},
		},
	},
	530: &SyscallData{
		Number: 530,
		Name: "set_robust_list",
		EntryPoint: "compat_sys_set_robust_list",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "struct compat_robust_list_head *",
					Name: "head",
				},
				{
					Ctype: "compat_size_t",
					Name: "len",
				},
			},
		},
	},
	531: &SyscallData{
		Number: 531,
		Name: "get_robust_list",
		EntryPoint: "compat_sys_get_robust_list",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "pid",
				},
				{
					Ctype: "compat_uptr_t *",
					Name: "head_ptr",
				},
				{
					Ctype: "compat_size_t *",
					Name: "len_ptr",
				},
			},
		},
	},
	532: &SyscallData{
		Number: 532,
		Name: "vmsplice",
		EntryPoint: "sys_vmsplice",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "iov",
				},
				{
					Ctype: "unsigned long",
					Name: "nr_segs",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	533: &SyscallData{
		Number: 533,
		Name: "move_pages",
		EntryPoint: "sys_move_pages",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "unsigned long",
					Name: "nr_pages",
				},
				{
					Ctype: "const void * *",
					Name: "pages",
				},
				{
					Ctype: "const int *",
					Name: "nodes",
				},
				{
					Ctype: "int *",
					Name: "status",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	534: &SyscallData{
		Number: 534,
		Name: "preadv",
		EntryPoint: "compat_sys_preadv64",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "loff_t",
					Name: "pos",
				},
			},
		},
	},
	535: &SyscallData{
		Number: 535,
		Name: "pwritev",
		EntryPoint: "compat_sys_pwritev64",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "loff_t",
					Name: "pos",
				},
			},
		},
	},
	536: &SyscallData{
		Number: 536,
		Name: "rt_tgsigqueueinfo",
		EntryPoint: "compat_sys_rt_tgsigqueueinfo",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "compat_pid_t",
					Name: "tgid",
				},
				{
					Ctype: "compat_pid_t",
					Name: "pid",
				},
				{
					Ctype: "int",
					Name: "sig",
				},
				{
					Ctype: "struct compat_siginfo *",
					Name: "uinfo",
				},
			},
		},
	},
	537: &SyscallData{
		Number: 537,
		Name: "recvmmsg",
		EntryPoint: "compat_sys_recvmmsg_time64",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct compat_mmsghdr *",
					Name: "mmsg",
				},
				{
					Ctype: "unsigned",
					Name: "vlen",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
				{
					Ctype: "struct __kernel_timespec *",
					Name: "timeout",
				},
			},
		},
	},
	538: &SyscallData{
		Number: 538,
		Name: "sendmmsg",
		EntryPoint: "compat_sys_sendmmsg",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "struct compat_mmsghdr *",
					Name: "mmsg",
				},
				{
					Ctype: "unsigned",
					Name: "vlen",
				},
				{
					Ctype: "unsigned int",
					Name: "flags",
				},
			},
		},
	},
	539: &SyscallData{
		Number: 539,
		Name: "process_vm_readv",
		EntryPoint: "sys_process_vm_readv",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "const struct iovec *",
					Name: "lvec",
				},
				{
					Ctype: "unsigned long",
					Name: "liovcnt",
				},
				{
					Ctype: "const struct iovec *",
					Name: "rvec",
				},
				{
					Ctype: "unsigned long",
					Name: "riovcnt",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	540: &SyscallData{
		Number: 540,
		Name: "process_vm_writev",
		EntryPoint: "sys_process_vm_writev",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "pid_t",
					Name: "pid",
				},
				{
					Ctype: "const struct iovec *",
					Name: "lvec",
				},
				{
					Ctype: "unsigned long",
					Name: "liovcnt",
				},
				{
					Ctype: "const struct iovec *",
					Name: "rvec",
				},
				{
					Ctype: "unsigned long",
					Name: "riovcnt",
				},
				{
					Ctype: "unsigned long",
					Name: "flags",
				},
			},
		},
	},
	541: &SyscallData{
		Number: 541,
		Name: "setsockopt",
		EntryPoint: "sys_setsockopt",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "int",
					Name: "level",
				},
				{
					Ctype: "int",
					Name: "optname",
				},
				{
					Ctype: "char *",
					Name: "optval",
				},
				{
					Ctype: "int",
					Name: "optlen",
				},
			},
		},
	},
	542: &SyscallData{
		Number: 542,
		Name: "getsockopt",
		EntryPoint: "sys_getsockopt",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "fd",
				},
				{
					Ctype: "int",
					Name: "level",
				},
				{
					Ctype: "int",
					Name: "optname",
				},
				{
					Ctype: "char *",
					Name: "optval",
				},
				{
					Ctype: "int *",
					Name: "optlen",
				},
			},
		},
	},
	543: &SyscallData{
		Number: 543,
		Name: "io_setup",
		EntryPoint: "compat_sys_io_setup",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned",
					Name: "nr_reqs",
				},
				{
					Ctype: "u32 *",
					Name: "ctx32p",
				},
			},
		},
	},
	544: &SyscallData{
		Number: 544,
		Name: "io_submit",
		EntryPoint: "compat_sys_io_submit",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "compat_aio_context_t",
					Name: "ctx_id",
				},
				{
					Ctype: "int",
					Name: "nr",
				},
				{
					Ctype: "u32 *",
					Name: "iocb",
				},
			},
		},
	},
	545: &SyscallData{
		Number: 545,
		Name: "execveat",
		EntryPoint: "compat_sys_execveat",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "int",
					Name: "dfd",
				},
				{
					Ctype: "const char *",
					Name: "filename",
				},
				{
					Ctype: "const compat_uptr_t *",
					Name: "argv",
				},
				{
					Ctype: "const compat_uptr_t *",
					Name: "envp",
				},
				{
					Ctype: "int",
					Name: "flags",
				},
			},
		},
	},
	546: &SyscallData{
		Number: 546,
		Name: "preadv2",
		EntryPoint: "compat_sys_preadv64v2",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "loff_t",
					Name: "pos",
				},
				{
					Ctype: "rwf_t",
					Name: "flags",
				},
			},
		},
	},
	547: &SyscallData{
		Number: 547,
		Name: "pwritev2",
		EntryPoint: "compat_sys_pwritev64v2",
		CompatEntry: "",
		Implementation: "x32",
		ArgsInfo: ArgsInfo{
			ArgsType: ArgNormal,
			Args: []*ArgData{
				{
					Ctype: "unsigned long",
					Name: "fd",
				},
				{
					Ctype: "const struct iovec *",
					Name: "vec",
				},
				{
					Ctype: "unsigned long",
					Name: "vlen",
				},
				{
					Ctype: "loff_t",
					Name: "pos",
				},
				{
					Ctype: "rwf_t",
					Name: "flags",
				},
			},
		},
	},
}
