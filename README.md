# The Morse Machine
An experimental, personal, twitter bot written in Go.

## Usage
To use this bot simply send a tweet to the handle `@TheMorseMachine` on twitter

It will take the contents of your tweet and reply back to you with that tweet translated to Morse Code.
For example:

> @TheMorseMachine SOS
>
>***************************
>
> To be translated: "SOS"
>
> Result:
> `...---...`

#### Information
This bot uses the Twitter api, a morse code translation api, and utilizes two Go libraries that make interfacing with the Twitter API a little bit easier.
Thanks to [Fallenstedt/twitter-stream](https://github.com/Fallenstedt/twitter-stream) and [michimani/gotwi](https://github.com/michimani/gotwi) for their work to make our lives easier. Go check out their Githubs!
