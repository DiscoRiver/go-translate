# go-translate
A multi-threaded translator using the Google API

# Disclaimer
Currently a proof-of-concept. More error sensitivity will be added. 

# Setup
Since the Google API for translation isn't free, the only additional thing you'll need is a Google credentials file, added as an environment variable like;

```
export GOOGLE_APPLICATION_CREDENTIALS="<path>/google-creds.json"
```

See here for further information: https://cloud.google.com/translate/docs/quickstart-client-libraries?authuser=1

# Usage

```
./go-translate <file> <languageID (e.g en, fr, ru)>
```
