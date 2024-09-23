package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	database "github.com/Eltanio-one/jumpin-go-rewrite/src/db"
)

type machineList struct {
	MachineOne   string `json:"machineOne"`
	MachineTwo   string `json:"machineTwo"`
	MachineThree string `json:"machineThree"`
	MachineFour  string `json:"machineFour"`
	MachineFive  string `json:"machineFive"`
}

// SessionPlan handles the matching algorithms if there are enough users within a lobby
func SessionPlan(w http.ResponseWriter, r *http.Request) {

	logger := log.New(os.Stdout, "sessionPlan", log.LstdFlags)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		logger.Printf("error=%q statuscode=%d message=%q", "invalid HTTP method", http.StatusMethodNotAllowed, "only accepts POST method")
		http.Error(w, "incorrect HTTP method used", http.StatusMethodNotAllowed)
		return
	}

	db, err := database.InitialiseConnection()
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "database connection error", http.StatusInternalServerError, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// check user is assigned to gym, if not err and close
	_, err = Store.Get(r, "login")
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "session error", http.StatusInternalServerError, "no session stored for the user")
		http.Error(w, "no session stored for the user", http.StatusMethodNotAllowed)
		return
	}
	// userID := session.Values["userID"]

	// db call to get gymid from user

	// check user has previous session, if so err and close

	// unmrshal machine list from req into struct
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to read request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var machines machineList

	err = json.Unmarshal(body, &machines)
	if err != nil {
		logger.Printf("error=%q statuscode=%d message=%q", "unable to unmarshal request body", http.StatusInternalServerError, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// validate all machines provided?

	// insert the session (userid, username, datetime and machine list) into userSessions table (needs creating, add session to list so we can check lobby)

	// from userSessions, pull and check if length >= 2, if not, they wait

	// if so, run matching algorithms

	// 1 - match based on provided machines and order
	// 2 - match users based on provided machines with no order, just how many they have in common

}

func userMatches(rows [][]string) (map[string]string, error) {
	matches := zeroArray(rows)
	for i := range rows {
		for j := range rows {
			if i == j {
				continue
			}
			matches[i][j] = len(intersect(rows[i], rows[j]))
		}
	}

	// output := make(map[string]string)
	// for i := range matches {
	// 	rowMax, err := maxOccurrence(matches[i])
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if rowMax == 0 {
	// 		continue
	// 	}
	// 	for j := range matches[i] {

	// 	}
	// }
	return nil, nil
}

func maxOccurrence(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("empty slice provided")
	}

	maxVal := nums[0]
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal, nil
}

// zeroArray initialises an empty 3D int matrix corresponding to number of users and number of machines selected
func zeroArray(rows [][]string) [][]int {
	matrix := make([][]int, len(rows))
	for i := range matrix {
		matrix[i] = make([]int, len(rows[0]))
	}
	return matrix
}

// intersect replicates the numpy.Intersect1d function, taking two user string slices containing machines
// returning a slice of the intersecting values
func intersect(userOne, userTwo []string) []string {
	elements := make(map[string]struct{})
	for _, val := range userOne {
		elements[val] = struct{}{}
	}

	var intersection []string
	for _, val := range userTwo {
		if _, found := elements[val]; found {
			intersection = append(intersection, val)
		}
	}

	return intersection
}
