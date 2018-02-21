var form = $('#trailname');
var resultStack = [];
var prevStartText = "";
$(form).submit(function(event){
	event.preventDefault();
	startText = $("#starttext").val();
	// If there are results from a previous query left in the stack, use em
	// Otherwise, submit a new query
	if (resultStack.length > 0 && startText == prevStartText) {
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
			//for(var i = 0; i < .length; i++){
			//namesContent += "<br>"+names[i];
			//}
			namesContent = "Your trail name is " + resultStack.pop();
			//$("#trailNames").html(responseData.result);
			$("#trailNames").empty().append(namesContent);

		}).fail(function(data) {
			if (data.responseText !== '') {
				$("#trailNames").empty().append(data.responseText);
			} else {
				$("#trailNames").empty().append("Oops! We couldn't grab your trail name at this time, please try again later!");
			}
		});
	}
});
