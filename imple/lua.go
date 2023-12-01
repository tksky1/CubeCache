package imple

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func ExecuteQueryLua(L *lua.LState, luaScript string, key string) ([]byte, error) {
	// execute query lua when key not in cache
	if err := L.DoString(luaScript); err != nil {
		return nil, fmt.Errorf("lua execute error: %v", err)
	}

	// 获取查询函数
	luaExecuteQuery := L.GetGlobal("executeQuery")
	if luaExecuteQuery == nil || L.GetTop() < 1 {
		return nil, fmt.Errorf("get lua query func fail for key %s", key)
	}

	// 调用查询函数
	if err := L.CallByParam(lua.P{
		Fn:      luaExecuteQuery,
		NRet:    1,
		Protect: true,
	},
		lua.LString(key),
	); err != nil {
		return nil, fmt.Errorf("查询执行失败: %v", err)
	}

	// 获取查询结果
	luaQueryResult := L.Get(-1)
	L.Pop(1)

	// 转换为[]byte类型
	queryResult := lua.LVAsString(luaQueryResult)

	return []byte(queryResult), nil
}
