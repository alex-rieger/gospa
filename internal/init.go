package internal

import (
	"flag"
	"html/template"
	"net/http"

	"github.com/alex-rieger/govite"
	"go.uber.org/zap"
)

type application struct {
	logger        *zap.Logger
	templateCache map[string]*template.Template
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

	vite := govite.New(govite.Config{
		DevServerEnabled:   *dev,
		DevServerProtocol:  "http",
		DevServerHost:      "localhost",
		DevServerPort:      "3000",
		WebSocketClientUrl: "@vite/client",
		AssetsPath:         "./web/view/assets",
		ManifestPath:       "./web/view/dist/manifest.json",
	})

	templateCache, err := newTemplateCache(templatesDir, template.FuncMap{
		"vite":  vite.TemplateTagViteClient,
		"asset": vite.TemplateTagAsset,
	})
	if err != nil {
		return err
	}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
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
