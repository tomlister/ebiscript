var currentscreen = "main"
var main = function() {
    if (currentscreen == "main") {
        currentscreen = "title"
    } else if (currentscreen == "title") {
        drawTitle()
    } else if (currentscreen == "planets") {
        drawPlanets()
    } else if (currentscreen == "market_earth") {
        drawMarkets()
    } else {
        clearObjects()
        drawText("screen doesn't exist", "#FF0000", 4, (320-105)/2, 240/2)
    }
}