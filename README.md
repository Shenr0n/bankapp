# Banking Backend

•	REST APIs for admin and user operations, including user and account management, and seamless transactions, using Go, Gin, PostgreSQL, sqlc.

• Authentication middleware to authorize API requests, with PASETO, and exhaustive unit tests.

• CI/CD pipelines for automated testing, database migration, containerization, and deployment, utilizing Github Actions, Docker, Kubernetes, Amazon Web Services (ECR, RDS, EKS and IAM).


## Tech-stack (Everything used in this project)
* `Go` - [Website](https://go.dev/)
* `Gin` - HTTP web framework written in Go  [Github](https://github.com/gin-gonic/gin)
`PostgreSQL` - Open-source relational database [PostgreSQL Website](https://www.postgresql.org/) [Docker Image Website](https://hub.docker.com/_/postgres)
* `sqlc` - SQL to type-safe code [Website](https://sqlc.dev/)
* `PASETO` - JOSE (JWT, JWE, JWS) without any of the design deficits in JOSE standards. [Website](https://paseto.io/)
* `Docker` - For containerization [Docker Website](https://www.docker.com/)
* `Kubernetes` - Container orchestration [Kubernetes Website](https://kubernetes.io/docs/home/)
* `AWS` - Used RDS, EKS, ECR, IAM [Website](https://aws.amazon.com/)
* `Viper` - Go configurations [Github](https://github.com/spf13/viper)
* `lib/pq` - Pure Go Postgres driver for database/sql [Website](https://github.com/lib/pq)
* `jwt-go` - JWT in Go [Github](github.com/dgrijalva/jwt-go)
* `TablePlus` - Database management with GUI [Website](https://tableplus.com/)
* `Postman` - Testing APIs [Website](https://www.postman.com/)
* `kubectl` - [Website](https://kubernetes.io/docs/tasks/tools/install-kubectl-macos/)
* `Nginx Ingress Controller` - [Website](https://docs.nginx.com/nginx-ingress-controller/)

### Running the project 
* Install Go, Docker Desktop, sqlc, TablePlus, Postman and other required tools and software mentioned.
* Make changes to the Makefile, Dockerfile, yml/yaml files for different settings for containerization, deployment, testing, AWS configurations.
* For running and testing on localhost, setup the postgreSQL docker image, `go mod tidy` and run the Makefile commands.
* The database can be viewed in TablePlus and tested on Postman.

---
**NOTE**
For deploying with Github Actions, AWS configurations have to be setup and the deploy.yml in workflows has to be moved into the .github/workflows directory.
For EKS operations, configurations can be managed with kubectl.
This has been disabled here due to AWS service costs of EKS and VPC.

---