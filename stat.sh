#!/bin/bash
#
export GOGC=200

function startStat() {
	/home/ubuntu/stats/build/pancake-statas --config-path /home/ubuntu/stats/config/config.json
}

function stopStat() {
    pid=`ps -ef | grep /home/ubuntu/stats/build/pancake-statas | grep -v grep | awk '{print $2}'`
    if [ -n "$pid" ]; then
        for((i=1;i<=4;i++));
        do
            kill $pid
            sleep 5
            pid=`ps -ef | grep  /home/ubuntu/stats/build/pancake-statas | grep -v grep | awk '{print $2}'`
            if [ -z "$pid" ]; then
                #echo "statas stoped"
                break
            elif [ $i -eq 4 ]; then
                kill -9 $kid
            fi
        done
    fi
}

CMD=$1

case $CMD in
-start)
    echo "start"
    startStat
    ;;
-stop)
    echo "stop"
    stopStat
    ;;
-restart)
    stopStat
    sleep 3
    startStat
    ;;
*)
    echo "Usage: stat.sh -start | -stop | -restart .Or use systemctl start | stop | restart stat.service "
    ;;
esac

