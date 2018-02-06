
➜  02-hands-on git:(master) ✗ GOOS=linux GOARCH=amd64 go build -o hands-on-binary
➜  02-hands-on git:(master) ✗ scp -i ~/.ssh/aws-kp-2018-02-05.pem hands-on-binary ubuntu@52.207.57.168
➜  02-hands-on git:(master) ✗ scp -i ~/.ssh/aws-kp-2018-02-05.pem hands-on-binary ubuntu@52.207.57.168:
hands-on-binary                                                                                                                                             100% 7985KB   3.9MB/s   00:02
➜  02-hands-on git:(master) ✗ ssh -i ~/.ssh/aws-kp-2018-02-05.pem ubuntu@52.207.57.168
Welcome to Ubuntu 16.04.3 LTS (GNU/Linux 4.4.0-1047-aws x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

0 packages can be updated.
0 updates are security updates.


Last login: Tue Feb  6 01:25:34 2018 from 64.62.224.29
ubuntu@ip-172-31-94-23:~$ ls
hands-on-binary  hello-binary
ubuntu@ip-172-31-94-23:~$ mkdir 01-hands-on
ubuntu@ip-172-31-94-23:~$ ls
01-hands-on  hands-on-binary  hello-binary
ubuntu@ip-172-31-94-23:~$ mv hands-on-binary 01-hands-on/
ubuntu@ip-172-31-94-23:~$ ls
01-hands-on  hello-binary
ubuntu@ip-172-31-94-23:~$ cd 01-hands-on/
ubuntu@ip-172-31-94-23:~/01-hands-on$ ls
hands-on-binary
ubuntu@ip-172-31-94-23:~/01-hands-on$ exit
logout
Connection to 52.207.57.168 closed.
➜  02-hands-on git:(master) ✗ scp -i ~/.ssh/aws-kp-2018-02-05.pem templates ubuntu@52.207.57.168:01-hands-on
templates: not a regular file
➜  02-hands-on git:(master) ✗ scp -r -i ~/.ssh/aws-kp-2018-02-05.pem templates ubuntu@52.207.57.168:01-hands-on
bar.gohtml                                                                                                                                                  100%  337     4.7KB/s   00:00
signup.gohtml                                                                                                                                               100%  663     7.8KB/s   00:00
login.gohtml                                                                                                                                                100%  362     5.0KB/s   00:00
index.gohtml                                                                                                                                                100%  416     5.8KB/s   00:00
➜  02-hands-on git:(master) ✗ ssh -i ~/.ssh/aws-kp-2018-02-05.pem ubuntu@52.207.57.168
Welcome to Ubuntu 16.04.3 LTS (GNU/Linux 4.4.0-1047-aws x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

  Get cloud support with Ubuntu Advantage Cloud Guest:
    http://www.ubuntu.com/business/services/cloud

0 packages can be updated.
0 updates are security updates.


Last login: Tue Feb  6 02:28:18 2018 from 64.62.224.29
ubuntu@ip-172-31-94-23:~$ cd 01-hands-on/
ubuntu@ip-172-31-94-23:~/01-hands-on$ ls -la
total 8000
drwxrwxr-x 3 ubuntu ubuntu    4096 Feb  6 02:31 .
drwxr-xr-x 5 ubuntu ubuntu    4096 Feb  6 02:29 ..
-rwxr-xr-x 1 ubuntu ubuntu 8176717 Feb  6 02:27 hands-on-binary
drwxr-xr-x 2 ubuntu ubuntu    4096 Feb  6 02:31 templates
ubuntu@ip-172-31-94-23:~/01-hands-on$ chmod 700 hands-on-binary
ubuntu@ip-172-31-94-23:~/01-hands-on$ ls -la
total 8000
drwxrwxr-x 3 ubuntu ubuntu    4096 Feb  6 02:31 .
drwxr-xr-x 5 ubuntu ubuntu    4096 Feb  6 02:29 ..
-rwx------ 1 ubuntu ubuntu 8176717 Feb  6 02:27 hands-on-binary
drwxr-xr-x 2 ubuntu ubuntu    4096 Feb  6 02:31 templates
ubuntu@ip-172-31-94-23:~/01-hands-on$ cd /
bin/        dev/        home/       lib64/      media/      opt/        root/       sbin/       srv/        tmp/        var/
boot/       etc/        lib/        lost+found/ mnt/        proc/       run/        snap/       sys/        usr/
ubuntu@ip-172-31-94-23:~/01-hands-on$ cd /etc/systemd/system/
ubuntu@ip-172-31-94-23:/etc/systemd/system$ sudo systemctl stop hello-binary.service
ubuntu@ip-172-31-94-23:/etc/systemd/system$ ls
cloud-init.target.wants  final.target.wants  graphical.target.wants  iscsi.service            network-online.target.wants  sockets.target.wants  sysinit.target.wants  timers.target.wants
default.target.wants     getty.target.wants  hello-binary.service    multi-user.target.wants  paths.target.wants           sshd.service          syslog.service
ubuntu@ip-172-31-94-23:/etc/systemd/system$ mv hello-binary.service _hello-binary.service
mv: cannot move 'hello-binary.service' to '_hello-binary.service': Permission denied
ubuntu@ip-172-31-94-23:/etc/systemd/system$ sudo mv hello-binary.service _hellow-binary.service
ubuntu@ip-172-31-94-23:/etc/systemd/system$ sudo vim hands-on-binary.service
ubuntu@ip-172-31-94-23:/etc/systemd/system$ sudo systemctl enable hands-on-binary.service
Created symlink from /etc/systemd/system/multi-user.target.wants/hands-on-binary.service to /etc/systemd/system/hands-on-binary.service.
ubuntu@ip-172-31-94-23:/etc/systemd/system$ sudo systemctl start hands-on-binary.service
ubuntu@ip-172-31-94-23:/etc/systemd/system$ sudo systemctl status hands-on-binary.service
● hands-on-binary.service - Go Server
   Loaded: loaded (/etc/systemd/system/hands-on-binary.service; enabled; vendor preset: enabled)
   Active: active (running) since Tue 2018-02-06 02:45:13 UTC; 6s ago
 Main PID: 2165 (hands-on-binary)
    Tasks: 3
   Memory: 920.0K
      CPU: 2ms
   CGroup: /system.slice/hands-on-binary.service
           └─2165 /home/ubuntu/01-hands-on/hands-on-binary

Feb 06 02:45:13 ip-172-31-94-23 systemd[1]: Started Go Server.
ubuntu@ip-172-31-94-23:/etc/systemd/system$