// go run main.go database.go databaseMaps.go mainFunctions.go messageTemplates.go rawData.go queueManagement.go
/*Prototyping app for Go Assignment 2*/

package main

import (
	"fmt"
	"os"
	"time"
)

var MyFoodListDB = InitMyFoodList()
var TransID int = 0
var QueueID int = 500
var SysQueue = InitSysQueue()

type CustomerData struct { //every customer data must be populated this way.
	Username   string  //their username
	Password   string  //their password
	Email      string  //their email
	TotalSpend float64 //their total spend with the kay cafe
}

type driverInfo struct {
	Fname         string
	Lname         string
	ContactNumber string
	Email         string
}

type FoodInfo struct { //struct type because we need to hold values

	FoodName         string
	MerchantName     string
	DetailedLocation string
	PostalCode       int
	Price            float64
	OpeningPeriods   OpeningPeriods
}

type FoodInfo2 struct { //struct type because we need to hold values

	FoodName         string
	MerchantName     string
	DetailedLocation string
	PostalCode       int
	Price            float64
	OpeningPeriods   OpeningPeriods
}

type ShoppingCartOrder struct {
	FoodName     string
	MerchantName string
	Quantity     float64
	Price        float64
	// OrderID      int
	Username string
	//datetime parementers can also be considered at a later time
	//Discounts	[]string //stacked discounts field
	//
}

type Checkout struct {
	FoodName     string
	MerchantName string
	Quantity     float64
	Price        float64
	OrderID      string
	Username     string
}

type UsernameCustom struct {
	UserName string
	// Password     string
	isAdmin bool //admins have true/ all others have false
	// email        string
	subscription bool
	isDispatch   bool //rider now has the option of dequeue priority queues
}

type KVorder struct {
	transID       []string
	username      string
	systemQueueID string
	priorityIndex int
}

type OrderStatuses struct {
	MerchantsStatus bool //check if all transactionsIDS from ALL merchants are completed. If yes, then send true
	OrderPickedUp   bool //once picked up by driver/ rider/ status becomes true
	DriverDelivered bool //once driver/ rider has delivered, this status becomes true
	CusDeliveryCfm  bool //once customer confirms delivery OR when a preset timeout expires from DriverDelivered = true, this status becomes true.
}
type SystemOrderInfo struct {
	KVorder
	DriverName string
	OrderStatuses
}

type OpeningPeriods map[string][]string

func main() {

	ch := make(chan string) //create a channel called c
	var usrnameInpt string
	var count int = 0
	var postalCodeInpt int
	currentTime := time.Now()
	usernameDS := InitUsernameTrie() //inits Trie data for username DS
	go AppendFakeQueueData()
	go CreateFoodList(ch) //newResult is a slice that is being returned byCreateFoodList function
	fmt.Println("\nSystem Message :", <-ch)
	go CreateFoodListMap(ch)
	fmt.Println("System Message :", <-ch)
	go MyFoodListDB.PreInsertTrie(FoodMerchantNameAddress, ch) //populates Trie Data for Food LIst
	go SysQueue.PrepreEnqueue(ch)                              //populates Queue weith predetermined data
	go usernameDS.PreInsertTrieUser(UsernameList2, ch)         //preInserts all existing user database
	myPostalCodesDB := InitPostalCode()                        //creates PostalCode BST DB
	myPostalCodesDB.PreInsertPostalCode()                      //preinset POSTAL Code DB
	fmt.Println("System Message :", <-ch)
	fmt.Println("System Message :", <-ch)
	fmt.Println("System Message :", <-ch)
	fmt.Println("System Message : System is Ready", currentTime.Format("2006-01-02 15:04:05"))

	PrintWelcomeMessage() //prints welcome message

	fmt.Println()
	for {
		fmt.Println("\nPlease enter username to proceed\n")
		fmt.Scanln(&usrnameInpt) //username checked against trie
		checkUsername := usernameDS.UserSearch(usrnameInpt)
		if checkUsername {
			PrintUserValidated(usrnameInpt)
			break
		} else {
			count++
			PrintUserNotValidated(usrnameInpt)
			if count >= 5 {
				PrintNoOfTriesExceeded()
				break //tested. breaks the main looping for loop
			}
		}
	}
	for { //display main meu and loop until you die.
		switchRS, err := displayMainMenu()
		if err != nil {
			fmt.Println("\nSystem Error:", err)
		} else {
			switch switchRS {
			case 1:
				for {
					fmt.Println("case 1 Access all food items")
					case1result := Case1DisplayAllFoodItems(FoodMerchantNameAddress)
					if ToQuit(case1result) {
						fmt.Println("Returned to previous menu!")
						break
					}
				}
			case 2:
				// fmt.Println("case 2 search and add item to cart")

				case2rseultstring, case2rseultint, errorSearchandATC := Case2DisplayAllSearchAndATC() //call sub menu function// case2result is a string
				// fmt.Println("recover to main")
				if errorSearchandATC != nil {
					fmt.Print("System Error:", errorSearchandATC)
				} else {

					case2SearchResults := MyFoodListDB.GetSuggestion(case2rseultstring, case2rseultint)
					editedSearch := PrintKeywordSearchResults(case2SearchResults, usrnameInpt)
					// fmt.Println("i got called")
					AddToCart(editedSearch, usrnameInpt)
				}
				break
			case 3:
				// fmt.Println("case 3 check merchant postal code validity")
				fmt.Println("\nEnter Postal Code to check\n")
				fmt.Scanln(&postalCodeInpt)
				checkPC, errPC := myPostalCodesDB.Search(postalCodeInpt)
				if errPC != nil {
					fmt.Println("System Error:", errPC)
				} else {
					if checkPC != nil {
						fmt.Println("We have merchants registered at this postal code. ")
					} else {
						fmt.Println("We have no merchants registered at this postal code. Please advise the sales team accordingly")
					}
				}
				break
			case 4:
				fmt.Println("\nHere are the Current Orders In Queue")
				SysQueue.PrintAllNodes()
				break

			case 5:
				checkUserDispatchResult := checkUserDispatch(usrnameInpt)

				if checkUserDispatchResult { //if loop runs if logged in user has rider/driver/dispatch rights
					displayDispatchMenu()

				} else {
					fmt.Println("\nAccess Denied. User Does Not Have Rider or Dispatch Rights")
					break
				}
				break

			case 6:
				fmt.Println("Display and Export Databases")
				DisplayAllDatabase()
				break

			case 7:
				editSystemOrderInforamtion()
				break

			case 8:
				os.Exit(1)
			}
		} // end switch else statement for error handling
	}

} // close main functioin
