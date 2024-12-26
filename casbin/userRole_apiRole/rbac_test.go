package userRole_apiRole

import (
	"fmt"
	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
	defaultrolemanager "github.com/casbin/casbin/v3/rbac/default-role-manager"
	"github.com/casbin/casbin/v3/util"
	"testing"
)

const rbac_models = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && regexMatch(r.act, p.act)
`

var requests = [][]interface{}{
	{"user", "/system/aaa", "read"},     // false
	{"user", "/servers/", "read"},       // true
	{"user", "/lists/devices", "read"},  // true
	{"user", "/lists/devices", "write"}, // true
	{"admin", "/system/bbb", "write"},   // true
	{"admin", "/status/", "write"},      // true
	{"admin", "/status/list", "read"},   // true
	{"audit", "/status/", "read"},       // true
	{"audit", "/system/log", "read"},    //true
	{"audit", "/lists/", "read"},        //true
	{"audit", "/users/", "write"},       //false
}

func TestAcl(t *testing.T) {
	m, _ := model.NewModelFromString(rbac_models)
	e, _ := casbin.NewEnforcer(m, false)

	groupPolicies := [][]string{
		/*
			用户和用户角色的映射， 字段: 用户---用户角色
			admin 同属于administrators
			user 同数据user 用户角色
		*/
		{"admin", "administrators"}, // admin 属于administrators 角色
		{"user", "users"},           //user 属于 users角色
		{"audit", "audit"},          // audit审计员属于audit角色
	}
	g2Polices := [][]string{
		/*
			普通资源角色， 字段： 资源---资源角色
		*/
		{"/status/*", "commons"},
		{"/servers/*", "commons"},
		{"/lists/*", "commons"},
		{"/system/log", "audit"}, //查看日志权限的资源角色

		/*
			管理员角色
		*/
		{"/*", "all"}, //管理员角色可以访问所有路径
	}
	policies := [][]string{
		/*
			角色和资源角色的映射,  用户角色---资源角色---权限
			administrators用户角色 可以访问 all资源角色
			users普通用户角色  可以访问 commons普通资源角色
		*/
		{"administrators", "all", ".*"},
		{"users", "commons", ".*"},
		{"audit", "audit", "read"}, // 审计员角色可以查看日志
		{"audit", "commons", "read"},
	}

	_, _ = e.AddPolicies(policies)
	_, _ = e.AddGroupingPolicies(groupPolicies)
	_, _ = e.AddNamedGroupingPolicies("g2", g2Polices)

	e.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc(util.KeyMatch)
	fmt.Println(e.GetImplicitRolesForUser("audit"))

	for _, request := range requests {
		res, err := e.Enforce(request...)
		if err != nil {
			t.Error(err)
		}
		if res == true {
			t.Logf("%s can %s %s.\n", request[0], request[2], request[1])
		} else {

			t.Logf("%s can't %s %s.\n", request[0], request[2], request[1])
		}
	}
}
