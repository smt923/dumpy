# dumpy
Run through an array of Linux commands and generate a simple report continaing the results

## Usage
Simply run the binary on the system to gather information from, the results are printed to stdout so you can pipe easily pipe it

```bash
./dumpy > /some/location/report.txt
```

## Customization
Editing the commands currently required recompiling the project after editing `commands.go`, I plan to allow some configuration via either a file or stdin in the future (and PRs or requests for more defaults are welcome)

## Example
(Some sections shortened (marked with `...`), the program will return the full result of every command)

```md
Dump started at: 2017-08-04T01:25:30Z
---------------------------------------------
# [System Time]:

Fri  4 Aug 01:25:31 UTC 2017

---------------------------------------------
# [System Uptime]:

 01:25:31 up  4:30,  1 user,  load average: 1.24, 1.25, 1.26

---------------------------------------------
# [System Information]:

Linux ... 4.10.0-28-generic #32~16.04.2-Ubuntu SMP Thu Jul 20 10:19:48 UTC 2017 x86_64 x86_64 x86_64 GNU/Linux

---------------------------------------------
# [Network Interfaces]:

lo        Link encap:Local Loopback  
          inet addr:127.0.0.1  Mask:255.0.0.0
          inet6 addr: ::1/128 Scope:Host
          UP LOOPBACK RUNNING  MTU:65536  Metric:1
          RX packets:6958 errors:0 dropped:0 overruns:0 frame:0
          TX packets:6958 errors:0 dropped:0 overruns:0 carrier:0
          collisions:0 txqueuelen:1000 
          RX bytes:846378 (846.3 KB)  TX bytes:846378 (846.3 KB)

...

---------------------------------------------
# [File System]:

Filesystem         1K-blocks      Used Available Use% Mounted on
udev                  730312         0    730312   0% /dev
tmpfs                 150516      8924    141592   6% /run
/dev/sda1           21542460   8703204  11721920  43% /
tmpfs                 752572       172    752400   1% /dev/shm
tmpfs                   5120         4      5116   1% /run/lock
tmpfs                 752572         0    752572   0% /sys/fs/cgroup
tmpfs                 150516        96    150420   1% /run/user/1000
/home/smt/vmshared 976760828 825678796 151082032  85% /home/smt/vmshared

---------------------------------------------
# [Opened files marked for deletion (Unlinked)]:

COMMAND     PID USER   FD   TYPE DEVICE SIZE/OFF NLINK  NODE NAME
gnome-ter 13341  smt   15u   REG    8,1   983040     0 14870 /tmp/#14870 (deleted)
gnome-ter 13341  smt   16u   REG    8,1   262144     0 14871 /tmp/#14871 (deleted)
unity-sco 23714  smt    9u   REG    8,1    65536     0 15478 /tmp/tmpfTmijDs (deleted)
unity-sco 23714  smt   10u   REG    8,1    32768     0 15479 /tmp/tmpfMrwq5L (deleted)
unity-sco 23714  smt   11u   REG    8,1    32768     0 15480 /tmp/tmpfYs7yx5 (deleted)
unity-sco 23714  smt   12u   REG    8,1    16384     0   999 /tmp/tmpf0fu8ah (deleted)
unity-sco 23714  smt   13u   REG    8,1    16384     0  1000 /tmp/tmpfuLT3ii (deleted)

---------------------------------------------
# [Logged In Users]:

 01:25:32 up  4:30,  1 user,  load average: 1.47, 1.30, 1.28
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
smt      tty7     :0               19:33    5:52m 10:29   1.22s /sbin/upstart --user

---------------------------------------------
# [List of user login/log outs]:

smt      tty7         :0               Thu Aug  3 19:33    gone - no logout
reboot   system boot  4.10.0-28-generi Thu Aug  3 19:33   still running
smt      tty7         :0               Thu Aug  3 19:22 - down   (00:10)
reboot   system boot  4.10.0-28-generi Thu Aug  3 19:21 - 19:32  (00:10)

wtmp begins Thu Aug  3 19:21:52 2017

---------------------------------------------
# [Open TCP sockets]:

Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State      
tcp        0      0 127.0.1.1:53            0.0.0.0:*               LISTEN     
tcp6       0      0 :::80                   :::*                    LISTEN     

...

---------------------------------------------
# [Open UDP sockets]:

Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State      
udp        0      0 0.0.0.0:59710           0.0.0.0:*                          
udp        0      0 0.0.0.0:631             0.0.0.0:*                          
udp        0      0 0.0.0.0:59253           0.0.0.0:*                          
udp        0      0 127.0.1.1:53            0.0.0.0:*                          

...

```