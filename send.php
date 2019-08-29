<?php
ini_set('display_errors', 1);
ini_set('display_startup_errors', 1);
error_reporting(E_ALL);

echo '<pre>';
$url = "http://localhost:8080/log";

$data = array
(
    'channel' => 'WS3',
    'remote_addr' => '172.21.0.1',
    'user_agent' => 'PostmanRuntime/7.6.1',
    'ota_action' => 'VehAvailRateRQ',
    'company_id' => '0',
    'agency_id' => '1884523',
    'rate_qualifier' => 'FTC13',
    'pickup_datetime' => '2019-10-24T10:00:00',
    'return_datetime' => '2019-10-25T10:00:00',
    'pickup_location' => 'GRU',
    'return_location' => 'GRU',
    'condutor' => 'DANIEL SILVA',
    'duration' => '0.78868699073792',
    'success' => '1',
    'total' => '0',
    'error_message' => '',
    'error_code' => '0'
);

$payload = json_encode($data);

//print_r($payload); exit;

$ch = curl_init();

curl_setopt( $ch, CURLOPT_URL, $url );
curl_setopt( $ch, CURLOPT_POST, true );
curl_setopt( $ch, CURLOPT_POSTFIELDS, $payload );
curl_setopt( $ch, CURLOPT_RETURNTRANSFER, true);

$result = curl_exec($ch);

print_r($result);


?>