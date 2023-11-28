package main

type srb2kart struct {
	Server server
	General general
	Cheats cheats
}

type server struct {
	AdvanceMap *string `json:"advanceMap"`
	AllowJoin *bool `json:"allowJoin"`
	BlameCFail *bool `json:"blameCFail"`
	Downloading *bool `json:"downloading"`
	DownloadSpeed *int `json:"downloadSpeed"`
	IntTime *int `json:"intTime"`
	JoinNextRound *bool `json:"joinNextRound"`
	JoinTimeOut *int `json:"joinTimeOut"`
	MaxPing *int `json:"maxPing"`
	MaxPlayers *int `json:"maxPlayers"`
	MaxSend *int `json:"maxSend"`
	Mute *bool `json:"mute"`
	NetTimeOut *int `json:"netTimeOut"`
	MaxDelayTimeout *int `json:"maxDelayTimeout"`
	MaxDelay *int `json:"maxDelay"`
	CpuSleep *int `json:"cpuSleep"`
	KickTime *int `json:"kickTime"`
	HttpSource *string `json:"httpSource"`
	KartVoices *string `json:"kartVoices"`
	DiscordInvites *string `json:"discordInvites"`
	PingMeasurement *string `json:"pingMeasurement"`
	Contact *string `json:"contact"`
	HolePunchServer *string `json:"holePunchServer"`
	NoticeDownload *bool `json:"noticeDownload"`
	PausePermission *string `json:"pausePermission"`
	ResynchAttempts *int `json:"resynchAttempts"`
	Name *string `json:"name"`
	ShowJoinAddress *bool `json:"showJoinAddress"`
	ExitMove *bool `json:"exitMove"`
	MasterServer masterServer `json:"masterServer"`
}

type masterServer struct {
	Url *string `json:"url"`
	UpdateRate *int `json:"updateRate"`
	Timeout *int `json:"timeout"`
	Debug *bool `json:"debug"`
	Token *string `json:"token"`
	NagAttempts *int `json:"nagAttempts"`
}

type general struct {
	AllowExitLevel *bool `json:"allowExitLevel"`
	AllowMLook *bool `json:"allowMLook"`
	AllowSeeNames *bool `json:"allowSeeNames"`
	AllowTeamChange *bool `json:"allowTeamChange"`
	AutoBalance *int `json:"autoBalance"`
	ForceSkin *int `json:"forceSkin"`
	FriendlyFire *bool `json:"friendlyFire"`
	Overtime *bool `json:"overtime"`
	PointLimit *int `json:"pointLimit"`
	RespawnDelay *int `json:"respawnDelay"`
	RespawnItem *bool `json:"respawnItem"`
	ScrambleOnChange *string `json:"scrambleOnChange"`
	SkipMapCheck *bool `json:"skipMapCheck"`
	TeamScramble *string `json:"teamScramble"`
	Timelimit *int `json:"timelimit"`
	CountDownTime *int `json:"countDownTime"`
	NumLaps *int `json:"numLaps"`
}

type cheats struct {
	RespawnItemTime *int `json:"respawnItemTime"`
	RestrictSkinChange *bool `json:"restrictSkinChange"`
}
