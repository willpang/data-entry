#!/bin/bash

# Find the PID of the process running on port 8080
PID=$(lsof -t -i:8080)

# If a process is found, kill it
if [ -n "$PID" ]; then
    echo "Killing process on port 8080 with PID: $PID"
    kill $PID
else
    echo "No process running on port 8080"
fi