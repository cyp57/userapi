# User API

## Description

API for a basic membership registration system, integrating FusionAuth for authentication (login, password handler) and authorization (role assignment). The main functions include:

- `/login`
- `/signup`
- `/refresh/token`
- `/forgot/password`
- `/change/password`
- CRUD operations for user information

## FusionAuth

[FusionAuth](https://fusionauth.io/) is a customer authentication and authorization platform that provides developers with a customizable, scalable solution for managing user access. For more information, refer to [FusionAuth Documentation](https://fusionauth.io/docs/quickstarts/).

## Dependencies

- Programming Language: Go
- API Framework: Gin Framework
- Database: MongoDB
- Library: mongo-driver
- Containerization: Docker
- API Integration: FusionAuth (with PostgreSQL database)

Other Dependencies:

- github.com/FusionAuth/go-client
- github.com/coolbed/mgo-oid
- github.com/gin-gonic/gin
- github.com/go-ini/ini
- github.com/joho/godotenv
- github.com/sirupsen/logrus
- github.com/spf13/viper
- go.mongodb.org/mongo-driver
- golang.org/x/crypto

## Project Structure

## /app
   - **/api/v1**
     - api handler

   - **/controller/v1**
     - business logic

   - **/middleware**
     - middleware handler

   - **/response**
     - response handler

## /assets
   - **/log**
     - contain log request and response that write by /pkg/logger

## /cnst
   - contain constant

## /config
   - contain config env , yaml

## /model
   - contain struct

## /pkg
   - other packages used in this api

   - **/fusionauth**
     - fusionauth-client

   - **/logger**
     - write log request, response to /assets/log

   - **/logrus**
     - custom logrus for print by specify level (/config/.env, logLevel=Debug , Info ,Warning, Error, Fatal)

   - **/mongodb**
     - code related to mongodb

## /route
   - route for api

## /setting
   - load .env

## /test
   - testing

## /utils
   - utilities code

main.go

## Authors

[chanyapatshell@gmail.com](mailto:chanyapatshell@gmail.com)

Fusionauth : https://fusionauth.io/

## Version History

- 1.0.0
  - Initial Release

## License

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)