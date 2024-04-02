<?php

require(__DIR__ . '/Html2Text.php');

$botApiToken = '';
$chatId = '';
$iesRssUrl = 'https://sms.schoolsoft.se/engelska/jsp/public/right_public_parent_rss.jsp?key=XXX&key2=YYY';

$lastReadGuidFile = __DIR__ . '/feed_' . md5($iesRssUrl);
$lastReadGuid = '';
if (file_exists($lastReadGuidFile)) {
    $lastReadGuid = file_get_contents($lastReadGuidFile);
}

$feed = simplexml_load_file($iesRssUrl);
if (empty($feed)) {
    exit;
}

$newGuid = null;
$i = 0;
$messagesToSend = [];
foreach ($feed->channel->item as $item) {
    if ((string)$item->guid === $lastReadGuid) {
        break;
    }

    if ($i === 0) {
        $newGuid = (string)$item->guid;
    }
    $i++;

    $message = (new DateTime($item->pubDate))->format('Y-m-d H:i:s') . "\n*" . (string)$item->title . '*';
    if (strlen($item->description) > 0) {
        $html = new \Html2Text\Html2Text($item->description);
        $message .= "\n\n" . $html->getText();
    }
    $message = str_replace(['_', '*', '`', '['], ['\_', '\*', '\`', '\['], $message);
    $messagesToSend[] = $message;
}

$success = true;

$rpsLimiter = 0;
for ($i = count($messagesToSend) - 1; $i >= 0; $i--) {
    $rpsLimiter++;
    if ($rpsLimiter > 5) {
        sleep(60);
        $rpsLimiter = 0;
    }

    $query = http_build_query([
        'chat_id' => $chatId,
        'parse_mode' => 'markdown',
        'text' => $messagesToSend[$i],
    ]);
    $iesRssUrl = "https://api.telegram.org/bot{$botApiToken}/sendMessage?{$query}";
    $data = file_get_contents($iesRssUrl);
    if (json_decode($data, true)['ok'] !== true) {
        $success = false;
        var_dump($iesRssUrl);
        var_dump($data);
        exit;
    }
}

if ($success && $newGuid !== null) {
    file_put_contents($lastReadGuidFile, $newGuid);
}
