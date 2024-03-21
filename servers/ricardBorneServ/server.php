<?php 

set_time_limit ( 0 );
ini_set('memory_limit', '-1');

require __DIR__.'/vendor/autoload.php';

use Ratchet\Server\IoServer;
use Ratchet\Http\HttpServer;
use Ratchet\WebSocket\WsServer;
use Ratchet\MessageComponentInterface;
use Ratchet\ConnectionInterface;

define('APP_PORT', 8837);

class ServerImpl implements MessageComponentInterface {
    protected $clients;
	
	//protected $player;
	
    public function __construct() {
        $this->clients = new \SplObjectStorage;
		//$this->player=Array();
    }

    public function onOpen(ConnectionInterface $conn) {
        $this->clients->attach($conn);
        echo "New connection! ({$conn->resourceId}).\n";
		
		//$this->player[$conn->resourceId]=[$conn,$conn->resourceId, "", "-400,0" ];
		
    }

    public function onMessage(ConnectionInterface $conn, $msg) {
        echo sprintf("New message from '%s': %s\n\n\n", $conn->resourceId, $msg);
        //GOGAME|idBORNE|IdQR|
		
		$Bmesg="";
		$Pmesg="";
		$PmesgId=0;
		
		$spitMsg = explode("|", $msg);
		
		if(count($spitMsg)<=1)
		{
			$spitMsg[]="NOTING";
		}
		
		switch ($spitMsg[0]) 
		{
			case 'UPDATE':
				
				//$Pmesg="PONG";
				$Bmesg= $msg;
				
			break;
			
			
			default:
				echo "unable to interprete : ".$msg;
				$Bmesg="unable to interprete :".$msg;
			
		}
		
	
		
		if($Bmesg!="")
		{
			echo "send BROADCAST: ".$Bmesg;
			foreach ($this->clients as $client) { // BROADCAST
				if ($conn !== $client) {
					$client->send($Bmesg);
				}
			}
		}
		
		if($Pmesg!="")
		{
			//echo "send Private to : ".$this->player[$PmesgId][1];
			$conn->send($Pmesg);
			
		}
		
		
		
    }
	
	
	

    public function onClose(ConnectionInterface $conn) {
		
		
		
		
		
        $this->clients->detach($conn);
		
		//unset($this->player[$conn->resourceId]);  
        echo "Connection {$conn->resourceId} is gone.\n";
    }

    public function onError(ConnectionInterface $conn, \Exception $e) {
        echo "An error occured on connection {$conn->resourceId}: {$e->getMessage()}\n\n\n";
        $conn->close();
    }
}

$server = IoServer::factory(
    new HttpServer(
        new WsServer(
            new ServerImpl()
        )
    ),
    APP_PORT
);
echo "Server created on port " . APP_PORT . "\n\n";
$server->run();
