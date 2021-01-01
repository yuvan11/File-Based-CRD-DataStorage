## Prerequisites 

1.Golang version go1.15 required

Before Run the program, please follow the steps,

1. Install golang  (if your already install golang please skip the step)
    Note  - User - root
   
    Download  Golang package 
		https://golang.org/doc/install
	
   	   1. unzip the download folder 
   	   2. tar -xzf golangpackage name 
	   3. Copy go folder 
	   4. cp -rf  go /user/local
	   5. Set golang path 
 	   6. vim ~/.bashrc
	   
	   Add the below line in the bashrc 
		1. export GOROOT=/usr/local/go
		2. export GOPATH=$HOME/Goprojects
		3 .export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
		
	   7. Save the file
	   8.source  ~/.bashrc

	Check go version in comment line
		
		go version 

	 Output 
		go version go1.15 linux/amd64

	
2. Run the program  
    
     Go to project src folder run the program 
     go run main.go

    Example,

root@yuvi:~/Goprojects/src/File-Based-CRD/src# go run main.go


## System Requirement
    
  Operating System : linux 
 

