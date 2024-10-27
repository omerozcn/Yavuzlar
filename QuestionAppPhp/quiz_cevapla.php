<?php
session_start();
include('veritabani_ayarlari.php');
$user_id = $_SESSION['user_id'];
$soru_id = $_POST['soru_id'];
$cevap = $_POST['cevap'];
$stmt = $db->prepare("SELECT dogru_cevap FROM sorular WHERE id = ?");
$stmt->execute([$soru_id]);
$dogru_cevap = $stmt->fetchColumn();
if ($cevap === $dogru_cevap) {
    $db->prepare("UPDATE users SET puan = puan + 10 WHERE id = ?")->execute([$user_id]);
    $db->prepare("INSERT INTO submissions (user_id, soru_id, dogru_mu) VALUES (?, ?, 1)")->execute([$user_id, $soru_id]);
    echo "Tebrikler, cevabınız doğru!";
} else {
    $db->prepare("INSERT INTO submissions (user_id, soru_id, dogru_mu) VALUES (?, ?, 0)")->execute([$user_id, $soru_id]);
    echo "Maalesef, cevabınız yanlış.";
}
?>
