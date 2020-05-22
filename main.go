package main

import (
	"encoding/json"
	"gorm_poc/connector"
	"gorm_poc/domain"
	"gorm_poc/router"
	"gorm_poc/sql"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/lib/pq"
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
	sql.Test()
	wg.Add(1)
	dbConfig := connector.ReadInDbConfig()
	connString := dbConfig.GetConnectionString()
	log.Println(connString)
	db, err := gorm.Open("postgres", dbConfig.GetConnectionString())

	if err != nil {
		log.Println("err: %v", err)
		log.Fatal("Error connecting to DB")
	}
	db.LogMode(true)
	defer db.Close()

	go runServerAsync()

	defer func() {
		log.Println("*** Finishing all threads ***")
		wg.Wait()
	}()

	log.Println("*** GORM Examples ***")
	/*
		EXAMPLE 1:
		- insert a simple object into a database
		- how to manage a JSONB column (the Postgresql binary JSON column type).
	*/
	reg1 := domain.Region{}
	reg1.Name = "us-west-1"
	reg1.CreatedAt = time.Now()
	jsonbDescription1 := json.RawMessage(`{"city":"San Francisco","active":"true","state":"California"}`)
	reg1.Description = postgres.Jsonb{RawMessage: jsonbDescription1}
	////note: we do not need to initialize the PK field (ID)
	db.Create(&reg1)

	////note: the second instance (the same approach)
	reg1.Description = postgres.Jsonb{RawMessage: jsonbDescription1}
	reg2 := domain.Region{}
	reg2.Name = "us-east-1"
	reg2.CreatedAt = time.Now()
	jsonbDescription2 := json.RawMessage(`{"city":"New York","active":"true","state":"New York"}`)
	reg2.Description = postgres.Jsonb{RawMessage: jsonbDescription2}
	db.Create(&reg2)
	log.Println("Example 1: %+v", reg1.ID) ////note IDs are added after Create()
	log.Println("Example 1: %+v", reg2.ID)
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
		- how to handle json arrays in Postgres
		(specifically `json[]` not an array of json within the json data type)
	*/
	////note: hrDocs are of type json and we use json.RawMessage to handle it
	hrDocs := json.RawMessage(`{"status":"classified", "content":"some documentation that is very interesting"}`)
	////note: finDocsStringArray (a seed for FinancialDocumentation column)
	////is handled pq.StringArray but we need to convert it to a slice of string as an intermediary step
	finDocsStringArray := []string{`{"id":"1","name":"AAAA1234dfsdf"}`, `{"id":"2","name":{"lang":"IT", "name":"Gianpaolo"}}`}
	dep1 := domain.Department{
		Name:                   "Accounting",
		CreatedAt:              time.Now(),
		HrDocumentation:        postgres.Jsonb{RawMessage: hrDocs},
		FinancialDocumentation: finDocsStringArray,
	}
	dep2 := domain.Department{
		Name:                   "Logistics",
		CreatedAt:              time.Now(),
		HrDocumentation:        postgres.Jsonb{RawMessage: hrDocs},
		FinancialDocumentation: finDocsStringArray,
	}
	dep3 := domain.Department{
		Name:                   "Trading",
		CreatedAt:              time.Now(),
		HrDocumentation:        postgres.Jsonb{RawMessage: hrDocs},
		FinancialDocumentation: finDocsStringArray,
	}
	db.Create(&dep1) ////save in DB
	db.Create(&dep2)
	db.Create(&dep3)

	////note: this is how we handle json[] in a raw SQL query
	db.Exec(`INSERT INTO "Departments"(name,"financialDocumentation") VALUES('name', ARRAY['{"AAA":"1"}'::json,'{"BBB":""}'::json])`)

	log.Println("Example 3 /FinancialDocumentation: dep1, dep2, dep3/: %v || %v || %v",
		dep1.FinancialDocumentation, dep2.FinancialDocumentation, dep3.FinancialDocumentation)

	/*
		EXAMPLE 4:
		- how to handle associations MANY-TO-MANY
	*/
	// team1 := domain.Team{
	// 	Name:       "Equity",
	// 	Department: &dep3,
	// 	Regions:    []*domain.Region{&reg1},
	// }
	// team2 := domain.Team{
	// 	Name:       "Equity",
	// 	Department: &dep3,
	// 	Regions:    []*domain.Region{&reg2},
	// }
	// log.Println("Example 4 /reg1.id/: %v", reg1.ID)
	// log.Println("Example 4 /team1.id/: %v", team1.ID)
	// // log.Println("Example 4 /team1.id/: %v", team1.RegionId)
	// log.Println("Example 4 /reg2.id/: %v", reg2.ID)
	// log.Println("Example 4 /team2.id/: %v", team2.ID)
	// // log.Println("Example 4 /team2.id/: %v", team2.RegionId)
	// reg1.Teams = []*domain.Team{&team1}
	// reg2.Teams = []*domain.Team{&team2}
	// // db.Save(&reg1)
	// db.Model(&reg1).Association("Teams").Append(&team1)

	// db.Save(&reg1)
	// db.Save(&team1)
	// db.Save(&team2)
	log.Println("**********")
}
