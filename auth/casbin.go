package auth

import (
	"github.com/casbin/casbin/v2"
)

func NewEnforcer() (*casbin.Enforcer, error) {
	enforcer, _ := casbin.NewEnforcer("/model.conf", "/policy.csv")
	return enforcer, nil
}
