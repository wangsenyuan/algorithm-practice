# Little Lesha and Music

Little Lesha loves listening to music via his smartphone. Because the phone doesn't have much memory, he streams his favorite songs through the social network InTalk.

Unfortunately, the internet is slow in Ekaterinozavodsk, so downloading a song takes time. The song lasts for `T` seconds. Lesha downloads the first `S` seconds and starts playback. Whenever playback reaches a point that has not been downloaded yet, he immediately restarts the song from the beginning. The downloaded portion remains on the phone, and downloading continues from where it left off. This repeats until the song finishes downloading and Lesha listens to it completely. During `q` seconds of real time, the connection downloads `q - 1` seconds of the track.

Determine how many times Lesha starts the song, including the first start.

## Input

A single line with three integers `T`, `S`, `q` (`2 ≤ q ≤ 10^4`, `1 ≤ S < T ≤ 10^5`).

## Output

Output a single integer — the number of times the song restarts.

## Examples

```
Input
5 2 2
Output
2
```

```
Input
5 4 7
Output
1
```

```
Input
6 2 3
Output
1
```

## Note

In the first test, the song plays twice as fast as it downloads. During the first four seconds playback catches up to the download, forcing a restart. After two more seconds, the song finishes downloading, so Lesha starts it twice.

In the second test, the song is almost fully downloaded, so Lesha starts it only once.

In the third test, downloading and playback finish simultaneously, so the song is not restarted.


### ideas
1. 一开始5秒的，下载了2秒，然后开始play，然后播放了2秒，此时下载了2-1秒，然后，所以还可以听一秒，但是这时候，没有下载; 所以在3秒的时候，又要重新开始，到3秒的时候正好下载完
2. 第二个例子，[5, 4, 7], 在4秒的时候开始播放，还剩下一秒的时间，如果用时7秒的时候，可以下载6秒
3. 那么用时4秒，4 * 6 / 7 秒？
4. 