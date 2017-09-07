#define LED 13

void setup()
{
    Serial.begin(9600);
    pinMode(LED, OUTPUT);
}

void loop()
{
    if (Serial.available() > 0) {
        String received = Serial.readString();
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

        Serial.println(toSend);
    }
}
