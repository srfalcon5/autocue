'use strict';

// Configuration
const speeds = [
	"256", // slow
	"128", // normal
	"32", // fast
];
const fontSizes = [
	"1rem", // small
	"3rem", // normal
	"5rem" // large
];
const align = [
	"left",
	"center",
	"right"
];

// // // // //

let play = false;
let curScroll = 0;
let speed = speeds[1];
console.log("built scroller")
let scroller = setInterval(marquee, speed);

function marquee() {
	if(play == true) {
		const obj = document.getElementById("marquee");
		curScroll = obj.scrollTop;
		if (curScroll + obj.clientHeight == document.clientHeight) {
			console.log("0 0");
			obj.scrollTo(0, 0);
		} else {
			obj.scrollTo({
				top: curScroll + 5,
				behavior: "smooth"
			});
		}
	}
}

// Handle scrolling start and stop
document.getElementById("toggle").addEventListener("click", () => {
	play = !play;
	if (play) {
		console.log("scrolling");
	} else {
		console.log("not scrolling");
	}
	return false;
}, false);

// Handle speed changes
document.getElementById("speed").addEventListener("change", () => {
	const index = document.getElementById("speed").selectedIndex;
	speed = speeds[index];
	console.log("destroy scroller")
	clearInterval(scroller);
	scroller = setInterval(marquee, speed);
	console.log("built scroller")
	return false;
}, false);

// Handle text size changes
document.getElementById("size").addEventListener("change", () => {
	const index = document.getElementById("size").selectedIndex;
	document.getElementById('marquee').style.fontSize = fontSizes[index];
	return false;
}, false);

// Handle alignment changes
document.getElementById("align").addEventListener("change", () => {
	const index = document.getElementById("align").selectedIndex;
	document.getElementById("marquee").style.textAlign = align[index];
	return false;
}, false);

// Handle mirror changes
document.getElementById("mirror").addEventListener("change", () => {
	const index = document.getElementById("mirror").selectedIndex;
	const obj = document.getElementById("marquee"); 
	if (index == 0) {
		obj.style.transform = "scale(1,1)";
	} else if (index == 1) {
		obj.style.transform = "scale(-1,1)";
	}
	return false;
}, false);
