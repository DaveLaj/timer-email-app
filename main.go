package main

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	TimerEmailLoop(cfg)
	RunForever()
}

func RunForever() {
	select {}
}
