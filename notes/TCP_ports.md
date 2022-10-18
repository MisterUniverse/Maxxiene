# TCP ports 

On your web server the service that you are providing should be on the port that it goes with ex:

```
{
    Incoming E-mail => Port: 25
    Login => Port: 23
    Web Server => Port: 80 <- for unsecure http mostly used in dev mode
    Web Server => Port: 443 <- for secure https mostly used in production code
    SMTP Mail Box server => Port: 109/110 etc..
}

```