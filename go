<html>
<head>
	<meta charset='utf-8'>
  <title>바둑판</title>
	<style>
		button {width:300px; height:100px; float:right; font-size:70px; margin:40px;}

		table#pan {float:left; border-spacing: 0; border-collapse: collapse; position:absolute; left:25px; top:25px;}
		table#pan td {width:30px; height:33px; border:1px solid #888888; background-color:#ddcc00;}

		table#al {float:left; border-spacing: 0; border-collapse: collapse; position:absolute; z-index:2;}
		table#al td {width:30px; height:33px; border:1px solid rgba(255,255,255,0); border-radius:50px 50px; background-color:#000000; opacity:0;}
		table#al td:hover{background-color:#777777; opacity:0.6;}
		table#al td.b {background-color:black; opacity:1;}
		table#al td.w {background-color:#eeeeee; opacity:1;}
	</style>

	<script>
		var cc=0;
		function c(event){
			var t = event.target;
			if(t.classList.contains('b') || t.classList.contains('w')){
				t.classList.remove('b');
				t.classList.remove('w');
				cc = ++cc%2;
				return;
			}
			if(cc === 0) t.classList.add("b");
			else t.classList.add("w");
			cc = ++cc%2;

		}

		function reset(){
			var t = document.getElementsByTagName('td');
			for(var i=0;i<t.length;i++){
				t[i].classList.remove('b');
				t[i].classList.remove('w');
			}
			cc=0;
		}
	</script>
</head>

<body>
	<div id='container'>
		<table id='pan'>
			<tbody id='t'>
			</tbody>
		</table>
		
		<table id='al'>
			<tbody id='alt'>
			</tbody>
		</table>
		
		
		<button id='reset' value='reset' onclick='reset()'>RESET</button>
	</div>




	<script>
		var t = document.getElementById('t');
		var i,j, h="";
		for(i=0;i<18;i++){
			h += "<tr>";
			for(j=0;j<18;j++) h+="<td></td>";
			h += "</tr>";
		}
		t.innerHTML = h;

		t = document.getElementById('alt');
		h="";
		for(i=0;i<19;i++){
			h += "<tr>";
			for(j=0;j<19;j++) h+="<td onclick='c(event)'></td>";
			h += "</tr>";
		}
		t.innerHTML = h;
	</script>
</body>
</html>
