### mdtoHTMLemail

This little project is a way to send an email from a markdown file. It converts the file to HTML before sending via smtp. I've only used it to send mail from gmail.

#### Single Email
 Included is an example file easyemail.md that I used for testing. To send one email run:

```zsh
go run mdtoHTMLemail.go -filename=easyemail.md -username=johnny@gmail.com -password=secret -smtphost=smtp.gmail.com -destination=johnnyduderino@gmail.com
```

#### Multiple Emails
In order to send multiple emails at a time, for instance to upload a big directory of markdown files and send them all as email, we have to be mindful of the limitations of our smtp service. Gmail seems to disconnect our connection after about 70 emails, so I decided to split up a directory of 387 markdown files into subfolders and send each one independently. Probably a good idea to do this programmatically in go at somepoint.

Although it seems that after a couple time block sending 50 emails gmail still locks me up...

I'm in zsh so I did the following to split it up, then used mdtoHTMLemail_loop.go to send each folder independently. The mv commands won't work in regular bash.

```zsh
ls | wc -l
  387
```

387/50 = 7.74 so I should split up all these files into 8 directories.

```zsh
mkdir 1 2 3 4 5 6 7 8
mv -- *(^/[1,50]) 1/
mv -- *(^/[1,50]) 2/
mv -- *(^/[1,50]) 3/
mv -- *(^/[1,50]) 4/
mv -- *(^/[1,50]) 5/
mv -- *(^/[1,50]) 6/
mv -- *(^/[1,50]) 7/
mv -- *(^/[1,50]) 8/
```

#### Still To Do
 - Images are not attached to the email yet
 - programmatically split up large directory of md files
