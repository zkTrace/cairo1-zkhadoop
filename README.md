#  Multi-Processing Cairo1 for a Verifiable GPU
<img width="1150" alt="Screenshot 2024-02-24 at 11 06 08 AM" src="https://github.com/STARKScale/zkhadoop-cairo1/assets/47396265/f80e95fb-4d41-4509-83b0-caa99aa413c5">

---

This repository implements the logic of building the distributed system for sharding Cairo1 execution and recursive proving system for a comrpessed monolithic proof of this execution.

## Problem Statement

In the world of machine learning and data analytics, ensuring privacy while performing computations is crucial. However, current Zero-Knowledge Proof (ZKP) systems struggle with large-scale verifiable computations, making it challenging to achieve privacy without sacrificing efficiency.

Simply put, zkML is slow and immature. Think of it as the difference between manually tilling land 🌾 and using a high-powered tractor 🚜.

One leading frontier of improvement in this area currently is GizaTech, but one of the most glaring bottlenecks in their workflow of verifiable inferences is the inefficiency of large-scale verifiable computations like matrix multiplication. With the release of AI Actions SDK, the friction for developers to instantly create a zkML workflow has diminished, but the friction of proof generation and computation remains the same, especially when we're looking at datasets of 3,000,000+ inputs.

## Our Solution: Verifiable zk-SPARK / Hadoop MapReduce

We introduce a ZK-SPARK distributed system that parallelizes the matrix multiplication step of generating an AI inference and generates a proof of integrity of our work alongside it. Instead of for-loops in Cairo, we have created a CAIRO/GO implementation of Hadoop MapReduce that uses Recursive Proving and Herodotus to verify our work in logn time.

- **Coordinator:** Manager node that lives on a Go Server and facilitates this process.
- **Distributed File System:** Scaling data storage and memory capacities of inputs.
- **Mapper:** Worker node that executes subtasks of the computation in CAIRO.
- **Reducer:** Worker node that compiles the outputs of the CAIRO trace computations.
- **Recursive Prover:** Compiles proofs across all actions into one singular source of truth.

Our tech works not just as GizaTech's GPU, but as any computational force needing MapReduce algorithms: open source AI models, page rank, recommendation systems.

## What We Built

In this first prototype, we fully implemented the Hadoop MapReduce for matrix multiplication where a CAIRO trace of integrity of work and the computed output is returned. This can significantly speed up the matrix multiplication subroutine of GizaTech's verifiable inference generation as the work is parallelized and distributed.

- User inputs Matrix and Vector as a JSON file.
- Coordinator (Go server) splits task into subtasks for parallelized compute.
- Coordinator (Go server) gossips the subtasks into various Goroutines.
- Each lightweight thread injects distributed matrix and vector into CAIRO.
- Each lightweight thread executes a bash script to compute a CAIRO trace and computation output.
- Each lightweight thread notifies the coordinator of a completed task.
- Coordinator (Go server) gossips the received output to a reducer to aggregate together all the distributed values.
- Reducer threads (Go worker) aggregate / reduce the output and returns a finalized value.

For the sake of this demo, we are currently operating on one node.

## What is Next

For full cohesiveness, Herodotus's verifier plays a large role in verifying our verifiable inference computations. This is the next step for rounding out the process and fully integrating with GizaTech's mature tech stack. In the meantime, we aim to implement a recursive proof system so that we can compile a singular proof to represent all the work done in this complex system.

Although we started with MapReduce, we envision this zkSPARK system scaling big data analysis and zkML/algorithmic computation:

- **Feb 16:** MVP Distributed System, MapReduce Cairo Algorithm, Matrix Multiplication Support
- **Q2 2024:** Recursive Proof, Verifier settling on Starknet, IO File Execution, more algorithm support
- **Q4 2024:** SDK interface, Generalized algorithm computation support, dev tooling for relayers and indexers
- **Q2 2025:** SPARK Engine, Decentralized Worker Node Program, Decentralized L2 Network, zkML models on-chain execution

🚀 zkML is a pioneering field of innovation that opens doors to actions such as claiming IP on AI agents, ensuring honest AI computations, and decentralizing computing protocols. But for these possibilities, we believe that we need to lay the groundwork for scaling ZK computations horizontally through STARKScale.

## Running the Code: Docker Compose

# ====== hot reload docker compse ======

```sh
docker-compose up --build

# separate terminal
docker-compose exec app bash
cd server
go test -v ./mr/common -run TestAggregateCairo

cd server/
bash test-matmult.sh

#misc
docker-compose down
docker ps
```
