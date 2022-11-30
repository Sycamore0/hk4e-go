-- 基础信息
local base_info = {
	group_id = 133301367
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 367001, monster_id = 28010203, pos = { x = -826.764, y = 38.196, z = 3323.700 }, rot = { x = 0.000, y = 27.312, z = 0.000 }, level = 30, drop_tag = "采集动物", area_id = 23 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
}

-- 区域
regions = {
}

-- 触发器
triggers = {
	{ config_id = 1367002, name = "GROUP_LOAD_367002", event = EventType.EVENT_GROUP_LOAD, source = "", condition = "condition_EVENT_GROUP_LOAD_367002", action = "action_EVENT_GROUP_LOAD_367002", trigger_count = 0 }
}

-- 变量
variables = {
}

--================================================================
-- 
-- 初始化配置
-- 
--================================================================

-- 初始化时创建
init_config = {
	suite = 1,
	end_suite = 2,
	rand_suite = false
}

--================================================================
-- 
-- 小组配置
-- 
--================================================================

suites = {
	{
		-- suite_id = 1,
		-- description = ,
		monsters = { },
		gadgets = { },
		regions = { },
		triggers = { "GROUP_LOAD_367002" },
		rand_weight = 100
	},
	{
		-- suite_id = 2,
		-- description = ,
		monsters = { 367001 },
		gadgets = { },
		regions = { },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================

-- 触发条件
function condition_EVENT_GROUP_LOAD_367002(context, evt)
	if ScriptLib.GetDeathZoneStatus(context,0) == 0 then
	
	return true
	
	end
	
	
	return false
end

-- 触发操作
function action_EVENT_GROUP_LOAD_367002(context, evt)
		-- 重新生成指定group，指定suite
		if 0 ~= ScriptLib.RefreshGroup(context, { group_id = 133301367, suite = 2 }) then
	    ScriptLib.PrintContextLog(context, "@@ LUA_WARNING : refresh_group_to_suite")
			return -1
		end
	
	return 0
end

require "V3_0/DeathFieldStandard"