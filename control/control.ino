#define LED 13

String received;
bool transmitted;

void setup()
{
    Serial.begin(9600);
    pinMode(LED, OUTPUT);
    received.reserve(256);
    transmitted = false;
}

void loop()
{
    if (transmitted) {
        String toSend = "what?";
        received.trim();

        if (received.equals("/on")) {
            toSend = "turning on";
            digitalWrite(LED, HIGH);
        }
        else if (received.equals("/off")) {
            toSend = "turning off";
            digitalWrite(LED, LOW);
        }

        long written = Serial.println(toSend);

        received = "";
        transmitted = false;
    }
}

void serialEvent()
{
    while (Serial.available() > 0)
    {
        char incoming = (char) Serial.read();
        received += incoming;
        if (incoming == '\n') {
            transmitted = true;
        }
    }
}
