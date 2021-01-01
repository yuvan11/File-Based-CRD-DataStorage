package CRDFunctions

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Structure to define value as json type
type KeyValue struct {
	Value          string `json:"value"`
	Created        string `json:"createdAt"`
	TimetoLiveFlag bool   `json:"liveFlagStatus"`
}

// setting custom timetolive key interval for each user
var TimetoLiveProperty int64 = 20000

// setting filesizelimit 1GB (1024 * 1024 * 1024 ) in bytes
var filesizelimit int64 = 1073741824
var splittedkeyValue []string

//created a map to hold key value pair
var dataStorageMap = make(map[string]KeyValue, 0)

// Default path location for the file
var dir_path = "/root/Goprojects/src/File-Based-CRD/src/"
var filename string

// function allows to create a new file for each specific user
func CreateFile() {

	CurrentTime := time.Now().Local()
	CurrentDate := CurrentTime.Format("2006_01_02_15_04_05")
	filename = dir_path + "user" + CurrentDate + ".txt"
	fmt.Println("File",filename)

    // Create a new file 
	_, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	return
}


// function allows to create a new key value pair
func Create() {

    // flag value to track present in the data store
	var createflag bool = false
	var keyvalue KeyValue
	var createMap = make(map[string]KeyValue, 0)

   

    // Calculate the file size
	fi, _ := os.Stat(filename)
	

    // Calculate start time for each user
    startTime := time.Now().Local()
    starttimeStampString := startTime.Format("2006-01-02 15:04:05")
    unixNano := startTime.UnixNano()
    umillisec := unixNano / 1000000


    // scanning the input from user
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter Key Value Pair e.g --> (1 , hello)")
	scanner.Scan()
	input := scanner.Text()
    // Provide input splitted by comma
	splittedkeyValue = strings.Split(input, ",")
    // splittedkeyValue[0] is a key and splittedkeyValue[1] is value
	keyvalue.Value = splittedkeyValue[1]
	input_key := splittedkeyValue[0]

    // Check if dataStorageMap is empty
	if len(dataStorageMap) == 0 {
		//fmt.Println("Data Storage is empty")
		createflag = false

	} else {

		for keys, _ := range dataStorageMap {
			if keys == input_key {
				createflag = true
				break
			}

		}
	}

	if createflag == true {
		fmt.Println("Key Exists, Can't add, Try with diferent Key ")
		return
	} else {

        // Calculating end time for each user
		endTime := time.Now().Local()
		_ = endTime.Format("2006-01-02 15:04:05")
		unixNano = endTime.UnixNano()
		umillisec1 := unixNano / 1000000

        /*
         Find differece betweeen start time and end time to check the 
        whether user key generated key is  expired or not
        */
		diff_milli_seconds := umillisec1 - umillisec
		createflag = false
		//fmt.Println("Data Added Successfully")


        // if expiry time is greater than timeToliveProperty time
        // user cannot be allowed to read or write the specific key
		if diff_milli_seconds > TimetoLiveProperty {

			fmt.Println("Time Limit exceeded ")
			var fetchValue = dataStorageMap[input_key]
            // timeToliveProperty flag is assgined to true to track the expiry key
			fetchValue.TimetoLiveFlag = true
			dataStorageMap[input_key] = fetchValue

			keyvalue.Created = starttimeStampString
			keyvalue.TimetoLiveFlag = true
			dataStorageMap[input_key] = keyvalue
			createMap[input_key] = keyvalue

            // Open file
			d, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}

            // Marshal map data in to JSON 
			jsonCreateData, err := json.Marshal(createMap)

			if err != nil {
				log.Println(err)
			}

            // Write JSON Data in to the file
			d.WriteString(string(jsonCreateData) + "\n")

		} else {

        // if expiry time is not greater than timeToliveProperty time
        // user be allowed to read or write the specific key
			fmt.Println("You're in time")

			var fetchValue = dataStorageMap[input_key]

            // timeToliveProperty flag is assgined to true to track the expiry key
			fetchValue.TimetoLiveFlag = false
			dataStorageMap[input_key] = fetchValue

			keyvalue.Created = starttimeStampString
			keyvalue.TimetoLiveFlag = false
			dataStorageMap[input_key] = keyvalue
			createMap[input_key] = keyvalue

            // open file
			d, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}

            // Marshal map data in to JSON 
			jsonCreateData, err := json.Marshal(createMap)

			if err != nil {
				log.Println(err)
			}

            // Check file custom limit, if exceeded not allowed to write in file
			if fi.Size() < filesizelimit {
				d.WriteString(string(jsonCreateData) + "\n")
			} else {

				fmt.Println("The size of the file storing data  1GB is exceed, cannot insert anything")
			}

		}

	}

}

func Read() {

	var readkey string
    //Flag is assigned to keep track of the key presence
	var readflag bool = false
	fmt.Print("Enter the Key to read json value:\n")

	// Read key from user to display corresponding value Json
	fmt.Scanln(&readkey)

    // Read file
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

      // Check if dataStorageMap is empty
	if len(dataStorageMap) == 0 {
		readflag = false
	} else {

		if dataStorageMap[readkey].TimetoLiveFlag {
			fmt.Println("You're not allowed to read the expired key", readkey)
		} else {
			for keys, _ := range dataStorageMap {
				if keys == readkey {
					readflag = true
					break
				}
			}
		}
	}

	if readflag == true {

		var readValue = dataStorageMap[readkey]
        // print the exact value for the user entered key
		fmt.Println(readValue.Value)
		fmt.Println("Data Fetched Successfully")
		return
	} else {
		fmt.Println("Key not available")
	}
}

func Delete() {

	var delKey string
	var delflag bool = false

	fmt.Print("Enter Key to delete: ")

	fmt.Scanln(&delKey)


	if len(dataStorageMap) == 0 {
		delflag = false

	} else {

		if dataStorageMap[delKey].TimetoLiveFlag {

			fmt.Println("You're not allowed to delete the expired key", delKey)
		} else {
			for keys, _ := range dataStorageMap {
				if keys == delKey {
					delflag = true

				}

			}
		}

	}

	if delflag == true {

		fmt.Println("Key Exists, We can Delete")
		delete(dataStorageMap, delKey)
		fmt.Println("Key Deleted Successfully")

		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
		if err != nil {
			panic(err)
		}

		jsonDataDel, err := json.Marshal(dataStorageMap)

		if err != nil {
			log.Println(err)
		}

		f.WriteString(string(jsonDataDel) + "\n")

	} else {

		fmt.Println("Key not available")

	}

}

// function to display data storage as key value pair
func DisplayAll() {

	if len(dataStorageMap) == 0 {
		fmt.Println("Data Storage Empty!!! No data Found !!!")
	} else {
		for key, val := range dataStorageMap {
			fmt.Println(key, "-->", val)
		}
	}
}
