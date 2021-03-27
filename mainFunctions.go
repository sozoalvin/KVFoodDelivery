package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var FoodMerchantNameAddress []string
var FoodMerchantNameAddress1 []string
var KVorderforQueue []KVorder

func checkUsernameStructure(s string) bool {

	for _, v := range s {
		switch string(v) {

		case "0":
			fallthrough
		case "1":
			fallthrough
		case "2":
			fallthrough
		case "3":
			fallthrough
		case "4":
			fallthrough
		case "5":
			fallthrough
		case "6":
			fallthrough
		case "7":
			fallthrough
		case "8":
			fallthrough
		case "9":
			fmt.Println("\nYour username should not contain any numbers. Please try again")
			return false
		default:
			continue
		}
	}
	return true
}

func displayMainMenu() (int, error) {

	// var usrInpt int
	var usrInpt string
	fmt.Println("\nPlease select from the following menu\n")
	fmt.Println("1. Access Food Items")
	fmt.Println("2. Search and Add Items to Cart")
	fmt.Println("3. Check Merchant Postal Code Validity")
	fmt.Println("4. Check on Current Order Queue")
	fmt.Println("5. Dispatch Order")
	fmt.Println("6. Display Databases")
	fmt.Println("7. Edit and Delete System Order Information")
	fmt.Println("8. End Program")
	fmt.Scanln(&usrInpt)
	// fmt.Println("value", usrInpt)
	// if usrInpt <= 0 || usrInpt > 7 {
	// 	return -1, errors.New("Input cannot be negative or more than the number of options provided")
	// }
	matched, _ := regexp.MatchString(`^[0-8]`, usrInpt) //please change back to 7 , after you are done. menu option number 8 only enabled for testing purposes

	if !matched {
		return -1, errors.New("Input cannot be negative or more than the number of options provided")
	}
	newusrInt, _ := strconv.Atoi(usrInpt)
	return newusrInt, nil
}

func concatenateFoodList() { //return you a slice of strings
	fmt.Println(V)
}

func CreateFoodList(ch chan string) {
	// func CreateFoodList() []string {

	for _, v := range V {
		FoodMerchantNameAddress = append(FoodMerchantNameAddress, v.FoodName+" - "+v.MerchantName+" - "+v.DetailedLocation)
	}
	sort.Strings(FoodMerchantNameAddress)
	ch <- "Mandatory - Food List Data Generated"
	// return FoodMerchantNameAddress
}

// func CreateFoodList1() []string { //function for template testing.

// 	for _, v := range V_2 {
// 		FoodMerchantNameAddress1 = append(FoodMerchantNameAddress1, v.FoodName+" - "+v.MerchantName+" - "+v.DetailedLocation)
// 	}
// 	sort.Strings(FoodMerchantNameAddress1)
// 	return FoodMerchantNameAddress1
// }

func PrintSliceinLines(s []string) {
	sum := 0

	fmt.Println("You are now browsing all available items regardless of your postal code")

	for i, v := range s {
		fmt.Printf("%d. %s\n", i+1, v)
		sum++
	}
	fmt.Printf("\nThere are a total of %d food items available for order\n", sum)
	fmt.Println("You can order them in the add to cart option")
}

func Case1DisplayAllFoodItems(s []string) string {
	var usrInpt string
	fmt.Println("You are now browsing all available items regardless of your postal code")
	fmt.Println("Please use the search function on the previous menu if you need to search for something\n")
	PrintSliceinLinesFoodListGeneral(s)
	fmt.Println("\nEnter the letter Q to exit to previous menu")
	fmt.Scanln(&usrInpt)
	return usrInpt
}

func PrintSliceinLinesFoodListGeneral(s []string) {

	for i, v := range s {
		fmt.Printf("%d. %s\n", i+1, v)
	}
}

func PrintSliceinLinesGeneral(s []string) {

	for i, v := range s {
		fmt.Printf("%d. %s\n", i+1, v)
	}

}

func PrintSliceinLinesGeneralSearch(s []string) int {

	for i, v := range s {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	return len(s)
}

func ToQuit(s string) bool { //use this function when u want to check for input = q . usually used to return to main menu

	r := strings.ToLower(s)
	if r == "q" {
		return true
	}
	return false
}

func Case2DisplayAllSearchAndATC() (string, int, error) {

	var usrInpt string
	var usrInpt2 int

	fmt.Println("\nPlease Search for an Item(lower case alphabets, including numbers 0-9)\n")
	fmt.Scanln(&usrInpt)
	fmt.Println("\nPlease Indicate the Number of Similar Results You Want to Display\n")
	fmt.Scanln(&usrInpt2)
	if usrInpt2 <= 0 {
		return "", -1, errors.New("Number of Search Results to Display Should be a positive number. i.e. 10 ")
	}
	return usrInpt, usrInpt2, nil

}

func AddToCart(s []string, un string) { // s is the search result
	var usrInptChoice string
	var usrInptQty string
	fmt.Println("\nPlease Indicate the item you would like to place an order for:")
	for {
		fmt.Scanln(&usrInptChoice)
		x, _ := strconv.Atoi(usrInptChoice)
		if x <= 0 || x > len(s) {
			fmt.Println("\nInvalid Input. Please try again. Please enter your choice again")
			// return
		} else {
			break
		}
	}
	fmt.Println("\nPlease Indicate the quantity of the item you would like to place an order for:")
	fmt.Scanln(&usrInptQty)
	MatchUsrInptToSlice(s, usrInptChoice, usrInptQty, un) //item Properties should return you item FULL NAME and price with reference to map values

}

func MatchUsrInptToSlice(s []string, s1 string, s2 string, un string) {

	// noOfChoices := len(s) //if search results only have 2 elements; this returns 2.
	s1n, _ := strconv.Atoi(s1)
	s1nn := s1n - 1
	// fmt.Println(s1nn)
	s2n, _ := strconv.ParseFloat(s2, 32)

	// var usrInpt int

	fmt.Println("\n===================================================\n")
	fmt.Printf("\nThis Order Will be Tagged to User: %s. Please exit program to relogin if your login ID is incorrect\n", un)
	fmt.Printf("\nYour Option: %s Has Been Selected\n\n", s[s1nn])
	fmt.Printf("The Name of the Selected Food Item is: %s\n", MyFoodListMap[s[s1nn]].FoodName)
	fmt.Printf("\nThe Total Quantity Ordered: %0.0f\n", s2n)
	fmt.Printf("\nThe Total Cost for this is: $%0.2f\n", s2n*MyFoodListMap[s[s1nn]].Price)
	fmt.Println("\nPlease select save shopping cart if you want to checkout the above items. ")

	fmt.Println("\n===================================================\n")

	fmt.Println("Please choose from the following options\n")
	SearchSaveCheckOut(un, s[s1nn], s2n, s2n*MyFoodListMap[s[s1nn]].Price, MyFoodListMap[s[s1nn]].MerchantName, MyFoodListMap[s[s1nn]].FoodName) //passes username, the complete food dish, the total quantity, the price, the merchant name, the food name

	// return "", 0
}

func SearchSaveCheckOut(un string, s string, d float64, f float64, s2 string, s3 string) { //passes username, the complete food dish, the total quantity, the price, the merchant name, the food name

	var usrInpt int

	fmt.Println("1. Search for new item to add to cart")
	fmt.Println("2. Update Shopping Cart")
	fmt.Println("3. Clear Shopping Cart")
	fmt.Println("4. Checkout Shopping Cart")
	fmt.Println("5. Return to Previous Menu")
	fmt.Scanln(&usrInpt)

	switch usrInpt {
	case 1:
		case2rseultstring, case2rseultint, errorSearchandATC := Case2DisplayAllSearchAndATC()
		if errorSearchandATC != nil {
			fmt.Print("System Error:", errorSearchandATC)
		} else {
			case2SearchResults := MyFoodListDB.GetSuggestion(case2rseultstring, case2rseultint)
			// PrintKeywordSearchResults(case2SearchResults, un)
			// AddToCart(case2SearchResults, un)
			editedSearch := PrintKeywordSearchResults(case2SearchResults, un)
			AddToCart(editedSearch, un)
		}
		break

	case 2:
		AddOrdertoCart(un, s, d, f, s2, s3) //append item to cart list
		SearchSaveCheckOut(un, s, d, f, s2, s3)
		break

	case 3:
		ClearShoppingCartAndCheckoutinformation()
		break

	case 4:
		CheckoutConfirm(un) //transform MyShoppingCart into checkout with transactionID
		break

	case 5:
		break

	}
	// return
}

func CheckoutConfirm(s string) { //s in this case is your username

	fmt.Println("\n========================================================")
	fmt.Println("\n\t\tCheckout Confirmed")
	fmt.Println("\n========================================================\n")

	var PriorityIndex = 1
	var usrInpt int
	totalTransPerSession := []string{} //everytime this function is freshly called, total transaction per session resets to 0 which is correct.

	for _, v := range MyShoppingCart {
		TransID++                         //everytime there is a new shopping cart; we need to have a new transactioID for EACH registered mercahnt
		ns := strconv.Itoa(TransID)       //converts transaction ID to string
		transactionID := "MC" + ns + "KV" //generates transaction ID
		MyCheckoutTranID[transactionID] = Checkout{v.FoodName, v.MerchantName, v.Quantity, v.Price, transactionID, s}
		MyCheckoutIDUsername[s] = Checkout{v.FoodName, v.MerchantName, v.Quantity, v.Price, transactionID, s}
		totalTransPerSession = append(totalTransPerSession, transactionID) //if there 5 different transactions for 5 different merchats, all will be appended to this variable
	}
	QueueID++
	ns2 := strconv.Itoa(QueueID)       //converts queue ID to string
	systemQueueID := "OS" + ns2 + "KV" //generates system queue ID. NOT queue for merchant but on overall system level

	checkUserAdminResult := checkUserAdmin(s) //returns T or F
	if checkUserAdminResult {                 //this allows an admin user to push a queue to priority
		fmt.Printf("\nHi user: %s,! You are authorised as a customer service officer!\n\n", s)
		fmt.Println("Please enter priority number more than 0 if an order is an order is meant for service recovery")
		fmt.Println("Enter 0 for default if is not for service recovery\n")
		fmt.Scanln(&usrInpt)
		PriorityIndex = usrInpt
	}
	// KVorderforQueue = append(KVorderforQueue, KVorder{totalTransPerSession, s, systemQueueID, PriorityIndex})
	KVorderforQueue := KVorder{totalTransPerSession, s, systemQueueID, PriorityIndex}
	SysQueue.Enqueue(KVorderforQueue, PriorityIndex) // can we change KVorderforQueue STRUCT to take in priority index also?
	//store system queue number queue and other  information into map
	SystemQueueDB[systemQueueID] = &SystemOrderInfo{KVorder{totalTransPerSession, s, systemQueueID, PriorityIndex}, "", OrderStatuses{false, false, false, false}} //have to put in empty string for driver information
	// SystemQueueDB["OS123KV"] = SystemOrderInfo{KVorder{[]string{"MC8912KV", "MC9123KV"}, "testadmin", "OS123KV", 7}, "driver1"}

	fmt.Println("\nNext Order in Queue successfully processed and dispatched. What would you like to do next?\n")
	ClearShoppingCartAndCheckoutinformation()
	return
}

func checkUserAdmin(s string) bool {
	for _, v := range UsernameList2 {
		if v.UserName == s {
			if v.isAdmin == true {
				return true
			}
		}
	}
	return false
}

func checkUserDispatch(s string) bool {
	for _, v := range UsernameList2 {
		if v.UserName == s {
			if v.isDispatch == true {
				return true
			}
		}
	}
	return false
}

func AddOrdertoCart(un string, s string, d float64, f float64, s2 string, s3 string) { //passes username, the complete food dish, the total quantity, the price, the merchant name.

	ok := RepeatedShoppingCartCheck(s2, s3) // check for 100% && match with FOOD NAME as well as merchant name
	if ok {
		MyShoppingCart = append(MyShoppingCart, ShoppingCartOrder{s3, s2, d, f, un})
		fmt.Println("Order Successfully Added to Cart!")
	} else {
		fmt.Println("\nCart Already contains similar items. Please consider adding something else into cart")
	}
	PrintShoppingCartItems(un, s, d, f, s2, s3)
}

func RepeatedShoppingCartCheck(s2 string, s3 string) bool {

	for _, v := range MyShoppingCart {
		if v.FoodName == s3 && v.MerchantName == s2 {
			return false
		}
	}
	return true
}

func PrintShoppingCartItems(un string, s string, d float64, f float64, s2 string, s3 string) { //passes username, the complete food dish, the total quantity, the price, the merchant name.
	var totalQty float64
	var totalSum float64

	if len(MyShoppingCart) == 0 {
		//do nothing
	} else {
		fmt.Println("\n========================================================")
		fmt.Println("\n\tYour Current Shopping Basket Items")
		// fmt.Println("\n================================================")
		for i, v := range MyShoppingCart {

			fmt.Println("\n========================================================")
			fmt.Printf("\n%d.\n", i+1)
			fmt.Printf("\tMerchant Name: \t\t\t%s", v.MerchantName)
			fmt.Printf("\n\n\tDish Name: \t\t\t%s", v.FoodName)
			fmt.Printf("\n\n\tQuantity: \t\t\t%0.0f", v.Quantity)
			fmt.Printf("\n\n\tUnit Price: \t\t\t%0.2f", v.Price)
			fmt.Printf("\n\n\tTotal Cost: \t\t\t%0.2f\n", v.Quantity*v.Price)
			// fmt.Printf("\n\n", s)
			// fmt.Printf("\n\nMerchant Name: %s", 2)
			totalQty += v.Quantity
			totalSum += v.Quantity * v.Price

		}
		fmt.Println("\n========================================================")
		fmt.Printf("\nTotal Number of Food Dishes Ordered: \t%0.0f\n", totalQty)
		fmt.Printf("Total Cost of all Food Dishes Ordered: \t$%0.2f\n", totalSum)
		fmt.Println("\n========================================================")

	}
}

func PrintKeywordSearchResults(ss []string, s string) []string { //new func for PrintKeywordSearchResults

	// var usrInpt string
	var found bool
	var Recoverycase2SearchResults []string
	if len(ss) != 0 {
		fmt.Println("\nHere Are the Search Results:\n")

		PrintSliceinLinesGeneral(ss)
		return ss

	} else { //we need to querty again since there are no search results related to the query:

		for found == false {

			fmt.Println("\nThere are no search results for your previous search terms. Please try again")
			//can call function to store rubbish keywords that people are seaching for to create new database
			case2rseultstring, case2rseultint, errorSearchandATC := Case2DisplayAllSearchAndATC()

			if errorSearchandATC != nil {

				fmt.Print("System Error:", errorSearchandATC)

			} else {

				Recoverycase2SearchResults = MyFoodListDB.GetSuggestion(case2rseultstring, case2rseultint)

				if len(Recoverycase2SearchResults) != 0 { //finally have search results

					PrintSliceinLinesGeneral(Recoverycase2SearchResults)
					found = true
					break
				}
			}
		}
		return Recoverycase2SearchResults
	}
}

func ClearShoppingCartAndCheckoutinformation() {
	MyShoppingCart = []ShoppingCartOrder{} //clears shopping cart
	KVorderforQueue = []KVorder{}          //
	// return
}

func DisplayAllDatabase() {

	var usinpt int
	i := 1
	j := 1

	fmt.Println("\nPlease Select from the following options")
	fmt.Println("\n========================================================")

	fmt.Println("1. View all Data regarding Transaction ID(s)")
	fmt.Println("2. Export all Data regarding Transaction ID(s)")
	fmt.Println("3. View all Data related to usernames")
	fmt.Println("4. Export all Data related to usernames")
	fmt.Println("5. Export all Data related to the queue")
	fmt.Scanln(&usinpt)
	fmt.Println("\n========================================================\n")

	switch usinpt {

	case 1:

		// MyCheckoutTranID[transactionID]
		fmt.Println("All data related to Transaction IDs\n")

		for k, _ := range MyCheckoutTranID {

			fmt.Printf("\n%d. Order ID: %s \t\t\t\n", i, k)
			fmt.Printf("\nOrdered Dish Name:\t\t\t%s\n", MyCheckoutTranID[k].FoodName)
			fmt.Printf("\nFulfiled by Merchant: \t\t\t%s\n", MyCheckoutTranID[k].MerchantName)
			fmt.Printf("\nOrdered Quantity:%0.0f\t\t\tUnit Price %0.2f\n", MyCheckoutTranID[k].Quantity, MyCheckoutTranID[k].Price)
			fmt.Printf("\nOrder tagged to username: \t\t%s\n\n", MyCheckoutTranID[k].Username)
			fmt.Printf("Total Order Value: \t\t\t%0.2f\n", MyCheckoutTranID[k].Quantity*MyCheckoutTranID[k].Price)
			fmt.Println("\n============================================================")
			i++
		}
	case 2:
		fmt.Println("Exporting of Data regarding Transaction ID(s) has been started. You can find the file name transactionIDsData.json in the root folder. What else would you like to do next?\n")
		ExportMyCheckoutTranID()

	case 3:
		// MyCheckoutIDUsername[s]
		fmt.Println("All data related to usernames\n")

		for k, _ := range MyCheckoutIDUsername {
			fmt.Printf("\n\n%d. Order ID: %s \t\t\t\n", j, k)
			fmt.Printf("\nOrdered Dish Name:\t\t\t%s\n", MyCheckoutIDUsername[k].FoodName)
			fmt.Printf("\nFulfiled by Merchant: \t\t\t%s\n", MyCheckoutIDUsername[k].MerchantName)
			fmt.Printf("\nOrdered Quantity:%0.0f\t\t\tUnit Price %0.2f\n", MyCheckoutIDUsername[k].Quantity, MyCheckoutIDUsername[k].Price)
			fmt.Printf("\nOrder tagged to username: \t\t%s\n\n", MyCheckoutIDUsername[k].Username)
			fmt.Printf("Total Order Value: \t\t\t%0.2f\n", MyCheckoutIDUsername[k].Quantity*MyCheckoutIDUsername[k].Price)
			fmt.Println("\n============================================================")
			j++
		}

	case 4:
		fmt.Println("Exporting of Data regarding usernames has been started. You can find the file name You can find the file name usernameData.json in the root folder. What else would you like to do next?\n")
		ExportMyCheckoutIDUsername()

	case 5:

		fmt.Println("Show All Past Queue Data")
		// AppendFakeQueueData()
		// fmt.Println(SystemQueueDB)
		fmt.Println("\n============================================================")

		for l, _ := range SystemQueueDB {
			count := 1
			fmt.Printf("\nSystem Queue ID:\t%s\n", l)
			fmt.Printf("\n\nOrder ID(s) Associated with the order: \t\t\t\n")
			fmt.Println("\n============================================================\n")
			for k, v := range SystemQueueDB[l].transID {
				fmt.Printf("%d. \t%s\n", k+1, v) //prints index number and number of transID
			}
			fmt.Println("\n============================================================")
			fmt.Printf("\nSystem Queue tagged to username:\t%s\n", SystemQueueDB[l].username)
			fmt.Printf("System Queue with priority index:\t%d\n", SystemQueueDB[l].priorityIndex)
			if SystemQueueDB[l].DriverName == "" {
				fmt.Println("**There is no driver assigned for this order. Please assign!")
			} else {
				fmt.Printf("System Queue tagged to driver:\t\t%s\n", SystemQueueDB[l].DriverName)
			}
			fmt.Println("\n============================================================")
			count++
		}
	}
}

func ExportMyCheckoutTranID() {

	file, _ := json.MarshalIndent(MyCheckoutTranID, "", " ")

	_ = ioutil.WriteFile("transactionIDsData.json", file, 0644)

}

func ExportMyCheckoutIDUsername() {

	file, _ := json.MarshalIndent(MyCheckoutIDUsername, "", " ")
	_ = ioutil.WriteFile("usernameData.json", file, 0644)

}

func AppendFakeQueueData() {
	SystemQueueDB["OS120KV"] = &SystemOrderInfo{FirstQueueValue, "", OrderStatuses{false, false, false, false}}
	SystemQueueDB["OS1201KV"] = &SystemOrderInfo{SecondQueueValue, "driver2", OrderStatuses{false, false, false, false}}
	SystemQueueDB["OS123KV"] = &SystemOrderInfo{ThirdQueueValue, "driver1", OrderStatuses{false, false, false, false}}
}

func checkValidDriver(s string) bool {

	if _, ok := DriverDB[s]; ok {
		return true
	}
	return false
}

func displayDispatchMenu() {

	var driverInput string
	fmt.Println("Please Enter Delivery Partner Detail's username")
	fmt.Scanln(&driverInput)
	if checkValidDriver(driverInput) {
		fmt.Printf("\nOrder Number %s (The Next Order in Line) Successfully dispatched. Please check current order queue again to check latest queues.\n", SysQueue.front.item.systemQueueID)
		s := SysQueue.front.item.systemQueueID
		fmt.Printf("\nThe Order is tagged to: %s\n", driverInput)
		PushDriverDatatoMap(s, driverInput)
		SysQueue.Dequeue()

	} else {
		fmt.Println("\nDelivery Partner Username Invalid. Please try again later")
	}
	// break
}

func PushDriverDatatoMap(sqid string, dn string) { // system queue ID/ driver name
	SystemQueueDB[sqid].DriverName = dn
}

func editSystemOrderInforamtion() {

	var sqid string
	var orderToDelete string
	var tempSearchResults []string
	var ai int = -1
	var newSlice []string
	fmt.Println("Please enter System Queue Number")
	fmt.Scanln(&sqid)
	sc := strings.ToUpper(sqid) //converts whatever stirng value to UPPER case to match whatever that is stored in SystemQueueDB

	if _, ok := SystemQueueDB[sc]; ok {
		fmt.Println("\nPlease See All Data Available for this System Queue Number:\n")
		for k, v := range SystemQueueDB[sc].transID {
			fmt.Printf("%d. Transaction ID:\t%s\n", k+1, v)
			tempSearchResults = append(tempSearchResults, v) //this will give us a short liste of transaction to work with. Saving on memory consumption
			// fmt.Println("see below new slice data")

		}
		// fmt.Println(tempSearchResults)

		fmt.Println("\nPlease Enter Transaction ID to Cancel. Once cancelled, order number will be deleted form System Queue Map Database.")
		fmt.Scanln(&orderToDelete)
		// fmt.Printf("\n%T%[1]v\n", orderToDelete)
		//check if order to delete is available in temp search results
		for k, v := range tempSearchResults {
			// fmt.Println(v)
			// fmt.Println(strings.ToUpper(orderToDelete))

			if strings.ToUpper(orderToDelete) == v {
				ai = k // assigns affected index = k
				break  // break once value found
			}
		} //end of range loop. Continue procees since keyed in value is a valid order numer

		if ai == -1 {
			fmt.Println("Transaction ID not found in this System Queue Number. Please ensure you are searching in the correct System Queue ID")
			return
		}

		x := len(tempSearchResults) //gives us total len

		if ai == 0 {
			newSlice = tempSearchResults[1:] // deletes first order number

		} else if ai == x+1 { //last element
			newSlice = tempSearchResults[:x] // deletes the last order number

		} else {
			newSlice = tempSearchResults[:ai] //will capture first array all the way until affected array -1 . if data is 2, 3, 4, 5, and affected order is 5, this slice captures up till data 4 only. EXCLUDING 5
			newSlice2 := tempSearchResults[ai+1:]
			for _, v := range newSlice2 {
				newSlice = append(newSlice, v) //this will give me a new slice to work with
			}
		}

		SystemQueueDB[sc].transID = newSlice
		fmt.Printf("\nOrder: %s has been deleted\n", orderToDelete)
		fmt.Printf("\nOrder Numbers in System Queue Number: %s Sucesfully Updated\n", sc)
		// fmt.Println(SystemQueueDB[sc].transID)

		return
	} else {
		fmt.Println("\nThe System Queue Number you've entered is invalid. Please ensure you've typed input the correct information")
	}

}
