var drawMarkets = function() {
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
                currentscreen = "planets"
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
}