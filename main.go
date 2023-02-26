package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"questioner/config"
)

func main() {
	fmt.Println(`   _,-----,_   `)
	fmt.Println(`  /  ,---,  \  `)
	fmt.Println(` /  /     \  \       .-')                    ('-.    .-')    .-') _                            .-') _   ('-.  _  .-')   `)
	fmt.Println(`|__|       |  |    .(  OO)                 _(  OO)  ( OO ). (  OO) )                          ( OO ) )_(  OO)( \( -O )  `)
	fmt.Println(`           |  |   (_)---\_)   ,--. ,--.   (,------.(_)---\_)/     '._ ,-.-')  .-'),-----. ,--./ ,--,'(,------.,------.  `)
	fmt.Println(`          /  /    '  .-.  '   |  | |  |    |  .---'/    _ | |'--...__)|  |OO)( OO'  .-.  '|   \ |  |\ |  .---'|   /'. ' `)
	fmt.Println(`         /  /    ,|  | |  |   |  | | .-')  |  |    \  :' '. '--.  .--'|  |  \/   |  | |  ||    \|  | )|  |    |  /  | | `)
	fmt.Println(`        /  /    (_|  | |  |   |  |_|( OO )(|  '--.  '..'''.)   |  |   |  |(_/\_) |  |\|  ||  .     |/(|  '--. |  |_.' | `)
	fmt.Println(`       /  /       |  | |  |   |  | | '-' / |  .--' .-._)   \   |  |  ,|  |_.'  \ |  | |  ||  |\    |  |  .--' |  .  '.' `)
	fmt.Println(`      (  (        '  '-'  '-.('  '-'(_.-'  |  '---.\       /   |  | (_|  |      ''  '-'  '|  | \   |  |  '---.|  |\  \  `)
	fmt.Println(`      |__|         '-----'--'  '-----'     '------' '-----'    '--'   '--'        '-----' '--'  '--'  '------''--' '--' `)
	fmt.Println(`       __      `)
	fmt.Println(`      /  \     `)
	fmt.Println(`      \__/     `)

	var configFileName string
	flag.StringVar(&configFileName, "config", "config.json", "configuration file")

	configFile, configFileErr := os.Open(configFileName)

	if configFileErr != nil {
		log.Fatalf("couldn't open configuration file %v: %v", configFileName, configFileErr)
	}

	defer configFile.Close()

	var config config.Root

	if decodeErr := json.NewDecoder(configFile).Decode(&config); decodeErr != nil {
		log.Fatalf("could not read configuration in %v: %v", configFileName, decodeErr)
	}

	log.Printf("configuration loaded")
	log.Printf("%v tasks available", len(config.Tasks))
	log.Printf("server address: %v", config.ListenAddress)
	log.Printf("server port: %v", config.ListenPort)

	os.Exit(0)
}
