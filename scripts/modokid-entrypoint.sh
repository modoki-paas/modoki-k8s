#!/bin/sh

usage() {
    echo "Usage: $0 [-m] [-q] [-w]" 1>&2
    echo "Options: " 1>&2
    echo "-m: Run migration" 1>&2
    echo "-q: Quit without running server" 1>&2
    echo "-w: Wait for database to start" 1>&2
    exit 1
}

while getopts :mqwh OPT
do
    case $OPT in
    m)  MIGRATION=1
        ;;
    q)  QUIT=1
        ;;
    w)  WAIT=1
        ;;
    h)  usage
        ;;
    \?) usage
        ;;
    esac
done

if [ "$WAIT" = "1" ]; then
    echo "Waiting for db"
    dockerize -wait tcp://$DB_HOST -timeout 60s
fi

if [ "$MIGRATION" = "1" ]; then
    echo "Running migration"
    cat /schema/*.sql | mysqldef --user=$DB_USER --password=$DB_PASSWORD --host=$DB_HOST $DB_DATABASE
fi

if [ "$QUIT" = "1" ]; then
    exit 0
fi

modokid --db="$DB_USER:$DB_PASSWORD@$DB_HOST/$DB_DATABASE?parseTime=true"
