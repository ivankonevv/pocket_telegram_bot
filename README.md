# Pocket Bot


It is a project written in Golang language.
<p>Pocket-bot is a bot for conveniently saving useful links to your Pocket account. 
The bot was written from the video <a href="https://github.com/zhashkevych" target="blank">@zhashkevych</a>. 
The bot uses the <a href="https://github.com/zhashkevych/go-pocket-sdk" target="blank">SDK library</a> to work with pocket api. 
Boltdb is used as a database.
<p><h3>Installation:</h3>
<p>Note: To use, you need a telegram bot token, pocket account and a consumer key. Pocket api <a href="https://getpocket.com/developer/?src=footer_v2" target="blank">documentation</a>.</p>
<p>1. Install the required modules:</p>
<code>go get all</code>
<p>2. Create an .env file in your project directory and put your values there.</p><b>Structure:</b>
<pre><code>TOKEN="your bot token"
CONSUMER_KEY="your consumer key"
AUTH_SERVER_URL="your auth redirect url"</code>
</pre>
<p>3. After this point, enter the command:</p>
<code>make run</code>
<h3>Done</h3>


