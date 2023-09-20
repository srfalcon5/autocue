'use strict';

// Configuration //
var speed = 128; // default scroll speed
const fontSizes = [
	"1rem", // tiny
	"3rem", // small
	"5rem", // normal
];
const align = ["left", "center", "right"];

// // // // //

const obj = document.getElementById("marquee");
let play = false;
let curScroll = 0;
let scroller = setInterval(marquee, speed);

function marquee() {
	if(play) {
		curScroll = obj.scrollTop;
		if (curScroll + obj.clientHeight == document.clientHeight) {
			obj.scrollTo(0, 0);
		} else {
			obj.scrollTo({
				top: curScroll + 5,
				behavior: "smooth",
			});
		}
	}
}

document.getElementById("toggle").addEventListener("click", () => {
	// Visual button state
	if(!play) {
		document.getElementById("toggle").innerHTML = "Stop";
		document.getElementById("toggle").style.background = "#F34";
	}
	else {
		document.getElementById("toggle").innerHTML = "Start";
		document.getElementById("toggle").style.background = "#2B3";
	}

	// Actually start marquee
	play = !play;
}, false); // Handle scrolling start and stop
document.addEventListener("wheel", (e) => {
	if(play) {
		e.preventDefault();
		speed += e.deltaY * 10;
		if (speed > 3000) { speed = 3000; }
		if (speed < 10) { speed = 10; }
		console.log(speed);
		console.log("destroy scroller");
		clearInterval(scroller);
		scroller = setInterval(marquee, speed);
		console.log("built scroller");
	}
})

/*document.getElementById("speed").addEventListener("change", () => {
	const i = document.getElementById("speed").selectedIndex;
	speed = speeds[i];
	console.log("destroy scroller");
	clearInterval(scroller);
	scroller = setInterval(marquee, speed);
	console.log("built scroller");
}, false); // Handle speed changes*/

document.getElementById("size").addEventListener("change", () => {
	const i = document.getElementById("size").selectedIndex;
	obj.style.fontSize = fontSizes[i];
}, false); // Handle text size changes

document.getElementById("align").addEventListener("change", () => {
	const i = document.getElementById("align").selectedIndex;
	obj.style.textAlign = align[i];
}, false); // Handle alignment changes

document.getElementById("mirror").addEventListener("change", () => {
	const i = document.getElementById("mirror").selectedIndex;
	if (i == 0) {
		obj.style.transform = "scale(1,1)";
	} else if (i == 1) {
		obj.style.transform = "scale(-1,1)";
	}
}, false); // Handle mirror changes
