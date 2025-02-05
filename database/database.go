package database

import (
	logger "go-fiber-jwt-crud/log"
	"go-fiber-jwt-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	logger.Success("✅ Database connected successfully!")
	/*সংযোগ সফল হলে "✅ Database connected successfully!" মেসেজ প্রিন্ট করবে।*/

	//drop model
	DB.Migrator().DropTable(&models.User{})
	// Auto Migrate the User model
	DB.AutoMigrate(&models.User{})
	// fmt.Println("✅ Database migration completed!")
	/*
		AutoMigrate() ফাংশন models.User{} কে অটো-মাইগ্রেট করবে।
		যদি users টেবিল না থাকে, তাহলে এটি স্বয়ংক্রিয়ভাবে তৈরি করবে।
		সফল হলে "✅ Database migration completed!" প্রিন্ট করবে।
	*/
}
