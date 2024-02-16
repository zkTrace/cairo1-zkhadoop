STARKScale ~ Let us be your GPU for Verifiable zkML Inferences 🚀
Problem Statement
In the world of machine learning and data analytics, ensuring privacy while performing computations is crucial. However, current Zero-Knowledge Proof (ZKP) systems struggle with large-scale verifiable computations, making it challenging to achieve privacy without sacrificing efficiency. 

Simply put, zkML is slow and immature. Think of it as the difference between manually tilling land 🌾 and using a high-powered tractor 🚜. 

One leading frontier of improvement in this currently is GizaTech, but one of the most glaring bottlenecks in their workflow of verifiable inferences is the inefficiency of large-scale verifiable computations like matrix multiplicaiton. With the release of AI Actions SDK, the friction of developers to instantly create an zkml workflow has diminished, but the friction of proof generation and computation remains the same, especially when we're looking at datasets of 3,000,000+ inputs.



Our Solution: Verifiable zk-SPARK / Hadoop MapReduce
We introduce a ZK-SPARK distributed system that parallelizes the matrix multiplication step of generating an ai inference and generates a proof of integrity of our work alongside it. Instead of for-loops in Cairo, we have created a CAIRO/GO implementation of Hadoop MapReduce that uses Recursive Proving and Herodotus to verify our work in logn time.

Coordinator: manager node that lives on a Go Server and facilitates this process
Distributed File System: scaling data storage and memory capacities of inputs
Mapper: worker node that executes subtasks of the computation in CAIRO
Reducer: worker node that compiles the outputs of the CAIRO trace computations
Recursive Prover: compiles proofs across all actions into one singular source of truth
Our tech works not just as GizaTech's GPU, but as any computational force needing MapReduce algorithms: open source ai models, page rank, recommendation systems. 



What We Built
In this first prototype, we fully implemented the Hadoop MapReduce for matrix multiplication where a CAIRO trace of integrity of work and the computed output is returned. This can significantly speed up the matrix multiplication subroutine of GizaTech's verifiable inference generation as the work is paralellized and distributed.

User inputs Matrix and Vector as a json file
Coordinator (go server) splits task into subtasks for parallelized compute
Coordinator (go server) gossips the subtasks into various Goroutine
Each lightweight thread injects distributed matrix and vector into CAIRO
Each lightweight thread  executes a bash script to compute a CAIRO trace and computation output
Each lightweight thread notifies the coordinator of a completed task
Coordinate (go server) gossips the received output to a reducer to aggregator together all the distributed values
Reducer threads (go worker) aggregate / reduce the output and returns a finalized value
For the sake of this demo, we are currently operating on one node.



What is Next
For full cohesiveness, Herodotus's verifier plays a large role in verifying our verifiable inference computations. This is the next step for rounding out the process and fully integrating with GizaTech's mature tech stack. In the meantime, we aim to implement a recursive proof system so that we can compile a singular proof to represent all the work done in this complex system. 

Although we started with MapReduce, we envision this zkSPARK system scaling big data analysis and zkML/algorithmic computation:  

Feb 16: MVP Distributed System, MapReduce Cairo Algorithm, Matrix Multiplication Support
Q2 2024: Recursive Proof, Verifier settling on Starknet, IO File Execution, more algorithm support
Q4 2024: SDK interface, Generalized algorithm computation support, dev tooling for relayers and indexers
Q2 2025: SPARK Engine, Decentralized Worker Node Program, Decentralized L2 Network, zkML models on-chain execution


🚀 zkML is a pioneering field of innovation that opens doors to actions such as claiming IP on AI agents, ensuring honest AI computations, and decentralizing computing protocols. But for these possibility, we believe that we need to lay the groundwork for scaling ZK computations horizontally through STARKScale.

## Docker

build docker with

```sh
docker build -t zkscales-server .
```

run docker with

```sh
docker run -it zkscales-server
```

In the docker terminal
```sh
cd server/
bash test-matmult.sh
```
