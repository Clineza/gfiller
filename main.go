/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "main/cmd"

func main() {
	cmd.Execute()
}

/*
	Let's plan this program!! What do we want this program to do?

	Ultimately, this program should allow users to create filler text that's repeated and be able to copy it to clipboard.
	I do want the program to create txt files or some data system to allow users to access their filler text.
	which allows them to always access it without having to recreate it.

	potential commands ---
	filler // this is the actual program name

	filler load mytext // this will look for a txt file called mytext.txt and load the file to read and write from.
		load will copy the text from mytext.txt to the temp.txt.. or create a new temp.txt if it doesn't exist.
		other commands won't work if temp.txt doesn't exist.. so create an error that says temp.txt doesn't exist. load your file again to ensure that it can be recreated.. something like that.
	filler save mytext // this will allow users to save the txt file that is currently loaded
		if this file exists. then it will create a new one..
		also note: if you load a food.txt and the save it to a different file name.. then it will copy its contents to that file
	filler copy // this allows the user to copy to clipboard from the currently loaded txt
	filler generate 1000 "this is the string" // generate will create that line the amount specified by the number flag.
		add a flag to allow for same line or new line capabilities.

	--maybe have a file name temp.txt.. this will be what's used as a buffer.. and don't let user's change it directly.
	it's not going to cause a problem if they rename it or anything.. if the program detects it no longer exists.. it will create a new temp.txt file.. and during each load of the program.. it will clear out what's currenly there.. or maybe it won't.. we'll see.
*/

/*
	UPDATES

	main command "gfiller"
	add commands --

	add, delete, ls, words,

*/
