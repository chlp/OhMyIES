<?php

namespace app;

class Telegram
{
    private const SEND_MSG_URL_TEMPLATE = "https://api.telegram.org/bot%s/sendMessage?%s";
    private readonly string $token;

    public function __construct()
    {
        $this->token = (string)(Conf::loadApiConf()['telegram-token'] ?? '');
        if ($this->token === '') {
            Utils::error('Wrong Telegram token');
        }
    }

    public function sendMsg(string $chatId, string $msg): bool
    {
        $query = http_build_query([
            'chat_id' => $chatId,
            'parse_mode' => 'markdown',
            'text' => $msg,
        ]);
        $url = sprintf(self::SEND_MSG_URL_TEMPLATE, $this->token, $query);
        $data = file_get_contents($url);
        if (json_decode($data, true)['ok'] !== true) {
            Utils::log(json_encode(['Telegram_sendMsg_error', $url]));
            return false;
        }
        return true;
    }
}