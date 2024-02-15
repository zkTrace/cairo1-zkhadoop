#!/bin/bash

#
# basic matrix vector multiplication test
#

# prints current matrix and vector json input
cat "../data/input.json"

# run the test in a fresh sub-directory.
rm -rf ../data/mr-tmp
mkdir ../data/mr-tmp || exit 1
cd ../data/mr-tmp || exit 1
rm -f mr-*

# make sure software is freshly built.
(cd .. && go build $RACE mrcoordinator.go) || exit 1
(cd .. && go build $RACE mrworker.go) || exit 1
(cd .. && go build $RACE mrsequential.go) || exit 1

# generate correct output
# TODO

echo '***' Starting matmul vec test.

timeout -k 2s 180s ../mrcoordinator ../data/input.json &
pid=$!

# give the coordinator time to create the sockets.
sleep 1

# start multiple workers. (1 for now)
timeout -k 2s 180s ../mrworker &

# wait for the coordinator to exit.
wait $pid

# since workers are required to exit when a job is completely finished,
# and not before, that means the job has finished.
sort mr-out* | grep . > mr-matmul-all
cat mr-matmul-all

echo '***' FINISHED RUNNING

exit 0