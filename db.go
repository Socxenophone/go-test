// YApi QuickType插件生成，具体参考文档:https://plugins.jetbrains.com/plugin/18847-yapi-quicktype/documentation

type DB struct {
	CouncilMembers []CouncilMember `json:"councilMembers"`
	CityName       string          `json:"cityName"`
	Exists         bool            `json:"exists"`
	Hero           string          `json:"hero"`
	Formed         int64           `json:"formed"`
}

type CouncilMember struct {
	Skills []string `json:"skills"`
	Name   string   `json:"name"`
	Class  string   `json:"class"`
	Age    int64    `json:"age"`
}
