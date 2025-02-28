# Lotr-Companion-App CLI
A command-line interface (CLI) application written in Go that functions as a companion app. This application allows users to retrieve information about Lord of the Rings movies, books, and characters using The-One-Api (https://the-one-api.dev/).

<p>
    <a href="https://github.com/mechamjonathan/lotr-companion-app/workflows"><img src="https://github.com/mechamjonathan/lotr-companion-app/workflows/ci.yml/badge.svg" alt="CI Status"></a>
</p>

## Features 
- Retrieve a a list of _Lord of the Rings_ books, movies, characters.
- Get details about specific books, movies, characters.
- Browse character quotes with pagination
- Cache frequently accessed data for faster subsequent retrievals.

## Requirements
- Go (version 1.18 or later)
- Internet connection (to fetch data from The-One-Api)
  
## Installation 
1. Clone the Repository:
```
git clone https://github.com/MechamJonathan/lotr-companion-app
cd lotr-companion-app
go build
```

## Run Application
In the lotr-companion-app directory run the following command:
```
./lotr-companion-app
```
# Commands

| Command     | Description                                        |
| ----------- | -----------                                        |
| books       | Lists all books                                    |
| characters  | Lists all characters or group of characters        |
| details     | Return details about specific character, movie, or book     |
| exit        | Exit the program                                   |
| help        | Display help message and all available commands    |
| movies      | List all movies                                    |
| quotes      | View next page of a character's quotes             |
| quotesb     | View previous page of a character's quotes         |

## Example
```
Lotr-Companion-App > movies

  Movie Details       
  --------------------
   - The Lord of the Rings Series
   - The Hobbit Series
   - The Unexpected Journey
   - The Desolation of Smaug
   - The Battle of the Five Armies
   - The Two Towers
   - The Fellowship of the Ring
   - The Return of the King

Lotr-Companion-App > details The Two Towers

  Movie Details       
  --------------------
   - Name: The Two Towers
   - Runtime: 179 min
   - Budget: $94.00M
   - Box Office: $926.00M
   - Awards: 6 nominations, 2 wins
   - Rotten Tomatoes: 96.0%

Lotr-Companion-App > characters
  usage: characters <all> | <fellowship> | <hobbits> | <men> | <elves> | <dwarves> | <orcs> | <wizards> | <creatures>

Lotr-Companion-App > characters fellowship

  Fellowship Members  
  --------------------
   - Aragorn II Elessar
   - Boromir
   - Frodo Baggins
   - Gandalf
   - Gimli
   - Legolas
   - Meriadoc Brandybuck
   - Peregrin Took
   - Samwise Gamgee

Lotr-Companion-App > details Frodo Baggins

  Character Details   
  --------------------
   - Name: Frodo Baggins
   - WikiURL: http://lotr.wikia.com//wiki/Frodo_Baggins
   - Race: Hobbit
   - Birth: 22 September ,TA 2968
   - Gender: Male
   - Death: Unknown (Last sighting ,September 29 ,3021,) (,SR 1421,)
   - Hair: Brown
   - Height: 1.06m (3'6")
   - Realm: 
   - Spouse:

Lotr-Companion-App > quotes Frodo Baggins

  ------------------------------------------
  "Gandalf?"
  - Frodo Baggins
  ------------------------------------------
  
  ------------------------------------------
  "Oooohhh!"
  - Frodo Baggins
  ------------------------------------------
  
  ------------------------------------------
  "Gimli!"
  - Frodo Baggins
  ------------------------------------------
  
  ------------------------------------------
  "No, it isn't. It isn't midday yet. The days are growing darker."
  - Frodo Baggins
  ------------------------------------------

```
