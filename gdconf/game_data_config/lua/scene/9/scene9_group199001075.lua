-- 基础信息
local base_info = {
	group_id = 199001075
}

-- DEFS_MISCS
local shadowConfigIDList = {75002,75003}

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
	{ config_id = 75001, gadget_id = 70500033, pos = { x = -345.783, y = 232.465, z = -648.707 }, rot = { x = 0.000, y = 47.210, z = 0.000 }, level = 20, arguments = { 43 }, area_id = 400 },
	{ config_id = 75002, gadget_id = 70500048, pos = { x = -350.084, y = 230.210, z = -644.887 }, rot = { x = 0.000, y = 268.917, z = 0.000 }, level = 20, area_id = 400 },
	{ config_id = 75003, gadget_id = 70500052, pos = { x = -341.734, y = 230.876, z = -647.751 }, rot = { x = 0.000, y = 289.255, z = 0.000 }, level = 20, area_id = 400 }
}

-- 区域
regions = {
	{ config_id = 75004, shape = RegionShape.SPHERE, radius = 12, pos = { x = -345.017, y = 231.454, z = -646.102 }, area_id = 400 }
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
		gadgets = { 75001 },
		regions = { 75004 },
		triggers = { },
		rand_weight = 100
	}
}

--================================================================
-- 
-- 触发器
-- 
--================================================================

require "V2_8/EchoConch"