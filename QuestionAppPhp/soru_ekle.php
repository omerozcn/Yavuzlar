<?php
session_start();
include('veritabani_ayarlari.php');
if ($_SESSION['rol'] !== 'admin') {
    header("Location: yetkisiz.php");
    exit();
}
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $soru_metin = $_POST['soru_metin'];
    $siklar = json_encode([
        'A' => $_POST['sikA'],
        'B' => $_POST['sikB'],
        'C' => $_POST['sikC'],
        'D' => $_POST['sikD']
    ]);
    $dogru_cevap = $_POST['dogru_cevap'];
    $stmt = $db->prepare("INSERT INTO sorular (soru_metin, siklar, dogru_cevap) VALUES (?, ?, ?)");
    $stmt->execute([$soru_metin, $siklar, $dogru_cevap]);
}
?>
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Soru Ekle</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="login-container">
        <form method="POST">
            <h2>Soru Ekle</h2>
            <input type="text" name="soru_metin" class="input-bar" placeholder="Soru metnini girin">
            <input type="text" name="sikA" class="input-bar" placeholder="Şık A">
            <input type="text" name="sikB" class="input-bar" placeholder="Şık B">
            <input type="text" name="sikC" class="input-bar" placeholder="Şık C">
            <input type="text" name="sikD" class="input-bar" placeholder="Şık D">
            <label for="dogru_cevap">Doğru Cevap:</label>
            <select name="dogru_cevap" class="input-bar">
                <option value="A">Şık A</option>
                <option value="B">Şık B</option>
                <option value="C">Şık C</option>
                <option value="D">Şık D</option>
            </select>
            <button type="submit" class="main-button">Soru Ekle</button>
            <button class="main-button" onclick="window.location.href='admin_panel.php'">Yönetim Paneline Dön</button>
        </form>
    </div>
</body>
</html>
