<?php

namespace app;

static $loaded = false;
if ($loaded) {
    return true;
}

require __DIR__ . '/Utils.php';
Utils::timer();
require __DIR__ . '/Conf.php';
require __DIR__ . '/Db.php';
require __DIR__ . '/Feed.php';
require __DIR__ . '/Telegram.php';
