# A easy devops. Hook scripts for work. The project is created to convenient and simplify for my work.

## Installing and Getting started

1. Clone the repository.

```bash
git clone git@github.com:guaidashu/hook_scripts.git
```

2. Usage.

    (1) Modify hook_scripts/config/config.yml. If the file is not exists, please creating it and adding the following.
    
        hookpath:
          "/home/scripts/"
        app:
          host: "127.0.0.1"
          port: 8099

    (2) First you should create a shell script or other scripts which can be executed. /home/scripts/test.sh
        
        #!/bin/bash
        
        LOGPATH="$1"logs/project_log.log
        
        date +"%Y-%m-%d %H:%M:%S" >> $LOGPATH
        
        cd project_root_dir
        git pull >> $LOGPATH 2>&1
        echo "" >> $LOGPATH
        sudo supervisorctl stop project_name >> $LOGPATH 2>&1
        
        echo "" >> $LOGPATH
        
        export GOPATH=/home/ubuntu/
        export GOCACHE="/home/ubuntu/.cache/go-build"
        go build >> $LOGPATH 2>&1
        
        sudo supervisorctl start project_name >> $LOGPATH 2>&1
        
        echo "" >> $LOGPATH
        
        echo "" >> $LOGPATH
        
    (3) Modify permissions
    
        chmod 755 /home/scripts/test.sh
        
    (4) Run hook_scripts
    
        go build
        ./hook_scripts 
    
## FAQ

Contact to me with email "1023767856@qq.com" or "song42960@gmail.com"

## Running Tests

curl 127.0.0.1:8099/hook/test

## Finally Thanks

Thanks for your support.