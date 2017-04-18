<html>
<head>
<title>Welcome to Library of The Seven Kingdom</title>
</head>
<body>
<h1>Book Store</h1>
<form action="/libraryList" method="post">
    Book Name : <input type="text" name="bookName">
    <input type="submit" value="Search">
    {{range .}}
        <br><br>
        <a href="https://en.wikipedia.org/wiki/{{.Name}}">{{.Name}}</a> - <a href="https://en.wikipedia.org/wiki/{{.Author}}">{{.Author}}</a> (<a href="https://en.wikipedia.org/wiki/{{.GenreS}}">{{.GenreS}}</a>)
    {{end}}
</form>
</body>
</html>