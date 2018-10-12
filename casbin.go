package yugle

import (
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
)

var enforcer *casbin.Enforcer

func GetEnforcer() *casbin.Enforcer {
	if enforcer == nil {
		adapter := gormadapter.NewAdapter("mysql", "root:sql123@tcp(127.0.0.1:3306)/")
		enforcer = casbin.NewEnforcer("./rbac_model.conf", adapter)
	}
	return enforcer
}
