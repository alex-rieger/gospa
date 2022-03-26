package internal

import (
	"encoding/json"
	"flag"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/alex-rieger/govite"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

type sharedConfig struct {
	DefaultLocale  string
	EnabledLocales []string
}

type application struct {
	logger        *zap.Logger
	templateCache map[string]*template.Template
	config        *sharedConfig
	i18n          *i18n.Bundle
}

func Init() error {
	addr := flag.String("addr", ":8080", "http network address")
	dev := flag.Bool("dev", true, "run in localmode?")
	flag.Parse()

	logger, err := initLogger()
	if err != nil {
		return err
	}
	defer logger.Sync()

	config, err := initSharedConfig()
	if err != nil {
		return err
	}

	// handle vite
	vite := govite.New(govite.Config{
		DevServerEnabled:   *dev,
		DevServerProtocol:  "http",
		DevServerHost:      "localhost",
		DevServerPort:      "3000",
		WebSocketClientUrl: "@vite/client",
		AssetsPath:         "./web/view/assets",
		ManifestPath:       "./web/view/dist/manifest.json",
	})

	// handle templates
	templateCache, err := newTemplateCache(templatesDir, template.FuncMap{
		"vite":  vite.TemplateTagViteClient,
		"asset": vite.TemplateTagAsset,
	})
	if err != nil {
		return err
	}

	// handle translations
	i18nBundle := i18n.NewBundle(language.AmericanEnglish)
	i18nBundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	for _, enabledLocale := range config.EnabledLocales {
		i18nBundle.LoadMessageFile("./configs/lang/" + enabledLocale + ".json")
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
		config:        config,
		i18n:          i18nBundle,
	}

	app.logger.Info("application started")

	return http.ListenAndServe(*addr, app.routes())
}

func initLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return logger, nil
}

func initSharedConfig() (*sharedConfig, error) {
	config := &sharedConfig{}
	configBytes, err := ioutil.ReadFile(sharedConfigPath)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
