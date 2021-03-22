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
	isRider      bool //rider now has the option of dequeue priority queues
}

type KVorder struct {
	transID       []string
	username      string
	systemQueueID string
	// priority      int
}

type OpeningPeriods map[string][]string

// var UsernameList = []string{"1", "admin", "estella", "alvin", "customer", "customer1", "estella45", "admin6"} //username prepopulated but can grow accordingly to the business's requirements
var UsernameList2 = []UsernameCustom{{"1", true, false, false}, {"admin", true, false, false}, {"estella", false, false, false}, {"alvin", false, false, false}, {"customer", false, false, false}, {"customer1", false, false, false}, {"estella45", false, false, false}, {"admin6", true, true, true}, {"rider", false, true, true}, {"supervisor", false, true, true}, {"pris", true, true, true}}

func main() {

	ch := make(chan string) //create a channel called c
	var usrnameInpt string
	var count int = 0
	var postalCodeInpt int
	currentTime := time.Now()

	usernameDS := InitUsernameTrie() //inits Trie data for username DS
	go CreateFoodList(ch)            //newResult is a slice that is being returned byCreateFoodList function
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
			fmt.Println("System Error:", err)
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
				fmt.Println("Here are the current orders in queue")
				SysQueue.PrintAllNodes()
				break

			case 5:
				fmt.Println("Order Successfully dispatched. Please check current order queue again to check latest queues.")
				SysQueue.Dequeue()
				break

			case 6:
				fmt.Println("Display and Export Databases")
				DisplayAllDatabase()
				break

			case 7:
				os.Exit(1)

			default:
				break
			}
		} // end switch else statement for error handling
	}

} // close main functioin
