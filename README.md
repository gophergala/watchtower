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

#### Message formats

##### TCP/IP & UDP
s
	Sender (uint32) - channel (uint32) - reserved (uint16) - message length (uint32) - message (bytes)

Everything encoded as little-endian (where applicable)

