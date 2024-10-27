<?php
session_start();
include('veritabani_ayarlari.php');
if ($_SESSION['rol'] !== 'admin') {
    header("Location: yetkisiz.php");
    exit();
}
$stmt = $db->query("SELECT * FROM sorular");
$sorular = $stmt->fetchAll(PDO::FETCH_ASSOC);
?>
<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Soruları Listele</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="login-container">
        <h2>Soruları Listele</h2>
        <table>
            <thead>
                <tr>
                    <th>Soru Metni</th>
                    <th>Şıklar</th>
                    <th>Doğru Cevap</th>
                </tr>
            </thead>
            <tbody>
                <?php foreach ($sorular as $soru): ?>
                <tr>
                    <td><?= htmlspecialchars($soru['soru_metin']) ?></td>
                    <td><?= htmlspecialchars($soru['siklar']) ?></td>
                    <td><?= htmlspecialchars($soru['dogru_cevap']) ?></td>
                </tr>
                <?php endforeach; ?>
            </tbody>
        </table>
        <button class="main-button" onclick="window.location.href='admin_panel.php'">Yönetim Paneline Dön</button>
    </div>
</body>
</html>
