package data

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pedro-git-projects/call-of-gopher/internal/data/validator"
)

type Investigator struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Residence   string `json:"residence"`
	Birthplace  string `json:"birthplace"`
	Occupation  string `json:"occupation"`
	Str         int    `json:"str"`
	Con         int    `json:"con"`
	Pow         int    `json:"pow"`
	Dex         int    `json:"dex"`
	App         int    `json:"app"`
	Siz         int    `json:"siz"`
	Int         int    `json:"int"`
	Edu         int    `json:"edu"`
	Luck        int    `json:"luck"`
	Mp          int    `json:"mp"`
	Db          int    `json:"db"`
	Build       int    `json:"build"`
	Hp          int    `json:"hp"`
	San         int    `json:"san"`
	Mv          int    `json:"mv"`
	Description Description
}

func (i *Investigator) DetermineCharacteristics() {
	i.Str = rollStr()
	i.Con = rollCon()
	i.Dex = rollDex()
	i.App = rollApp()
	i.Pow = rollPow()
	i.setSanity()
	i.setMp()
	i.Siz = rollSiz()
	i.setHp()
	i.Int = rollInt()
	i.Edu = rollEdu()
	i.Luck = rollLuck()
	i.setMoveRate()
}

func (i *Investigator) SetName(name string) {
	i.Name = name
}

func (i *Investigator) SetOccupation(occupation string) {
	i.Occupation = occupation
}

func (i *Investigator) SetBirthplace(birthplace string) {
	i.Birthplace = birthplace
}

func (i *Investigator) SetResidence(residence string) {
	i.Residence = residence
}

func (i *Investigator) SetAge(age int) {
	i.Age = age
}

func (i *Investigator) setSanity() {
	i.San = i.Pow
}

func (i *Investigator) setMp() {
	i.Mp = (i.Pow / 5)
}

func (i *Investigator) setHp() {
	i.Hp = (i.Siz + i.Con) / 10
}

func (i *Investigator) setMoveRate() {
	if i.Dex < i.Siz && i.Str < i.Siz {
		i.Mv = 7
	} else if i.Str >= i.Siz || i.Dex >= i.Siz {
		i.Mv = 8
	} else if i.Str > i.Siz && i.Dex > i.Siz {
		i.Mv = 9
	}
}

func rollStr() int {
	d6 := Dice{NumberOfSides: D6}
	str := (3 * d6.Roll()) * 5
	return str
}

func rollCon() int {
	d6 := Dice{NumberOfSides: D6}
	con := (3 * d6.Roll()) * 5
	return con
}

func rollDex() int {
	d6 := Dice{NumberOfSides: D6}
	dex := (3 * d6.Roll()) * 5
	return dex
}

func rollApp() int {
	d6 := Dice{NumberOfSides: D6}
	app := (3 * d6.Roll()) * 5
	return app
}

func rollPow() int {
	d6 := Dice{NumberOfSides: D6}
	pow := (3 * d6.Roll()) * 5
	return pow
}

func rollSiz() int {
	d6 := Dice{NumberOfSides: D6}
	siz := ((2*d6.Roll() + 6) * 5)
	return siz
}

func rollInt() int {
	d6 := Dice{NumberOfSides: D6}
	intelligence := ((2*d6.Roll() + 6) * 5)
	return intelligence
}

func rollEdu() int {
	d6 := Dice{NumberOfSides: D6}
	edu := ((2*d6.Roll() + 6) * 5)
	return edu
}

func rollLuck() int {
	d6 := Dice{NumberOfSides: D6}
	luck := (3 * d6.Roll()) * 5
	return luck
}

func (i *Investigator) eduImprovementCheck() {
	d100 := Dice{NumberOfSides: D100}
	roll := d100.Roll()
	tenthOfRoll := (roll / 10)
	if roll > i.Edu && (i.Edu+tenthOfRoll <= 99) {
		i.Edu += tenthOfRoll
	}
}

func isGreaterThanZero(status int, debuff int) bool {
	if status-debuff >= 0 {
		return true
	}
	return false
}

// finds the highest non-negative debuff value within range
func findOptimumDebuff(maxDebuff int, status int) int {
	dynamicDebuff := 0
	for (dynamicDebuff <= maxDebuff) && (status-dynamicDebuff >= 0) {
		dynamicDebuff++
	}
	// by one correction because of loop break
	return dynamicDebuff - 1
}

func (i *Investigator) youngModifier() {
	strDebuff := findOptimumDebuff(2, i.Str)
	i.Str = i.Str - strDebuff

	sizDebuff := findOptimumDebuff(2, i.Siz)
	i.Siz = i.Siz - sizDebuff
}

func (i *Investigator) fortiesModifier() {
	strDebuff := findOptimumDebuff(1, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(2, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(2, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(1, i.Mv)
	i.Mv = i.Mv - mvDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) fifthtiesModifier() {
	strDebuff := findOptimumDebuff(3, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(3, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(3, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(2, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(10, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) sixtiesModifier() {

	strDebuff := findOptimumDebuff(7, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(1, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(7, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(3, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(15, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) seventiesModifier() {

	strDebuff := findOptimumDebuff(13, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(13, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(13, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(4, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(20, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) eightiesModifier() {

	strDebuff := findOptimumDebuff(13, i.Str)
	i.Str = i.Str - strDebuff

	conDebuff := findOptimumDebuff(13, i.Con)
	i.Con = i.Con - conDebuff

	dexDebuff := findOptimumDebuff(13, i.Dex)
	i.Dex = i.Dex - dexDebuff

	mvDebuff := findOptimumDebuff(5, i.Mv)
	i.Mv = i.Mv - mvDebuff

	appDebuff := findOptimumDebuff(25, i.App)
	i.App = i.App - appDebuff

	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
	i.eduImprovementCheck()
}

func (i *Investigator) AccountForAgeModifiers() {
	switch i.Age >= 15 {
	case i.Age >= 15 && i.Age <= 19:
		i.youngModifier()
	case i.Age >= 20 && i.Age <= 39:
		i.eduImprovementCheck()
	case i.Age >= 40 && i.Age <= 49:
		i.fortiesModifier()
	case i.Age >= 50 && i.Age <= 59:
		i.fifthtiesModifier()
	case i.Age >= 60 && i.Age <= 69:
		i.sixtiesModifier()
	case i.Age >= 70 && i.Age <= 79:
		i.seventiesModifier()
	case i.Age >= 80:
		i.eightiesModifier()
	}
}

func (i *Investigator) SetDamageBonusAndBuild() {
	db := i.Str + i.Siz
	d4 := Dice{NumberOfSides: D4}
	d6 := Dice{NumberOfSides: D6}
	switch db >= 2 {
	case db <= 64:
		i.Db = -2
		i.Build = -2
	case db >= 65 && db <= 84:
		i.Db = -1
		i.Build = -1
	case db >= 85 && db <= 124:
		i.Db = 0
		i.Build = 0
	case db >= 125 && db <= 164:
		i.Db = d4.Roll()
		i.Build = 1
	case db >= 165 && db <= 204:
		i.Db = d6.Roll()
		i.Build = 2
	case db >= 205 && db <= 284:
		i.Db = 2 * d6.Roll()
		i.Build = 3
	case db >= 285 && db <= 364:
		i.Db = 3 * d6.Roll()
		i.Build = 4
	case db >= 365 && db <= 444:
		i.Db = 4 * d6.Roll()
		i.Build = 5
	case db >= 445:
		i.Db = 5 * d6.Roll()
		i.Build = 6
	}
}

func PrintInvestigator(i Investigator) {
	fmt.Printf("Your name is %s\n", i.Name)
	fmt.Printf("Your age is %d\n", i.Age)
	fmt.Printf("Your birthplace is %s\n", i.Birthplace)
	fmt.Printf("Your residence  is %s\n", i.Residence)
	fmt.Printf("Your occupation  is %s\n", i.Occupation)
	fmt.Printf("Your strengh is %d\n", i.Str)
	fmt.Printf("Your constitution is %d\n", i.Con)
	fmt.Printf("Your power is %d\n", i.Pow)
	fmt.Printf("Your dexterity is %d\n", i.Dex)
	fmt.Printf("Your appearence is %d\n", i.App)
	fmt.Printf("Your size is %d\n", i.Siz)
	fmt.Printf("Your intelligence is %d\n", i.Int)
	fmt.Printf("Your education %d\n", i.Edu)
	fmt.Printf("Your luck is %d\n", i.Luck)
	fmt.Printf("Your magic points are %d\n", i.Mp)
	fmt.Printf("Your damage bonus is %d\n", i.Db)
	fmt.Printf("Your hit points are %d\n", i.Hp)
	fmt.Printf("Your sanity is %d\n", i.San)
	fmt.Printf("Your move rate is %d\n", i.Mv)
	fmt.Printf("With regards to strenght you are %s\n", i.Description.StrDescription)
	fmt.Printf("You%s\n", i.Description.ConDescription)
	fmt.Printf("With regards to appearence you are %s\n", i.Description.AppDescription)
	fmt.Printf("You %s\n", i.Description.IntDescription)
	fmt.Printf("You are%s\n", i.Description.SizDescription)
	fmt.Printf("You%s\n", i.Description.PowDescription)
	fmt.Printf("You%s\n", i.Description.DexDescription)
	fmt.Printf("You%s\n", i.Description.EduDescription)
}

func (i Investigator) WriteInvetigatorWeb(w http.ResponseWriter) {
	fmt.Fprintf(w, "Your name is %s\n", i.Name)
	fmt.Fprintf(w, "Your age is %d\n", i.Age)
	fmt.Fprintf(w, "Your birthplace is %s\n", i.Birthplace)
	fmt.Fprintf(w, "Your residence  is %s\n", i.Residence)
	fmt.Fprintf(w, "Your occupation  is %s\n", i.Occupation)
	fmt.Fprintf(w, "Your strengh is %d\n", i.Str)
	fmt.Fprintf(w, "Your constitution is %d\n", i.Con)
	fmt.Fprintf(w, "Your power is %d\n", i.Pow)
	fmt.Fprintf(w, "Your dexterity is %d\n", i.Dex)
	fmt.Fprintf(w, "Your appearence is %d\n", i.App)
	fmt.Fprintf(w, "Your size is %d\n", i.Siz)
	fmt.Fprintf(w, "Your intelligence is %d\n", i.Int)
	fmt.Fprintf(w, "Your education %d\n", i.Edu)
	fmt.Fprintf(w, "Your luck is %d\n", i.Luck)
	fmt.Fprintf(w, "Your magic points are %d\n", i.Mp)
	fmt.Fprintf(w, "Your damage bonus is %d\n", i.Db)
	fmt.Fprintf(w, "Your build is %d\n", i.Build)
	fmt.Fprintf(w, "Your hit points are %d\n", i.Hp)
	fmt.Fprintf(w, "Your sanity is %d\n", i.San)
	fmt.Fprintf(w, "Your move rate is %d\n", i.Mv)
	fmt.Fprintf(w, "With regards to strenght you are %s\n", i.Description.StrDescription)
	fmt.Fprintf(w, "You%s\n", i.Description.ConDescription)
	fmt.Fprintf(w, "With regards to appearence you are %s\n", i.Description.AppDescription)
	fmt.Fprintf(w, "You %s\n", i.Description.IntDescription)
	fmt.Fprintf(w, "You are%s\n", i.Description.SizDescription)
	fmt.Fprintf(w, "You%s\n", i.Description.PowDescription)
	fmt.Fprintf(w, "You%s\n", i.Description.DexDescription)
	fmt.Fprintf(w, "You%s\n", i.Description.EduDescription)
}

func CreateInvestigator(i *Investigator, name string, age int, birth string, residence string, occupation string) *Investigator {
	i.DetermineCharacteristics()
	i.SetName(name)
	i.SetAge(age)
	i.SetBirthplace(birth)
	i.SetResidence(residence)
	i.SetOccupation(occupation)
	i.AccountForAgeModifiers()
	i.SetDamageBonusAndBuild()
	i.SetDescription()
	return i
}

func (i *Investigator) WriteHandler(w http.ResponseWriter, r *http.Request) {
	i.WriteInvetigatorWeb(w)
}

func ValidateInvestigator(v *validator.Validator, investigator *Investigator) {
	v.Check(investigator.Name != "", "name", "the name field must be provided")
	v.Check(len(investigator.Name) <= 500, "name", "the name field must not be greater than 500 bytes")

	v.Check(validator.Matches(strconv.Itoa(investigator.Age), validator.AgeRX), "age", "age must be in the closed 1-99 range")
	//	v.Check(investigator.Age != 0, "age", "must be provided")
	//  v.Check(investigator.Age >= 0, "age", "invalid age, please enter in the range 1-99")
	//  v.Check(investigator.Age <= 99, "age", "invalid age, please enter in the range 1-99")

	v.Check(len(investigator.Residence) <= 500, "residence", "the residence field must not be greater than 500 bytes")
	v.Check(len(investigator.Birthplace) <= 500, "birthplace", "the birthplace field must not be greater than 500 bytes")
	v.Check(len(investigator.Occupation) <= 500, "occupation", "the occupation field must not be greater than 500 bytes")

}
