# Rosey the Mastodon Bot
This is some skeleton code for building a bot for Mastodon. The idea was to allow the bot to take in different data streams from files, network signals, and Mastodon messages to then make a toot or complete an action based on the data. Currently the only finished feature of the bot is tooting one line at a time from a file based on a cron timer.

I do not intend to do much more work on this, but if I do, it will be to build out a feature that allows Rosey to monitor recent posts from accounts she follows and repost content with selected tags. 

# Features
- [x] Read lines sequentially from a file on a timer
- [x] Repost manual input from user DMs
- [ ] Relay data as it arrives to a TCP port with authentication
- [ ] Find posts with specific tags and repost them
- [ ] Remote administration

# Goodbye
Goodbye for now, Rosey...

![Rosey standing alone at a bus stop.](https://tf-cmsv2-smithsonianmag-media.s3.amazonaws.com/legacy_blog/rosey-470x251.jpg)
