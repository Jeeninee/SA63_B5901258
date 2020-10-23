package main

 
import (
	"context"
	"log"
  
	"github.com/oreo/app/controllers"
	_ "github.com/oreo/app/docs"
	"github.com/oreo/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
 )
 
type Province struct {
	Name string
}

type Provinces struct {
	Province []Province
}

type District struct {
	District string
}

type Districts struct {
	District []District
}

type Subdistrict struct {
	Subdistrict string
}

type Subdistricts struct {
	Subdistrict []Subdistrict
}

type Student struct {
	Student string
}

type Students struct {
	Student []Student
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewProvinceController(v1, client)
	controllers.NewDistrictController(v1, client)
	controllers.NewSubdistrictController(v1, client)
	controllers.NewStudentController(v1, client)

	

	//Set District Data
	districts := Districts{
		District: []District{
			District{"เมืองนครราชสีมา"},
			District{"เมืองนครราชสีมา"},
		},
	}
	for _, di := range districts.District {
		client.District.
			Create().
			SetDistrict(di.District).
			Save(context.Background())
	}

	//Set Subdistrict Data
	subdistricts := Subdistricts{
		Subdistrict: []Subdistrict{
			Subdistrict{"ในเมือง"},
			Subdistrict{"ในเมือง"},
		},
	}
	for _, sd := range subdistricts.Subdistrict {
		client.Subdistrict.
			Create().
			SetSubdistrict(sd.Subdistrict).
			Save(context.Background())
	}

	//Set Student Data
	students := Students{
		Student: []Student{
			Student{"จีนิณีย์"},
			Student{"จีนิณีย์"},
		},
	}
	for _, st := range students.Student {
		client.Student.
			Create().
			SetStudent(st.Student).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
