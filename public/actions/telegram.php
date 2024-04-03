<?php

$input = require __DIR__ . '/actions.php';

\app\Utils::log(json_encode(['input' => $input, 'query' => $_GET]));

$action = (string)($input['action'] ?? '');
switch ($action) {
    case 'getUpdate':
        echo 'update';
        break;
    default:
        error('wrong action', 404);
}
