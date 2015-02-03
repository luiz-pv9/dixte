Dixte Analytics
===============

Events and profiles analytics written in Go using PostgreSQL. The goal of this
project is 

### Overview


### Architecture

### Packages

* `jsonvalidation`

**Why?**
Go's JSON unmarshal works great with struct type, and you can validate the arguments by simply checking some `nil` values or empty strings to see if there are some missing fields. But when the structure is unkown, (e.g event and profile properties) it becomes hard to validate some unauthorized values, and easy to forget some case.

**How it works?**
The main function of `jsonvalidation` is `CleanObject(data, rules)` that returns a "safe" object with only the attributes that passed in the specified rules.

* `databasemanager`
* `environment`
* `properties`
* `apps`

* `events` and `profiles`
	
	In development