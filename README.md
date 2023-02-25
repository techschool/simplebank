# Simple Bank

This repository contains the codes of the [Backend Master Class](https://bit.ly/backendmaster) course by [TECH SCHOOL](https://bit.ly/m/techschool).

![Backend master class](backend-master.png)

You can also find it on Udemy at [this link](https://bit.ly/backendudemy).

And don't hesitate to [join Tech School's Discord group](https://bit.ly/techschooldc) to chat directly with me and other students.

In this course, you will learn step-by-step how to design, develop and deploy a backend web service from scratch. I believe the best way to learn programming is to build a real application. Therefore, throughout the course, you will learn how to build a backend web service for a simple bank. It will provide APIs for the frontend to do the following things:

- Create and manage bank accounts.
- Record all balance changes to each of the accounts.
- Perform a money transfer between 2 accounts.

The programming language we will use to develop the service is Golang, but the course is not just about coding in Go. You will learn a lot of different topics regarding backend web development. They are presented in 5 sections:

1. In the 1st section, you will learn deeply about how to design the database, generate codes to talk to the DB in a consistent and reliable way using transactions, understand the DB isolation levels, and how to use it correctly in production. Besides the database, you will also learn how to use docker for local development, how to use Git to manage your codes, and how to use GitHub Action to run unit tests automatically.

2. In the 2nd section, you will learn how to build a set of RESTful HTTP APIs using Gin - one of the most popular Golang frameworks for building web services. This includes everything from loading app configs, mocking DB for more robust unit tests, handling errors, authenticating users, and securing the APIs with JWT and PASETO access tokens.  

3. In the 3rd section, you will learn how to build your app with Docker and deploy it to a production Kubernetes cluster on AWS. The lectures are very detailed with a step-by-step guide, from how to build a minimal docker image, set up a free-tier AWS account, create a production database, store and retrieve production secrets, create a Kubernetes cluster with EKS, use GitHub Action to automatically build and deploy the image to the EKS cluster, buy a domain name and route the traffics to the service, secure the connection with HTTPS and auto-renew TLS certificate from Let's Encrypt.

4. In the 4th section, we will discuss several advanced backend topics such as managing user sessions, building gRPC APIs, using gRPC gateway to serve both gRPC and HTTP requests at the same time, embedding Swagger documentation as part of the backend service, partially updating a record using optional parameters, and writing structured logger HTTP middlewares and gRPC interceptors.

5. Then the 5th section will introduce you to asynchronous processing in Golang using background workers and Redis as its message queue, and how to gracefully shut down the server to protect your processing resources. As this part is still a work in progress, we will keep making and uploading new videos about new topics in the future, such as: sending emails, gracefully shutting down servers, CORS, bulk inserts, etc. So please come back here to check them out from time to time.

This course is designed with a lot of details, so that everyone, even with very little programming experience can understand and do it by themselves. I strongly believe that after the course, you would be able to work much more confidently and effectively on your projects.

## Course videos

### Section 1: Working with database [Postgres]

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

### Section 2: Building RESTful HTTP JSON API [Gin]

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

### Section 3: Deploying the application to production [Kubernetes + AWS]

- Lecture #23: [Build a minimal Golang Docker image with a multistage Dockerfile](https://www.youtube.com/watch?v=p1dwLKAxUxA&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=23)
- Lecture #24: [How to use docker network to connect 2 stand-alone containers](https://www.youtube.com/watch?v=VcFnqQarpjI&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=24)
- Lecture #25: [How to write docker-compose file and control service start-up orders with wait-for.sh](https://www.youtube.com/watch?v=jf6sQsz0M1M&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=25)
- Lecture #26: [How to create a free tier AWS account](https://www.youtube.com/watch?v=4UqN1P8pIkM&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=26)
- Lecture #27: [Auto build & push docker image to AWS ECR with Github Actions](https://www.youtube.com/watch?v=3M4MPmSWt9E&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=27)
- Lecture #28: [How to create a production DB on AWS RDS](https://www.youtube.com/watch?v=0EaG3T4Q5fQ&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=28)
- Lecture #29: [Store & retrieve production secrets with AWS secrets manager](https://www.youtube.com/watch?v=3i1mQ_Ye8jE&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=29)
- Lecture #30: [Kubernetes architecture & How to create an EKS cluster on AWS](https://www.youtube.com/watch?v=TxnCMhYhqRU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=30)
- Lecture #31: [How to use kubectl & k9s to connect to a kubernetes cluster on AWS EKS](https://www.youtube.com/watch?v=hwMevai3_wQ&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=31)
- Lecture #32: [How to deploy a web app to Kubernetes cluster on AWS EKS](https://www.youtube.com/watch?v=PH-Mcd0Rs1w&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=32)
- Lecture #33: [Register a domain name & set up A-record using Route53](https://www.youtube.com/watch?v=-JF2ukmW3i8&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=33)
- Lecture #34: [How to use Ingress to route traffics to different services in Kubernetes](https://www.youtube.com/watch?v=lBrqP6FkNsU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=34)
- Lecture #35: [Automatic issue TLS certificates in Kubernetes with Let's Encrypt](https://www.youtube.com/watch?v=nU4FTjrgSKI&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=35)
- Lecture #36: [Automatic deploy to Kubernetes with Github Action](https://www.youtube.com/watch?v=GVY-zze0V_U&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=36)

### Section 4: Advanced Backend Topics [Sessions + gRPC]

- Lecture #37: [How to manage user session with refresh token - Golang](https://www.youtube.com/watch?v=rT20ylRLm5U&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=37)
- Lecture #38: [Generate DB documentation page and schema SQL dump from DBML](https://www.youtube.com/watch?v=dGfVwsPr-IU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=38)
- Lecture #39: [Introduction to gRPC](https://www.youtube.com/watch?v=mRGnA3wPxMM&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=39)
- Lecture #40: [Define gRPC API and generate Go code with protobuf](https://www.youtube.com/watch?v=mVWgEmyAhvM&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=40)
- Lecture #41: [How to run a golang gRPC server and call its API](https://www.youtube.com/watch?v=BkfBJIS0_ro&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=41)
- Lecture #42: [Implement gRPC API to create and login users in Go](https://www.youtube.com/watch?v=7xiWqyZW9lE&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=42)
- Lecture #43: [Write code once, serve both gRPC & HTTP requests](https://www.youtube.com/watch?v=3FfDH3d0aHs&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=43)
- Lecture #44: [How to extract info from gRPC metadata](https://www.youtube.com/watch?v=Sno10WQ21Zs&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=44)
- Lecture #45: [Automatic generate & serve Swagger docs from Go server](https://www.youtube.com/watch?v=Uwkxxee7tvk&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=45)
- Lecture #46: [Embed static frontend files inside Golang backend server's binary](https://www.youtube.com/watch?v=xNgOIm86N5Q&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=46)
- Lecture #47: [Validate gRPC parameters and send human/machine friendly response](https://www.youtube.com/watch?v=CxZ9hMtmZtc&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=47)
- Lecture #48: [Run DB migrations directly inside Golang code](https://www.youtube.com/watch?v=TG43cMpaxlI&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=48)
- Lecture #49: [Partial update DB record with SQLC nullable parameters](https://www.youtube.com/watch?v=I2sbw1PzzW0&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=49)
- Lecture #50: [Build gRPC update API with optional parameters](https://www.youtube.com/watch?v=ygqSHIEc8sc&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=50)
- Lecture #51: [Add authorization to protect gRPC API](https://www.youtube.com/watch?v=_jqNs3d99ps&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=51)
- Lecture #52: [Write structured logs for gRPC APIs](https://www.youtube.com/watch?v=tTAxLGrDmPo&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=52)
- Lecture #53: [How to write HTTP logger middleware in Go](https://www.youtube.com/watch?v=Lbiz-PZNiU0&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=53)

### Section 5: Asynchronous processing with background workers [Asynq + Redis]

- Lecture #54: [Implement background worker in Go with Redis and Asynq](https://www.youtube.com/watch?v=XOXdYs8mKkI&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=54)
- Lecture #55: [Integrate async worker to Go web server](https://www.youtube.com/watch?v=eXYKGPEXocM&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=55)
- Lecture #56: [Send async tasks to Redis within a DB transaction](https://www.youtube.com/watch?v=ZfFxdPbgN88&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=56)
- Lecture #57: [How to handle errors and print logs for Go Asynq workers](https://www.youtube.com/watch?v=YgfmPIJRg2U&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=57)
- Lecture #58: [A bit of delay might be good for your async tasks](https://www.youtube.com/watch?v=ILNiZgseLUI&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=58)
- Lecture #59: [How to send emails in Go via Gmail](https://www.youtube.com/watch?v=L9TbZxpykLQ&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=59)
- Lecture #60: [How to skip test in Go and config test flag in vscode](https://www.youtube.com/watch?v=0UwZGM9iqTE&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=60)

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

- [DB Docs](https://dbdocs.io/docs)

    ```bash
    npm install -g dbdocs
    dbdocs login
    ```

- [DBML CLI](https://www.dbml.org/cli/#installation)

    ```bash
    npm install -g @dbml/cli
    dbml2sql --version
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

### Documentation

- Generate DB documentation:

    ```bash
    make db_docs
    ```

- Access the DB documentation at [this address](https://dbdocs.io/techschool.guru/simple_bank). Password: `secret`

### How to generate code

- Generate schema SQL file with DBML:

    ```bash
    make db_schema
    ```

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
    make new_migration name=<migration_name>
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
