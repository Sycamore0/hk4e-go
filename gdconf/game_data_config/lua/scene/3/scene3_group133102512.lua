-- 基础信息
local base_info = {
	group_id = 133102512
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
	{ config_id = 512001, monster_id = 21010201, pos = { x = 1431.304, y = 223.136, z = 93.631 }, rot = { x = 352.309, y = 71.869, z = 11.291 }, level = 16, drop_tag = "丘丘人", disableWander = true, pose_id = 9003, area_id = 5 },
	{ config_id = 512003, monster_id = 21010201, pos = { x = 1434.458, y = 223.684, z = 94.661 }, rot = { x = 10.304, y = 260.369, z = 6.902 }, level = 16, drop_tag = "丘丘人", disableWander = true, pose_id = 9003, area_id = 5 }
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 512004, gadget_id = 70310006, pos = { x = 1432.590, y = 223.334, z = 94.057 }, rot = { x = 355.229, y = 359.728, z = 6.528 }, level = 16, area_id = 5 },
	{ config_id = 512005, gadget_id = 70220013, pos = { x = 1437.420, y = 222.839, z = 91.438 }, rot = { x = 0.000, y = 45.576, z = 0.000 }, level = 16, area_id = 5 },
	{ config_id = 512007, gadget_id = 70220026, pos = { x = 1436.017, y = 222.759, z = 90.516 }, rot = { x = 0.000, y = 350.652, z = 0.000 }, level = 16, area_id = 5 }
}

-- 区域
regions = {
}

-- 触发器
triggers = {
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
	end_suite = 0,
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
		monsters = { 512001, 512003 },
		gadgets = { 512004, 512005, 512007 },
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