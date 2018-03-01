var form = $('#trailname');
var resultStack = [];
var prevStartText = "";

//Poll the API on load and enable submit after we've recieved response
$( window  ).on( "load", function() {
	pollApi();	    
});

function pollApi() {
	console.log("polling api");
	// Create JSON data to submit to api
	var jsonObj = new Object();
	jsonObj.poll = "true";
	var jsonStr = JSON.stringify(jsonObj);
	console.log(JSON.parse(jsonStr));
	// Create ajax request and submit with long timeout
	$.ajax({
		type: 'POST',
		url: $(form).attr('action'),
		data: jsonStr,
		contentType: "application/json; charset=utf-8"
	}).done(function(responseData) {
		console.log(responseData);
		var pollResult = responseData.result;
		console.log(pollResult)
		if (pollResult && pollResult == true) {
			// API poll reports API is g2g, enable submit and remove loader
			$("#submit").removeClass('disable')
			$("#submit").empty().append("Go")
		} else {
				$("#trailNames").empty().append("Oops! Our trail name bots aren't home, please try again later!");
			}
	}).fail(function(data) {
		if (data.responseText !== '') {
				$("#trailNames").empty().append(data.responseText);
			} else {
				$("#trailNames").empty().append("Oops! Our trail name bots aren't home, please try again later!");
			}
	});
	// On done() remove disable class from submit and replace btn-loader div with "Go"
}

$(form).submit(function(event){
	event.preventDefault();
	startText = $("#starttext").val();
	// If there are results from a previous query left in the stack, use em
	// Otherwise, submit a new query
	if ($('#submit').hasClass('disable')) {
		return false;
	}
	else if (resultStack.length > 0 && startText == prevStartText) {
		namesContent = "Your trail name is " + resultStack.pop();
		$("#trailNames").empty().append(namesContent);
	} else {
		var formData = $(form).serialize();
		// Submit the form using AJAX.
		$.ajax({
			type: 'POST',
			url: $(form).attr('action'),
			data: formData
		}).done(function(responseData) {
			// you can see the result from the console
			// tab of the developer tools
			console.log(responseData);
			var namesContent = ""
			prevStartText = startText
			resultStack = responseData.result
			if (resultStack && resultStack.length > 0) {
				namesContent = "Your trail name is " + resultStack.pop();
				$("#trailNames").empty().append(namesContent);
			} else {
				$("#trailNames").empty().append("Oops! We couldn't grab your trail name at this time, please try again later!");
			}	
		}).fail(function(data) {
			if (data.responseText !== '') {
				$("#trailNames").empty().append(data.responseText);
			} else {
				$("#trailNames").empty().append("Oops! We couldn't grab your trail name at this time, please try again later!");
			}
		});
	}
});
