# Simple Bank

This repository contains the codes of the [Backend master class](https://bit.ly/backendmaster) course by [TECH SCHOOL](https://dev.to/techschoolguru).

![Backend master class](backend-master.png)

In this backend master class, we’re going to learn everything about how to design, develop, and deploy a complete backend system from scratch using Golang, PostgreSQL, and Docker.

## Course videos

- Lecture #1: [Design DB schema and generate SQL code with dbdiagram.io](https://www.youtube.com/watch?v=rx6CPDK_5mU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=1)
- Lecture #2: [Install & use Docker + Postgres + TablePlus to create DB schema](https://www.youtube.com/watch?v=Q9ipbLeqmQo&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=2)
- Lecture #3: [How to write & run database migration in Golang](https://www.youtube.com/watch?v=0CYkrGIJkpw&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=3)
- Lecture #4: [Generate CRUD Golang code from SQL | Compare db/sql, gorm, sqlx & sqlc](https://www.youtube.com/watch?v=prh0hTyI1sU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=4)
- Lecture #5: [Write unit tests for database CRUD with random data in Golang](https://www.youtube.com/watch?v=phHDfOHB2PU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=5)
- Lecture #6: [A clean way to implement database transaction in Golang](https://www.youtube.com/watch?v=gBh__1eFwVI&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=6)
- Lecture #7: [DB transaction lock & How to handle deadlock in Golang](https://www.youtube.com/watch?v=G2aggv_3Bbg&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=7)
- Lecture #8: [How to avoid deadlock in DB transaction? Queries order matters!](https://www.youtube.com/watch?v=qn3-5wdOfoA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=8)
- Lecture #9: [Deeply understand transaction isolation levels & read phenomena in MySQL & PostgreSQL](https://www.youtube.com/watch?v=4EajrPgJAk0&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=9)
- Lecture #10: [Setup Github Actions for Golang + Postgres to run automated tests](https://www.youtube.com/watch?v=3mzQRJY1GVE&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=10)
- Lecture #11: [Implement RESTful HTTP API in Go using Gin](https://www.youtube.com/watch?v=n_Y_YisgqTw&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=11)
- Lecture #12: [Load config from file & environment variables in Go with Viper](https://www.youtube.com/watch?v=n5p8HkO6bnE&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=12)
- Lecture #13: [Mock DB for testing HTTP API in Go and achieve 100% coverage](https://www.youtube.com/watch?v=rL0aeMutoJ0&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=13)
- Lecture #14: [Implement transfer money API with a custom params validator](https://www.youtube.com/watch?v=5q_wsashJZA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=14)
- Lecture #15: [Add users table with unique & foreign key constraints in PostgreSQL](https://www.youtube.com/watch?v=D4VtNC3vQUs&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=15)
- Lecture #16: [How to handle DB errors in Golang correctly](https://www.youtube.com/watch?v=mJ8b5GcvoxQ&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=16)
- Lecture #17: [How to securely store passwords? Hash password in Go with Bcrypt!](https://www.youtube.com/watch?v=B3xnJI2lHmc&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=17)
- Lecture #18: [How to write stronger unit tests with a custom gomock matcher](https://www.youtube.com/watch?v=DuzBE0jKOgE&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=18)
- Lecture #19: [Why PASETO is better than JWT for token-based authentication?](https://www.youtube.com/watch?v=nBGx-q52KAY&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=19)
- Lecture #20: [How to create and verify JWT & PASETO token in Golang](https://www.youtube.com/watch?v=Oi4FHDGILuY&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=20)
- Lecture #21: [Implement login user API that returns PASETO or JWT access token in Go](https://www.youtube.com/watch?v=lnHbZ9GOGAs&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=21)
- Lecture #22: [Implement authentication middleware and authorization rules in Golang using Gin](https://www.youtube.com/watch?v=Pw8fVBRS4jA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=22)
- Lecture #23: [Build a minimal Golang Docker image with a multistage Dockerfile](https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=23)
- Lecture #24: [How to use docker network to connect 2 stand-alone containers](https://www.youtube.com/watch?v=VcFnqQarpjI&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=24)
- Lecture #25: [How to write docker-compose file and control service start-up orders with wait-for.sh](https://www.youtube.com/watch?v=jf6sQsz0M1M&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=25)
- Lecture #26: [How to create a free tier AWS account](https://www.youtube.com/watch?v=4UqN1P8pIkM&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=26)
- Lecture #27: [Auto build & push docker image to AWS ECR with Github Actions](https://www.youtube.com/watch?v=3M4MPmSWt9E&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=27)
- Lecture #28: [How to create a production DB on AWS RDS](https://www.youtube.com/watch?v=0EaG3T4Q5fQ&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=28)
- Lecture #29: [Store & retrieve production secrets with AWS secrets manager](https://www.youtube.com/watch?v=3i1mQ_Ye8jE&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=29)

## Simple bank service

The service that we’re going to build is a simple bank. It will provide APIs for the frontend to do following things:

1. Create and manage bank accounts, which are composed of owner’s name, balance, and currency.
2. Record all balance changes to each of the account. So every time some money is added to or subtracted from the account, an account entry record will be created.
3. Perform a money transfer between 2 accounts. This should happen within a transaction, so that either both accounts’ balance are updated successfully or none of them are.

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

- [Gomock](https://github.com/golang/mock)

    ``` bash
    go install github.com/golang/mock/mockgen@v1.6.0
    ```

### Setup infrastructure

- Create the bank-network

    ``` bash
    make network
    ```

- Start postgres container:

    ```bash
    make postgres
    ```

- Create simple_bank database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```

### How to generate code

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

- Create a new db migration:

    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```

## Deploy to kubernetes cluster

- [Install nginx ingress controller](https://kubernetes.github.io/ingress-nginx/deploy/#aws):

    ```bash
    kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.48.1/deploy/static/provider/aws/deploy.yaml
    ```

- [Install cert-manager](https://cert-manager.io/docs/installation/kubernetes/):

    ```bash
    kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.4.0/cert-manager.yaml
    ```
