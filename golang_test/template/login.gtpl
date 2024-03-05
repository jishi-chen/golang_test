<html>
<head>
<title></title>
</head>
<body>
<form action="/login" method="post">
    使用者名稱:<input type="text" name="username">
    密碼:<input type="password" name="password">
    年齡:<input type="text" name="age">
    <div>
        <select name="fruit">
            <option value="apple">apple</option>
            <option value="pear">pear</option>
            <option value="banana">banana</option>
        </select>
    </div>
    <div>
        <input type="radio" name="gender" value="1">男
        <input type="radio" name="gender" value="2">女
    </div>
    <div>
        <input type="checkbox" name="interest" value="football">足球
        <input type="checkbox" name="interest" value="basketball">籃球
        <input type="checkbox" name="interest" value="tennis">網球
    </div>
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登入">
</form>
</body>
</html>