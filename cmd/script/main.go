package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	
	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	gormdb, _ := gorm.Open(postgres.Open("host=192.168.3.30 user=gauss password=Gauss666 dbname=bit_mail port=5432 sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	g.UseDB(gormdb) // reuse your gorm db
	
	// Generate basic type-safe DAO API for struct `model.User` following conventions
	
	g.ApplyBasic(
		// Generate struct `User` based on table `users`
		g.GenerateModel("users"),
		
		// Generate struct `Employee` based on table `users`
		g.GenerateModelAs("users", "User"),
		
		
		// Generate struct `User` based on table `users` and generating options
		// g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	
	)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
