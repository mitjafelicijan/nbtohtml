run:
	go build . && duckdb -json ~/.newsboat/cache.db "SELECT * FROM rss_item WHERE unread=1;" | ./nbtohtml
