package main

import (
	"github.com/kerrgavin/srb2kart-go/config"
)

func mapSrb2kart(srb2kart *srb2kart) string {
	var configHelper config.ConfigHelper
	configHelper.Comment("main configs")
	mapServer(&srb2kart.Server, &configHelper)
	mapGeneral(&srb2kart.General, &configHelper)
	mapCheats(&srb2kart.Cheats, &configHelper)
	return configHelper.String()
}

func mapServer(server *server, configHelper *config.ConfigHelper) {
	configHelper.ConfigString("advancemap", server.AdvanceMap)
	configHelper.ConfigBool("allowjoin", server.AllowJoin)
	configHelper.ConfigBool("blamecfail", server.BlameCFail)
	configHelper.ConfigBool("downloading", server.Downloading)
	configHelper.ConfigInt("downloadspeed", server.DownloadSpeed)
	configHelper.ConfigInt("inttime", server.IntTime)
	configHelper.ConfigBool("joinnextround", server.JoinNextRound)
	configHelper.ConfigInt("jointimeout", server.JoinTimeOut)
	configHelper.ConfigInt("maxping", server.MaxPing)
	configHelper.ConfigInt("maxplayers", server.MaxPlayers)
	configHelper.ConfigInt("maxsend", server.MaxSend)
	configHelper.ConfigBool("mute", server.Mute)
	configHelper.ConfigInt("nettimeout", server.NetTimeOut)
	configHelper.ConfigInt("maxdelaytimeout", server.MaxDelayTimeout)
	configHelper.ConfigInt("maxdelay", server.MaxDelay)
	configHelper.ConfigInt("cpusleep", server.CpuSleep)
	configHelper.ConfigInt("kicktime", server.KickTime)
	configHelper.ConfigString("http_source", server.HttpSource)
	configHelper.ConfigString("kartvoices", server.KartVoices)
	configHelper.ConfigString("discordinvites", server.DiscordInvites)
	configHelper.ConfigString("pingmeasurement", server.PingMeasurement)
	configHelper.ConfigString("server_contact", server.Contact)
	configHelper.ConfigString("holepunchserver", server.HolePunchServer)
	configHelper.ConfigBool("noticedownload", server.NoticeDownload)
	configHelper.ConfigString("pausepermission", server.PausePermission)
	configHelper.ConfigInt("resynchattempts", server.ResynchAttempts)
	configHelper.ConfigString("servername", server.Name)
	configHelper.ConfigBool("showjoinaddress", server.ShowJoinAddress)
	configHelper.ConfigBool("exitmove", server.ExitMove)
	mapMasterServer(&server.MasterServer, configHelper)
}

func mapMasterServer(masterServer *masterServer, configHelper *config.ConfigHelper) {
	configHelper.ConfigString("masterserver", masterServer.Url)
	configHelper.ConfigInt("masterserver_update_rate", masterServer.UpdateRate)
	configHelper.ConfigInt("masterserver_timeout", masterServer.Timeout)
	configHelper.ConfigBool("masterserver_debug", masterServer.Debug)
	configHelper.ConfigString("masterserver_token", masterServer.Token)
	configHelper.ConfigInt("masterserver_nagattempts", masterServer.NagAttempts)
}

func mapGeneral(general *general, configHelper *config.ConfigHelper) {
	configHelper.ConfigBool("allowexitlevel", general.AllowExitLevel)
	configHelper.ConfigBool("allowmlook", general.AllowMLook)
	configHelper.ConfigBool("allowseenames", general.AllowSeeNames)
	configHelper.ConfigBool("allowteamchange", general.AllowTeamChange)
	configHelper.ConfigInt("autobalance", general.AutoBalance)
	configHelper.ConfigInt("forceskin", general.ForceSkin)
	configHelper.ConfigBool("friendlyfire", general.FriendlyFire)
	configHelper.ConfigBool("overtime", general.Overtime)
	configHelper.ConfigInt("pointlimit", general.PointLimit)
	configHelper.ConfigInt("respawndelay", general.RespawnDelay)
	configHelper.ConfigBool("respawnitem", general.RespawnItem)
	configHelper.ConfigString("scrambleonchange", general.ScrambleOnChange)
	configHelper.ConfigBool("skipmapcheck", general.SkipMapCheck)
	configHelper.ConfigString("teamscramble", general.TeamScramble)
	configHelper.ConfigInt("timelimit", general.Timelimit)
	configHelper.ConfigInt("countdowntime", general.CountDownTime)
	configHelper.ConfigInt("numlaps", general.NumLaps)
}

func mapCheats(cheats *cheats, configHelper *config.ConfigHelper) {
	configHelper.ConfigInt("respawnitemtime", cheats.RespawnItemTime)
	configHelper.ConfigBool("restrictskinchange", cheats.RestrictSkinChange)
}
