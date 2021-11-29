/*
make_lib.go
Description:
	An implementation of the file make_lib.go written entirely in go.
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MakeLibFlags struct {
	GurobiHome  string // Directory where Gurobi is installed
	GoFilename  string // Name of the Go File to
	PackageName string // Name of the package

}

type GurobiVersionInfo struct {
	MajorVersion    int
	MinorVersion    int
	TertiaryVersion int
}

type LPSolveVersionInfo struct {
	MajorVersion int
	MinorVersion int
}

func GetDefaultMakeLibFlags() (MakeLibFlags, error) {
	// Create Default Struct
	mlf := MakeLibFlags{
		GurobiHome:  "/Library/gurobi903/mac64",
		GoFilename:  "solvers/lib.go",
		PackageName: "solvers",
	}

	// Search through Mac Library for all instances of Gurobi
	libraryContents, err := os.ReadDir("/Library")
	if err != nil {
		return mlf, err
	}
	gurobiDirectories := []string{}
	for _, content := range libraryContents {
		if content.IsDir() && strings.Contains(content.Name(), "gurobi") {
			fmt.Println(content.Name())
			gurobiDirectories = append(gurobiDirectories, content.Name())
		}
	}

	// Convert Directories into Gurobi Version Info
	gurobiVersionList, err := StringsToGurobiVersionInfoList(gurobiDirectories)
	if err != nil {
		return mlf, err
	}

	fmt.Println(gurobiVersionList)

	highestVersion, err := FindHighestVersion(gurobiVersionList)
	if err != nil {
		return mlf, err
	}

	fmt.Println(highestVersion)

	// Write the highest version's directory into the GurobiHome variable
	mlf.GurobiHome = fmt.Sprintf("/Library/gurobi%v%v%v/mac64", highestVersion.MajorVersion, highestVersion.MinorVersion, highestVersion.TertiaryVersion)

	return mlf, nil

}

/*
StringToGurobiVersionInfo
Assumptions:
	Assumes that a valid gurobi name is given.
*/
func StringToGurobiVersionInfo(gurobiDirectoryName string) (GurobiVersionInfo, error) {
	//Locate major and minor version indices in gurobi directory name
	majorVersionAsString := string(gurobiDirectoryName[len("gurobi")])
	minorVersionAsString := string(gurobiDirectoryName[len("gurobi")+1])
	tertiaryVersionAsString := string(gurobiDirectoryName[len("gurobi")+2])

	// Convert using strconv to integers
	majorVersion, err := strconv.Atoi(majorVersionAsString)
	if err != nil {
		return GurobiVersionInfo{}, err
	}

	minorVersion, err := strconv.Atoi(minorVersionAsString)
	if err != nil {
		return GurobiVersionInfo{}, err
	}

	tertiaryVersion, err := strconv.Atoi(tertiaryVersionAsString)
	if err != nil {
		return GurobiVersionInfo{}, err
	}

	return GurobiVersionInfo{
		MajorVersion:    majorVersion,
		MinorVersion:    minorVersion,
		TertiaryVersion: tertiaryVersion,
	}, nil

}

/*
StringsToGurobiVersionInfoList
Description:
	Receives a set of strings which should be in the format of valid gurobi installation directories
	and returns a list of GurobiVersionInfo objects.
Assumptions:
	Assumes that a valid gurobi name is given.
*/
func StringsToGurobiVersionInfoList(gurobiDirectoryNames []string) ([]GurobiVersionInfo, error) {

	// Convert Directories into Gurobi Version Info
	gurobiVersionList := []GurobiVersionInfo{}
	for _, directory := range gurobiDirectoryNames {
		tempGVI, err := StringToGurobiVersionInfo(directory)
		if err != nil {
			return gurobiVersionList, err
		}
		gurobiVersionList = append(gurobiVersionList, tempGVI)
	}
	// fmt.Println(gurobiVersionList)

	return gurobiVersionList, nil

}

/*
// Iterate through all versions in gurobiVersionList and find the one with the largest major or minor version.
*/
func FindHighestVersion(gurobiVersionList []GurobiVersionInfo) (GurobiVersionInfo, error) {

	// Input Checking
	if len(gurobiVersionList) == 0 {
		return GurobiVersionInfo{}, errors.New("No gurobi versions were provided to FindHighestVersion().")
	}

	// Perform search
	highestVersion := gurobiVersionList[0]
	if len(gurobiVersionList) == 1 {
		return highestVersion, nil
	}

	for _, gvi := range gurobiVersionList {
		// Compare Major version numbers
		if gvi.MajorVersion > highestVersion.MajorVersion {
			highestVersion = gvi
			continue
		}

		// Compare minor version numbers
		if gvi.MinorVersion > highestVersion.MinorVersion {
			highestVersion = gvi
			continue
		}

		// Compare tertiary version numbers
		if gvi.TertiaryVersion > highestVersion.TertiaryVersion {
			highestVersion = gvi
			continue
		}
	}

	return highestVersion, nil

}

func ParseMakeLibArguments(mlfIn MakeLibFlags) (MakeLibFlags, error) {
	// Iterate through any arguments with mlfIn as the default
	mlfOut := mlfIn

	// Input Processing
	argIndex := 1 // Skip entry 0
	for argIndex < len(os.Args) {
		// Share parsing data
		fmt.Println("- Parsed input: %v", os.Args[argIndex])

		// Parse Inputs
		switch {
		case os.Args[argIndex] == "--gurobi-home":
			mlfOut.GurobiHome = os.Args[argIndex+1]
			argIndex += 2
		case os.Args[argIndex] == "--go-fname":
			mlfOut.GoFilename = os.Args[argIndex+1]
			argIndex += 2
		case os.Args[argIndex] == "--pkg":
			mlfOut.PackageName = os.Args[argIndex+1]
			argIndex += 2
		default:
			fmt.Printf("Unrecognized input: %v", os.Args[argIndex])
			argIndex++
		}

	}

	return mlfOut, nil
}

/*
CreateCXXFlagsDirective
Description:
	Creates the CXX Flags directive in the  file that we will use in lib.go.
*/
func CreateCXXFlagsDirective(mlfIn MakeLibFlags) (string, error) {
	// Create Statement
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	gurobiCXXFlagsString := fmt.Sprintf("// #cgo CXXFLAGS: --std=c++11 -I%v/include -I%v/include\n", mlfIn.GurobiHome, pwd)
	lpSolveCXXFlagsString := "// #cgo CXXFLAGS: -I/usr/local/opt/lp_solve/include\n" // Works as long as lp_solve was installed with Homebrew

	return fmt.Sprintf("%v%v", gurobiCXXFlagsString, lpSolveCXXFlagsString), nil
}

/*
CreatePackageLine
Description:
	Creates the "package" directive in the  file that we will use in lib.go.
*/
func CreatePackageLine(mlfIn MakeLibFlags) (string, error) {

	return fmt.Sprintf("package %v\n\n", mlfIn.PackageName), nil
}

/*
CreateLDFlagsDirective
Description:
	Creates the LD_FLAGS directive in the file that we will use in lib.go.
*/
func CreateLDFlagsDirective(mlfIn MakeLibFlags) (string, error) {
	// Constants
	AsGVI, err := mlfIn.ToGurobiVersionInfo()
	if err != nil {
		return "", err
	}

	// Locate the desired files for mac in the directory.
	// libContent, err := os.ReadDir(mlfIn.GurobiHome)
	// if err != nil {
	// 	return "", err
	// }

	ldFlagsDirective := fmt.Sprintf("// #cgo LDFLAGS: -L%v/lib", mlfIn.GurobiHome)

	targetedFilenames := []string{"gurobi_c++", fmt.Sprintf("gurobi%v%v", AsGVI.MajorVersion, AsGVI.MinorVersion)}

	for _, target := range targetedFilenames {
		ldFlagsDirective = fmt.Sprintf("%v -l%v", ldFlagsDirective, target)
	}
	ldFlagsDirective = fmt.Sprintf("%v \n", ldFlagsDirective)

	// Write the lp_solve LD Flags line.
	lvi, err := DetectLPSolveVersion()
	if err != nil {
		return "", err
	}
	ldFlagsDirective = fmt.Sprintf("%v// #cgo LDFLAGS: -L/usr/local/opt/lp_solve/lib -llpsolve%v%v\n", ldFlagsDirective, lvi.MajorVersion, lvi.MinorVersion)

	return ldFlagsDirective, nil
}

func (mlf *MakeLibFlags) ToGurobiVersionInfo() (GurobiVersionInfo, error) {
	// Split the GurobiHome variable by the name gurobi
	GurobiWordIndexStart := strings.Index(mlf.GurobiHome, "gurobi")
	GurobiDirNameIndexEnd := len(mlf.GurobiHome) - len("/mac64") - 1

	return StringToGurobiVersionInfo(string(mlf.GurobiHome[GurobiWordIndexStart : GurobiDirNameIndexEnd+1]))

}

/*
HeaderNameToLPSolveVersionInfo
Description:
	Converts the header file (like liblpsolve55.a) into an LPSolveVersionInfo object which can be used later.
*/
func HeaderNameToLPSolveVersionInfo(lpsolveHeaderName string) (LPSolveVersionInfo, error) {
	//Locate major and minor version indices in gurobi directory name
	majorVersionAsString := string(lpsolveHeaderName[len("liblpsolve")])
	minorVersionAsString := string(lpsolveHeaderName[len("liblpsolve")+1])

	// Convert using strconv to integers
	majorVersion, err := strconv.Atoi(majorVersionAsString)
	if err != nil {
		return LPSolveVersionInfo{}, err
	}

	minorVersion, err := strconv.Atoi(minorVersionAsString)
	if err != nil {
		return LPSolveVersionInfo{}, err
	}

	return LPSolveVersionInfo{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
	}, nil

}

func GetAHeaderFilenameFrom(dirName string) (string, error) {
	// Constants

	// Algorithm

	// Search through dirName directory for all instances of .a files
	libraryContents, err := os.ReadDir(dirName)
	if err != nil {
		return "", err
	}
	headerNames := []string{}
	for _, content := range libraryContents {
		if content.Type().IsRegular() && strings.Contains(content.Name(), ".a") {
			fmt.Println(content.Name())
			headerNames = append(headerNames, content.Name())
		}
	}

	return headerNames[0], nil

}

func DetectLPSolveVersion() (LPSolveVersionInfo, error) {
	// Constants
	homebrewLPSolveDirectory := "/usr/local/opt/lp_solve"

	// Algorithm
	headerFilename, err := GetAHeaderFilenameFrom(fmt.Sprintf("%v/lib/", homebrewLPSolveDirectory))
	if err != nil {
		return LPSolveVersionInfo{}, err
	}

	return HeaderNameToLPSolveVersionInfo(headerFilename)

}

func WriteLibGo(mlfIn MakeLibFlags) error {
	// Constants

	// Algorithm

	// First Create all Strings that we would like to write to lib.go
	// 1. Create package definition
	packageDirective, err := CreatePackageLine(mlfIn)
	if err != nil {
		return err
	}

	// 2. Create CXX_FLAGS argument
	cxxDirective, err := CreateCXXFlagsDirective(mlfIn)
	if err != nil {
		return err
	}

	// 3. Create LDFLAGS Argument
	ldflagsDirective, err := CreateLDFlagsDirective(mlfIn)
	if err != nil {
		return err
	}

	// Now Write to File
	f, err := os.Create("solvers/lib.go")
	if err != nil {
		return err
	}
	defer f.Close()

	// Write all directives to file
	_, err = f.WriteString(fmt.Sprintf("%v%v%v import \"C\"\n", packageDirective, cxxDirective, ldflagsDirective))
	if err != nil {
		return err
	}

	return nil

}
func main() {

	mlf, err := GetDefaultMakeLibFlags()

	// Next, parse the arguments to make_lib and assign values to the mlf appropriately
	mlf, err = ParseMakeLibArguments(mlf)

	fmt.Println(mlf)
	fmt.Println(err)

	// Write File
	err = WriteLibGo(mlf)
	if err != nil {
		fmt.Println(err)
	}

}
