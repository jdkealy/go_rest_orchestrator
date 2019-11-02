package types


type Schema struct {
	StructFields []string
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
	JsModelPath string
	JsModelsPath string
	JsViewsPath string
	JsPageListPath string
	JsPageNewPath string
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
	Type string
	Gorm string
	Json string
}

