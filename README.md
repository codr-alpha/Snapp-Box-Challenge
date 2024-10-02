# Snapp-Box-Challenge

we have 4 directories

	-in input we put our input files and our program reads from this directory
	-in output we put our output files and our program writes in this directory
	-in generator there is main.go which generates random csv input file in input directory
	 it gets two number i n from terminal and generates a random csv input file with n rows and
	 the name input$i.csv in input directory
	 and there is a bash script which we can run main.go and generate random tests with it
	-in project we have 3 directories
		-in struct_and_constants we have our structs and constants and a method
		and a unit test to test it
		-in calculations we have the part of our code which is resposible for processing input 
		and writing output and a unit test to test it
		-in main we have the main part of our code which the program starts from it and read from input file
		and send it to Process function in calculation
		and a unit test to test it
		main reads a string from terminal for example str and looks for the input/$str.csv to read from
		in main we read from csv input file and parse each line to Point type and send it to a buffered channel where the Process in calculation get Points from this channel and if the id is different than the previous Point it replace it and if they are the same it calculates the fare between two Points and if the second Point is valid it replaces it and updates the fare of that id
		in Process when we finish our job with a id we send our information of that id to a channel which writeToCSV function reads from it and writes it to our output file

		i test the code with 4 Giga byte input csv file and it finished successfuly and take around 1 minute
		and test it with 16 Giga byte input csv file and it finished successfuly and take around 4 minutes