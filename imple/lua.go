package imple

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

// RegisterLuaFunc register the lua script to LState
func RegisterLuaFunc(L *lua.LState, luaScript string) error {
	if err := L.DoString(luaScript); err != nil {
		return fmt.Errorf("register lua func error: %v", err)
	}
	return nil
}

func ExecuteSetterLua(L *lua.LState, key string, value []byte) error {
	luaExecuteSetter := L.GetGlobal("setter")
	if luaExecuteSetter == nil || L.GetTop() < 1 {
		return fmt.Errorf("get lua setter func fail for key %s", key)
	}
	if err := L.CallByParam(lua.P{
		Fn:      luaExecuteSetter,
		NRet:    1,
		Protect: true,
	},
		lua.LString(key),
		lua.LString(value),
	); err != nil {
		return fmt.Errorf("setter lua for key %s run fail: %v", key, err)
	}
	return nil
}

// ExecuteGetterLua execute query lua when key not in cache
func ExecuteGetterLua(L *lua.LState, funcName string, key string) ([]byte, error) {
	// 获取查询函数
	luaExecuteQuery := L.GetGlobal(funcName)
	if luaExecuteQuery == nil || L.GetTop() < 1 {
		return nil, fmt.Errorf("get lua getter func fail for key %s", key)
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
