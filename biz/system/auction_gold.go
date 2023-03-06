package system

/*
1. 拍卖行 定期 创建 表
*/

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
	"tool/mods/db"
)

// Auction 表格复制，最后加上 年月, 比如 table => table_202303
type Auction struct {
	DBKey         string
	CopySrcTables []string
	Wg            *sync.WaitGroup
	Quit          chan bool
}

func NewAuction() *Auction {
	return &Auction{
		DBKey:         "TestMysql1",
		CopySrcTables: []string{"auction_history", "auction_history_buyer"},
		Quit:          make(chan bool),
	}
}

func (a *Auction) Run() {
	go a.work()
}

func (a *Auction) Stop() {
	a.Quit <- true
}

func (a *Auction) SetWaitGroup(wg *sync.WaitGroup) {
	a.Wg = wg
}

func (a *Auction) work() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	defer a.Wg.Done()

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			nowStr := now.Format("200601")
			dbx := db.DBPools.Get(a.DBKey)
			if dbx == nil {
				log.Warn("auction not working")
				break
			}

			for _, table := range a.CopySrcTables {
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
			fmt.Println("auction is checking")

		case <-a.Quit:
			return
		}
	}

}
