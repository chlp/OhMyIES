<?php

namespace app;

use DateTime;

class Utils
{
    private const ID_LENGTH = 32;

    static public function genUuid(): string
    {
        return sprintf('%04x%04x%04x%04x%04x%04x%04x%04x',
            mt_rand(0, 0xffff), mt_rand(0, 0xffff),
            mt_rand(0, 0xffff),
            mt_rand(0, 0x0fff) | 0x4000,
            mt_rand(0, 0x3fff) | 0x8000,
            mt_rand(0, 0xffff), mt_rand(0, 0xffff), mt_rand(0, 0xffff)
        );
    }

    static public function isUuid(string $id): bool
    {
        return preg_match('/^[a-z0-9]{' . self::ID_LENGTH . '}$/', $id) === 1;
    }

    static public function timer(): float
    {
        static $start = null;
        if ($start == null) {
            $start = microtime(true);
            return 0;
        }
        return $start;
    }

    static public function log(string $str): void
    {
        $msg = (DateTime::createFromFormat('U.u', microtime(true)))->format('Y-m-d H:i:s.u');
        $timer = self::timer();
        if ($timer !== 0.0) {
            $msg .= ' (' . number_format(microtime(true) - $timer, 3, '.', '') . ')';
        }
        $msg .= ': ';
        $msg .= $str;
        error_log($msg);
    }
}
