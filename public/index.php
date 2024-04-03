<!DOCTYPE html>
<html lang="en">
<head>
    <title>OhMyIES</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="icon" type="image/png" href="favicon.ico">
    <link rel="stylesheet" href="style.css?<?= md5_file(__DIR__ . '/style.css') ?>">
</head>
<body>

<?php
require __DIR__ . '/../app/bootstrap.php';
var_dump((new \app\Telegram())->sendMsg('-4135383525', 'test msg'));
?>

<h1>OhMyIES</h1>

<div>
    Want some API?
</div>

</body>
</html>
