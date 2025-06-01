package main

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm/dao"
	"log"
	"math/rand"
)

// go标准库
//var godb = dao.GoDB

// sqlx
var dbx = dao.DBx

// gorm
var db = dao.DB

func gorm1() {
	// 写入新纪录
	zhangsan := dao.Student{
		Name:  "张三",
		Age:   20,
		Grade: "三年级",
	}
	db.Create(&zhangsan)

	// 查询大于18岁所有学生
	var students []dao.Student
	db.Model(&students).
		Where("age > ?", 18).
		Find(&students)
	fmt.Println(students)

	// 更新信息 age > 18
	db.Model(&dao.Student{}).
		Where("age > ?", 18).
		Update("grade", "四年级")

	// 删除信息 age < 15
	db.Where("age < ?", 15).
		Delete(&dao.Student{})
}

func transferMoney(fromID, toID uint, money float64) (string, bool) {

	var tMsg string
	var fromAccount, toAccount dao.Account

	err := db.Transaction(func(tx *gorm.DB) error {
		// 加锁查询
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Take(&fromAccount, fromID).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 加锁失败.", fromID)
			return err
		}

		// 检查余额
		if fromAccount.Balance < money {
			tMsg = fmt.Sprintf("账户: %d 余额不足.", fromID)
			return errors.New(fmt.Sprintf("用户%d 余额不足", fromAccount.ID))
		}

		// 扣款
		if err := tx.Model(&fromAccount).
			Update("balance", gorm.Expr("balance - ?", money)).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 扣款失败.", fromID)
			return err
		}

		// 转账
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Take(&toAccount, toID).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 加锁失败.", toID)
			return err
		}

		if err := tx.Model(&toAccount).
			Update("balance", gorm.Expr("balance + ?", money)).
			Error; err != nil {
			tMsg = fmt.Sprintf("账户: %d 加款失败.", toID)
			return err
		}
		return nil
	})
	if err != nil {
		return tMsg, false
	}
	return "转账成功", true
}

func gorm2() {
	// 创建用户
	//accountA := dao.Account{
	//	Balance: 100,
	//}
	//accountB := dao.Account{
	//	Balance: 80,
	//}

	//db.Create(&accountA)
	//db.Create(&accountB)
	//fmt.Println(accountA)
	//fmt.Println(accountB)

	// 转账 account2 -> account1 100  失败
	msg, ok := transferMoney(2, 1, 100)
	fmt.Println(msg, ok)

	// 转账 account1 -> account2 100  成功
	msg1, ok1 := transferMoney(1, 2, 100)
	fmt.Println(msg1, ok1)

}

type employee struct {
	ID         uint    `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func gorm3() {

	// 写入测试数据
	var tmpE []employee
	departments := []string{"技术部", "企划部", "采购部", "人事部"}

	for i := 0; i < 20; i++ {
		tmpE = append(tmpE, employee{
			Name:       fmt.Sprintf("员工%d", i),
			Department: departments[rand.Intn(len(departments))],
			Salary:     rand.Float64(),
		})
	}

	res, err := dbx.NamedExec(`
		INSERT INTO
			employees(name, department, salary)
		VALUES (:name, :department, :salary)`, tmpE)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())

	var employees []employee
	err = dbx.Select(&employees, `
		SELECT
		    id,
		    name,
		    department,
		    salary
		FROM employees
		WHERE department = "技术部"`)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(employees)

	var employeeOne employee
	errOne := dbx.Get(&employeeOne, `
		SELECT
		    id,
		    name,
		    department,
		    salary
		FROM employees
		ORDER BY 
		    salary DESC LIMIT 1`)
	if errOne != nil {
		log.Fatal(errOne)
	}
	fmt.Println(employeeOne)

}

type book struct {
	ID     uint    `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func gorm4() {
	//var insertBooks []book
	//for i := 0; i < 20; i++ {
	//	insertBooks = append(insertBooks, book{
	//		Title:  fmt.Sprintf("book%d", i),
	//		Author: fmt.Sprintf("author%d", i),
	//		Price:  float64(rand.Intn(180)),
	//	})
	//}
	//
	//res, insertErr := dbx.NamedExec(`
	//	INSERT INTO
	//		books(title, author, price)
	//	VALUES (:title, :author, :price)`, insertBooks)
	//if insertErr != nil {
	//	log.Fatal(insertErr)
	//}
	//fmt.Println(res.RowsAffected())

	var books []book
	err := dbx.Select(&books, `
		SELECT
		    *
		FROM
		    books
		WHERE 
		    price > ?`, float64(50))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(books)
}

func gorm5() {
	var users []dao.User
	for i := 0; i < 10; i++ {
		var tmpPosts []dao.Post
		for j := 0; j < rand.Intn(5); j++ {
			var tmpComments []dao.Comment
			for k := 0; k < rand.Intn(10); k++ {
				tmpComments = append(tmpComments, dao.Comment{
					Content: fmt.Sprintf("评论1%d", i),
				})
			}
			tmpPosts = append(tmpPosts, dao.Post{
				Title:    fmt.Sprintf("标题1%d", i),
				Comments: tmpComments,
			})
		}
		users = append(users, dao.User{
			Name:  fmt.Sprintf("用户%d", i),
			Posts: tmpPosts})

	}
	res := db.CreateInBatches(&users, 5)
	if res.Error != nil {
		log.Fatal(res)
	}
	fmt.Println(res.RowsAffected)
}

func gorm6() {
	var user dao.User
	db.Preload("Posts.Comments").Take(&user, 1)
	fmt.Println(user)

	var maxCount struct {
		Count  int
		PostID uint
	}
	db.Model(&dao.Comment{}).
		Select("post_id, Count(1) count").
		Group("post_id").
		Order("count DESC").
		Limit(1).
		Scan(&maxCount)
	fmt.Println(maxCount)

	var post dao.Post
	db.Preload("User").
		Preload("Comments").
		Where("id = ?", maxCount.PostID).
		Find(&post)
	fmt.Println(post)
}

func gorm7() {
	var user dao.User
	db.Take(&user, 1)

	cdb := context.WithValue(context.Background(), "WordCount", 678)
	db.WithContext(cdb).
		Set("WordCount", 678).
		Create(&dao.Post{
			Title:    "新文章1",
			UserID:   user.ID,
			Comments: []dao.Comment{},
		})

	comment := dao.Comment{
		PostID: 12,
		ID:     23,
	}

	db.Delete(&comment)
}

func main() {
	gorm1()

	gorm2()

	gorm3()

	gorm4()

	gorm5()

	gorm6()

	gorm7()

}
