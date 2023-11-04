#!/bin/bash

# Run the binary
./bin/main &

# Get the PID of the last background process
PID=$!

# Kill the process when the script exits
trap "kill $PID" EXIT

# Open the web page
open http://localhost:8080

# Wait for the process to finish
wait $PID