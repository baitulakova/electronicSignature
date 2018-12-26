package vars

var LoginPage=[]byte(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>Регистрация</title>
	</head>
	<body>
	<h3>Форма регистрации</h3>
	<form action="/member" method="post">
		Имя:
		<input type="text" name="fname">
		<br><br>
		Фамилия:
		<input type="text" name="lname">
		<br><br>
		<input type="submit" value="Зарегистрировать">
	</form>
	</body>
	</html>
`)

var RegisterAnswer =string(
	`<!DOCTYPE html>
<html>
<head>
<title>Регистрация</title>
</head>
<body>
Успешно создан участник: 
<br>
Имя: {{.FName}} 
<br>
Фамилия: {{.LName}}
<br>
<br>
Ваш публичный ключ: {{.PubKey}}
</body>
</html>
`)

var CipherPage=[]byte(`
	<!DOCTYPE html>
	<html>
	<head>
	<title>Шифрование</title>
	</head>
	<body>
	<h3>Форма шфирования</h3>
	<form action="/decrypt" method="post">
		Введите текст:
		<input type="text" name="plainText">
		<br><br>
		Введите имя получателя:
		<input type="text" name="targetName">
		<br><br>
		<input type="submit" value="Зашифровать">
	</form>
	</body>
	</html>
`)

var CipherText =string(
	`<!DOCTYPE html>
<html>
<head>
<title>Результат шифрования</title>
</head>
<body>
Шифрованный текст: 
<br>
{{.cText}}
</body>
</html>
`)





