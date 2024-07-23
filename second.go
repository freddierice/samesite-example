package main

import (
	"fmt"
	"net/http"
)

func serveSecondPage(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>Second Page</title>
</head>
<body>
    <h1>Second Page</h1>
    <button onclick="fetchNumber()">Fetch Number</button>
    <p>Number: <span id="numberDisplay"></span></p>

    <script>
    function fetchNumber() {
        var req = new XMLHttpRequest();
        req.open("POST", "https://first.trends.stream/number", true);
        req.withCredentials = true;
        req.onload = function() {
            if (req.status == 200) {
                var jsonResponse = JSON.parse(req.responseText);
                document.getElementById("numberDisplay").textContent = jsonResponse.number;
            } else {
                console.log("Error: " + req.status);
            }
        };
        req.send();
    }
    </script>
</body>
</html>
`
	fmt.Fprint(w, html)
}
