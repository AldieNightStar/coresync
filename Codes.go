package coresync

// ResponseStatusOk - when everything is ok
const ResponseStatusOk = 1

// ResponseStatusUnknownError - When error, but not sure what the reason
const ResponseStatusUnknownError = 2

// ResponseStatusNotFound - Not found some element (For example document or file)
const ResponseStatusNotFound = 3

// ResponseStatusCanceled - When something was failed to do (For example started writing message and failed before finish)
const ResponseStatusCanceled = 4

// ResponseStatusNotDone - When request was not done. No any error. Just not done. Ignored
const ResponseStatusNotDone = 5

// ResponseStatusNotAllowed - Command not allowed for user/group etc
const ResponseStatusNotAllowed = 6

// ResponseStatusNoSuchCommand - Such command not found.
const ResponseStatusNoSuchCommand = 7

// ResponseStatusNotEnoughArguments - Not enough arguments for command
const ResponseStatusNotEnoughArguments = 8

// ResponseStatusWrongArguments - Some arguments (or all) are invalid (For example some argument should be number, but text)
const ResponseStatusWrongArguments = 9

// ResponseStatusNotAuth - when User has no Auth data (Need to login first)
const ResponseStatusNotAuth = 10

// ResponseStatusAuthSetup - Server send NEW auth string for the user (Then user can do request with that string)
const ResponseStatusAuthSetup = 11
