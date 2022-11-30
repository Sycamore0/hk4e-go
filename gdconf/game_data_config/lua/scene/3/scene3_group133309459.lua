-- 基础信息
local base_info = {
	group_id = 133309459
}

--================================================================
-- 
-- 配置
-- 
--================================================================

-- 怪物
monsters = {
}

-- NPC
npcs = {
}

-- 装置
gadgets = {
	{ config_id = 459007, gadget_id = 70500000, pos = { x = -2104.885, y = -34.999, z = 5771.089 }, rot = { x = 0.000, y = 359.634, z = 0.000 }, level = 32, point_type = 1003, area_id = 27 },
	{ config_id = 459008, gadget_id = 70500000, pos = { x = -2108.451, y = -35.626, z = 5769.099 }, rot = { x = 335.506, y = 112.638, z = 343.857 }, level = 32, point_type = 1003, area_id = 27 },
	{ config_id = 459009, gadget_id = 70500000, pos = { x = -2113.593, y = -37.768, z = 5766.968 }, rot = { x = 11.532, y = 223.398, z = 340.308 }, level = 32, point_type = 1003, area_id = 27 },
	{ config_id = 459013, gadget_id = 70500000, pos = { x = -2124.095, y = -44.026, z = 5943.997 }, rot = { x = 335.517, y = 38.934, z = 351.644 }, level = 32, point_type = 1003, area_id = 27 },
	{ config_id = 459014, gadget_id = 70500000, pos = { x = -2126.604, y = -43.065, z = 5946.564 }, rot = { x = 10.282, y = 139.203, z = 327.391 }, level = 32, point_type = 1003, area_id = 27 }
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
		monsters = { },
		gadgets = { 459007, 459008, 459009, 459013, 459014 },
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