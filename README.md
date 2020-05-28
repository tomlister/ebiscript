# Commercium

## This is the client for my trading game.

The main premise of the game is that you are a trader in space but your trades and market data are delayed due to latency.

### I wrote a custom hot reloadable engine for this game in order to improve development.

It features:

* Asset manifest system
* VM Entitlements (security)
* JS & Go Engine bindings
* Instant hot reloading
* More to come.

### VM Quirks/Conventions
In order to conserve persistence of program states across hot reloads I've included state headers and footers.
```js
#state
var foo = "bar";
#endstate
```
Anything inbetween these two comments will only be loaded on vm init and won't be reset if hot reloaded.