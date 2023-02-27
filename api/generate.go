package api

//go:generate rm -rf models restapi
//go:generate bash -c "~/go/bin/swagger generate server --exclude-main --flag-strategy=flag -A quizzer -f api.yaml"
