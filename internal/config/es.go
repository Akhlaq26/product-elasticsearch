package config

import (
	"github.com/elastic/go-elasticsearch/v8"
	fiberlog "github.com/gofiber/fiber/v3/log"
)

func InitES(cfg *Config) *elasticsearch.Client {
	esCfg := elasticsearch.Config{
		Addresses: []string{
			cfg.EsHost,
		},
		Username: cfg.EsUser,
	}
	es, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		fiberlog.Fatal(err)
	}
	res, err := es.Info()
	if err != nil {
		fiberlog.Fatal(err)
	}

	fiberlog.Infof("connected to elasticsearch : %v", res.String())
	return es
}
