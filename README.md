# CoreSync - API helper

```go
import "github.com/AldieNightStar/coresync"

// Create Command registry
reg := make(coresync.CommandRegistry)

// Add some commands
// Command has:
//   auth - user token/name to allow request secured commands (Which are not allowed for non-login users)
//   args - array of strings. Arguments for your api command
// Command returns:
//   status (int)      - status code
//   messsage (string) - message (Response text)
reg["abc"] = func(auth string, args []string) (int, string) {
    return len(args), strings.Join(args, "; ")
}

// Start the Server
// You can choose which type you want to serve: HTTP, TCP, WebSocket (Soon), Custom

// HTTP Server
coresync.ServeHttp("0.0.0.0", 8080, reg)

// TCP Server
coresync.ServeSocket("0.0.0.0", 4444, reg)

// WebSocket (Not available currently)
// Do not use it for now
coresync.ServeWebSocket("0.0.0.0", 4444, reg)


// Raw Server (Without any implementation)
// Could be used to embed in your custom protocol server
// Be it console or even telnet. Or you want to transfer API via Sound :D
// Returns:
//   c  - CallBack function with CommandDTO as request and ResponseDTO as a response
c := coresync.ServeRaw(reg)

c(&coresync.CommandDTO{"commandName", []string{"Arg1", "Arg2"}, "AUTH"})
```

# Login / Register

* You need to create your own login/register system
* How it should work?
    * Create for example `register` command which returns status: `coresync.ResponseStatusAuthSetup` and message with new `Auth` string
    * That status tells client to change current `Auth` string
    * Then you can validate that `Auth` string any way you want. Be it login/password or hex token

# Register / login sample
```go
// In this code sample shows to write Auth string like login::passw
// It is not recommended as it very easy to hack.
// Use generated tokens and validate that tokens with DB ones.
// Also that token can be regenerated each login to keep logined limited range of devices

reg["reigster"] = func(auth string, args []string) (int, string) {
    if len(args) < 2 {
        return coresync.ResponseStatusNotEnoughArguments, "Args len < 2"
    }
    login := args[0]
    passw := args[1]
    // Some DB api
    // You can have any you want. Just a simple example
    if !DB.IsExist(login) {
        DB.Register(login, passw)
        // Very important for status to be ResponseStatusAuthSetup
        // Because any other statuses will not change Auth string for the client
        return coresync.ResponseStatusAuthSetup, login+"::"+passw
    } else {
        return coresync.ResponseStatusNotDone, "This user is registered"
    }
}

reg["login"] = func(auth string, args []string) (int, string) {
    if len(args) < 2 {
        return coresync.ResponseStatusNotEnoughArguments, "Args len < 2"
    }
    login := args[0]
    passw := args[1]
    // Some DB api
    // You can have any you want. Just a simple example
    ok := DB.LogIn(login, passw)
    if !ok {
        return coresync.ResponseStatusNotDone, "Wrong creds"
    } else {
        return coresync.ResponseStatusAuthSetup, login+"::"+passw
    }
}
```