<!DOCTYPE html>
<html lang="tr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Scoreboard</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="login-container">
        <h2>Puan Tablosu</h2>
        <table>
            <thead>
                <tr>
                    <th>İsim</th>
                    <th>Puan</th>
                </tr>
            </thead>
            <tbody>
                <?php
                include('veritabani_ayarlari.php');
                $stmt = $db->query("SELECT isim, puan FROM users WHERE rol = 'ogrenci' ORDER BY puan DESC");
                while ($row = $stmt->fetch(PDO::FETCH_ASSOC)) {
                    echo "<tr><td>{$row['isim']}</td><td>{$row['puan']}</td></tr>";
                }
                ?>
            </tbody>
        </table>
        <button class="main-button" onclick="window.location.href='index.php'">Ana Sayfaya Dön</button>
    </div>
</body>
</html>
