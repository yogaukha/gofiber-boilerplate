# Go Boilerplate using GoFiber Framework

## Description
Golang RESTFul Application Boilerplate is built with GoFiber Framework that implementing clean architecture.
This boilerplate is inspired by @bxcodec's [Go Clean Architecture V3](https://github.com/bxcodec/go-clean-arch/tree/v3).

## Used Tools
GoFiber Boilerplate was made by using several tools, such as:
- GoFiber itself, make this app is served like ExpressJS
- Viper, to set the configuration
- gORM, to communicate with DB (I use postgreSQL in this repo)
- Morkid Paginate, to set pagination
- Air, to hot reload the code
- Other libraries that listed in [go.mod](https://github.com/yogaukha/gofiber-boilerplate/blob/main/go.mod)

## Directory Structure
1. app
   - All of modules shipped in here. These modules should rely on a file in Domain directory (Abstraction). Each modules contains 3 package, such as: Delivery, Usecase and Repository.
   - Delivery is controller-like in MVC Pattern. It has responsibility to handle all of requests.
   - Usecase is acted like service. It would communicate with Delivery and Repository Layer. All logics shipped in here.
   - Repository has responsibility to communicate with database, such as: start transaction, inserting row(s), update data, (soft) delete data, selecting data, etc.
3. configs
   - All of configurations shipped in here, DB Config & App Config
4. domain
   - Domain is same as Model in MVC Pattern. All of models rely here
5. internal
   - All of helpers file shipped in here
6. routes
   - Routing are declared in this directory
7. server
   - This directory contains file that will set you up a restful API server
