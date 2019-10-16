PROJECTNAME=project_name

PROJECTPATH=project_path

if [ -z "$1" ]
then
        echo "Not a root path."
        exit 1
fi

LOGPATH="$1"logs/"$PROJECTNAME".log

function log() {
        echo $* >> $LOGPATH
}

log $( date +"%Y-%m-%d %H:%M:%S" 2>&1 )

cd /home/ubuntu/"$PROJECTPATH"
log $( git pull 2>&1 )
log ""
log $( sudo supervisorctl stop $PROJECTNAME 2>&1 )

log ""

export GOPATH=/home/ubuntu/
export GOCACHE="/home/ubuntu/.cache/go-build"
log $( go build  2>&1 )

log $( sudo supervisorctl start $PROJECTNAME 2>&1 )

log ""

log ""