package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"quizzer/api"
	"quizzer/api/restapi"
	"quizzer/config"

	"github.com/AlphaOne1/templig"
)

func main() {
	fmt.Println(`   _,-----,_   `)
	fmt.Println(`  /  ,---,  \  `)
	fmt.Println(` /  /     \  \      .-')                            .-') _    .-') _   ('-.  _  .-')  `)
	fmt.Println(`|__|       |  |   .(  OO)                          (  OO) )  (  OO) )_(  OO)( \( -O ) `)
	fmt.Println(`           |  |  (_)---\_)   ,--. ,--.    ,-.-') ,(_)----. ,(_)----.(,------.,------. `)
	fmt.Println(`          /  /   '  .-.  '   |  | |  |    |  |OO)|       | |       | |  .---'|   /'. '`)
	fmt.Println(`         /  /   ,|  | |  |   |  | | .-')  |  |  \'--.   /  '--.   /  |  |    |  /  | |`)
	fmt.Println(`        /  /   (_|  | |  |   |  |_|( OO ) |  |(_/(_/   /   (_/   /  (|  '--. |  |_.' |`)
	fmt.Println(`       /  /      |  | |  |   |  | | '-' /,|  |_.' /   /___  /   /___ |  .--' |  .  '.'`)
	fmt.Println(`      (  (       '  '-'  '-.('  '-'(_.-'(_|  |   |        ||        ||  '---.|  |\  \ `)
	fmt.Println(`      |__|        '-----'--'  '-----'     '--'   '--------''--------''------''--' '--'`)
	fmt.Println(`       __      `)
	fmt.Println(`      /  \     `)
	fmt.Println(`      \__/     `)

	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "configuration file")

	cfg, cfgErr := templig.FromFile[config.Root](configFileName)

	if cfgErr != nil {
		slog.Error("could not read configuration file", slog.String("error", cfgErr.Error()))
		os.Exit(1)
	}

	slog.Info("configuration loaded")
	slog.Info("tasks available", slog.Int("num_tasks", len(cfg.Get().Tasks)))
	slog.Info("server config",
		slog.String("address", cfg.Get().ListenAddress),
		slog.Int("port", cfg.Get().ListenPort),
	)

	handler := &api.Handler{
		Tasks: cfg.Get().Tasks,
	}

	service, serviceErr := restapi.NewServer(handler)

	if serviceErr != nil {
		slog.Error("could not start API server", slog.String("error", serviceErr.Error()))
		os.Exit(1)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Get().ListenAddress, cfg.Get().ListenPort),
		Handler: service,
	}

	defer func() { _ = server.Shutdown(context.Background()) }()

	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			slog.Error("server shutdown unexpectedly")
			os.Exit(1)
		}

		slog.Info("server shutdown")
	}
}
