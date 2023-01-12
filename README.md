# auto-otp

**still not functional**

This is an application made in Go that provides auto-typing of OTP values based on a system tray selection


## Planned Implementation

### Mockup screenshot
This is how the app is being planned to look like:

![Image showing a mockup desktop resembling MacOS with a browser window showing a website on an authentication screen asking for an OTP code of six digits to be typed. On the OS system tray, there is an item with the text "Auto OTP" and a context menu open and showing three entries: "Demo - 123-456", "Example - 777-888" and "Any Site - 000-000". On top of the OTP entries there is a title in a smaller font and grayed out displaying the text "Click to type OTP". Over the last digit of one of the entries there is a mouse pointer.](mockup.jpeg)

### Mockup configuration file

The configuration file will probably look like this:
```yaml
# config.yaml

entries:
  - title: "Demo"
    totp:
      secret: "lkaskdiuysdnn89kl23kjh5"
  - title: "Example"
    totp:
      secret: "2501254asiojnsad980cn29"
  - title: "Any Site"
    totp:
      secret: "l345232ai789jkwjklzxncz"
```