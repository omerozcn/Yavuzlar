<?php
session_start();
if ($_SESSION['rol'] !== 'admin') {
    header("Location: yetkisiz.php");
    exit();
}
?>
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Yönetim Paneli</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="login-container">
        <button class="main-button" onclick="window.location.href='soru_ekle.php'">Soru Ekle</button>
        <button class="main-button" onclick="window.location.href='soru_listele.php'">Soruları Listele</button>
        <button class="main-button" onclick="window.location.href='index.php'">Ana Sayfaya Dön</button>
    </div>
</body>
</html>
