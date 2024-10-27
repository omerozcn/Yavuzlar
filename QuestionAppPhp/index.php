<?php
session_start();
?>
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Quest App</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="login-container">
        <h1>Hoşgeldiniz</h1>
        <button class="main-button" onclick="window.location.href='login.php'">Giriş Yap</button>
        <button class="main-button" onclick="window.location.href='register.php'">Kayıt Ol</button>
        <button class="main-button" onclick="window.location.href='scoreboard.php'">Puan Tablosu</button>
    </div>
</body>
</html>
