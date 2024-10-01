# SAP Event Broker Take Home Problem

The goal of this project was to implement a web project allowing users to search for trails based on various filters. The project was done in Go language and with HTML. 

## Files

-   go.mod

The go.mod file describes the module’s properties. It usually includes its dependencies on other modules and of versions of Go. In my go.mod file, the module path and the minimum version of Go required are listed. In the future, this would be where to put new dependencies on other modules if needed.

-   BoulderTrailHeads.csv

This is the original csv file containing all the different hike trails as well as their properties, such as whether there are restrooms, fishing, bike racks… etc. Every column is separated by a comma. This is where to add any new hike trails in the future.

-   main.go

This is the file where the main package is initialized as an executable command must always use the package main, and where the function main() is invoked. Most of the logic goes here for this project. 
The imports are at the top of the file, and this is where they can be updated if needed. For this specific project, a structure was created to handle easily the properties of a trail, called HIkingTrail. This needs to be updated if there are new criteria we want to use for the filtering of the trails. 

func init():
The init function is invoked only once and loads the html template to display the filtering UI. It also calls getTrailValues() and saves the data in a global variable.

getTrailsValues():
This function opens and reads the csv file containing the hiking Trails. With the data read, a map is created to obtain the column titles, which correspond to the properties’ names. A check is done to validate that all the properties are there in the csv file. This can be updated if the format of the csv changes or some field names are modified. An array is then created with all the different hiking trails from the csv file.

main():
This binds the path to localhost:3000 where the function homeHandler() will be invoked.

homeHandler():
This is the function getting the information from the user input, filtering it against the available hiking trails and then executing the HTML template to display the results.

filterTrails():
This function takes the array of trails saved from the csv file and compares each one with the list of criteria specified by the user. If the criteria is not specified by the user, the trail is valid for that property. If the criteria is specified, the trail’s property is then checked to validate if it corresponds. The same thing is done to each criteria, and a trail is only valid if all the criteria match. 

filterTrashCans():
This function takes the Trash Cans field which is a string and convert it to an int. It then compares to the criteria to ensure that the trail is valid if there is more or equal trash cans than required. 

filterDifficulty():
As the difficulties of the trails were not consistent, a check is done to see whether Easy, Moderate or Difficult is contained in the string. That way, for a search including the “Easy” difficulty, the results will include Easy, Easy-Moderate and East to Difficult. 

-   search.html

This is the HTML template that is used to display the web app. Some styling is included and more can be added in the future. The filtering display uses a form tag which allows to easily send and receive the user input. A table to display the results is only visible if there are any valid results.

## Improvements

In the future, a better UI can be created with a React Native project instead of a simple HTML template. A script to sort and clean the csv file is also another one. Based on customers’ needs, better criteria can be selected. The Difficulty field can also be improved once we know what the users want. Tests can and should be added, to ensure that the application is functional. 
