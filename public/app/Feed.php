<?php

class Feed
{
    private Db $db;

    public function __construct(
    )
    {
        $this->db = Db::get();
    }
}