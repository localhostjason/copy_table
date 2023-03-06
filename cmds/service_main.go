package cmds

import (
	"errors"
	"os"
	"syscall"
	"tool/biz/system"
	"tool/mods/db"
	"tool/mods/logger"
	"tool/mods/svc"

	log "github.com/sirupsen/logrus"
)

type MainProc struct {
	singleMode bool
	quit       chan bool
}

func (m *MainProc) Stop() {
	m.quit <- true
}

func NewProc(singleMode bool) *MainProc {
	return &MainProc{singleMode: singleMode, quit: make(chan bool)}
}

func (m *MainProc) Run(svc *svc.Svc) {
	if err := startAgent(true); err != nil {
		return
	}

	tm := system.NewTaskManage()
	tm.Add(system.NewAuction())
	tm.Add(system.NewGold())
	tm.Run()
	<-m.quit
	tm.Stop()

}

func (m *MainProc) SigHandlers() map[os.Signal]svc.SignalHandlerFunc {
	return map[os.Signal]svc.SignalHandlerFunc{
		syscall.SIGTERM: m.handleSigTerm,
		os.Interrupt:    m.handleSigTerm,
	}
}

func (m *MainProc) handleSigTerm(sig os.Signal) (err error) {
	m.quit <- true
	return errors.New("quit by signal " + sig.String())
}

func startAgent(toConsole bool) error {
	err := logger.SetLogConfig(toConsole)
	if err != nil {
		log.Fatalln("failed to set log:", err)
	}

	if err = db.Connect(); err != nil {
		log.Fatalln(err)
	}
	if err = db.InitData(); err != nil {
		log.Fatalln(err)
	}
	return nil
}
