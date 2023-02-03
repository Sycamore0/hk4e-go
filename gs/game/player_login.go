package game

import (
	"time"

	"hk4e/common/constant"
	"hk4e/common/mq"
	"hk4e/gdconf"
	"hk4e/gs/model"
	"hk4e/pkg/logger"
	"hk4e/pkg/reflection"
	"hk4e/protocol/cmd"
	"hk4e/protocol/proto"

	pb "google.golang.org/protobuf/proto"
)

func (g *GameManager) PlayerLoginReq(userId uint32, clientSeq uint32, gateAppId string, payloadMsg pb.Message) {
	logger.Info("user login req, uid: %v, gateAppId: %v", userId, gateAppId)
	req := payloadMsg.(*proto.PlayerLoginReq)
	logger.Debug("login data: %v", req)
	g.OnLogin(userId, clientSeq, gateAppId)
}

func (g *GameManager) SetPlayerBornDataReq(userId uint32, clientSeq uint32, gateAppId string, payloadMsg pb.Message) {
	logger.Info("user reg req, uid: %v, gateAppId: %v", userId, gateAppId)
	req := payloadMsg.(*proto.SetPlayerBornDataReq)
	logger.Debug("reg data: %v", req)
	g.OnReg(userId, clientSeq, gateAppId, req)
}

func (g *GameManager) OnLogin(userId uint32, clientSeq uint32, gateAppId string) {
	logger.Info("user login, uid: %v", userId)
	player, asyncWait := USER_MANAGER.OnlineUser(userId, clientSeq, gateAppId)
	if !asyncWait {
		g.OnLoginOk(userId, player, clientSeq, gateAppId)
	}
}

func (g *GameManager) OnLoginOk(userId uint32, player *model.Player, clientSeq uint32, gateAppId string) {
	if player == nil {
		g.SendMsgToGate(cmd.DoSetPlayerBornDataNotify, userId, clientSeq, gateAppId, new(proto.DoSetPlayerBornDataNotify))
		return
	}
	player.OnlineTime = uint32(time.Now().UnixMilli())
	player.Online = true

	player.GateAppId = gateAppId

	// 初始化
	player.InitAll()

	// 确保玩家位置安全
	player.Pos.X = player.SafePos.X
	player.Pos.Y = player.SafePos.Y
	player.Pos.Z = player.SafePos.Z

	if player.SceneId > 100 {
		player.SceneId = 3
		player.Pos = &model.Vector{X: 2747, Y: 194, Z: -1719}
		player.Rot = &model.Vector{X: 0, Y: 307, Z: 0}
	}

	player.CombatInvokeHandler = model.NewInvokeHandler[proto.CombatInvokeEntry]()
	player.AbilityInvokeHandler = model.NewInvokeHandler[proto.AbilityInvokeEntry]()

	if userId < 100000000 {
		return
	}

	g.LoginNotify(userId, player, clientSeq)

	MESSAGE_QUEUE.SendToAll(&mq.NetMsg{
		MsgType: mq.MsgTypeServer,
		EventId: mq.ServerUserOnlineStateChangeNotify,
		ServerMsg: &mq.ServerMsg{
			UserId:   userId,
			IsOnline: true,
		},
	})

	TICK_MANAGER.CreateUserGlobalTick(userId)
	TICK_MANAGER.CreateUserTimer(userId, UserTimerActionTest, 100)
}

func (g *GameManager) OnReg(userId uint32, clientSeq uint32, gateAppId string, payloadMsg pb.Message) {
	logger.Debug("user reg, uid: %v", userId)
	req := payloadMsg.(*proto.SetPlayerBornDataReq)
	logger.Debug("avatar id: %v, nickname: %v", req.AvatarId, req.NickName)

	exist, asyncWait := USER_MANAGER.CheckUserExistOnReg(userId, req, clientSeq, gateAppId)
	if !asyncWait {
		g.OnRegOk(exist, req, userId, clientSeq, gateAppId)
	}
}

func (g *GameManager) OnRegOk(exist bool, req *proto.SetPlayerBornDataReq, userId uint32, clientSeq uint32, gateAppId string) {
	if exist {
		logger.Error("recv reg req, but user is already exist, userId: %v", userId)
		return
	}

	nickName := req.NickName
	mainCharAvatarId := req.GetAvatarId()
	if mainCharAvatarId != 10000005 && mainCharAvatarId != 10000007 {
		logger.Error("invalid main char avatar id: %v", mainCharAvatarId)
		return
	}

	player := g.CreatePlayer(userId, nickName, mainCharAvatarId)
	if player == nil {
		logger.Error("player is nil, uid: %v", userId)
		return
	}
	USER_MANAGER.AddUser(player)

	g.SendMsgToGate(cmd.SetPlayerBornDataRsp, userId, clientSeq, gateAppId, new(proto.SetPlayerBornDataRsp))
	g.OnLogin(userId, clientSeq, gateAppId)
}

func (g *GameManager) OnUserOffline(userId uint32, changeGsInfo *ChangeGsInfo) {
	logger.Info("user offline, uid: %v", userId)
	player := USER_MANAGER.GetOnlineUser(userId)
	if player == nil {
		logger.Error("player is nil, userId: %v", userId)
		return
	}
	TICK_MANAGER.DestroyUserGlobalTick(userId)
	world := WORLD_MANAGER.GetWorldByID(player.WorldId)
	if world != nil {
		g.UserWorldRemovePlayer(world, player)
	}
	player.OfflineTime = uint32(time.Now().Unix())
	player.Online = false
	player.TotalOnlineTime += uint32(time.Now().UnixMilli()) - player.OnlineTime
	USER_MANAGER.OfflineUser(player, changeGsInfo)
}

func (g *GameManager) LoginNotify(userId uint32, player *model.Player, clientSeq uint32) {
	g.SendMsg(cmd.PlayerDataNotify, userId, clientSeq, g.PacketPlayerDataNotify(player))
	g.SendMsg(cmd.StoreWeightLimitNotify, userId, clientSeq, g.PacketStoreWeightLimitNotify())
	g.SendMsg(cmd.PlayerStoreNotify, userId, clientSeq, g.PacketPlayerStoreNotify(player))
	g.SendMsg(cmd.AvatarDataNotify, userId, clientSeq, g.PacketAvatarDataNotify(player))
	g.SendMsg(cmd.OpenStateUpdateNotify, userId, clientSeq, g.PacketOpenStateUpdateNotify())
	g.GCGLogin(player) // 发送GCG登录相关的通知包
	playerLoginRsp := &proto.PlayerLoginRsp{
		IsUseAbilityHash: true,
		AbilityHashCode:  -228935105,
		GameBiz:          "hk4e_cn",
		IsScOpen:         false,
		RegisterCps:      "taptap",
		CountryCode:      "CN",
		Birthday:         "2000-01-01",
		TotalTickTime:    1185941.871788,
	}
	g.SendMsg(cmd.PlayerLoginRsp, userId, clientSeq, playerLoginRsp)
}

func (g *GameManager) PacketPlayerDataNotify(player *model.Player) *proto.PlayerDataNotify {
	playerDataNotify := &proto.PlayerDataNotify{
		NickName:          player.NickName,
		ServerTime:        uint64(time.Now().UnixMilli()),
		IsFirstLoginToday: true,
		RegionId:          player.RegionId,
		PropMap:           make(map[uint32]*proto.PropValue),
	}
	for k, v := range player.PropertiesMap {
		propValue := &proto.PropValue{
			Type:  uint32(k),
			Value: &proto.PropValue_Ival{Ival: int64(v)},
			Val:   int64(v),
		}
		playerDataNotify.PropMap[uint32(k)] = propValue
	}
	return playerDataNotify
}

func (g *GameManager) PacketStoreWeightLimitNotify() *proto.StoreWeightLimitNotify {
	storeWeightLimitNotify := &proto.StoreWeightLimitNotify{
		StoreType: proto.StoreType_STORE_PACK,
		// 背包容量限制
		WeightLimit:         30000,
		WeaponCountLimit:    2000,
		ReliquaryCountLimit: 1500,
		MaterialCountLimit:  2000,
		FurnitureCountLimit: 2000,
	}
	return storeWeightLimitNotify
}

func (g *GameManager) PacketPlayerStoreNotify(player *model.Player) *proto.PlayerStoreNotify {
	playerStoreNotify := &proto.PlayerStoreNotify{
		StoreType:   proto.StoreType_STORE_PACK,
		WeightLimit: 30000,
	}
	itemDataMapConfig := gdconf.CONF.ItemDataMap
	for _, weapon := range player.WeaponMap {
		pbItem := &proto.Item{
			ItemId: weapon.ItemId,
			Guid:   weapon.Guid,
			Detail: nil,
		}
		itemData, ok := itemDataMapConfig[int32(weapon.ItemId)]
		if !ok {
			logger.Error("config is nil, itemId: %v", weapon.ItemId)
			return nil
		}
		if uint16(itemData.Type) != constant.ItemTypeConst.ITEM_WEAPON {
			continue
		}
		affixMap := make(map[uint32]uint32)
		for _, affixId := range weapon.AffixIdList {
			affixMap[affixId] = uint32(weapon.Refinement)
		}
		pbItem.Detail = &proto.Item_Equip{
			Equip: &proto.Equip{
				Detail: &proto.Equip_Weapon{
					Weapon: &proto.Weapon{
						Level:        uint32(weapon.Level),
						Exp:          weapon.Exp,
						PromoteLevel: uint32(weapon.Promote),
						AffixMap:     affixMap,
					},
				},
				IsLocked: weapon.Lock,
			},
		}
		playerStoreNotify.ItemList = append(playerStoreNotify.ItemList, pbItem)
	}
	for _, reliquary := range player.ReliquaryMap {
		pbItem := &proto.Item{
			ItemId: reliquary.ItemId,
			Guid:   reliquary.Guid,
			Detail: nil,
		}
		if uint16(itemDataMapConfig[int32(reliquary.ItemId)].Type) != constant.ItemTypeConst.ITEM_RELIQUARY {
			continue
		}
		pbItem.Detail = &proto.Item_Equip{
			Equip: &proto.Equip{
				Detail: &proto.Equip_Reliquary{
					Reliquary: &proto.Reliquary{
						Level:            uint32(reliquary.Level),
						Exp:              reliquary.Exp,
						PromoteLevel:     uint32(reliquary.Promote),
						MainPropId:       reliquary.MainPropId,
						AppendPropIdList: reliquary.AffixIdList,
					},
				},
				IsLocked: reliquary.Lock,
			},
		}
		playerStoreNotify.ItemList = append(playerStoreNotify.ItemList, pbItem)
	}
	for _, item := range player.ItemMap {
		pbItem := &proto.Item{
			ItemId: item.ItemId,
			Guid:   item.Guid,
			Detail: nil,
		}
		itemDataConfig := itemDataMapConfig[int32(item.ItemId)]
		if itemDataConfig != nil && uint16(itemDataConfig.Type) == constant.ItemTypeConst.ITEM_FURNITURE {
			pbItem.Detail = &proto.Item_Furniture{
				Furniture: &proto.Furniture{
					Count: item.Count,
				},
			}
		} else {
			pbItem.Detail = &proto.Item_Material{
				Material: &proto.Material{
					Count:      item.Count,
					DeleteInfo: nil,
				},
			}
		}
		playerStoreNotify.ItemList = append(playerStoreNotify.ItemList, pbItem)
	}
	return playerStoreNotify
}

func (g *GameManager) PacketAvatarDataNotify(player *model.Player) *proto.AvatarDataNotify {
	chooseAvatarId := player.MainCharAvatarId
	avatarDataNotify := &proto.AvatarDataNotify{
		CurAvatarTeamId:   uint32(player.TeamConfig.GetActiveTeamId()),
		ChooseAvatarGuid:  player.AvatarMap[chooseAvatarId].Guid,
		OwnedFlycloakList: player.FlyCloakList,
		// 角色衣装
		OwnedCostumeList: player.CostumeList,
		AvatarList:       make([]*proto.AvatarInfo, 0),
		AvatarTeamMap:    make(map[uint32]*proto.AvatarTeam),
	}
	for _, avatar := range player.AvatarMap {
		pbAvatar := g.PacketAvatarInfo(avatar)
		avatarDataNotify.AvatarList = append(avatarDataNotify.AvatarList, pbAvatar)
	}
	for teamIndex, team := range player.TeamConfig.TeamList {
		var teamAvatarGuidList []uint64 = nil
		for _, avatarId := range team.GetAvatarIdList() {
			teamAvatarGuidList = append(teamAvatarGuidList, player.AvatarMap[avatarId].Guid)
		}
		avatarDataNotify.AvatarTeamMap[uint32(teamIndex)+1] = &proto.AvatarTeam{
			AvatarGuidList: teamAvatarGuidList,
			TeamName:       team.Name,
		}
	}
	return avatarDataNotify
}

func (g *GameManager) PacketOpenStateUpdateNotify() *proto.OpenStateUpdateNotify {
	openStateUpdateNotify := &proto.OpenStateUpdateNotify{
		OpenStateMap: make(map[uint32]uint32),
	}
	openStateConstMap := reflection.ConvStructToMap(constant.OpenStateConst)
	// 先暂时开放全部功能模块
	for _, v := range openStateConstMap {
		openStateUpdateNotify.OpenStateMap[uint32(v.(uint16))] = 1
	}
	return openStateUpdateNotify
}

func (g *GameManager) CreatePlayer(userId uint32, nickName string, mainCharAvatarId uint32) *model.Player {
	player := new(model.Player)
	player.PlayerID = userId
	player.NickName = nickName
	player.Signature = ""
	player.MainCharAvatarId = mainCharAvatarId
	player.HeadImage = mainCharAvatarId
	player.Birthday = []uint8{0, 0}
	player.NameCard = 210001
	player.NameCardList = make([]uint32, 0)
	player.NameCardList = append(player.NameCardList, 210001, 210042)

	player.FriendList = make(map[uint32]bool)
	player.FriendApplyList = make(map[uint32]bool)

	player.RegionId = 1
	player.SceneId = 3

	player.PropertiesMap = make(map[uint16]uint32)
	// 初始化所有属性
	propList := reflection.ConvStructToMap(constant.PlayerPropertyConst)
	for fieldName, fieldValue := range propList {
		// 排除角色相关的属性
		if fieldName == "PROP_EXP" ||
			fieldName == "PROP_BREAK_LEVEL" ||
			fieldName == "PROP_SATIATION_VAL" ||
			fieldName == "PROP_SATIATION_PENALTY_TIME" ||
			fieldName == "PROP_LEVEL" {
			continue
		}
		value := fieldValue.(uint16)
		player.PropertiesMap[value] = 0
	}
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_PLAYER_LEVEL] = 1
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_PLAYER_WORLD_LEVEL] = 0
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_IS_SPRING_AUTO_USE] = 1
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_SPRING_AUTO_USE_PERCENT] = 100
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_IS_FLYABLE] = 1
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_IS_TRANSFERABLE] = 1
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_MAX_STAMINA] = 24000
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_CUR_PERSIST_STAMINA] = 24000
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_PLAYER_RESIN] = 160
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_PLAYER_MP_SETTING_TYPE] = 2
	player.PropertiesMap[constant.PlayerPropertyConst.PROP_IS_MP_MODE_AVAILABLE] = 1

	player.FlyCloakList = make([]uint32, 0)
	player.FlyCloakList = append(player.FlyCloakList, 140001)
	player.FlyCloakList = append(player.FlyCloakList, 140002)
	player.FlyCloakList = append(player.FlyCloakList, 140003)
	player.FlyCloakList = append(player.FlyCloakList, 140004)
	player.FlyCloakList = append(player.FlyCloakList, 140005)
	player.FlyCloakList = append(player.FlyCloakList, 140006)
	player.FlyCloakList = append(player.FlyCloakList, 140007)
	player.FlyCloakList = append(player.FlyCloakList, 140008)
	player.FlyCloakList = append(player.FlyCloakList, 140009)
	player.FlyCloakList = append(player.FlyCloakList, 140010)

	player.CostumeList = make([]uint32, 0)
	player.CostumeList = append(player.CostumeList, 200301)
	player.CostumeList = append(player.CostumeList, 201401)
	player.CostumeList = append(player.CostumeList, 202701)
	player.CostumeList = append(player.CostumeList, 204201)
	player.CostumeList = append(player.CostumeList, 200302)
	player.CostumeList = append(player.CostumeList, 202101)
	player.CostumeList = append(player.CostumeList, 204101)
	player.CostumeList = append(player.CostumeList, 204501)
	player.CostumeList = append(player.CostumeList, 201601)
	player.CostumeList = append(player.CostumeList, 203101)
	player.CostumeList = append(player.CostumeList, 200201)
	player.CostumeList = append(player.CostumeList, 200601)
	player.CostumeList = append(player.CostumeList, 208201)

	player.SafePos = &model.Vector{X: 2747, Y: 194, Z: -1719}
	player.Pos = &model.Vector{X: 2747, Y: 194, Z: -1719}
	player.Rot = &model.Vector{X: 0, Y: 307, Z: 0}

	player.ItemMap = make(map[uint32]*model.Item)
	player.WeaponMap = make(map[uint64]*model.Weapon)
	player.ReliquaryMap = make(map[uint64]*model.Reliquary)
	player.AvatarMap = make(map[uint32]*model.Avatar)
	player.GameObjectGuidMap = make(map[uint64]model.GameObject)
	player.DropInfo = model.NewDropInfo()
	player.ChatMsgMap = make(map[uint32][]*model.ChatMsg)
	player.GCGInfo = model.NewGCGInfo()

	// 添加选定的主角
	player.AddAvatar(mainCharAvatarId)
	// 添加初始武器
	avatarDataConfig, ok := gdconf.CONF.AvatarDataMap[int32(mainCharAvatarId)]
	if !ok {
		logger.Error("config is nil, mainCharAvatarId: %v", mainCharAvatarId)
		return nil
	}
	weaponId := uint64(g.snowflake.GenId())
	player.AddWeapon(uint32(avatarDataConfig.InitialWeapon), weaponId)
	// 角色装上初始武器
	player.WearWeapon(mainCharAvatarId, weaponId)

	player.AddReliquary(24825, uint64(g.snowflake.GenId()), 15007)

	player.TeamConfig = model.NewTeamInfo()
	player.TeamConfig.GetActiveTeam().SetAvatarIdList([]uint32{mainCharAvatarId})

	return player
}