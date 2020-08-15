# backend-bff API

## Description
### List API（API一覧）
1. development：クエリ練習用 API
2. production：本番環境用 graphQL API

## How to run
#### development API
**ローカルで実行**  
- 実行コマンド： `make local-dev-run`
  
ブラウザで http://localhost:9020 にアクセスする  
※ `go 1.13`の環境で実行可能  
**Use Docker**  
- ビルドコマンド： `make docker-dev-build`
- 実行コマンド： `make docker-dev-run`

#### production API
- 実行コマンド： 

**Use Docker**  
- ビルドコマンド： 
- 実行コマンド： 
## Detail Query
### development
#### List Query
|   Type   |     Name     |            Argument            |    Return    |
| :------: | :----------: | :----------------------------: |:-----------: |
| Query    | odd          | number:Int!                    | Judge!       |
| Query    | even         | number:Int!                    | Judge!      |
| Query    | city         | id:Int                         | CityData!    |
| Query    | allCity      |                                | [CityData!]! |
| Mutation | addArticle   | index:Int, article: String     | Article!     |
| Mutation | postUserData | emotion: Float!, city: String! | CityData!    |

#### Data Type
**CityData**
|    Name   |  Type   |
| :-------: | :-----: |
| name      | String! |
| latitude  | Float!  |
| longitude | Float!  |

**Judge**
|  Name  |   Type   |
| :----: | :------: |
| number | Int!     |
| judge  | Boolean! |

**Article**
|   Name  |   Type  |
| :-----: | :-----: |
| index   | Int!    |
| article | String! |

#### production
#### List Query
#### Data Type
