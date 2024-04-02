<?php

$input = require __DIR__ . '/actions.php';

$action = (string)($input['action'] ?? '');
switch ($action) {
    case 'getUpdate':
        echo 'update';
        break;
    default:
        error('wrong action', 404);
}
