<?php
session_start();
include('veritabani_ayarlari.php');
$user_id = $_SESSION['user_id'];
$stmt = $db->prepare("
    SELECT * FROM sorular 
    WHERE id NOT IN (SELECT soru_id FROM submissions WHERE user_id = ?)
    ORDER BY RANDOM() LIMIT 1
");
$stmt->execute([$user_id]);
$soru = $stmt->fetch(PDO::FETCH_ASSOC);
echo json_encode($soru);
?>
