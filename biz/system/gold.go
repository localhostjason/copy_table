package system

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
	"tool/mods/db"
)

/*
1. 金币 定期 创建 表
*/

type Gold struct {
	DBKey         string
	CopySrcTables []string
	Wg            *sync.WaitGroup
	Quit          chan bool
}

func NewGold() *Gold {
	return &Gold{
		DBKey:         "TestMysql2",
		CopySrcTables: []string{"user"},
		Quit:          make(chan bool),
	}
}

func (g *Gold) Run() {
	go g.work()
}

func (g *Gold) Stop() {
	g.Quit <- true
}

func (g *Gold) SetWaitGroup(wg *sync.WaitGroup) {
	g.Wg = wg
}

func (g *Gold) work() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	defer g.Wg.Done()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			nowStr := now.Format("200601")
			dbx := db.DBPools.Get(g.DBKey)
			if dbx == nil {
				log.Warn("gold not working")
				break
			}
			for _, table := range g.CopySrcTables {
				distTable := fmt.Sprintf("%s_%s", table, nowStr)

				ok := dbx.Migrator().HasTable(distTable)
				if ok {
					continue
				}

				if err := dbx.Debug().Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (LIKE %s);", distTable, table)).Error; err != nil {
					log.Warn("auction CREATE TABLE  err:", err)
					break
				}
			}
			fmt.Println("gold is checking")
		case <-g.Quit:
			return
		}
	}

}
