<?php

require __DIR__ . '/bootstrap.php';

// 1. get all feeds
// 2. for a feed:
// 2.1. read rss
// 2.2. if exist msg with new id and dt
// 2.3. load chats for this feed
// 2.4. for a msg:
// 2.4.1. filter
// 2.4.2. send msg to filtered chats
// 2.4.3. update chat: last_msg_error, last_msg_dt, msg_count
// 2.4.4. on success update feed: rss_last_id, rss_last_dt