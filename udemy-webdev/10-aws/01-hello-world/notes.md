--------------------------------------------
Personal Notes when creating AWS EC2 Server
--------------------------------------------
1. Created an Ubuntu t2.micro server using the free tier
    - Left mostly defaults, added HTTP to open ports
    - created a Key-Value Pair .pem file
    - launched instance and activated billing alerts

2. Moved my .pem to a safe place
    mv <SRC> <DEST>
    mv ~/Downloads/aws-kp-2018-02-05.pem ~/.ssh

3. Change permissions of .pem file
    chmod <PERMISSIONS> <FILE>
    chmod 400 ~/.ssh/aws-kp-2018-02-05.pem

4. Build our binary for a linux system
    GOOS=<OPERATING SYSTEM> GOARCH=<ARCHITECTURE> go build -o <FILENAME>
    GOOS=linux GOARCH=amd64 go build -o hello-binary

5. Copy binary to the AWS server and say yes to "Authenticity of host... cannot be established"
    scp -i <PEM FILE> <BINARY> <USER>@<PUBLIC-DNS>:
    scp -i ~/.ssh/aws-kp-2018-02-05.pem hello-binary ubuntu@ec2-52-207-57-168.compute-1.amazonaws.com:

6. SSH into your server
    ssh -i <PEM FILE> <USER>@<PUBLIC-DNS>
    ssh -i ~/.ssh/aws-kp-2018-02-05.pem ubuntu@ec2-52-207-57-168.compute-1.amazonaws.com

7. Run our code:
    - Change permissions of our binary:
        sudo chmod 700 hello-binary
    - Run the binary
        sudo ./hello-binary
    - Visit our webserver using a browser
        on Chrome, went to: 52.207.57.168
    - Close terminal and check server again
        * It is dead now.

8. Persist Code:
    - Log into our remote AWS EC2 Server
        ssh -i <PEM FILE> <USER>@<PUBLIC-DNS> OR ssh -i <PEM FILE> <USER>@<IP-ADDRESS>
        ssh -i ~/.ssh/aws-kp-2018-02-05.pem ubuntu@52.207.57.168
    - Navigate to /etc/systemd/system
        cd /etc/systemd/system/
    - Create configuration file
        vim hello-binary.service
            [Unit]
            Description=Go Server

            [Service]
            ExecStart=/home/ubuntu/hello-binary
            User=root
            Group=root
            Restart=always

            [Install]
            WantedBy=multi-user.target
    - Add the service to systemd.
        sudo systemctl enable hello-binary.service
    - Activate the service.
        sudo systemctl start hello-binary.service
    - Check if systemd started it.
        sudo systemctl status hello-binary.service
    - If we want to stop the systemd:
        sudo systemctl stop hello-binary.service

    
