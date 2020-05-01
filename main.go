package main

import (
	"encoding/json"
	"gorm_poc/connector"
	"gorm_poc/domain"
	"gorm_poc/router"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db *gorm.DB
	wg sync.WaitGroup
)

func runServerAsync() {
	defer wg.Done()
	log.Fatal(http.ListenAndServe(":8080", router.InitRouter()))
}

func main() {
	wg.Add(1)
	dbConfig := connector.ReadInDbConfig()
	connString := dbConfig.GetConnectionString()
	log.Println(connString)
	db, err := gorm.Open("postgres", dbConfig.GetConnectionString())

	if err != nil {
		log.Println("err: %v", err)
		log.Fatal("Error connecting to DB")
	}
	defer db.Close()

	go runServerAsync()

	defer func() {
		// log.Println("*** Finishing all threads ***")
		wg.Wait()
	}()

	log.Println("*** GORM Examples ***")

	/*
		EXAMPLE 1:
		- insert a simple object into a database
		- how to manage a JSONB column (the Postgresql binary JSON column type).
	*/
	reg1 := &domain.Region{}
	reg1.Name = "us-west-1"
	reg1.CreatedAt = time.Now()
	jsonbDescription1 := json.RawMessage(`{"city":"San Francisco","active":"true","state":"California"}`)
	reg1.Description = postgres.Jsonb{RawMessage: jsonbDescription1}
	////note: we do not need to initialize the PK field (ID)
	db.Create(reg1)

	////note: the second instance (the same approach)
	reg1.Description = postgres.Jsonb{RawMessage: jsonbDescription1}
	reg2 := &domain.Region{}
	reg2.Name = "us-east-1"
	reg2.CreatedAt = time.Now()
	jsonbDescription2 := json.RawMessage(`{"city":"New York","active":"true","state":"New York"}`)
	reg2.Description = postgres.Jsonb{RawMessage: jsonbDescription2}
	db.Create(reg2)
	log.Println("**********")
	/*
		EXAMPLE 2:
		- simple queries (First, Find, Where, ) in PLAIN SQL
	*/
	var fReg1 domain.Region
	var fRegs []domain.Region
	regName1 := reg1.Name
	db.Where("name = ?", regName1).First(&fReg1)
	log.Printf("Example 2 /First/: %+v\n", fReg1)
	db.Where("name = ?", regName1).Find(&fRegs)
	log.Printf("Example 2 /Find - slice/: %+v\n", len(fRegs))
	db.Where([]uint{1, 2, 3}).Find(&fRegs)
	log.Printf("Example 2 /Find - slice - where id in (1,2,3)/: %+v\n", len(fRegs))
	db.Where("name = ?", regName1).
		Not([]uint{1}).
		Where("description is not null").
		Find(&fRegs)
	log.Printf("Example 2 /Find - slice/: %+v\n", len(fRegs))
	log.Println("**********")
	/*
		EXAMPLE 3:
		- how to handle object relationships
		- insert an object with its children
	*/
	hrDocs := json.RawMessage(`{"status":"classified", "content":"some documentation that is very interesting"}`)
	dep1 := domain.Department{Name: "Accounting", CreatedAt: time.Now(), HrDocumentation: postgres.Jsonb{RawMessage: hrDocs}}
	dep2 := domain.Department{Name: "Logistics", CreatedAt: time.Now(), HrDocumentation: postgres.Jsonb{RawMessage: hrDocs}}
	dep3 := domain.Department{Name: "Trading", CreatedAt: time.Now(), HrDocumentation: postgres.Jsonb{RawMessage: hrDocs}}
	// db.Create(dep1)
	// db.Create(dep2)
	// db.Create(dep3)
	log.Println(dep1, dep2, dep3)
	team1 := domain.Team{Name: "Equity", Department: &dep3, Regions: []domain.Region{*reg1}}
	team2 := domain.Team{Name: "Equity", Department: &dep3, Regions: []domain.Region{*reg2}}
	db.Save(&team1)
	db.Save(&team2)
	log.Println("**********")
}
