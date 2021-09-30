// var canvas = document.getElementById("main-canvas");
// let deltaDeg = 1;
//
//
// let square = new Image();
// square.src = "/imgs/square.svg";
//
//
// let currentDeg = 0.5;
// if (canvas.getContext) {
// 	var ctx = canvas.getContext("2d");
// 	window.requestAnimationFrame(drawAnimation)
// }
//
// function drawAnimation()
// {
// 	ctx = document.getElementById('main-canvas').getContext('2d');
// 	ctx.imageSmoothingEnabled = false;
// 	ctx.clearRect(0,0, canvas.offsetWidth, canvas.offsetHeight);
//
// 	ctx.drawImage(square, 45, 25, 50, 50);
// 	// ctx.fillRect(45, 25, 50, 50); // Draw a rectangle with new settings
// 	// currentDeg += deltaDeg * 0.25;
// 	// if (currentDeg >= 360) {
// 	// 	currentDeg = currentDeg % 360;
// 	// }
// 	ctx.translate(45 + 25, 25 + 25);
// 	ctx.rotate((Math.PI / 180) * currentDeg);
//
// 	ctx.translate(-45 - 25, -25 - 25);
// 	ctx.save();                  // Save the current state
//
// 	window.requestAnimationFrame(drawAnimation);
// }