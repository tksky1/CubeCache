package example

const luaScript = `
function initDBConnection()
    -- 连接MySQL数据库
    local db, err = sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
    if err ~= nil then
        error("无法连接到数据库: " .. err.Error())
    end

    -- 设置数据库连接池的最大空闲连接数
    db:SetMaxIdleConns(10)

    -- 返回数据库连接对象
    return db
end

-- Lua脚本中的函数用于执行查询操作
function getter(db, query)
    -- 执行查询语句
    local rows, err = db:Query(query)
    if err ~= nil then
        error("查询执行失败: " .. err.Error())
    end

    -- 遍历查询结果
    for rows:Next() do
        local id = rows:Scan()
        print("ID:", id)
    end
end
`
