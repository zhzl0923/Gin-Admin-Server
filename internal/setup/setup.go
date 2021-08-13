package setup

import "log"

func Setup() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("setupLogger err: %v", err)
	}

	err = setupDB()
	if err != nil {
		log.Fatalf("setupDB err: %v", err)
	}

	err = setupTrans("zh")
	if err != nil {
		log.Fatalf("setup trans failed, err:%v\n", err)
	}

}
