## Watchtower

Watchtower is a service for easily interacting between many different kinds of devices in (almost) real-time. Think https://pusher.com but for devices ranging from PC’s all the way down to Arduinos with only the barest of network stacks.

It provides **channels** you can communicate on, which carry events from senders to any amount of receivers, so your devices can talk to each other in any kind of setup you choose. It supports endpoints ranging from UDP, TCP all the way up to HTTP - making it easy to choose the right one.

You get freedom to choose the format of your messages. You like JSON? You get JSON! You like binary? You get binary!

#### How should you use Watchtower?

	* You download it and set up a server
	* All your flower-watering Arduinos register and joins a channel
	* Your Macbook joins the same channel
	* Your Macbook starts playing Flight of the Valkyries on max volume
	* One by one, the Arduinos start spraying water all over
	* Your wife kicks you out for ruining the curtains for the umpteenth time
	* You are now living on the street, trading Go Programming for food (but you have an awesome flower watering system, if only one of those VC's would invest...)
	
#### So how do I actually use it?

	* Register as a new sender
	* Join some channels
	* Broadcast or send some messages
	* Act on / display received messages
	
The rest, as they say, is history.

#### Ideas

	* Connect all your embedded devices to different channels - you now have a communication protocol that ties everything together
	* Use it as an extremely lightweight IRC server, for chatting
	* Use it to launch nuclear weapons, but require two or more registered users to confirm the launch


### Examples (HTTP)

##### Register as a new sender (HTTP)

POST to http://your-watchtower-url/register

	{
	  "secret_key": "watchtowers_are_awesome"
	}

Watchtower will reply:

HTTP 200

	{
	  "id": 72
	}

Or any of the following

	* HTTP 401 - invalid or missing secret key

##### List channels

GET to http://your-watchtower-url/channels?sender=72

Watchtower will reply:

HTTP 200

	{
	  "channels" [5,6]
	}

Or any of the following

	* HTTP 401 - invalid or missing sender
	* HTTP 204 - no channels created yet


	
##### Join a channel (HTTP stream)

GET to http://your-watchtower-url/channels/join?channels=5,6&sender=72

Watchtower will open a HTTP streaming endpoint towards you, using '\n' as a keep-alive heartbeat. Messages will be sent as JSON.


##### Join a channel asynchronusly (HTTP callbacks)

POST to http://your-watchtower-url/channels/join/async

	{
	  "sender": 72,
	  "channels": [5],
	  "callback_url": http://client-url.com/callbacks
	}

Watchtower will reply:

HTTP 200 (channel joined or created)

Or any of the following:

	* HTTP 400 - "invalid callback URL"


##### Broadcast a message on a channel (HTTP)

POST to http://your-watchtower-url/broadcast

{
    "channels": [5],
    "sender": 72,
    "message": "all your base are belong to us"
}


Watchtower might reply HTTP 200 (message broadcasted)
	
Or any of the following

	* HTTP 400 - tried to send to unjoined channel
	* HTTP 400 - invalid callback URL
	* HTTP 500 - broadcast failed 

Assuming success, that same message will now be received by all subscribers on channel 5. 

##### Send a message on a channel to specific subscriber(s)

POST to http://your-watchtower-url/send

{
    "channels": [5],
    "sender": 72,
    "receivers": [52, 34],
    "message": "your specifical bases are belong to us"
}

Watchtower might reply HTTP 200 (message sent to all parties)

Or any of the following

	* HTTP 400 - tried to send to unjoined channel
	* HTTP 400 - one or more of the receivers have not joined the channel
	* HTTP 500 - send failed
