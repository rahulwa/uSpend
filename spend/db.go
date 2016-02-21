package main

import (
  "database/sql"
  "github.com/mattn/go-sqlite3"
  "log"
  "time"
  "fmt"
)

/*var (
  CreateWallet := "CREATE TABLE `wallet` (
	 `walletId`	INTEGER NOT NULL UNIQUE,
	 `name`	TEXT NOT NULL,
	 `currency`	TEXT NOT NULL,
	 `initialBalance`	INTEGER,
	 PRIMARY KEY(walletId));"
  CreateCatogary := "CREATE TABLE `catogary` (
	 `name`	TEXT NOT NULL,
	 `catogaryId`	INTEGER NOT NULL UNIQUE,
   `walletId` INTEGER NOT NULL,
	 `isExpenseType`	INTEGER NOT NULL,
	 PRIMARY KEY(catogaryId));"
  CreateTransation := "CREATE TABLE `transation` (
    `transationId` INTEGER NOT NULL UNIQUE,
    `walletId` INTERGER NOT NULL,
    `catogaryId` INTERGER NOT NULL,
    `timestamp` TIMESTAMP/DATETIME NOT NULL,
    `isExpense` INTERGER NOT NULL,
    `amount` INTERGER NOT NULL,
    `currency` TEXT NOT NULL,
    `note` TEXT,
    `tag` TEXT,
    PRIMARY KEY(transationId));"
) */

type Wallet struct {
  WalletId sql.NullInt64
  Name sql.NullString
  Currency sql.NullString
  InitialBalance sql.NullInt64
  //currentBalance sql.NullInt64
}

type Catogary struct {
  CatogaryId sql.NullInt64
  WalletId sql.NullInt64
  Name sql.NullString
  IsExpenseType sql.NullBool
}

type Transation struct {
  TransationId sql.NullInt64
  WalletId sql.NullInt64
  CatogaryId sql.NullInt64
  Timestamp time.Time
  IsExpense sql.NullBool
  Amount sql.NullInt64
  Currency sql.NullString
  Note sql.NullString
  Tag sql.NullString
 }

/*func setupDB() error {
  var DbDriver
  sql.Register(DbDriver, &sqlite3.SQLiteDriver{})
  database, err := sql.Open(DbDriver, "DbFile")
  if err != nil {
    fmt.Println("Failed to create the handle")
  }
  if err2 := database.Ping(); err2 != nil {
    fmt.Println("Failed to keep connection alive")
  }
} */

func main() {
  /*valueWallet := Wallet{
    WalletId : 1,
    Name : "My",
    Currency : "INR",
    InitialBalance : 5000,
  } */

  CreateWallet := "CREATE TABLE IF NOT EXISTS `wallet` (`walletId`	INTEGER NOT NULL UNIQUE,`name`	TEXT NOT NULL,`currency`	TEXT NOT NULL,`initialBalance`	INTEGER,PRIMARY KEY(walletId));"
  CreateCatogary := "CREATE TABLE IF NOT EXISTS `catogary` (`name`	TEXT NOT NULL,`catogaryId`	INTEGER NOT NULL UNIQUE,`walletId` INTEGER NOT NULL,`isExpenseType`	INTEGER NOT NULL,PRIMARY KEY(catogaryId));"
  CreateTransation := "CREATE TABLE IF NOT EXISTS `transation` (`transationId` INTEGER NOT NULL UNIQUE,`walletId` INTERGER NOT NULL,`catogaryId` INTERGER NOT NULL,`timestamp` DATETIME NOT NULL,`isExpense` INTERGER NOT NULL,`amount` INTERGER NOT NULL, `currency` TEXT NOT NULL,`note` TEXT,`tag` TEXT,PRIMARY KEY(transationId));"
  WalletInsert := "INSERT INTO wallet (walletId, name, currency, initialBalance) VALUES (?,?,?,?)"
  var DbDriver string
  sql.Register(DbDriver, &sqlite3.SQLiteDriver{})
  database, err := sql.Open(DbDriver, "DbFile")
  if err != nil {
    fmt.Println("Failed to create the handle")
  }
  if err1 := database.Ping(); err1 != nil {
    fmt.Println("Failed to keep connection alive")
  }
  _, err2 := database.Begin()
  if err2 != nil {
    log.Fatal(err2, "err2")
  }
  result, err3 := database.Exec(
    CreateWallet,
  )
  if err3 != nil {
    log.Fatal(err3,"err3")
  }
  result1, err4 := database.Exec(
    CreateCatogary,
  )
  if err4 != nil {
    log.Fatal(err4,"err4")
  }
  result2, err5 := database.Exec(
    CreateTransation,
  )
  if err5 != nil {
    log.Fatal(err5,"err5")
  }
  result3, err6 := database.Exec(
    WalletInsert,
    1,
    "My",
    "INR",
    50000,
  )
  if err6 != nil {
    log.Fatal(err5,"err6")
  }


  fmt.Println(result, result1, result2, result3)

}
