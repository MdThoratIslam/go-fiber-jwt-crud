### Go Fiber JWT Middleware সহ MVC Pattern ব্যবহার করে CRUD অপারেশন তৈরি করুন (Bangla Guide)

আমরা Go Fiber ব্যবহার করে MVC (Model-View-Controller) আর্কিটেকচার ফলো করে একটি CRUD অ্যাপ তৈরি করবো, যেখানে JWT Authentication থাকবে।
প্রজেক্টটি Router, Controller, Model এবং Middleware আলাদা করে ম্যানেজ করা হবে।

    📌 প্রজেক্ট স্ট্রাকচার
    প্রজেক্টের ফাইল গুলো নিম্নরূপ হবে:
```bash
go-fiber-jwt-crud/
│── main.go
│── go.mod
│── database/
│   ├── database.go
│── routes/
│   ├── routes.go
│── controllers/
│   ├── authController.go
│   ├── userController.go
│── models/
│   ├── user.go
│── middleware/
│   ├── jwtMiddleware.go
```

### স্টেপ ১: নতুন Go Fiber প্রজেক্ট সেটআপ করা
প্রথমে নতুন প্রজেক্ট তৈরি করুন:

```bash
mkdir go-fiber-jwt-crud
cd go-fiber-jwt-crud
go mod init go-fiber-jwt-crud
```
### Fiber এবং অন্যান্য প্যাকেজ ইনস্টল করুন:
```bash 
go get github.com/gofiber/fiber/v2
go get github.com/gofiber/jwt/v3
go get github.com/golang-jwt/jwt/v4
go get gorm.io/driver/mysql
go get gorm.io/gorm
```

go get github.com/gofiber/fiber/v2 কমান্ডটি Go প্রোগ্রামিং ভাষার Go Modules ব্যবস্থার মাধ্যমে Fiber v2 লাইব্রেরি ইনস্টল করতে ব্যবহৃত হয়।
ব্যাখ্যা:

    go get → Go মডিউল ব্যবস্থার মাধ্যমে নির্দিষ্ট লাইব্রেরি ডাউনলোড ও ইনস্টল করার জন্য ব্যবহৃত হয়।
    github.com/gofiber/fiber/v2 → এটি Fiber v2 প্যাকেজের সম্পূর্ণ import path, যা GitHub থেকে Fiber লাইব্রেরি ডাউনলোড ও ইনস্টল করে।

Fiber কী?

Fiber হল Go ভাষায় লেখা একটি দ্রুত, লাইটওয়েট এবং উচ্চ-পারফরম্যান্স ওয়েব ফ্রেমওয়ার্ক, যা Express.js এর মতো API ডিজাইন ফলো করে। এটি fasthttp ইঞ্জিনের উপর ভিত্তি করে তৈরি হয়েছে, যা এটিকে খুব দ্রুত করে তোলে।
ব্যবহার:

এই কমান্ড চালানোর পর, Fiber লাইব্রেরিটি Go Modules (যদি go.mod ফাইল থাকে) এ অন্তর্ভুক্ত হয়ে যাবে এবং আপনি এটি ব্যবহার করতে পারবেন:

===================================================================================================
### go get github.com/gofiber/jwt/v3 কমান্ডের ব্যাখ্যা:

এই কমান্ডটি Go Fiber JWT (JSON Web Token) Middleware লাইব্রেরি ইনস্টল করার জন্য ব্যবহৃত হয়। এটি Fiber ফ্রেমওয়ার্কের জন্য JWT authentication middleware সরবরাহ করে, যা অ্যাপ্লিকেশনের সিকিউরিটি নিশ্চিত করতে সাহায্য করে।
কমান্ড বিশ্লেষণ:

    go get → Go মডিউল ব্যবস্থার মাধ্যমে নির্দিষ্ট লাইব্রেরি ডাউনলোড ও ইনস্টল করার জন্য ব্যবহৃত হয়।
    github.com/gofiber/jwt/v3 → এটি Fiber-এর JWT middleware লাইব্রেরির সম্পূর্ণ import path, যা GitHub থেকে ডাউনলোড ও ইনস্টল হয়।

এই লাইব্রেরি কেন দরকার?

Fiber-এর JWT middleware টোকেন-ভিত্তিক অথেনটিকেশন ব্যবস্থাপনা সহজ করে তোলে। এটি JWT টোকেন যাচাই (verify) করতে ব্যবহার করা হয়, যাতে API endpoints নিরাপদ থাকে এবং শুধুমাত্র অথেনটিকেটেড ব্যবহারকারীরা নির্দিষ্ট route বা resource অ্যাক্সেস করতে পারে।
ব্যবহার:

ইনস্টল করার পরে, আপনি Fiber অ্যাপে এটি ব্যবহার করতে পারেন:

===================================================================================================
### go get github.com/golang-jwt/jwt/v4 কমান্ডের ব্যাখ্যা:

এই কমান্ডটি Go JWT (JSON Web Token) লাইব্রেরি ইনস্টল করার জন্য ব্যবহৃত হয়। এটি Go প্রোগ্রামিং ভাষায় JWT টোকেন তৈরি এবং যাচাই করার জন্য ব্যবহার করা হয়।

কমান্ড বিশ্লেষণ:

    go get → Go মডিউল ব্যবস্থার মাধ্যমে নির্দিষ্ট লাইব্রেরি ডাউনলোড ও ইনস্টল করার জন্য ব্যবহৃত হয়।
    github.com/golang-jwt/jwt/v4 → এটি Go JWT লাইব্রেরির সম্পূর্ণ import path, যা GitHub থেকে ডাউনলোড ও ইন্সটল হয়। 

এই লাইব্রেরি কেন দরকার?
 
এই লাইব্রেরি ব্যবহার করে আপনি Go প্রোগ্রামিং ভাষায় JWT টোকেন তৈরি এবং যাচাই করতে পারেন। এটি একটি সুরক্ষিত এবং স্থিতিশীল প্রক্রিয়া সরবরাহ করে যাতে আপনি আপনার অ্যাপ্লিকেশনের সিকিউরিটি বাড়াতে পারেন।

===================================================================================================

### go get gorm.io/driver/mysql কমান্ডের ব্যাখ্যা:

এই কমান্ডটি Go GORM (Go Object Relational Mapping) লাইব্রেরির MySQL ড্রাইভার ইনস্টল করার জন্য ব্যবহৃত হয়। এটি Go প্রোগ্রামিং ভাষায় MySQL ডেটাবেসের সাথে সম্পর্ক স্থাপন করার জন্য ব্যবহার করা হয়।

কমান্ড বিশ্লেষণ:

    go get → Go মডিউল ব্যবস্থার মাধ্যমে নির্দিষ্ট লাইব্রেরি ডাউনলোড ও ইনস্টল করার জন্য ব্যবহৃত হয়।
    gorm.io/driver/mysql → এটি Go GORM লাইব্রেরির MySQL ড্রাইভারের সম্পূর্ণ import path, যা GitHub থেকে ডাউনলোড ও ইনস্টল হয়।

এই লাইব্রেরি কেন দরকার?

এই লাইব্রেরি ব্যবহার করে আপনি Go প্রোগ্রামিং ভাষায় MySQL ডেটাবেসের সাথে সম্পর্ক স্থাপন করতে পারেন। এটি একটি সুরক্ষিত এবং স্থিতিশীল প্রক্রিয়া সরবরাহ করে যাতে আপনি আপনার অ্যাপ্লিকেশনের ডেটা স্থায়ীভাবে সংরক্ষণ করতে পারেন।

================================================================================================

### go get gorm.io/gorm কমান্ডের ব্যাখ্যা:

এই কমান্ডটি Go GORM (Go Object Relational Mapping) লাইব্রেরি ইনস্টল করার জন্য ব্যবহৃত হয়। এটি Go প্রোগ্রামিং ভাষায় ডেটাবেস সাথে সম্পর্ক স্থাপন করার জন্য ব্যবহার করা হয়।

কমান্ড বিশ্লেষণ:

    go get → Go মডিউল ব্যবস্থার মাধ্যমে নির্দিষ্ট লাইব্রেরি ডাউনলোড ও ইনস্টল করার জন্য ব্যবহৃত হয়।
    gorm.io/gorm → এটি Go GORM লাইব্রেরির সম্পূর্ণ import path, যা GitHub থেকে ডাউনলোড ও ইন্সটল হয়।

এই লাইব্রেরি কেন দরকার?

এই লাইব্রেরি ব্যবহার করে আপনি Go প্রোগ্রামিং ভাষায় ডেটাবেস সাথে সম্পর্ক স্থাপন করতে পারেন। এটি একটি সুরক্ষিত এবং স্থিতিশীল প্রক্রিয়া সরবরাহ করে যাতে আপনি আপনার অ্যাপ্লিকেশনের ডেটা স্থায়ীভাবে সংরক্ষণ করতে পারেন।

================================================================================================

### স্টেপ ২: ডাটাবেস সেটআপ করা
প্রথমে ডাটাবেস তৈরি করুন এবং এটির সাথে সংযোগ করার জন্য একটি ইউজার তৈরি করুন। এখানে আমরা মাইক্রোসফট স্কিউল সার্ভার ম্যানেজমেন্ট স্টুডিও (MS SQL Server Management Studio) ব্যবহার করবো।

প্রথমে একটি ডাটাবেস তৈরি করুন এবং এটির নাম সেট করুন। এখানে আমরা ডাটাবেসের নাম হবে go_fiber_jwt_crud।

এখন একটি নতুন ইউজার তৈরি করুন এবং এটির জন্য সম্পূর্ণ অ্যাক্সেস প্রদান করুন। এখানে আমরা ইউজারের নাম হবে MdThoratIslam 
এবং পাসওয়ার্ড হবে 123456।

এখন এই ইউজারকে go_fiber_jwt_crud ডাটাবেসের জন্য সম্পূর্ণ অ্যাক্সেস প্রদান করুন।

এখন ডাটাবেস সেটআপ সম্পন্ন হয়েছে।

### স্টেপ ৩: ডাটাবেস সংযোগ করা

ডাটাবেস সংযোগ করার জন্য database/database.go ফাইলটি তৈরি করুন এবং নিম্নলিখিত কোডটি যোগ করুন:

```go
package database
import (
	"fmt"
	"go-fiber-mysql-jwt-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)
/*
package config: এই কোডটি config নামে একটি প্যাকেজের মধ্যে রাখা হয়েছে।
import: প্রয়োজনীয় লাইব্রেরি ইম্পোর্ট করা হয়েছে।
    fmt: টার্মিনালে লোগ মেসেজ প্রিন্ট করার জন্য।
    models: যেখানে User মডেল সংরক্ষিত আছে।
    gorm.io/driver/mysql: GORM এর জন্য MySQL ড্রাইভার।
    gorm.io/gorm: GORM লাইব্রেরি।
	log: লগ মেসেজ প্রিন্ট করার জন্য।
*/

var DB *gorm.DB //DB হলো গ্লোবাল ভেরিয়েবল, যা GORM এর মাধ্যমে ডাটাবেস সংযোগ রাখবে।

func ConnectDB() {
	// ✅ Use correct MySQL DSN format
	dsn := "root:@tcp(127.0.0.1:3306)/db_golang?charset=utf8mb4&parseTime=True&loc=Local"
	/*
		dsn (Data Source Name) হল MySQL কানেকশন স্ট্রিং, যেখানে:

		    root → MySQL ইউজার।
		    @tcp(127.0.0.1:3306) → লোকালহোস্ট (localhost) IP ও পোর্ট (3306)।
		    /db_golang → ডাটাবেসের নাম db_golang।
		    charset=utf8mb4 → ইউনিকোড সমর্থনের জন্য।
		    parseTime=True → টাইমস্ট্যাম্প ফরম্যাটিং ঠিক রাখতে।
		    loc=Local → লোকাল টাইমজোন সেট করতে।
	*/

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// ❌ যদি সংযোগ ব্যর্থ হয়, তাহলে panic দিয়ে প্রোগ্রাম থামিয়ে দাও
		panic("Failed to connect to the database: " + err.Error())
	}
	/*
		gorm.Open() → MySQL ডাটাবেস সংযোগ স্থাপন করা হচ্ছে।
		err চেক করা হচ্ছে, যদি সংযোগ ব্যর্থ হয়, তাহলে panic ফেলে প্রোগ্রাম থামিয়ে দেয়া হবে।
	*/

	fmt.Println("✅ Database connected successfully!")
	/*সংযোগ সফল হলে "✅ Database connected successfully!" মেসেজ প্রিন্ট করবে।*/

	// Auto Migrate the User model
	DB.AutoMigrate(&models.User{})
	fmt.Println("✅ Database migration completed!")
	/*
		AutoMigrate() ফাংশন models.User{} কে অটো-মাইগ্রেট করবে।
		যদি users টেবিল না থাকে, তাহলে এটি স্বয়ংক্রিয়ভাবে তৈরি করবে।
		সফল হলে "✅ Database migration completed!" প্রিন্ট করবে।
	*/
}
```
👉 MySQL সংযোগ সেটআপ করার জন্য dsn অনুযায়ী নিজের তথ্য পরিবর্তন করুন।
👉 MySQL ডাটাবেস তৈরি করুন:
```sql
CREATE DATABASE db_golang;
```
### 📌 স্টেপ ৩: User Model তৈরি করা

models/user.go
```go
package models

// User struct
type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
```
    package models → এটি নির্দেশ করে যে এই ফাইলটি models নামক প্যাকেজের অন্তর্গত।
    সাধারণত models প্যাকেজ ডাটাবেজ মডেল সংজ্ঞায়িত করতে ব্যবহৃত হয়।

// User struct
type User struct {

    type User struct {} → এটি User নামে একটি struct (structure) তৈরি করেছে।
    struct হল Go-তে একটি custom data type, যা বিভিন্ন field ধারণ করতে পারে।

স্ট্রাকচারের ফিল্ড ব্যাখ্যা:

	ID       uint   `json:"id"`

    ID uint → ID একটি unsigned integer (uint) টাইপের ফিল্ড, যা ব্যবহারকারীর অনন্য শনাক্তকারী (unique identifier) হিসাবে ব্যবহৃত হয়।
    json:"id" → এটি নির্দেশ করে যে JSON রেসপন্সে এই ফিল্ডের নাম id হবে।

	Name     string `json:"name"`

    Name string → Name ফিল্ডটি string টাইপের, যা ব্যবহারকারীর নাম সংরক্ষণ করবে।
    json:"name" → JSON ফরম্যাটে একে name নামে পাঠানো হবে।

	Email    string `json:"email"`

    Email string → এটি ব্যবহারকারীর ইমেইল ঠিকানা সংরক্ষণ করবে।
    json:"email" → JSON-এ এটি email নামে যাবে।

	Password string `json:"password"`

    Password string → এটি ব্যবহারকারীর পাসওয়ার্ড সংরক্ষণ করবে।
    json:"password" → JSON-এ এটি password নামে পাঠানো হবে।
    ⚠ সতর্কতা: প্রকৃত অ্যাপ্লিকেশনে কখনোই পাসওয়ার্ড plain text আকারে সংরক্ষণ করা উচিত নয়। বরং hashed (এনক্রিপ্টেড) করে সংরক্ষণ করা উচিৎ।

	Age      int    `json:"age"`

    Age int → এটি ব্যবহারকারীর বয়স (age) সংরক্ষণ করবে।
    json:"age" → JSON-এ এটি age নামে পাঠানো হবে।

পুরো কোডের সংক্ষিপ্ত ব্যাখ্যা:

    User নামের একটি struct তৈরি করা হয়েছে, যাতে ব্যবহারকারীর বিভিন্ন তথ্য সংরক্ষণ করা যায়।
    এটি json ট্যাগ ব্যবহার করছে, যা নির্দেশ করে কিভাবে JSON ডাটা রূপান্তরিত হবে।
    ডাটাবেজ মডেল বা API রেসপন্সের জন্য User struct সাধারণত ব্যবহৃত হয়।

JSON আউটপুট কেমন হবে?

যদি আমরা User struct এর একটি instance তৈরি করে JSON রূপান্তর করি, তাহলে এটি নিম্নলিখিত আউটপুট দেবে:

{
"id": 1,
"name": "Masud",
"email": "masud@example.com",
"password": "hashed_password",
"age": 25
}

⚠ নোট: প্রকৃত ব্যবহারে password ফিল্ড এনক্রিপ্ট করে সংরক্ষণ করা উচিত।

### 📌 স্টেপ ৪: JWT Middleware তৈরি করা

```go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v3"
)

var SecretKey = []byte("your_secret_key")

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: SecretKey},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}
```
    package middleware → এটি নির্দেশ করে যে এই কোড middleware নামক একটি প্যাকেজের অংশ।
    Middleware হল এমন একটি ফাংশন, যা অনুরোধ (request) প্রক্রিয়াকরণ চলাকালীন authentication, authorization, logging ইত্যাদির জন্য ব্যবহৃত হয়।

📌 লাইব্রেরি ইম্পোর্ট:

    import ("github.com/gofiber/fiber/v2"
    "github.com/gofiber/jwt/v3"
    )

github.com/gofiber/fiber/v2 → Fiber ফ্রেমওয়ার্ক ইম্পোর্ট করা হয়েছে।
github.com/gofiber/jwt/v3 → Fiber এর JWT Middleware ইম্পোর্ট করা হয়েছে, যা JWT টোকেন যাচাই (verify) করতে ব্যবহৃত হয়।

📌 সিক্রেট কী ডিফাইন করা হয়েছে:

    var SecretKey = []byte("your_secret_key")

SecretKey → এটি JWT টোকেন স্বাক্ষর করার জন্য ব্যবহৃত একটি গোপন চাবি (secret key)।
এই key ব্যবহার করে JWT token encode এবং decode করা হয়।
প্রকৃত প্রজেক্টে এটি .env ফাইলে সংরক্ষণ করা উচিত, সরাসরি কোডে রাখা উচিত নয়।

📌 Middleware ফাংশন (JWTMiddleware):

    func JWTMiddleware() fiber.Handler {

JWTMiddleware() → এটি একটি middleware function, যা Fiber-এর Handler (middleware) রিটার্ন করে।
এই middleware ব্যবহার করলে নির্দিষ্ট routes-এ JWT অথেনটিকেশন চেক করা হবে।

📌 JWT Middleware কনফিগারেশন:

    return jwtware.New(jwtware.Config{
    SigningKey: jwtware.SigningKey{Key: SecretKey},
    jwtware.New(jwtware.Config{}) → এটি একটি JWT middleware তৈরি করে।
    SigningKey: jwtware.SigningKey{Key: SecretKey}
এটি SecretKey ব্যবহার করে JWT টোকেন যাচাই (verify) করবে।
যদি JWT টোকেন বৈধ হয়, তাহলে রিকোয়েস্ট স্বাভাবিকভাবে প্রসেস হবে।
যদি টোকেন ভুল হয় বা অনুপস্থিত থাকে, তাহলে এটি Unauthorized (401) স্ট্যাটাস পাঠাবে।

📌 ত্রুটি হ্যান্ডলার (ErrorHandler)

    ErrorHandler: func(c *fiber.Ctx, err error) error {
    return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
    "error": "Unauthorized",
    })
    },

ErrorHandler → যদি JWT টোকেন অবৈধ হয়, তাহলে 401 Unauthorized রেসপন্স ফেরত দেবে।
এটি JSON ফরম্যাটে error message পাঠাবে:

### 📌 স্টেপ ৫: Authentication Controller (JWT Login & Register)

    📌 controllers/authController.go
```go
package controllers

import (
"time"
"go-fiber-jwt-crud/database"
"go-fiber-jwt-crud/models"
"go-fiber-jwt-crud/middleware"
"github.com/gofiber/fiber/v2"
"github.com/golang-jwt/jwt/v4"
"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
        return err
    }
	// Hash Password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	
	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(hashedPassword),
	}

	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
    var data map[string]string
    if err := c.BodyParser(&data); err != nil {
    return err
    }
	var user models.User
	database.DB.Where("email = ?", data["email"]).First(&user)
	// Validate Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Credentials"})
	}
	// Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	t, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{"token": t})
}
```
### 📌 স্টেপ ৬: User CRUD Controller তৈরি করা

    📌 controllers/userController.go
```go
package controllers

import (
"go-fiber-jwt-crud/database"
"go-fiber-jwt-crud/models"
"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
    var users []models.User
    database.DB.Find(&users)
    return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    database.DB.First(&user, id)
    return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    database.DB.First(&user, id)
    if err := c.BodyParser(&user); err != nil {
        return err
    }
	database.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")
    database.DB.Delete(&models.User{}, id)
    return c.SendString("User Deleted")
}
```

    📌 কোড বিশ্লেষণ:
    প্যাকেজ ডিক্লারেশন ও লাইব্রেরি ইম্পোর্ট
```go
package controllers

import (
"go-fiber-mysql-jwt-crud/database"
"go-fiber-mysql-jwt-crud/models"
"github.com/gofiber/fiber/v2"
)
```

package controllers → এই কোডটি controllers প্যাকেজের অংশ।
    
    ইম্পোর্ট করা লাইব্রেরি ও প্যাকেজ:
        go-fiber-mysql-jwt-crud/database → MySQL ডাটাবেজ সংযোগ সংরক্ষণ করে।
        go-fiber-mysql-jwt-crud/models → User মডেল ডিফাইন করা হয়েছে।
        github.com/gofiber/fiber/v2 → Fiber ফ্রেমওয়ার্ক ব্যবহার করা হয়েছে।

📌 ২️⃣ সকল ইউজার রিটার্ন করার ফাংশন
```go
func GetUsers(c *fiber.Ctx) error {
    var users []models.User
    database.DB.Find(&users)
    return c.JSON(users)
}
```

GetUsers(c *fiber.Ctx) error → সকল ইউজার (Users) রিটার্ন করার জন্য একটি Fiber handler function।
var users []models.User → users নামের একটি স্লাইস ডিক্লেয়ার করা হয়েছে।
database.DB.Find(&users) → MySQL ডাটাবেজ থেকে সকল ইউজার খুঁজে আনা হয়েছে।
return c.JSON(users) → JSON ফরম্যাটে রেসপন্স রিটার্ন করা হয়েছে।

✅ উদাহরণ রেসপন্স:
```json
[
    {
        "id": 1,
        "name": "Masud",
        "email": "masud@example.com",
        "age": 25
    },
    {
        "id": 2,
        "name": "Rahim",
        "email": "rahim@example.com",
        "age": 30
    }
]
```

📌 ৩️⃣ নির্দিষ্ট একজন ইউজার রিটার্ন করার ফাংশন
```go
func GetUser(c *fiber.Ctx) error {
id := c.Params("id")
var user models.User
database.DB.First(&user, id)
return c.JSON(user)
}
```

id := c.Params("id") → URL থেকে id প্যারামিটার পাওয়া হচ্ছে।
database.DB.First(&user, id) → MySQL থেকে নির্দিষ্ট ইউজার খুঁজে বের করা হচ্ছে।
return c.JSON(user) → JSON ফরম্যাটে রেসপন্স রিটার্ন করছে।

✅ উদাহরণ API কল:

    GET /users/1

✅ উদাহরণ JSON রেসপন্স:
```json
{
    "id": 1,
    "name": "Masud",
    "email": "masud@example.com",
    "age": 25
}
```
📌 ৪️⃣ ইউজার আপডেট করার ফাংশন
```go
func UpdateUser(c *fiber.Ctx) error {
    id := c.Params("id")
    var user models.User
    database.DB.First(&user, id)
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	database.DB.Save(&user)
	return c.JSON(user)
}
```
    কীভাবে কাজ করে?
        id := c.Params("id") → URL থেকে id সংগ্রহ করা হয়েছে।
        database.DB.First(&user, id) → id অনুসারে ইউজার খোঁজা হয়েছে।
        c.BodyParser(&user) → JSON ডাটা থেকে body প্যার্স করা হয়েছে।
        database.DB.Save(&user) → ডাটাবেজে ইউজার আপডেট করা হয়েছে।
        return c.JSON(user) → আপডেটেড ইউজার রেসপন্সে পাঠানো হয়েছে।

✅ উদাহরণ API কল:

    PUT /users/1
    Content-Type: application/json
```json
{
    "name": "Masud Updated",
    "email": "masud_updated@example.com",
    "age": 26
}
```

✅ JSON রেসপন্স:
```json
{
"id": 1,
"name": "Masud Updated",
"email": "masud_updated@example.com",
"age": 26
}
```
📌 ৫️⃣ ইউজার ডিলিট করার ফাংশন
```go
func DeleteUser(c *fiber.Ctx) error {
    id := c.Params("id")
    database.DB.Delete(&models.User{}, id)
    return c.SendString("User Deleted")
}
```

    id := c.Params("id") → URL থেকে id সংগ্রহ করা হয়েছে।
    database.DB.Delete(&models.User{}, id) → MySQL থেকে নির্দিষ্ট ইউজার ডিলিট করা হয়েছে।
    return c.SendString("User Deleted") → টেক্সট রেসপন্স পাঠানো হয়েছে।

✅ উদাহরণ API কল:

    DELETE /users/1

✅ রেসপন্স:
    User Deleted
📌 সম্পূর্ণ কোড সংক্ষেপে

<table>
    <tr>
        <th>ফাংশন</th>
        <th>উদ্দেশ্য</th>
        <th>HTTP মেথড</th>
        <th>এন্ডপয়েন্ট</th>
    </tr>
    <tr>
        <td>GetUsers</td>
        <td>সকল ইউজার দেখানো</td>
        <td>GET</td>
        <td>/users</td>
    </tr>
    <tr>
        <td>GetUser</td>
        <td>নির্দিষ্ট ইউজার দেখানো</td>
        <td>GET</td>
        <td>/users</td>
    </tr>
    <tr>
        <td>GetUsers</td>
        <td>সকল ইউজার দেখানো</td>
        <td>GET</td>
        <td>/users/:id</td>
    </tr>
    <tr>
        <td>UpdateUser</td>
        <td>ইউজার আপডেট করা</td>
        <td>PUT</td>
        <td>/users/:id</td>
    </tr>
    <tr>
        <td>DeleteUser</td>
        <td>ইউজার ডিলিট করা</td>
        <td>DELETE</td>
        <td>/users/:id</td>
    </tr>
</table>

### 📌 স্টেপ ৭: রাউট তৈরি করা

    📌 routes/routes.go
```go
package routes

import (
	"go-fiber-jwt-crud/controllers"
	"go-fiber-jwt-crud/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Authentication Routes
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	// User CRUD Routes (JWT Protected)
	api := app.Group("/api", middleware.JWTMiddleware())

	api.Get("/users", controllers.GetUsers)
	api.Get("/users/:id", controllers.GetUser)
	api.Put("/users/:id", controllers.UpdateUser)
	api.Delete("/users/:id", controllers.DeleteUser)
}
```
    📌 কোড বিশ্লেষণ:
    ইম্পোর্ট করা লাইব্রেরি ও প্যাকেজ
```go
package routes

import (
"go-fiber-jwt-crud/controllers"
"go-fiber-jwt-crud/middleware"
"github.com/gofiber/fiber/v2"
)
```

package routes → এই কোডটি routes প্যাকেজের অংশ।
    
    ইম্পোর্ট করা লাইব্রেরি ও প্যাকেজ:
        go-fiber-jwt-crud/controllers → কন্ট্রোলার ফাংশন ইম্পোর্ট করা হয়েছে।
        go-fiber-jwt-crud/middleware → JWT মিডলওয়্যার ইম্পোর্ট করা হয়েছে।
        github.com/gofiber/fiber/v2 → Fiber ফ্রেমওয়ার্ক ইম্পোর্ট করা হয়েছে।

📌 রাউট সেটআপ ফাংশন (SetupRoutes):
```go

func SetupRoutes(app *fiber.App) {
    // Authentication Routes
    app.Post("/register", controllers.Register)
    app.Post("/login", controllers.Login)

    // User CRUD Routes (JWT Protected)
    api := app.Group("/api", middleware.JWTMiddleware())

    api.Get("/users", controllers.GetUsers)
    api.Get("/users/:id", controllers.GetUser)
    api.Put("/users/:id", controllers.UpdateUser)
    api.Delete("/users/:id", controllers.DeleteUser)
}
```

    📌 কীভাবে কাজ করে?
        app.Post("/register", controllers.Register) → নতুন ইউজার রেজিস্টার করার জন্য রেজিস্টার রাউট সেট করা হয়েছে।
        app.Post("/login", controllers.Login) → ইউজার লগইন করার জন্য লগইন রাউট সেট করা হয়েছে।
        api := app.Group("/api", middleware.JWTMiddleware()) → JWT মিডলওয়্যার সহ একটি গ্রুপ তৈরি করা হয়েছে।
        api.Get("/users", controllers.GetUsers) → সকল ইউজার দেখানোর জন্য রাউট সেট করা হয়েছে।
        api.Get("/users/:id", controllers.GetUser) → নির্দিষ্ট ইউজার দেখানোর জন্য রাউট সেট করা হয়েছে।
        api.Put("/users/:id", controllers.UpdateUser) → ইউজার আপডেট করার জন্য রাউট সেট করা হয়েছে।
        api.Delete("/users/:id", controllers.DeleteUser) → ইউজার ডিলিট করার জন্য রাউট সেট করা হয়েছে।

📌 রাউট সেটআপ করা হয়েছে, এখন এই রাউটগুলো এপ্লিকেশনে যোগ করতে হবে।

### 📌 স্টেপ ৮: এপ্লিকেশন চালু করা

    📌 main.go
```go
package main

import (
	"go-fiber-jwt-crud/database"
	"go-fiber-jwt-crud/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to Database
	database.ConnectDB()

	// Setup Routes
	routes.SetupRoutes(app)

	// Start Server
	app.Listen(":3000")
}
```
    📌 কোড বিশ্লেষণ:
    ইম্পোর্ট করা লাইব্রেরি ও প্যাকেজ
```go
package main

import (
"go-fiber-jwt-crud/database"
"go-fiber-jwt-crud/routes"
"github.com/gofiber/fiber/v2"
)
```

package main → এই কোডটি main প্যাকেজের অংশ।
    
    ইম্পোর্ট করা লাইব্রেরি ও প্যাকেজ:
        go-fiber-jwt-crud/database → MySQL ডাটাবেজ সংযোগ স্থাপন করা হয়েছে।
        go-fiber-jwt-crud/routes → রাউট সেটআপ করা হয়েছে।
        github.com/gofiber/fiber/v2 → Fiber ফ্রেমওয়ার্ক ইম্পোর্ট করা হয়েছে। 

📌 মেইন ফাংশন (main):
```go

func main() {
    app := fiber.New()

    // Connect to Database
    database.ConnectDB()

    // Setup Routes
    routes.SetupRoutes(app)

    // Start Server
    app.Listen(":3000")
}
```

    📌 কীভাবে কাজ করে?
        app := fiber.New() → নতুন Fiber এপ্লিকেশন তৈরি করা হয়েছে।
        database.ConnectDB() → MySQL ডাটাবেজে সংযোগ স্থাপন করা হয়েছে।
        routes.SetupRoutes(app) → রাউট সেটআপ করা হয়েছে।
        app.Listen(":3000") → সার্ভার চালু করা হয়েছে পোর্ট 3000 এ।

📌 এপ্লিকেশন চালু করা হয়েছে, এখন এই এপ্লিকেশন চালু করতে হবে।

### 📌 স্টেপ ৯: এপ্লিকেশন চালু করা

    টার্মিনালে নিচের কমান্ড দিন:
```bash

go run main.go
```
    এপ্লিকেশন চালু হলে নিচের মত মেসেজ দেখা যাবে:
```bash
🚀 Fiber is running on :3000
```
    এপ্লিকেশন সফলভাবে চালু হয়েছে। এখন আমরা এপ্লিকেশনে রিকোয়েস্ট পাঠাতে পারি।

### 📌 স্টেপ ১০: এপ্লিকেশন টেস্ট করা

    এপ্লিকেশন চালু করার পর, আমরা এপ্লিকেশনে রিকোয়েস্ট পাঠাতে পারি।
    এপ্লিকেশনে রিকোয়েস্ট পাঠানোর জন্য আমরা পোস্টম্যান, ইনসম্যান, পুটম্যান ইত্যাদি ধরণের টুল ব্যবহার করতে পারি।
    এখানে আমরা পোস্টম্যান টুল ব্যবহার করব।

#### আপনার Go Fiber + MySQL + JWT Authentication + CRUD প্রজেক্ট চালানোর সময় golang.org/x/crypto/bcrypt প্যাকেজ মিসিং 
দেখাচ্ছে।
আপনার go.sum ফাইলটি আপডেট করার জন্য নিচের কমান্ড রান করুন:

go get golang.org/x/crypto
go mod tidy

📌 go mod tidy কি করে?

go mod tidy কমান্ডটি go.mod এবং go.sum ফাইলগুলো আপডেট করে এবং প্রয়োজনীয় সব dependencies যোগ করে।


##### Error Solving

📌 সমাধান ১: log ইম্পোর্ট করা হয়েছে কিন্তু ব্যবহার করা হয়নি

এই error মূলত database/database.go ফাইলে log প্যাকেজ ইম্পোর্ট করা হয়েছে কিন্তু কোথাও ব্যবহার করা হয়নি।
এটা ঠিক করতে fmt.Println ব্যবহার করুন বা log.Println ব্যবহার করুন।

📌 database/database.go (সংশোধিত)
```go
package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"go-fiber-jwt-crud/models"
)

// গ্লোবাল ডাটাবেস ভেরিয়েবল
var DB *gorm.DB

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.1:3306)/fiber_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("⛔ Database Connection Failed!")
		return
	}

	fmt.Println("✅ Connected to MySQL Database!")

	// মডেল অটো-মাইগ্রেশন
	db.AutoMigrate(&models.User{})

	DB = db
}
```
👉 log প্যাকেজ রিমুভ করে fmt.Println ব্যবহার করা হয়েছে।

##### 📌 সমাধান ২: jwtware.SigningKey আনডিফাইন্ড

এই error এসেছে কারণ jwtware.SigningKey পরিবর্তে এখন jwt.SigningKey ব্যবহার করতে হবে।

📌 middleware/jwtMiddleware.go (সংশোধিত)

```go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var SecretKey = []byte("your_secret_key")

func JWTMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: SecretKey}, // ✅ ঠিক করা হয়েছে
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
	})
}

```

👉 jwtware.SigningKey{Key: SecretKey} → পরিবর্তে jwt.SigningKey{Key: SecretKey} ব্যবহার করুন।

📌 ফাইনাল স্টেপ

এখন Go মডিউল আপডেট করুন:

```bash
go mod tidy
```

এই কমান্ডটি আপনার go.mod এবং go.sum ফাইলগুলো আপডেট করবে এবং প্রয়োজনীয় সব dependencies যোগ করবে।

তারপর প্রজেক্ট রান করুন:

```bash
go run main.go
```
