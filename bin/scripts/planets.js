var drawPlanets = function() {
    var mousex = getMouseX()
    var mousey = getMouseY()
    var xoffset = ((mousex-160)/320)*10
    var yoffset = ((mousey-120)/240)*10
    var buttonbgcolor = "#ffffff"
    var buttonx = 320-30
    var buttony = 10
    var buttonw = 20
    var buttonh = 20
    if (mousex >= buttonx && mousex <= buttonx+buttonw) {
        if (mousey >= buttony && mousey <= buttony+buttonh) {
            buttonbgcolor = "#5555ff"
            if (isLeftMouseDown()) {
                currentscreen = "title"
            }
        }
    }
    clearObjects()
    drawImage("bg0", 0+xoffset, 0+yoffset)
    drawText("IMT - Interplanetary Market Terminal", "#ffffff", 5, 10, 25)
    drawText("NetLink Online", "#55ff55", 4, 10, 40)
    drawText("QuantumLink Not Found", "#ff5555", 4, 10, 55)
    drawSolidImage(buttonbgcolor, 320-30, 10, 20, 20, 0.5)
    drawText("x", "#000000", 10, 320-28, 28)
    drawText("InterplanID - (X66urf)", "#ffffff", 4, 10, 240-10)
    drawPlanetCard(20, 120-(80/2), "Earth", "earth", 1000)
    drawPlanetCard(120, 120-(80/2), "Mars", "mars", 178120000-1000)
    drawPlanetCard(220, 120-(80/2), "Europa", "europa", 628300000-1000)
 }

var drawPlanetCard = function(x, y, name, img, distance) {
    var speedoflight = 299792.485 // km/s
    var mousex = getMouseX()
    var mousey = getMouseY()
    var cardbgcolor = "#ffffff"
    var cardtextcolor = "#ffffff"
    if (mousex >= x && mousex <= x+80) {
        if (mousey >= y && mousey <= y+80) {
            cardbgcolor = "#5555ff"
            cardtextcolor = "#5555ff"
            if (isLeftMouseDown()) {
            }
        }
    }
    drawSolidImage(cardbgcolor, x, y, 80, 80, 0.5)
    drawImage(img, x, y)
    drawText(name, cardtextcolor, 4, x, y-2)
    drawText(Math.round((distance/speedoflight)*1000)+"ms", cardtextcolor, 4, x, y+90)
}