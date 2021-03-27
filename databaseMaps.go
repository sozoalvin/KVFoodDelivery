package main

import "fmt"

var MyFoodListMap = make(map[string]FoodInfo)
var MyFoodListMap2 = make(map[string]FoodInfo)

var MyCheckoutTranID = make(map[string]Checkout)     //map with transaction ID as key
var MyCheckoutIDUsername = make(map[string]Checkout) //map with username as key

// var SystemQueueDB = make(map[string]KVorder)
var SystemQueueDB = make(map[string]*SystemOrderInfo)

func CreateFoodListMap(ch chan string) {

	for _, v := range V {

		keyValue := v.FoodName + " - " + v.MerchantName + " - " + v.DetailedLocation
		MyFoodListMap[keyValue] = FoodInfo{v.FoodName, v.MerchantName, v.DetailedLocation, v.PostalCode, v.Price, v.OpeningPeriods}

	}
	ch <- "Food List Map Data Completed"

}

func CreateFoodListMap2(f *FoodInfo) {

	keyValue := f.FoodName + " - " + f.MerchantName + " - " + f.DetailedLocation
	MyFoodListMap2[keyValue] = FoodInfo{f.FoodName, f.MerchantName, f.DetailedLocation, f.PostalCode, f.Price, f.OpeningPeriods}
	fmt.Println(MyFoodListMap2)
}

var DriverDB = map[string]driverInfo{

	"driver":  {"john", "doe", "9798545", "tokyodrifter@test.com"},
	"driver1": {"xiao", "hua", "9798545", "tokyodrifter@test.com"},
	"driver2": {"xiao", "hua", "92632482", "driver2@test.com"},
	"driver3": {"leng", "lui", "92125482", "driver3@test.com"},
	"driver4": {"leng", "zai", "93455482", "driver4@test.com"},
}
