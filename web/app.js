'use strict';

// Configuration //
var speed = 200; // default scroll speed
const fontSizes = [
	"24px", // tiny
	"48px", // small
	"72px", // normal
];
const align = ["left", "center", "right"];

// // // // //

const obj = document.getElementById("marquee");
let play = false;
let curScroll = 0;
let curSize = 48; // script size in pixels
let scroller = setInterval(marquee, speed);

function marquee() {
	if(!play) { return; }
	curScroll = obj.scrollTop;
	if (curScroll + obj.clientHeight == document.clientHeight) {
		obj.scrollTo(0, 0);
	} else {
		obj.scrollTo({
			top: curScroll + (5000/speed),
			behavior: "smooth",
		});
	}
}
obj.addEventListener("wheel", (e) => { if(play) { e.preventDefault(); } });

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
	if(!play) { return }
	e.preventDefault();
	if (e.deltaY) {
		speed += curSize*0.25;
	} else {
		speed -= curSize*0.25;
	}
	//speed += e.deltaY * 10;
	if (speed > 3000) { speed = 3000; }
	if (speed < 5) { speed = 5; }
	console.log(e.deltaY);
	console.log(speed);
	console.log("destroy scroller");
	clearInterval(scroller);
	scroller = setInterval(marquee, speed);
	console.log("built scroller");
});
document.getElementById("marquee_text").addEventListener("paste", (e) => {
	e.preventDefault();
    let text = (e.originalEvent || e).clipboardData.getData('text/plain');
    console.log(text);
    document.execCommand("insertHTML", false, text);
});

document.getElementById("size").addEventListener("change", () => {
	const i = document.getElementById("size").selectedIndex;
	obj.style.fontSize = fontSizes[i];
	curSize = fontSizes[i].replace("px", "");
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
