# midi-led-lights
Input midi on a raspberry pi and trigger led strips

This project was influenced by a modified Roland TD50 I saw on a Twitch, which had LEDs installed under each mesh pad. I wanted to see if I could improve upon the example I saw, and be able to trigger lighting effect according to what you play. I will be creating 

Things you will need to recreate this project:
- Something way to output a midi signal over BLE. My implementation uses the midi over bluetooth feature of the Roland TD17k module.
- A raspberry pi 3 or 4 (both have built in bluetooth capabilities).
- ws2812b LED strips of your prefered density and length. My implementation uses 60 pixel/meter strips with ~3m of strip total.
- A 5V power supply for the LED strips, Amperage will depend on the amount LEDs you use. each pixel will use between 0.1W and 0.3W depending on the color displayed.
- Wires to connect the raspberry pi and power supply to the LED strip. The gauge depends heavily on the distance of this connection.
- Soldering tools to connect your wires.

TODO: Fill out links for the above items.

TODO: Create some cool looking diagrams to demonstrate whats going on here.
