package main

// Command represents a Linux/Unix command to run, with description and optional args
type Command struct {
	Name string
	Desc string
	Args string
}

// Array of commands to run with description for output, optional Args as a string
var (
	Commands = []Command{
		{Name: "date", Desc: "System Time"},
		{Name: "uptime", Desc: "System Uptime"},
		{Name: "uname", Desc: "System Information", Args: "-a"},
		{Name: "ifconfig", Desc: "Network Interfaces"},
		{Name: "df", Desc: "File System"},
		{Name: "lsof", Desc: "Opened files marked for deletion (Unlinked)", Args: "+L1"},
		{Name: "w", Desc: "Logged In Users"},
		{Name: "last", Desc: "List of user login/log outs"},
		{Name: "netstat", Desc: "Open TCP sockets", Args: "-nat"},
		{Name: "netstat", Desc: "Open UDP sockets", Args: "-nau"},
	}
)
