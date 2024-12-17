package rbac

import (
	"github.com/casbin/casbin/v3"
	"github.com/casbin/casbin/v3/model"
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
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && keyMatch(r.act, p.act)
`

var requests = [][]interface{}{
	{"alice", "data2", "read"},
	{"alice", "data2", "write"},
	{"bob", "data2", "read"},
	{"admin", "data10", "write"},
	{"admin", "data10", "read"},
}

func TestAcl(t *testing.T) {
	m, _ := model.NewModelFromString(rbac_models)
	e, _ := casbin.NewEnforcer(m, false)

	policies := [][]string{
		{"alice", "data1", "read"},
		{"bob", "data2", "write"},
		{"alice", "data2", "read"},
		{"bob", "data2", "write"},
		{"admin", "*", "*"},
		{"admin", "*", "write"},
		{"administrators", "*", "*"},
		{"users", "data", "*"},
	}
	_, _ = e.AddPolicies(policies)

	for _, request := range requests {
		res, _ := e.Enforce(request...)
		if res == true {
			t.Logf("%s can %s %s.\n", request[0], request[2], request[1])
		}
	}
}
