<?php


error_reporting(E_ALL);
ini_set("display_errors", 1);


echo "SALUT TOTOT<br>\n";

echo file_get_contents("http://localhost:4365/verifyfrom?adr=0x4f60ADa30Ca1FF8a3bCA7A7893bb2B60157F95f8&network=polygon&ca=0x19d6cfb8c9f053f3a37e29370764884ee9411aa3");

?>