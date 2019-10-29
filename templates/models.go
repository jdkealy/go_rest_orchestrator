package templates

var ModelTemplate = `package models

type {{.Model}} struct {
	ID int ` + "`" + `gorm:"AUTO_INCREMENT"` + "`" + `
}

func (h *Model) Create{{.Model}}({{.LowerModel}} {{.Model}}) (*{{.Model}},  error ){
	a := h.Db.Create(&{{.LowerModel}})
	err := a.Error
	if err != nil {
		return nil, err
	}
	return &{{.LowerModel}}, nil
}

func (h *Model) All{{.PluralModel}}() ([]{{.Model}},  error ){
	var {{.PluralLowerModel}} []{{.Model}}
	res := h.Db.Find(&{{.PluralLowerModel}})
	err := res.Error
	if err != nil {
		return nil, err
	}
	return {{.PluralLowerModel}}, nil
}

func (h *Model) Update{{.Model}}({{.LowerModel}} {{.Model}}) (*{{.Model}}, error){
	ret := h.Db.Model(&{{.LowerModel}}).Update()
	err := ret.Error
	if err != nil {
		return nil, err
	}
	return &{{.LowerModel}}, nil
}

func (h *Model) Delete{{.Model}}({{.LowerModel}} {{.Model}}) (error){
	ret := h.Db.Delete(&{{.LowerModel}})
	err := ret.Error
	if err != nil {
		return  err
	}

	return nil
}`