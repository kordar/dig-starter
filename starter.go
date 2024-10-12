package dig_starter

import (
	goframeworkdig "github.com/kordar/goframework-dig"
	logger "github.com/kordar/gologger"
	"github.com/spf13/cast"
)

var (
	defaultNamespace string = "default"
)

func HasDigInstance(db string) bool {
	return goframeworkdig.HasDigInstance(db)
}

type DigModule struct {
	name string
	load func(moduleName string, itemId string, item map[string]interface{})
}

func NewDigModule(name string, load func(moduleName string, itemId string, item map[string]interface{})) *DigModule {
	return &DigModule{name, load}
}

func (m DigModule) Name() string {
	return m.name
}

func (m DigModule) _load(id string, cfg map[string]interface{}) {
	if id == "" {
		logger.Fatalf("[%s] the attribute id cannot be empty.", m.Name())
		return
	}

	isDefault := cast.ToBool(cfg["id_default"])
	if isDefault {
		defaultNamespace = id
	}

	if err := goframeworkdig.AddDigInstance(id); err != nil {
		logger.Fatalf("[%s] id=%s，err=%v", m.Name(), id, err)
		return
	}

	if m.load != nil {
		m.load(m.name, id, cfg)
		logger.Debugf("[%s] triggering custom loader completion", m.Name())
	}

	logger.Infof("[%s] loading module '%s' successfully", m.Name(), id)
}

func (m DigModule) Load(value interface{}) {
	items := cast.ToStringMap(value)
	if items["id"] != nil {
		defaultNamespace = cast.ToString(items["id"])
		m._load(defaultNamespace, items)
		return
	}
	for key, item := range items {
		m._load(key, cast.ToStringMap(item))
	}
}

func (m DigModule) Close() {
}
