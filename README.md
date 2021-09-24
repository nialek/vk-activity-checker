# VK activity checker
A very simple service that parses `m.vk.com/id...` page and returns activity information of that user.

Returns the result in Russian (for now, since I can't figure out how to change this).

Made for study, fun, and some other of my projects.

### Endpoints
`GET /activity` returns a json object `{"activity":"..."}` or an error message
(if the requested page is visible only to authenticated users)
