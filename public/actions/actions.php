<?php

require __DIR__ . '/../../app/bootstrap.php';

if ($_SERVER['REQUEST_METHOD'] !== 'POST') {
    \app\Utils::error('Method not allowed', 405);
}

return json_decode(file_get_contents('php://input'), true);