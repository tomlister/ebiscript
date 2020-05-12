var currentscreen = "main"
var main = function() {
    if (currentscreen == "main") {
        currentscreen = "title"
    } else if (currentscreen == "title") {
        drawTitle()
    } else if (currentscreen == "planets") {
        drawPlanets()
    }
}