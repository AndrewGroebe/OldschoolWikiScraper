# OldschoolWikiScraper

A web scraper that can be used to parse data from the wiki for one of my favorite childhood games, Oldschool RuneScape.

The scraper visits the raw html of each url (by appending ```?action=raw``` to the end of each url, and parses information therefrom.

The use of goroutines, one per each url, allows for pretty fast concurrent parsing. This application averages approximately 700ms per 10 urls, whereas sequential code (in Java, ..., or even just in Go), averages 4.5 seconds per 10 urls. There are currently 21,000+ item components in the OldSchool RuneScape cache (revision 181), and each with a corresponding wiki page, so performance matters.
