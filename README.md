## Juice SQL Mapper Framework For Golang

![Go Doc](https://pkg.go.dev/badge/github.com/eatmoreapple/juice)
![Go Report Card](https://goreportcard.com/badge/github.com/eatmoreapple/juice)
![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)

This is a SQL mapper framework for Golang. It is inspired by MyBatis.

Juice is a simple and lightweight framework. It is easy to use and easy to extend.

### Features

* Simple and lightweight
* Easy to use
* Easy to extend
* Support for multiple databases
* Dynamic SQL
* Result to entity mapping
* Generic type support
* Middleware support
* Todo support more

### support xml tags

* select
* insert
* update
* delete
* sql
* if
* where
* trim
* set
* foreach
* choose
* when
* otherwise
* include
* resultMap


### Condition Method
The condition method can be used with `if` or `when` tags.

For example:
```xml
<!--ids = []int{1,2,3}-->
<if test='len(ids) > 0 && substr("eatmoreapple", 0, 3) == "eat"'>
    your sql node here
</if>
```

It can register to the framework with your own condition method.

Here are some default condition methods.

* len: return the length of the given parameter
* strsub: return the substring of the given parameter
* join: join the given parameters with the given separator
* contains: return true if the given parameter contains the given element
* slice: return the slice of the given parameter

### Quick Start

#### Install

```bash
go get github.com/eatmoreapple/juice

go install github.com/eatmoreapple/juice/cmd/juice@latest
```

#### Example

```shell
touch config.xml
```

and write the following content into config.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<configuration>
    <environments default="prod">
        <environment id="prod">
            <dataSource>root:qwe123@tcp(localhost:3306)/database</dataSource>
            <driver>mysql</driver>
        </environment>
    </environments>


    <mappers>
        <mapper namespace="main.UserRepository">
            <select id="GetUserByID" paramName="id">
                select * from user where id = #{id} limit 1
            </select>
            <insert id="CreateUser">
                insert into user (name, age) values (#{name}, #{age})
            </insert>
            <delete id="DeleteUserByID" paramName="id">
                delete from user where id = #{id}
            </delete>
        </mapper>
    </mappers>
</configuration>
```


define your interface

```go
package main

import (
	"context"
	"database/sql"
	
	"github.com/eatmoreapple/juice"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate juice --type UserRepository --config config.xml --namespace main.UserRepository --output interface_impl.go
type UserRepository interface {
	// GetUserByID get user by id
	GetUserByID(ctx context.Context, id int64) (*User, error)
	// CreateUser create user
	CreateUser(ctx context.Context, user User) error
	// DeleteUserByID delete user by id
	DeleteUserByID(ctx context.Context, id int64) (sql.Result, error)
}

type User struct {
	Id   int64  `column:"id" param:"id"`
	Name string `column:"name" param:"name"`
	Age  int    `column:"age" param:"age"`
}
```

run `go generate` to generate the implementation of the interface

then you can see the interface implementation like this

```go
// Code generated by "juice --type UserRepository --config config.xml --namespace main.UserRepository --output interface_impl.go"; DO NOT EDIT.

package main

import (
	"context"
	"database/sql"
	"github.com/eatmoreapple/juice"
)

type UserRepositoryImpl struct{}

// GetUserByID get user by id
func (u UserRepositoryImpl) GetUserByID(ctx context.Context, id int64) (*User, error) {
	manager := juice.ManagerFromContext(ctx)
	var iface UserRepository = u
	executor := juice.NewGenericManager[*User](manager).Object(iface.GetUserByID)
	return executor.QueryContext(ctx, id)
}

// CreateUser create user
func (u UserRepositoryImpl) CreateUser(ctx context.Context, user User) error {
	manager := juice.ManagerFromContext(ctx)
	var iface UserRepository = u
	executor := manager.Object(iface.CreateUser)
	_, err := executor.ExecContext(ctx, user)
	return err
}

// DeleteUserByID delete user by id
func (u UserRepositoryImpl) DeleteUserByID(ctx context.Context, id int64) (sql.Result, error) {
	manager := juice.ManagerFromContext(ctx)
	var iface UserRepository = u
	executor := manager.Object(iface.DeleteUserByID)
	return executor.ExecContext(ctx, id)
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
```


how to use it

```go
package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eatmoreapple/juice"

	_ "github.com/go-sql-driver/mysql"
)

//go:generate juice --type UserRepository --config config.xml --namespace main.UserRepository --output interface_impl.go
type UserRepository interface {
	// GetUserByID get user by id
	GetUserByID(ctx context.Context, id int64) (*User, error)
	// CreateUser create user
	CreateUser(ctx context.Context, user User) error
	// DeleteUserByID delete user by id
	DeleteUserByID(ctx context.Context, id int64) (sql.Result, error)
}

type User struct {
	Id   int64  `column:"id" param:"id"`
	Name string `column:"name" param:"name"`
	Age  int    `column:"age" param:"age"`
}

var schema = `
CREATE TABLE IF NOT EXISTS user (
  id int(11) NOT NULL AUTO_INCREMENT,
  name varchar(255) COLLATE utf8mb4_bin NOT NULL,
  age int(11) NOT NULL,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
`

func main() {
	cfg, err := juice.NewXMLConfiguration("/Users/eatmoreapple/GolandProjects/juice/.example/config.xml")
	if err != nil {
		panic(err)
	}
	engine, err := juice.DefaultEngine(cfg)
	if err != nil {
		panic(err)
	}

	if _, err := engine.DB().Exec(schema); err != nil {
		panic(err)
	}

	ctx := juice.ContextWithManager(context.Background(), engine)

	repo := NewUserRepository()

	fmt.Println(repo.CreateUser(ctx, User{Name: "eatmoreapple", Age: 18}))
	fmt.Println(repo.GetUserByID(ctx, 1))
	fmt.Println(repo.DeleteUserByID(ctx, 1))
}

```

### Document

[Read the document](https://juice-doc.readthedocs.io/en/latest/index.html)

### License

Juice is licensed under the Apache License, Version 2.0. See LICENSE for the full license text.

### Contact

If you like this project, please give me a star. Thank you.
And If you have any questions, please contact me by WeChat: `eatmoreapple`.


### Invite the author to have a cup of coffee


<img width="210px"  src="https://github.com/eatmoreapple/eatMoreApple/blob/main/img/wechat_pay.jpg" align="left">
