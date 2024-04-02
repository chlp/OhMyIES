<?php

namespace app;

class Feed
{
    private readonly Db $db;

    public function __construct()
    {
        $this->db = Db::get();
    }

    public function getFeeds(): array
    {
        $query = "SELECT * FROM `feeds`";
        return $this->db->select($query);
    }
}