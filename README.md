# csv-filter-challenge-public

## How-To
First, build the project with `go build`

This CSV tool has three modes: 
- `IMPLICIT` - User does not have to declare field names. Regex filters are used in the order that they appear in the .csv file. Missing filters will match all.
- `EXPLICIT` - User declares each field name explicitly. Missing regex filters will match all.
- `PROMPT` - User is prompted to enter a regex expression for each field in the .csv file.

Usage: \
`./csv-filter-challenge -f=[CSV_PATH] [flags] [optional_flags]` 

Flags: \
`-i`, Runs in implicit mode \
`-e`, Runs in explicit mode \
`-p`, Runs in prompt mode 

Optional Flags: \
`[regex0] [regex1] ...`, For use in implicit mode, defines regex filters in order of .csv fields. \
`[FIELD_NAME0]=[regex0] [FIELD_NAME1]=[regex1] ...`, For use in explicit mode, defines regex filters for each respective .csv field 

More Examples: \
`./csv-filter-challenge -f=data.csv -i Ken Thompson 19430204` \
`./csv-filter-challenge -f=data.csv -e first_name=Ken last_name=Thomp*` \
`./csv-filter-challenge -f=data.csv -p `


## Assumptions
- CSV file is syntactically correct.
- CSV file always has three fields:
	- `first_name` (string)
	- `last_name` (string)
	- `dob` (string)
- CSV file always has data in each field (no empty string)
- Each CSV line is under 64kb due to `bufio`'s default Scanner size.
- Assumes user enters in flags correctly. Currently, no error handling for flags.

# Instructions
1. Click "Use this template" to create a copy of this repository in your personal github account.  
1. Using technology of your choice, complete assignment listed below (we use [GoLang](https://go.dev/) at Scoir, so if you know Go, show it off, but that's not required!).
1. Update the README in your new repo with:
    * a `How-To` section containing any instructions needed to execute your program.
    * an `Assumptions` section containing documentation on any assumptions made while interpreting the requirements.
1. Send an email to Scoir (sbeyers@scoir.com and msullivan@scoir.com) with a link to your newly created repo containing the completed exercise (preferably no later than one day before your interview).

## Expectations
1. This exercise is meant to drive a conversation. 
1. Please invest only enough time needed to demonstrate your approach to problem solving, code design, etc.
1. Within reason, treat your solution as if it would become a production system.

## Assignment
Create a command line application that parses a CSV file and filters the data per user input.

The CSV will contain three fields: `first_name`, `last_name`, and `dob`. The `dob` field will be a date in YYYYMMDD format.

The user should be prompted to filter by `first_name`, `last_name`, or birth year. The application should then accept a name or year and return all records that match the value for the provided filter. 

Example input:
```
first_name,last_name,dob
Bobby,Tables,19700101
Ken,Thompson,19430204
Rob,Pike,19560101
Robert,Griesemer,19640609
```
