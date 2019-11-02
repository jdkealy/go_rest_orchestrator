package types

type RoutesJsonField struct {
	Name string `json:"name"`
	Id string  `json:"id"`
	Cmp string `json:"cmp"`
}

type JsModelsJsonFields struct {
	Name string `json:"name"`
	Id string  `json:"id"`
	Label string `json:"label"`
}

type Schema struct {
	StructFields []string
	RoutesJsonFields []RoutesJsonField
	JsFieldsConfig string
	JsRoutesConfig string
	FullFilePath string
	DependencyPath string
	CmdPath string
	ModelPath string
	RoutesPath string
	TestsPath string
	DbPath string
	CachePath string
	GoRoot string
	GitPath string
	ProjectRoot string
	GitRoutesPath string
	GitDbPath string
	GitModelsPath string
	ViewPath string
	JsModelPath string
	JsModelsPath string
	JsViewsPath string
	JsPageListPath string
	JsPageFieldsPath string
	JsPageNewPath string
	JsRoutesConfigPath string
	Type string
	Model string
	LowerModel string
	PluralModel string
	PluralLowerModel string
	Owner string
	Project string
	Name string
}

type Fields struct {
	Name string
	Id string
	Cmp string
	Type string
	Gorm string
	Json string
}

