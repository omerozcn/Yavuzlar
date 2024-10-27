<?php
try {
    $db = new PDO('sqlite:veritabani.sqlite3');
    $db->setAttribute(PDO::ATTR_ERRMODE, PDO::ERRMODE_EXCEPTION);

    $db->exec("CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY,
        isim TEXT,
        rol TEXT,
        puan INTEGER DEFAULT 0
    )");

    $db->exec("CREATE TABLE IF NOT EXISTS sorular (
        id INTEGER PRIMARY KEY,
        soru_metin TEXT,
        siklar TEXT,
        dogru_cevap TEXT
    )");

    $db->exec("CREATE TABLE IF NOT EXISTS submissions (
        id INTEGER PRIMARY KEY,
        user_id INTEGER,
        soru_id INTEGER,
        dogru_mu INTEGER,
        tarih TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )");
    $dbFile = __DIR__ . '/veritabani.db';
} catch (PDOException $e) {
    echo "Veritabanı bağlantı hatası: " . $e->getMessage();
    exit();
}
?>
