// Code generated by "juicecli impl --type Interface --output interface_impl.go"; DO NOT EDIT.

package testcase

import (
	"context"
	"database/sql"
	"github.com/eatmoreapple/juice"
)

type InterfaceImpl struct{}

func (i InterfaceImpl) GetUserByID(ctx context.Context, id int64) (result0 User, result1 error) {
	manager := juice.ManagerFromContext(ctx)
	var iface Interface = i
	executor := juice.NewGenericManager[User](manager).Object(iface.GetUserByID)
	ret, err := executor.QueryContext(ctx, id)
	return ret, err
}

func (i InterfaceImpl) CreateUser(ctx context.Context, u map[string]*User) (result0 error) {
	manager := juice.ManagerFromContext(ctx)
	var iface Interface = i
	executor := manager.Object(iface.CreateUser)
	_, err := executor.ExecContext(ctx, u)
	return err
}

func (i InterfaceImpl) DeleteUserByID(ctx context.Context, id int64, name string) (result0 sql.Result, result1 error) {
	manager := juice.ManagerFromContext(ctx)
	var iface Interface = i
	executor := manager.Object(iface.DeleteUserByID)
	return executor.ExecContext(ctx, juice.H{"id": id, "name": name})
}

// NewInterface returns a new Interface.
func NewInterface() Interface {
	return &InterfaceImpl{}
}
