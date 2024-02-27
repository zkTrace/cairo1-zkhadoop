#!/bin/bash

#
# basic matrix vector multiplication test
#

# Define the project's root directory
PROJECT_ROOT="$(pwd)"

# run the test in a fresh sub-directory.
rm -rf ./data/mr-tmp
mkdir ./data/mr-tmp || exit 1

# make sure software is freshly built.
# takes too long, do it in the docker instead
(cd ./main && go build -race mrcoordinator.go) || exit 1
(cd ./main && go build -race mrworker.go) || exit 1
(cd ./main && go build -race mrsequential.go) || exit 1

# generate correct output
# TODO

echo '***' Starting matmul vec test.

echo Current input data:

# prints current matrix and vector json input
cat "./data/input.json"
echo -e "\n"

# Run the coordinator and workers from the correct directory
timeout -k 2s 180s "$PROJECT_ROOT/main/mrcoordinator" "$PROJECT_ROOT/data/input.json" &
pid=$!

# give the coordinator time to create the sockets.
sleep 1

# start multiple workers. (1 for now)
(cd "$PROJECT_ROOT" && timeout -k 2s 180s ./main/mrworker)

# wait for the coordinator to exit.
wait $pid

# Output results
# Make sure to adjust the path for mr-out* files if they are generated in a specific directory
# cat "$PROJECT_ROOT/data/mr-tmp/mr-out-0" 
# cat "$PROJECT_ROOT/data/mr-tmp/mr-out-0" | grep -E '{|}|\"' 
echo "Reducer Result for job_id_0:"
# awk '/{/,/}/{print}' "$PROJECT_ROOT/data/mr-tmp/mr-out-0"

# sort "$PROJECT_ROOT"/mr-out* | grep . > "$PROJECT_ROOT/mr-matmul-all"
# cat "$PROJECT_ROOT/mr-matmul-all"

echo '***' FINISHED RUNNING

exit 0
