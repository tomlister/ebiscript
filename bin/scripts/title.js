var drawTitle = function() {
   var mousex = getMouseX()
   var mousey = getMouseY()
   var xoffset = ((mousex-160)/320)*10
   var yoffset = ((mousey-120)/240)*10
   var buttonbgcolor = "#ffffff"
   var buttonx = 110
   var buttony = 175
   var buttonw = 100
   var buttonh = 35
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
   //drawDebugText("(x:"+mousex+",y:"+mousey+")")
   drawText("Commercium", "#ffffff", 9, 85, 120)
   drawSolidImage(buttonbgcolor, 110, 175, 100, 35, 0.5)
   drawText("Start", "#000000", 9, 128, 200)
}