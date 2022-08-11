'use strict';

// Configuration //
const speeds = [
	"256", // slow
	"128", // normal
	"32", // fast
];
const fontSizes = [
	"1rem", // small
	"3rem", // normal
	"5rem", // large
];
const align = ["left", "center", "right"];

// // // // //

const obj = document.getElementById("marquee");
let play = false;
let curScroll = 0;
let speed = speeds[1];
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
	play = !play;
}, false); // Handle scrolling start and stop

document.getElementById("speed").addEventListener("change", () => {
	const i = document.getElementById("speed").selectedIndex;
	speed = speeds[i];
	console.log("destroy scroller");
	clearInterval(scroller);
	scroller = setInterval(marquee, speed);
	console.log("built scroller");
}, false); // Handle speed changes

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
