package main

import (
	"fmt"
)

type Candidate struct {
	Name       string
	Party      string
	VoteCount  int
	Percentage float64
}

type Voter struct {
	Name  string
	ID    int
	Voted bool
}

type Time struct {
	checktime bool
	hours     int
}

type Election struct {
	Candidates []Candidate
	Voters     []Voter
	Times      []Time
}

//MAIN FUNCTION

func main() {
	var e Election
	var candidate Candidate
	var time Time

	Nama := []string{"Mario", "Jonas", "Jorji"}
	Partai := []string{"PVT", "Fretilin", "APMT"}
	Waktu := []int{8}
	CheckTime := []bool{true}
	time.hours = Waktu[0]
	time.checktime = CheckTime[0]
	e.Times = append(e.Times, time)

	for i := 0; i < len(Nama); i++ {
		candidate.Name = Nama[i]
		candidate.Party = Partai[i]
		e.Candidates = append(e.Candidates, candidate)
	}

	mainMenu(&e)

}

// FUNC  MAINMENU (MAIN MENU)
func mainMenu(e *Election) {
	var choice int

	fmt.Println("[<>]=====================================================================[<>]")
	fmt.Println(" ||                 WELCOME TO GENERAL ELECTION APPLICATION                ||")
	fmt.Println("[<>]=====================================================================[<>]")
	fmt.Println("LOGIN AS : ")
	fmt.Println("1. Voter")
	fmt.Println("2. Voting Staff")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		voterMenu(e)

	} else if choice == 2 {
		adminMenu(e)
	} else {
		fmt.Println("Thanks for using our Application")
	}
}

// VOTERMENU FUNCTION (FOR VOTER MENU)
func voterMenu(e *Election) {
	var choice int

	fmt.Println("[<>]=========================================[<>]")
	fmt.Println(" ||                 VOTING MENU               || ")
	fmt.Println("[<>]=========================================[<>]")
	fmt.Println("1. Start Voting")
	fmt.Println("2. Show Candidates")
	fmt.Println("3. Show Results")
	fmt.Println("4. Search Candidates")
	fmt.Println("5. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		startVoting(e, 1)
	} else if choice == 2 {
		showCandidates(e, 1)
	} else if choice == 3 {
		showResults(e, 1)
	} else if choice == 4 {
		searchCandidates(e)
	} else if choice == 5 {
		fmt.Println("Thanks for using our Application")
		mainMenu(e)
	} else {
		fmt.Println("Invalid choice.")
	}
}

// ADMINMENU FUNCTION (FOR VOTING STAFF MENU)
func adminMenu(e *Election) {
	var choice int

	fmt.Println("[<>]=========================================[<>]")
	fmt.Println(" ||                 ADMIN MENU                || ")
	fmt.Println("[<>]=========================================[<>]")
	fmt.Println("1. Add Candidate")
	fmt.Println("2. Edit Candidate")
	fmt.Println("3. Delete Candidate")
	fmt.Println("4. Add Voter")
	fmt.Println("5. Edit Voter")
	fmt.Println("6. Delete Voter")
	fmt.Println("7. Start Voting")
	fmt.Println("8. Show Candidates")
	fmt.Println("9. Show Results")
	fmt.Println("10. Setting time")
	fmt.Println("11. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		addCandidate(e)
	} else if choice == 2 {
		editCandidate(e)
	} else if choice == 3 {
		deleteCandidate(e)
	} else if choice == 4 {
		addVoter(e)
	} else if choice == 5 {
		editVoter(e)
	} else if choice == 6 {
		deleteVoter(e)
	} else if choice == 7 {
		startVoting(e, 2)
	} else if choice == 8 {
		showCandidates(e, 2)
	} else if choice == 9 {
		showResults(e, 2)
	} else if choice == 10 {
		setTime(e)
	} else if choice == 11 {
		fmt.Println("Thanks for using our Application")
		mainMenu(e)
	} else {
		fmt.Println("Invalid choice.")
	}
}

// FUNC VALIDID (TO IDENTIFY VOTER HAS A VALID ID OR NOT)
func validID(idNumber int, year int) bool {
	var validyear int
	var found bool = false

	for !found {
		validyear = (idNumber / 10000) % 10
		if idNumber >= 1000000000000000 && idNumber <= 9999999999999999 {
			if validyear == (year % 100) {
				fmt.Println("ID Valid")
				found = true
			}
		}
		if !found {
			fmt.Println("ID Invalid")
			fmt.Println("Please Re-enter validID number : ")
			fmt.Scan(&idNumber)
			fmt.Println("Please Enter your Birthyear : ")
			fmt.Scan(&year)
		}
	}
	return found

}

// FUNC ISREGISTERED (TO CHECK WHETHER THE VOTER ALREADY VOTE OR NOT)
func IsRegistered(idNumber int, found *bool, election *Election) {
	*found = false
	for i := 0; i < len(election.Voters); i++ {
		if election.Voters[i].ID == idNumber {
			*found = true
		}
	}

}

func computeThreshold(e *Election) float64 {
	threshold := 50.01
	return threshold
}

// FUNC ADDCANDIDATE (TO ADD A NEW CANDIDATE)
func addCandidate(e *Election) {
	var candidate Candidate
	var found bool
	fmt.Print("Enter candidate name: ")
	fmt.Scan(&candidate.Name)
	fmt.Print("Enter candidate party: ")
	fmt.Scan(&candidate.Party)
	i := 0
	found = false
	for !found && i < len(e.Candidates) {
		if candidate.Party == e.Candidates[i].Party {
			found = true
		} else {
			i++
		}
	}
	if !found {
		e.Candidates = append(e.Candidates, candidate)
		fmt.Println("Candidate added successfully.")
		adminMenu(e)
	} else {
		fmt.Println("Sorry, the party already exist")
		adminMenu(e)
	}

}

// FUNC EDITCANDIDATE (TO EDIT CANDIDATE)
func editCandidate(e *Election) {

	if len(e.Candidates) == 0 {
		fmt.Println("No candidates available.")
	}

	fmt.Println("Available Candidates:")
	for i := 0; i < len(e.Candidates); i++ {
		fmt.Printf("%d. %s (%s)\n", i+1, e.Candidates[i].Name, e.Candidates[i].Party)
	}

	var choice int
	fmt.Print("Enter the candidate number to edit: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(e.Candidates) {
		fmt.Println("Invalid candidate number.")
	}

	candidate := &e.Candidates[choice-1]

	fmt.Print("Enter new name: ")
	fmt.Scan(&candidate.Name)
	fmt.Print("Enter new party: ")
	fmt.Scan(&candidate.Party)

	fmt.Println("Candidate edited successfully.")
	adminMenu(e)
}

// FUNC DELETECANDIDATE (TO DELETE CANDIDATE)
func deleteCandidate(e *Election) {
	if len(e.Candidates) == 0 {
		fmt.Println("No candidates available.")
	}

	fmt.Println("Available Candidates:")
	for i := 0; i < len(e.Candidates); i++ {
		fmt.Printf("%d. %s (%s)\n", i+1, e.Candidates[i].Name, e.Candidates[i].Party)
	}

	var choice int
	fmt.Print("Enter the candidate number to delete: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(e.Candidates) {
		fmt.Println("Invalid candidate number.")
	}

	// Shift candidates after the deleted candidate to the left
	for i := choice - 1; i < len(e.Candidates)-1; i++ {
		e.Candidates[i] = e.Candidates[i+1]
	}

	// Truncate the last candidate
	e.Candidates = e.Candidates[:len(e.Candidates)-1]

	fmt.Println("Candidate deleted successfully.")
	adminMenu(e)
}

// FUNC ADD VOTER (TO ADD VOTER)
func addVoter(e *Election) {
	var voter Voter
	fmt.Print("Enter voter name: ")
	fmt.Scan(&voter.Name)
	fmt.Print("Enter voter ID: ")
	fmt.Scan(&voter.ID)
	e.Voters = append(e.Voters, voter)
	fmt.Println("Voter added successfully.")
	adminMenu(e)
}

// FUNC EDIT VOTER (TO EDIT VOTER)
func editVoter(e *Election) {
	if len(e.Voters) == 0 {
		fmt.Println("No voters available.")
		adminMenu(e)
	}

	fmt.Println("Available Voters:")
	for i := 0; i < len(e.Voters); i++ {
		fmt.Printf("%d. %s (%d)\n", i+1, e.Voters[i].Name, e.Voters[i].ID)
	}

	var choice int
	fmt.Print("Enter the voter number to edit: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(e.Voters) {
		fmt.Println("Invalid voter number.")
	}

	voter := &e.Voters[choice-1]

	fmt.Print("Enter new name: ")
	fmt.Scan(&voter.Name)
	fmt.Print("Enter new ID: ")
	fmt.Scan(&voter.ID)

	fmt.Println("Voter edited successfully.")
	adminMenu(e)
}

// FUNC DELETEVOTER (TO DELETE VOTER)
func deleteVoter(e *Election) {
	if len(e.Voters) == 0 {
		fmt.Println("No voters available.")
		adminMenu(e)
	}

	fmt.Println("Available Voters:")
	for i := 0; i < len(e.Voters); i++ {
		fmt.Printf("%d. %s (%d)\n", i+1, e.Voters[i].Name, e.Voters[i].ID)
	}

	var choice int
	fmt.Print("Enter the voter number to delete: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(e.Voters) {
		fmt.Println("Invalid voter number.")
		adminMenu(e)
	}

	// Shift elements after the chosen voter to the left(therefore 1 is missing in the last element)
	for i := choice - 1; i < len(e.Voters)-1; i++ {
		e.Voters[i] = e.Voters[i+1]
	}

	// update new len e.voters
	e.Voters = e.Voters[0 : len(e.Voters)-1]

	fmt.Println("Voter deleted successfully.")
	adminMenu(e)
}

// FUNC STARTVOTING (TO START VOTING)
func startVoting(e *Election, n int) {
	var idNumber, year int
	var voting, found bool
	var respond string
	var voter Voter

	fmt.Print("Please Enter validID number : ")
	fmt.Scan(&idNumber)
	fmt.Println("Please Enter your Birthyear : ")
	fmt.Scan(&year)
	found = validID(idNumber, year)
	fmt.Println("Do you want to continue to voting? Y/N")
	fmt.Scan(&respond)
	if respond == "N" {
		mainMenu(e)
	} else {
		if e.Times[0].checktime {
			fmt.Println("Voting in progress...")
			IsRegistered(idNumber, &voting, e)

			if !voting && found {
				for !voting {
					for i := 0; i < len(e.Candidates); i++ {
						candidate := e.Candidates[i]
						fmt.Printf("%d. %s (%s)\n", i+1, candidate.Name, candidate.Party)
					}

					fmt.Print("Select your candidate Number: ")
					var vote int
					fmt.Scan(&vote)

					if vote < 1 || vote > len(e.Candidates) {
						fmt.Println("Invalid candidate number.")
						continue
					}

					candidate := &e.Candidates[vote-1]
					candidate.VoteCount++

					fmt.Println("Vote cast successfully.")
					ID := []int{idNumber}
					for i := 0; i < len(ID); i++ {
						voter.ID = ID[i]
						e.Voters = append(e.Voters, voter)
					}
					voting = true
					for i := 0; i < len(e.Candidates); i++ {
						e.Candidates[i].Percentage = (float64(e.Candidates[i].VoteCount) / float64(len(e.Voters))) * 100
					}
				}
				if n == 1 {
					voterMenu(e)
				} else {
					adminMenu(e)
				}
			} else {
				fmt.Println("Sorry, you already voted !")
				if n == 1 {
					voterMenu(e)
				} else {
					adminMenu(e)
				}
			}

		} else {
			fmt.Println("The time is out of Voting period, The time should be >= 8 and <= 17", e.Times[0].hours)
			if n == 1 {
				voterMenu(e)
			} else {
				adminMenu(e)
			}
		}

	}

}

// FUNC SEARCHCANDIDATES (TO LIST CANDIDATES BY CERTAIN PARTY, OR BY NAMES)
func searchCandidates(e *Election) {
	var choice, i int
	var search string

	fmt.Println("Search candidates by:")
	fmt.Println("1. Party")
	fmt.Println("2. Name")
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	i = 0

	if choice == 1 {
		fmt.Print("Enter the candidate's party name: ")
		fmt.Scan(&search)
		found := false
		for i < len(e.Candidates) {
			if search == e.Candidates[i].Party {
				fmt.Println("-------------------------------")
				fmt.Printf("Name: %s\n", e.Candidates[i].Name)
				fmt.Printf("Party: %s\n", e.Candidates[i].Party)
				fmt.Println("-------------------------------")
				found = true
			}
			i++
		}
		if !found {
			fmt.Println("There is no party named", search)
		}
	} else if choice == 2 {
		fmt.Print("Enter the candidate's name: ")
		fmt.Scan(&search)
		found := false
		for i < len(e.Candidates) {
			if search == e.Candidates[i].Name {
				fmt.Println("-------------------------------")
				fmt.Printf("Name: %s\n", e.Candidates[i].Name)
				fmt.Printf("Party: %s\n", e.Candidates[i].Party)
				fmt.Println("-------------------------------")
				found = true
			}
			i++
		}
		if !found {
			fmt.Println("There is no name called", search)
		}
	} else {
		fmt.Println("Invalid choice")
		voterMenu(e)
	}

	voterMenu(e)
}

// FUNC SHOWCANDIDATES (TO SHOW LIST OF CANDIDATES)
func showCandidates(e *Election, n int) {

	if len(e.Candidates) == 0 {
		fmt.Println("No candidates available.")
	}

	fmt.Println("\n------- LIST OF CANDIDATES -------")
	for i := 0; i < len(e.Candidates); i++ {
		fmt.Printf("Number: %d\n", i+1)
		fmt.Printf("Name: %s\n", e.Candidates[i].Name)
		fmt.Printf("Party: %s\n", e.Candidates[i].Party)
		fmt.Println("-------------------------------")
	}
	if n == 1 {
		voterMenu(e)
	} else {
		adminMenu(e)
	}

}

// FUNC SHOWRESULTS (TO SHOW RESULTS OF ELECTION)
func showResults(e *Election, n int) {
	if len(e.Candidates) == 0 {
		fmt.Println("No candidates available.")
	}

	// Sort candidates by vote count
	sortCandidatesByVoteCount(e.Candidates)

	fmt.Println("\n------- ELECTION RESULTS -------")
	for i := 0; i < len(e.Candidates); i++ {
		fmt.Printf("Name: %s\n", e.Candidates[i].Name)
		fmt.Printf("Party: %s\n", e.Candidates[i].Party)
		fmt.Printf("Vote Count: %d\n", e.Candidates[i].VoteCount)
		fmt.Printf("Vote percentage: %.2f%%\n", e.Candidates[i].Percentage)
		fmt.Println("-------------------------------")
	}
	idx, max := 0, e.Candidates[0].Percentage
	for i := 1; i < len(e.Candidates); i++ {
		if max < e.Candidates[i].Percentage {
			max = e.Candidates[i].Percentage
			idx = i
		}
	}
	if computeThreshold(e) <= max {
		fmt.Println("The Winner is ", e.Candidates[idx].Party, "With percentage", max, "%")
	} else {
		fmt.Println("There is no winner, because the persentage didn't reach the", computeThreshold(e), "%")
	}
	if n == 1 {
		voterMenu(e)
	} else {
		adminMenu(e)
	}
}

// FUNC SETTIME (FOR SET THE TIME OF VOTING, IF THE TIME OUT OF EXPETING/SET TIME, THE VOTING PROCESS WILL BE ENDED)
func setTime(e *Election) {

	fmt.Println("Available Hours:")
	for i := 0; i < len(e.Times); i++ {
		fmt.Printf("%d. %d (%v)\n", i+1, e.Times[i].hours, e.Times[i].checktime)
	}

	waktu := &e.Times[0]

	fmt.Print("Enter new hours: ")
	fmt.Scan(&waktu.hours)
	fmt.Print("Enter new checktime: ")
	fmt.Scan(&waktu.checktime)

	fmt.Println("Time edited successfully.")
	adminMenu(e)
}

// FUNC SORTCANDIDATE (SORTING BY VOTE COUNT)
func sortCandidatesByVoteCount(candidates []Candidate) {
	length := len(candidates)
	for i := 0; i < length-1; i++ {
		maxIndex := i
		for j := i + 1; j < length; j++ {
			if candidates[j].VoteCount > candidates[maxIndex].VoteCount {
				maxIndex = j
			}
		}
		if maxIndex != i {
			temp := candidates[i]
			candidates[i] = candidates[maxIndex]
			candidates[maxIndex] = temp
		}
	}
}
