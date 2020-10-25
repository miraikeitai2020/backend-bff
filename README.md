# backend-bff

<p align="left">
  <a href="https://github.com/actions/setup-node/actions?query=workflow%3Abuild-test"><img alt="build-test status" src="https://github.com/miraikeitai2020/backend-bff/workflows/Docker Image CI/badge.svg"></a>
</p>

<img src="./img/logo.png" width="600" alt="architecture" />

The `backend-bff` API is one of the services of [mirai-cluster](https://github.com/miraikeitai2020/mirai-cluster) developed by Future Mobile Phone Project 2020.  


## Description
### Endpoints
|Method|URL|Description|
|:-:|:-:|:-|
|GET|`/`||
|POST|`/query`|resolve client query handler|

### How to run
#### Build & Run
**◎Local**  
Advance preparation command: `make`  
Create files  
- `pkg/bff/bff.go`
- `pkg/server/model/model.go`
- `private.key`

Run API command: `make run`  

**◎Docker**  
Build command: `make docker-build`  
Run command: `make docker-run`  

#### Access
Access the link below from your browser.  
`http://localhost:9020`

## Other
- [Commit Rules](./docs/commit_rules.md)
- [Branch Rules](./docs/branch_rules.md)
- [Issue Rules](./docs/issue_rules.md)
