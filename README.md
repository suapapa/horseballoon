# horseballoon

![horseballoon](resource/horseballoon.png)

[개인방송](https://twitch.tv/suapapa)에서 사용하고 있는 한-영 음성 자동 변환기

# build and run

Dependency:

    portaudio

Clone and build:

    $ git clone https://github.com/suapapa/horseballoon
    $ go build

Fill your `API_KEY` of [kakao developers](https://developers.kakao.com/) to `env.sh`:

    $ source ./env.sh
    $ ./horseballoon
